package client

import (
	"bytes"
	"context"
	"encoding/json"
	goerrors "errors"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/util"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/kubectl/pkg/cmd/apply"
	cmddelete "k8s.io/kubectl/pkg/cmd/delete"
	"k8s.io/kubectl/pkg/cmd/diff"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/scheme"
	"sigs.k8s.io/yaml"
)

var (
	DiffTypeAdded    DiffType = "A"
	DiffTypeDeleted  DiffType = "D"
	DiffTypeModified DiffType = "M"
)

type DiffType string

type Diff struct {
	Resource  string
	Namespace string
	DiffText  string
	DiffType  DiffType
}

type NoChangesError struct{}

var _ error = NoChangesError{}

func (e NoChangesError) Error() string {
	return "no changes were made"
}

func (client *KubernetesClient) Apply(
	documents []string,
	applySetParentName string,
	applySetParentNamespace string,
	skipConfirmation bool,
	forceConflicts bool,
	timeout time.Duration,
) error {
	applySetParentRef, err := client.getApplySetParentRef(applySetParentName, applySetParentNamespace)
	if err != nil {
		return err
	}

	validationDirective := metav1.FieldValidationIgnore
	schema, err := client.Factory.Validator(validationDirective)
	if err != nil {
		return fmt.Errorf("failed to get schema validator: %w", err)
	}

	tooling := getTooling()
	restClient, err := client.getApplySetRestClient(applySetParentRef)
	if err != nil {
		return err
	}

	applySet := apply.NewApplySet(applySetParentRef, tooling, client.Mapper, restClient)

	// Use Kubectl's ApplySet feature to handle the apply operation.
	// This allows us to use server-side apply and manage the apply set lifecycle.
	applyOptions := &apply.ApplyOptions{
		DeleteOptions: &cmddelete.DeleteOptions{
			DynamicClient:     client.DynamicClient,
			CascadingStrategy: metav1.DeletePropagationForeground,
			GracePeriod:       -1, // Use default grace period.
			Quiet:             false,
			Output:            "name",
			Timeout:           timeout,
			WaitForDeletion:   true,
			IgnoreNotFound:    true,
			IOStreams:         genericclioptions.IOStreams{Out: bytes.NewBuffer(nil)},
		},
		ApplySet:            applySet,
		FieldManager:        applySetParentRef.Name,
		ServerSideApply:     true,
		ForceConflicts:      forceConflicts,
		Prune:               true,
		DryRunStrategy:      cmdutil.DryRunNone,
		ValidationDirective: validationDirective,
		PrintFlags:          genericclioptions.NewPrintFlags("created").WithTypeSetter(scheme.Scheme),
		IOStreams:           genericclioptions.IOStreams{Out: bytes.NewBuffer(nil), ErrOut: bytes.NewBuffer(nil)},

		Mapper:        client.Mapper,
		DynamicClient: client.DynamicClient,
		Validator:     schema,

		VisitedUids:       sets.New[types.UID](),
		VisitedNamespaces: sets.New[string](),

		Recorder: &genericclioptions.NoopRecorder{},
		ToPrinter: func(operation string) (printers.ResourcePrinter, error) {
			return printers.NewDiscardingPrinter(), nil
		},
	}

	// Create a buffer to store serialized documents.
	buffer := bytes.NewBuffer(nil)
	for _, document := range documents {
		fmt.Fprintf(buffer, "---\n%s\n", document)
	}

	// Configure builder and get objects.
	builder := client.Factory.NewBuilder().
		ContinueOnError().
		Flatten().
		NamespaceParam(applySetParentRef.Namespace).
		DefaultNamespace().
		Unstructured().
		Schema(schema).
		Stream(bytes.NewBufferString(buffer.String()), "")

	applyOptions.Builder = builder

	// Get the resources to apply.
	infos, err := builder.Do().Infos()
	if err != nil {
		return fmt.Errorf("failed to build resource infos: %w", err)
	}

	// Add managed-by label to all resources.
	extraLabels := map[string]string{
		ManagedByLabel: tooling.Name,
	}

	for _, info := range infos {
		addExtraLabels(info, extraLabels)
	}

	applyOptions.SetObjects(infos)

	// Add apply set labels to the resources.
	err = applySet.AddLabels(infos...)
	if err != nil {
		return err
	}

	// Add custom pre-processor to compute diffs and confirm changes.
	applyOptions.PreProcessorFn = func() error {
		return client.preprocess(infos, applyOptions, skipConfirmation)
	}

	// Kubectl's prune post-processor will handle the deletion of objects not present in the apply set.
	applyOptions.PostProcessorFn = applyOptions.PrintAndPrunePostProcessor()

	// Apply sets are currently behind a feature gate, so we need to set the environment variable to enable them.
	// This is a temporary workaround until the feature is stable.
	err = os.Setenv("KUBECTL_APPLYSET", "true")
	if err != nil {
		return fmt.Errorf("failed to set KUBECTL_APPLYSET environment variable: %w", err)
	}

	defer os.Unsetenv("KUBECTL_APPLYSET")

	// Run the apply operation.
	if err := applyOptions.Run(); err != nil {
		if _, ok := err.(NoChangesError); ok {
			return err
		}

		return err
	}

	// Add managed-by label to the apply set parent resource. This label will be used when listing the apply sets.
	return updateApplySetParentLabels(restClient, applySetParentRef, extraLabels, applyOptions.FieldManager)
}

