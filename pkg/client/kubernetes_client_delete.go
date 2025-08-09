package client

import (
	"context"
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"text/tabwriter"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/kubectl/pkg/cmd/apply"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func (client *KubernetesClient) Delete(
	applySetParentName string,
	applySetParentNamespace string,
	skipConfirmation bool,
) error {
	applySetParentRef, err := client.getApplySetParentRef(applySetParentName, applySetParentNamespace)
	if err != nil {
		return err
	}

	applySetRestClient, err := client.getApplySetRestClient(applySetParentRef)
	if err != nil {
		return err
	}

	// Check if the apply set parent exists.
	applySetHelper := resource.NewHelper(applySetRestClient, applySetParentRef.RESTMapping)
	_, err = applySetHelper.Get(applySetParentRef.Namespace, applySetParentRef.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			return fmt.Errorf("apply set %s is not found in namespace %s", applySetParentRef.Name, applySetParentRef.Namespace)
		}

		return err
	}

	tooling := getTooling()
	applySet := apply.NewApplySet(applySetParentRef, tooling, client.Mapper, applySetRestClient)

	// Fetch the resources associated with the apply set to find the objects to prune.
	err = applySet.BeforeApply(nil, cmdutil.DryRunClient, metav1.FieldValidationIgnore)
	if err != nil {
		return err
	}

	objectsToDelete, err := applySet.FindAllObjectsToPrune(context.TODO(), client.DynamicClient, sets.New[types.UID]())
	if err != nil {
		return fmt.Errorf("failed to find objects to delete: %w", err)
	}

	// Sort the objects to delete by namespace and name for better readability.
	sort.Slice(objectsToDelete, func(i, j int) bool {
		if objectsToDelete[i].Namespace != objectsToDelete[j].Namespace {
			return objectsToDelete[i].Namespace < objectsToDelete[j].Namespace
		}

		return objectsToDelete[i].Name < objectsToDelete[j].Name
	})

	if len(objectsToDelete) == 0 {
		slog.Info("No objects to delete. Will delete the apply set parent only.")
	} else {
		slog.Info("Deleting objects for apply set ${name}:", slog.String("name", cleanupApplySetParentName(applySetParentRef.Name)))

		builder := &strings.Builder{}
		writer := tabwriter.NewWriter(builder, 1, 1, 2, ' ', 0)
		fmt.Fprintf(writer, "\x1b[00m%s\t%s\t%s\n", "NAMESPACE", "RESOURCE", "NAME")

		for _, object := range objectsToDelete {
			fmt.Fprintf(writer, "\x1b[31m%s\t%s\t%s\x1b[0m\n", object.Namespace, object.Mapping.Resource.Resource, object.Name)
		}

		writer.Flush()

		for _, line := range strings.Split(builder.String(), "\n") {
			if line == "" {
				continue
			}

			slog.Info(line)
		}
	}

	if !skipConfirmation {
		confirmed, err := confirmChanges()
		if err != nil {
			return err
		}

		if !confirmed {
			return fmt.Errorf("aborting delete operation due to user confirmation")
		}
	}

	for _, object := range objectsToDelete {
		restClient, err := client.getRestClient(object.Mapping)
		if err != nil {
			return err
		}

		deletePropagation := metav1.DeletePropagationForeground

		helper := resource.NewHelper(restClient, object.Mapping)
		_, err = helper.DeleteWithOptions(object.Namespace, object.Name, &metav1.DeleteOptions{
			PropagationPolicy: &deletePropagation,
		})
		if err != nil {
			if errors.IsNotFound(err) {
				slog.Warn("Object not found, skipping deletion: ${resource}/${name}, namespace: ${namespace}",
					slog.String("resource", object.Mapping.Resource.Resource),
					slog.String("name", object.Name),
					slog.String("namespace", object.Namespace))

				continue
			}

			return fmt.Errorf(
				"failed to delete object %s/%s, namespace: %s, %w",
				object.Mapping.Resource.Resource,
				object.Name,
				object.Namespace, err)
		}

		slog.Info("Successfully deleted object ${resource}, namespace: ${namespace}",
			slog.String("resource", getDiffColored(fmt.Sprintf("%s/%s", object.Mapping.Resource.Resource, object.Name), DiffTypeDeleted)),
			slog.String("namespace", getDiffColored(object.Namespace, DiffTypeDeleted)))
	}

	// Finally, delete the apply set parent.
	_, err = applySetHelper.Delete(applySetParentRef.Namespace, applySetParentRef.Name)
	if err != nil {
		return fmt.Errorf(
			"deleted all objects, but failed to delete apply set parent %s/%s: %w",
			applySetParentRef.Namespace,
			applySetParentRef.Name, err)
	}

	return nil
}
