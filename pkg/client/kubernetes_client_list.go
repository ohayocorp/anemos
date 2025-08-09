package client

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"text/tabwriter"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client *KubernetesClient) List(applySetParentNamespace string) error {
	namespace, _, _ := client.Factory.ToRawKubeConfigLoader().Namespace()
	if applySetParentNamespace != "" {
		namespace = applySetParentNamespace
	}

	// Apply sets are stored as secrets with a specific label.
	listOfSecrets, err := client.CoreClient.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", ManagedByLabel, getTooling().Name),
	})
	if err != nil {
		return fmt.Errorf("failed to list secrets in namespace %q: %w", namespace, err)
	}

	sort.Slice(listOfSecrets.Items, func(i, j int) bool {
		return listOfSecrets.Items[i].Name < listOfSecrets.Items[j].Name
	})

	if len(listOfSecrets.Items) == 0 {
		slog.Info("No apply sets found.")
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	fmt.Fprintln(writer, "NAME\tCREATED AT")
	fmt.Fprintln(writer, "----\t----------")

	for _, secret := range listOfSecrets.Items {
		// Remove the fqdn prefix from the secret name.
		fmt.Fprintf(writer, "%s\t%s\n", cleanupApplySetParentName(secret.Name), secret.CreationTimestamp.Format("2006-01-02 15:04:05"))
	}

	writer.Flush()

	return nil
}