func (client *KubernetesClient) preprocess(
	infos []*resource.Info,
	applyOptions *apply.ApplyOptions,
	skipConfirmation bool,
) error {
	jobsToRecreate := newApplyJobRecreateSet()
	visitedUids := sets.New[types.UID]()
	diffs := []Diff{}

	for _, info := range infos {
		helper := resource.NewHelper(info.Client, info.Mapping).
			DryRun(true).
			WithFieldManager(applyOptions.FieldManager)

		local := info.Object.DeepCopyObject()

		// Don't use info.Get, as it will override the object with the live state and when we apply the patch,
		// it will not include the local changes.
		// Instead, we use the helper to get the live object.
		live, err := helper.Get(info.Namespace, info.Name)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				return err
			}

			// Object does not exist, treat it as a new object.
			live = nil
		}

		data, err := runtime.Encode(unstructured.UnstructuredJSONScheme, local)
		if err != nil {
			return err
		}

		options := metav1.PatchOptions{
			Force:        core.Pointer(true),
			FieldManager: applyOptions.FieldManager,
		}

		// Get the merged object by applying the patch on server-side with dry-run.
		merged, err := helper.Patch(
			info.Namespace,
			info.Name,
			types.ApplyPatchType,
			data,
			&options,
		)
		if err != nil {
			if isJobResourceInfo(info) && isImmutableFieldError(err) {
				if isJobRecreateOnImmutableEnabled(info) {
					jobsToRecreate.Add(info)
					// If we're going to recreate the Job, treat the local object as the merged object
					// for diff/confirmation purposes.
					merged = local
				} else {
					// Only Jobs can be recreated, and only when explicitly enabled.
					return err
				}
			} else {
				return err
			}
		}

		uid := getUID(live)
		if uid != nil {
			visitedUids.Insert(*uid)
		}

		// No need to compare managed fields, as they are not part of the manifests.
		omitManagedFields(live)
		omitManagedFields(merged)

		// Generation field is incremented for each update, it is redundant in the diff.
		// Set the generation of the live object to the merged object's generation.
		setGeneration(live, merged)

		// Mask secret values if object is V1Secret.
		if gvk := merged.GetObjectKind().GroupVersionKind(); gvk.Version == "v1" && gvk.Kind == "Secret" {
			m, err := diff.NewMasker(live, merged)
			if err != nil {
				return err
			}

			liveIsNil := live == nil
			live, merged = m.From(), m.To()

			if liveIsNil {
				// Masker returns an object that has nil data, but we need live object to be actually nil.
				live = nil
			}
		}

		liveYamlString := ""
		mergedYamlString := ""

		// If live is nil, it means the object does not exist in the cluster.
		// We currently don't show the diff for new objects, as they are created with the local YAML, but
		// we may want to change this in the future.
		if live != nil {
			liveYaml, err := yaml.Marshal(live)
			if err != nil {
				return err
			}

			mergedYaml, err := yaml.Marshal(merged)
			if err != nil {
				return err
			}

			liveYamlString = string(liveYaml)
			mergedYamlString = string(mergedYaml)
		} else {
			// For new objects, we use the local YAML representation as the merged YAML.
			// Using the merged YAML from the server would include unnecessary metadata
			// such as UID and creation timestamp.
			localYaml, err := yaml.Marshal(local)
			if err != nil {
				return err
			}

			mergedYamlString = string(localYaml)
		}

		// Get the diff text between the live and merged objects.
		diff, err := getDiffText(liveYamlString, mergedYamlString)
		if err != nil {
			return fmt.Errorf("failed to compute diff for %s/%s: %w", info.Namespace, info.Name, err)
		}

		if diff == "" {
			slog.Info(fmt.Sprintf(
				"No changes for %s",
				getDiffColored(fmt.Sprintf("%s/%s", info.Mapping.Resource.Resource, info.Name), DiffTypeAdded)))

			continue
		}

		diffType := DiffTypeModified
		if live == nil {
			diffType = DiffTypeAdded
		}

		diffs = append(diffs, Diff{
			Resource:  fmt.Sprintf("%s/%s", info.Mapping.Resource.Resource, info.Name),
			Namespace: info.Namespace,
			DiffText:  diff,
			DiffType:  diffType,
		})
	}

	// After collecting all diffs, we need to confirm the changes with the user.
	// We call BeforeApply to fetch the resources associated with the apply set
	// and find the objects to prune.
	applyOptions.ApplySet.BeforeApply(nil, cmdutil.DryRunClient, applyOptions.ValidationDirective)
	objectsToPrune, err := applyOptions.ApplySet.FindAllObjectsToPrune(context.TODO(), client.DynamicClient, visitedUids)
	if err != nil {
		return fmt.Errorf("failed to find objects to prune: %w", err)
	}

	for _, object := range objectsToPrune {
		omitManagedFields(object.Object)

		objectYaml, err := yaml.Marshal(object.Object)
		if err != nil {
			return err
		}

		// We don't show the diff for deleted objects, same as for new objects.
		diff, err := getDiffText(string(objectYaml), "")
		if err != nil {
			return fmt.Errorf("failed to compute diff for %s/%s: %w", object.Namespace, object.Name, err)
		}

		diffs = append(diffs, Diff{
			Resource:  fmt.Sprintf("%s/%s", object.Mapping.Resource.Resource, object.Name),
			Namespace: object.Namespace,
			DiffText:  diff,
			DiffType:  DiffTypeDeleted,
		})
	}

	if len(diffs) == 0 {
		// Return a specific type of error to indicate no changes were made.
		return NoChangesError{}
	}

	printChanges(diffs)

	// Lastly, we need to confirm the changes with the user.
	if !skipConfirmation {
		confirmed, err := confirmChanges()
		if err != nil {
			return err
		}
		if !confirmed {
			return fmt.Errorf("aborting apply operation due to user confirmation")
		}
	}

	if err := client.recreateJobsIfNeeded(context.TODO(), jobsToRecreate, applyOptions.DeleteOptions.Timeout); err != nil {
		return err
	}

	return nil
}

