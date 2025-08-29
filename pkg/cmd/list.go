package cmd

import (
	_ "embed"
	"fmt"

	"github.com/ohayocorp/anemos/pkg/client"
	"github.com/ohayocorp/anemos/pkg/util"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

type listContext struct {
	program   *AnemosProgram
	namespace string
}

func getListCommand(program *AnemosProgram) *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List apply sets",
		Long: util.Dedent(`
			List apply sets in the Kubernetes cluster.
			`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListCommand(cmd, program)
		},
		Args: cobra.ExactArgs(0),
	}

	command.Flags().StringP("namespace", "", "", "Namespace to apply the manifests to")

	return command
}

func runListCommand(cmd *cobra.Command, program *AnemosProgram) error {
	namespace := cmdutil.GetFlagString(cmd, "namespace")

	listContext := &listContext{
		program:   program,
		namespace: namespace,
	}

	return listManifests(listContext)
}

func listManifests(context *listContext) error {
	kubernetesClient, err := client.NewKubernetesClient()
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	err = kubernetesClient.List(context.namespace)
	if err != nil {
		return err
	}

	return nil
}
