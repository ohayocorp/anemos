package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/sergi/go-diff/diffmatchpatch"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
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
	documents []*core.Document,
	applySetParentName string,
	applySetParentNamespace string,
	skipConfirmation bool,
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
			CascadingStrategy: metav1.DeletePropagationBackground,
			GracePeriod:       -1, // Use default grace period.
			Quiet:             false,
			Output:            "name",
			Timeout:           0, // No timeout.
			WaitForDeletion:   true,
			IgnoreNotFound:    true,
			IOStreams:         genericclioptions.IOStreams{Out: bytes.NewBuffer(nil)},
		},
		ApplySet:            applySet,
		FieldManager:        applySetParentRef.Name,
		ServerSideApply:     true,
		ForceConflicts:      false,
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
		serializedDocument := core.SerializeToYaml(document)
		fmt.Fprintf(buffer, "---\n%s\n", serializedDocument)
	}

	// Configure builder and get objects.
	builder := client.Builder.
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
			if !errors.IsNotFound(err) {
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
			Force:        core.Pointer(false),
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
			return err
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
			live, merged = m.From(), m.To()
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
		diff := getDiffText(liveYamlString, mergedYamlString)
		if diff == "" {
			fmt.Printf("\nNo changes for %s\n", getDiffColored(fmt.Sprintf("%s/%s", info.Mapping.Resource.Resource, info.Name), DiffTypeModified))
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
		diff := getDiffText(string(objectYaml), "")
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

	return nil
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
				fmt.Printf("\nChanges to be applied:\n\n")
				printedLabel = true
			}

			fmt.Printf("%s:\n  %s\n", getDiffColored(diff.Resource, diff.DiffType), core.Indent(diff.DiffText, 2))
		}
	}

	fmt.Printf("Summary of changes:\n\n")

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

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

func getDiffText(left, right string) string {
	differ := diffmatchpatch.New()

	// Compute a line by line diff.
	// https://github.com/sergi/go-diff/issues/69#issuecomment-688602689
	leftChars, rightChars, lines := differ.DiffLinesToChars(left, right)
	diffs := differ.DiffMain(leftChars, rightChars, false)
	diffs = differ.DiffCharsToLines(diffs, lines)
	diffs = differ.DiffCleanupSemantic(diffs)

	if len(diffs) == 0 || (len(diffs) == 1 && diffs[0].Type == diffmatchpatch.DiffEqual) {
		return ""
	}

	return diffPrettyText(diffs)
}

func diffPrettyText(diffs []diffmatchpatch.Diff) string {
	var buff bytes.Buffer
	for _, diff := range diffs {
		text := diff.Text

		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			lines := strings.Split(text, "\n")
			for i, line := range lines {
				_, _ = buff.WriteString("\x1b[32m")
				_, _ = buff.WriteString(line)
				if i < len(lines)-1 {
					_, _ = buff.WriteString("\x1b[0m\n")
				} else {
					_, _ = buff.WriteString("\x1b[0m")
				}
			}

		case diffmatchpatch.DiffDelete:
			lines := strings.Split(text, "\n")
			for i, line := range lines {
				_, _ = buff.WriteString("\x1b[31m")
				_, _ = buff.WriteString(line)
				if i < len(lines)-1 {
					_, _ = buff.WriteString("\x1b[0m\n")
				} else {
					_, _ = buff.WriteString("\x1b[0m")
				}
			}
		case diffmatchpatch.DiffEqual:
			shortenEqualDiffText(&diff, 0, len(diffs))
			_, _ = buff.WriteString(diff.Text)
		}
	}

	return buff.String()
}

func shortenEqualDiffText(diff *diffmatchpatch.Diff, index int, length int) {
	// Limit the length of equal text to 3 lines to avoid flooding the output.
	contextLines := 3

	lines := strings.Split(diff.Text, "\n")
	if len(lines) <= contextLines {
		return
	}

	switch index {
	case 0:
		if len(lines) > contextLines {
			// Take the last n lines of the equal text if there is no change before this one.
			diff.Text = "...\n" + strings.Join(lines[len(lines)-contextLines-1:], "\n")
		}
	case length - 1:
		if len(lines) > contextLines {
			// Take the first n lines of the equal text if there is no change after this one.
			diff.Text = strings.Join(lines[:contextLines+1], "\n") + "\n...\n"
		}
	default:
		if len(lines) > contextLines*2 {
			// Take the first n and last n lines of the equal text if there are changes in between.
			diff.Text = strings.Join(lines[:contextLines+1], "\n") + "\n...\n" + strings.Join(lines[len(lines)-contextLines-1:], "\n")
		}
	}
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