const (
	JobRecreateOnImmutableFieldsAnnotation = "anemos.sh/recreate-on-immutable-fields-change"
)

type applyJobRecreateKey struct {
	gvr       schema.GroupVersionResource
	name      string
	namespace string
}

type applyJobRecreateSet struct {
	items map[applyJobRecreateKey]struct{}
}

func newApplyJobRecreateSet() *applyJobRecreateSet {
	return &applyJobRecreateSet{items: map[applyJobRecreateKey]struct{}{}}
}

func (s *applyJobRecreateSet) Add(info *resource.Info) {
	if info == nil || info.Mapping == nil {
		return
	}

	key := applyJobRecreateKey{
		gvr:       info.Mapping.Resource,
		name:      info.Name,
		namespace: info.Namespace,
	}
	if key.name == "" {
		return
	}

	s.items[key] = struct{}{}
}

func (s *applyJobRecreateSet) Len() int {
	return len(s.items)
}

func isJobResourceInfo(info *resource.Info) bool {
	if info == nil || info.Mapping == nil {
		return false
	}

	gvk := info.Mapping.GroupVersionKind
	if gvk.Kind != "Job" {
		return false
	}

	// Jobs are in the "batch" group (batch/v1).
	return gvk.Group == "batch"
}

func isJobRecreateOnImmutableEnabled(info *resource.Info) bool {
	u, ok := info.Object.(*unstructured.Unstructured)
	if !ok || u == nil {
		return false
	}

	annotations := u.GetAnnotations()
	if annotations == nil {
		return false
	}

	value := strings.TrimSpace(strings.ToLower(annotations[JobRecreateOnImmutableFieldsAnnotation]))
	return value == "true"
}

