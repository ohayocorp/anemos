package client

import (
	"context"
	"fmt"
	"os"
	"sort"
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
		fmt.Println("No objects to delete. Will delete the apply set parent only.")
	} else {
		fmt.Printf("Deleting objects for apply set %s:\n\n", cleanupApplySetParentName(applySetParentRef.Name))

		writer := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
		fmt.Fprintf(writer, "\x1b[00m%s\t%s\t%s\n", "NAMESPACE", "RESOURCE", "NAME")

		for _, object := range objectsToDelete {
			fmt.Fprintf(writer, "\x1b[31m%s\t%s\t%s\x1b[0m\n", object.Namespace, object.Mapping.Resource.Resource, object.Name)
		}

		writer.Flush()
	}

	fmt.Println()

	confirmed, err := confirmChanges()
	if err != nil {
		return fmt.Errorf("failed to confirm changes: %w", err)
	}

	if !confirmed {
		return fmt.Errorf("aborting delete operation due to user confirmation")
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
				fmt.Printf(
					"Object not found, skipping deletion: %s/%s, namespace: %s\n",
					object.Mapping.Resource.Resource,
					object.Name,
					object.Namespace)

				continue
			}

			return fmt.Errorf(
				"failed to delete object %s/%s, namespace: %s, %w",
				object.Mapping.Resource.Resource,
				object.Name,
				object.Namespace, err)
		}

		fmt.Printf(
			"Successfully deleted object %s, namespace: %s\n",
			getDiffColored(fmt.Sprintf("%s/%s", object.Mapping.Resource.Resource, object.Name), DiffTypeDeleted),
			getDiffColored(object.Namespace, DiffTypeDeleted))
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
