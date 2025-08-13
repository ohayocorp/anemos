package client

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/ohayocorp/anemos/pkg/core"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"sigs.k8s.io/cli-utils/pkg/kstatus/polling"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/aggregator"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/collector"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/engine"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/event"
	"sigs.k8s.io/cli-utils/pkg/kstatus/polling/statusreaders"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/cli-utils/pkg/object"
)

func (client *KubernetesClient) Wait(objects object.ObjMetadataSet, status status.Status, timeout time.Duration) error {
	if len(objects) == 0 {
		return nil
	}

	poller, err := polling.NewStatusPollerFromFactory(client.Factory, polling.Options{
		CustomStatusReaders: []engine.StatusReader{
			statusreaders.NewGenericStatusReader(client.Mapper, readStatus),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create status poller: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	eventsChannel := poller.Poll(ctx, objects, polling.PollOptions{
		PollInterval: 2 * time.Second,
	})

	initialObjects := append(make([]object.ObjMetadata, 0, len(objects)), objects...)
	statusCollector := collector.NewResourceStatusCollector(objects)
	done := statusCollector.ListenWithObserver(eventsChannel, statusObserver(initialObjects, cancel, status))
	<-done

	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("timed out waiting for resources to reach desired state after %v", timeout)
	}

	if statusCollector.Error != nil {
		return statusCollector.Error
	}

	return nil
}

func statusObserver(initialObjects object.ObjMetadataSet, cancel context.CancelFunc, desired status.Status) collector.ObserverFunc {
	successfulResources := make(map[string]bool)

	return func(statusCollector *collector.ResourceStatusCollector, _ event.Event) {
		var rss []*event.ResourceStatus
		var nonDesiredResources = make(map[string]*event.ResourceStatus)

		for _, rs := range statusCollector.ResourceStatuses {
			if rs == nil {
				continue
			}
			// If a resource is already deleted before waiting has started, it will show as unknown
			// this check ensures we don't wait forever for a resource that is already deleted
			if rs.Status == status.UnknownStatus && desired == status.NotFoundStatus {
				continue
			}

			if rs.Status != status.UnknownStatus {
				initialObjects = slices.DeleteFunc(initialObjects, func(identifier object.ObjMetadata) bool {
					return identifier == rs.Identifier
				})
			}

			rss = append(rss, rs)
			if rs.Status != desired {
				nonDesiredResources[rs.Identifier.String()] = rs
			} else if !successfulResources[rs.Identifier.String()] {
				// Deletion is already waited for with foreground propagation, no need to log again.
				if rs.Status != status.NotFoundStatus {
					slog.Info(
						"Resource ${kind}/${name} is in desired state: ${message}",
						slog.String("name", rs.Identifier.Name),
						slog.String("kind", rs.Identifier.GroupKind.Kind),
						slog.String("message", rs.Message),
					)
				}

				// Mark the resource as successful to avoid logging it again
				// in the next iteration.
				successfulResources[rs.Identifier.String()] = true
			}
		}

		if aggregator.AggregateStatus(rss, desired) == desired {
			cancel()
			return
		}

		maps.DeleteFunc(nonDesiredResources, func(identifier string, status *event.ResourceStatus) bool {
			return slices.Contains(initialObjects, status.Identifier)
		})

		if len(nonDesiredResources) > 0 {
			for _, key := range core.SortedKeys(nonDesiredResources) {
				value := nonDesiredResources[key]

				slog.Info(
					"Waiting for resource ${kind}/${name}, expected=${expectedStatus}, actual=${actualStatus}: ${message}",
					slog.String("name", value.Identifier.Name),
					slog.String("kind", value.Identifier.GroupKind.Kind),
					slog.Any("expectedStatus", desired),
					slog.Any("actualStatus", value.Status),
					slog.String("message", value.Message))
			}
		}
	}
}

func (client *KubernetesClient) WaitDocuments(documents []*core.Document, sts status.Status, timeout time.Duration) error {
	// Create a buffer to store serialized documents.
	buffer := bytes.NewBuffer(nil)
	for _, document := range documents {
		serializedDocument := core.SerializeToYaml(document)
		fmt.Fprintf(buffer, "---\n%s\n", serializedDocument)
	}

	namespace, _, err := client.Factory.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return fmt.Errorf("failed to get namespace from kubeconfig: %w", err)
	}

	builder := client.Factory.NewBuilder().
		ContinueOnError().
		Flatten().
		NamespaceParam(namespace).
		DefaultNamespace().
		Unstructured().
		Stream(bytes.NewBufferString(buffer.String()), "")

	infos, err := builder.Do().Infos()
	if err != nil {
		return fmt.Errorf("failed to build resource infos: %w", err)
	}

	identifiers := []object.ObjMetadata{}
	for _, info := range infos {
		identifier := object.ObjMetadata{
			Namespace: info.Namespace,
			Name:      info.Name,
			GroupKind: info.Mapping.GroupVersionKind.GroupKind(),
		}

		identifiers = append(identifiers, identifier)
	}

	return client.Wait(identifiers, sts, timeout)
}

func readStatus(u *unstructured.Unstructured) (*status.Result, error) {
	result, err := status.Compute(u)
	if err != nil {
		return nil, err
	}

	gvk := u.GroupVersionKind()
	if gvk.Kind == "Job" && gvk.Group == "batch" {
		setJobStatus(result)
	}

	return result, nil
}

func setJobStatus(result *status.Result) {
	if result.Status != status.CurrentStatus {
		return
	}

	// Current status can mean the job is running/suspended/completed. We only want to accept
	// completed jobs as current and wait on other conditions.
	if strings.Contains(result.Message, "suspended") {
		result.Status = status.InProgressStatus
	}

	if strings.Contains(result.Message, "in progress") {
		result.Status = status.InProgressStatus
	}
}