func isImmutableFieldError(err error) bool {
	if err == nil {
		return false
	}

	var statusErr *apierrors.StatusError
	if goerrors.As(err, &statusErr) {
		msg := strings.ToLower(statusErr.ErrStatus.Message)
		return strings.Contains(msg, "field is immutable")
	}

	return strings.Contains(strings.ToLower(err.Error()), "field is immutable")
}

func (client *KubernetesClient) recreateJobsIfNeeded(ctx context.Context, jobsToRecreate *applyJobRecreateSet, timeout time.Duration) error {
	if jobsToRecreate == nil || jobsToRecreate.Len() == 0 {
		return nil
	}

	for key := range jobsToRecreate.items {
		slog.Warn(
			`Recreating Job due to immutable field change: "${namespace}/${name}" since annotation "${annotation}" is set to true`,
			slog.String("namespace", key.namespace),
			slog.String("name", key.name),
			slog.String("annotation", JobRecreateOnImmutableFieldsAnnotation),
		)

		if err := client.deleteAndWaitForResourceGone(ctx, key.gvr, key.namespace, key.name, timeout); err != nil {
			return err
		}
	}

	return nil
}

func (client *KubernetesClient) deleteAndWaitForResourceGone(
	ctx context.Context,
	gvr schema.GroupVersionResource,
	namespace string,
	name string,
	timeout time.Duration,
) error {
	propagation := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{PropagationPolicy: &propagation}

	err := client.DynamicClient.Resource(gvr).Namespace(namespace).Delete(ctx, name, deleteOptions)
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}

	return wait.PollUntilContextTimeout(ctx, 1*time.Second, timeout, true, func(ctx context.Context) (bool, error) {
		_, getErr := client.DynamicClient.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
		if apierrors.IsNotFound(getErr) {
			return true, nil
		}
		return false, getErr
	})
}

func printChanges(diffs []Diff) {
	sort.Slice(diffs, func(i, j int) bool {
		// Sort by diff type first. The order is: Modified, Deleted, Added.
		if diffs[i].DiffType != diffs[j].DiffType {
			return diffs[i].DiffType > diffs[j].DiffType
		}

		// Then sort by namespace.
		if diffs[i].Namespace != diffs[j].Namespace {
			return diffs[i].Namespace < diffs[j].Namespace
		}

		// Finally, sort by resource name which is a combination of resource type and name.
		return diffs[i].Resource < diffs[j].Resource
	})

	// Print the changes.
	printedLabel := false
	for _, diff := range diffs {
		// Print the diff for modified resources.
		if diff.DiffType == DiffTypeModified {
			if !printedLabel {
				slog.Info("Changes to be applied:")
				printedLabel = true
			}

			slog.Info(fmt.Sprintf(
				"%s:\n  %s",
				getDiffColored(diff.Resource, diff.DiffType), util.Indent(diff.DiffText, 2)))
		}
	}

	slog.Info("Summary of changes:")

	builder := &strings.Builder{}
	w := tabwriter.NewWriter(builder, 1, 1, 2, ' ', 0)

	// Print the header with color codes. Labels won't have actual colors,
	// but when color codes are not added, tabwriter will not align the columns correctly.
	fmt.Fprintf(
		w,
		"%s\t%s\t%s\n",
		getDiffColored("OP", ""),
		getDiffColored("NAMESPACE", ""),
		getDiffColored("RESOURCE", ""),
	)

	for _, diff := range diffs {
		fmt.Fprintf(
			w,
			"%s\t%s\t%s\n",
			getDiffColored(string(diff.DiffType), diff.DiffType),
			getDiffColored(diff.Namespace, diff.DiffType),
			getDiffColored(diff.Resource, diff.DiffType))
	}

	fmt.Fprintf(w, "\n")

	w.Flush()

	for _, line := range strings.Split(builder.String(), "\n") {
		if line == "" {
			continue
		}

		slog.Info(line)
	}
}

func getDiffColored(text string, diffType DiffType) string {
	switch diffType {
	case DiffTypeAdded:
		return fmt.Sprintf("\x1b[32m%s\x1b[0m", text) // Green for added.
	case DiffTypeModified:
		return fmt.Sprintf("\x1b[33m%s\x1b[0m", text) // Yellow for modified.
	case DiffTypeDeleted:
		return fmt.Sprintf("\x1b[31m%s\x1b[0m", text) // Red for deleted.
	default:
		// Double 00 is used to fix tabwriter alignment issues.
		return fmt.Sprintf("\x1b[00m%s\x1b[0m", text) // Default color for unknown types.
	}
}

func getDiffText(left, right string) (string, error) {
	edits := myers.ComputeEdits("", left, right)
	edits = gotextdiff.LineEdits(left, edits)
	unified := gotextdiff.ToUnified("original", "modified", left, edits)

	if len(unified.Hunks) == 0 {
		return "", nil
	}

	var buff bytes.Buffer

	write := func(format string, args ...interface{}) {
		_, _ = buff.WriteString(fmt.Sprintf(format, args...))
	}

	for _, hunk := range unified.Hunks {
		fromCount, toCount := 0, 0
		for _, l := range hunk.Lines {
			switch l.Kind {
			case gotextdiff.Delete:
				fromCount++
			case gotextdiff.Insert:
				toCount++
			default:
				fromCount++
				toCount++
			}
		}

		red := "\x1b[31m"
		green := "\x1b[32m"
		reset := "\x1b[0m"

		write("%s%s", reset, "@@")

		if fromCount > 1 {
			write("%s -%d,%d", red, hunk.FromLine, fromCount)
		} else {
			write("%s -%d", red, hunk.FromLine)
		}

		if toCount > 1 {
			write("%s +%d,%d", green, hunk.ToLine, toCount)
		} else {
			write("%s +%d", green, hunk.ToLine)
		}

		write("%s @@\n", reset)

		for _, l := range hunk.Lines {
			switch l.Kind {
			case gotextdiff.Delete:
				write("%s-%s", red, l.Content)
			case gotextdiff.Insert:
				write("%s+%s", green, l.Content)
			default:
				write("%s %s", reset, l.Content)
			}
		}
	}

	return buff.String(), nil
}

func updateApplySetParentLabels(
	restClient resource.RESTClient,
	applySetParentRef *apply.ApplySetParentRef,
	extraLabels map[string]string,
	fieldManager string,
) error {
	patch := &metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			Kind:       applySetParentRef.GroupVersionKind.Kind,
			APIVersion: applySetParentRef.GroupVersionKind.GroupVersion().String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      applySetParentRef.Name,
			Namespace: applySetParentRef.Namespace,
			Labels:    extraLabels,
		},
	}

	data, err := json.Marshal(patch)
	if err != nil {
		return fmt.Errorf("failed to encode patch for ApplySet parent: %w", err)
	}

	helper := resource.NewHelper(restClient, applySetParentRef.RESTMapping)

	options := &metav1.PatchOptions{
		FieldManager: fieldManager,
	}

	_, err = helper.Patch(
		applySetParentRef.Namespace,
		applySetParentRef.Name,
		types.ApplyPatchType,
		data,
		options,
	)

	return err
}
