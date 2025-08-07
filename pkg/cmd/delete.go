package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/ohayocorp/anemos/pkg/client"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

type deleteContext struct {
	program          *AnemosProgram
	applySetName     string
	namespace        string
	skipConfirmation bool
	timeout          time.Duration
}

func getDeleteCommand(program *AnemosProgram) *cobra.Command {
	command := &cobra.Command{
		Use:   "delete [apply-set-name]",
		Short: "Delete manifests from the Kubernetes cluster for an apply set",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDeleteCommand(cmd, args, program)
		},
		Args: cobra.ExactArgs(1),
	}

	command.Flags().StringP("namespace", "", "", "Namespace of the apply set")
	command.Flags().BoolP("yes", "y", false, "Skip confirmation prompt and delete changes directly")
	command.Flags().StringP("timeout", "t", "5m0s", "Timeout for the delete operation")

	return command
}

func runDeleteCommand(cmd *cobra.Command, args []string, program *AnemosProgram) error {
	skipConfirmation := cmdutil.GetFlagBool(cmd, "yes")
	namespace := cmdutil.GetFlagString(cmd, "namespace")
	timeoutString := cmdutil.GetFlagString(cmd, "timeout")
	if timeoutString == "" {
		timeoutString = "5m0s"
	}

	timeout, err := time.ParseDuration(timeoutString)
	if err != nil {
		return fmt.Errorf("invalid timeout value: %w", err)
	}

	// Check if we should skip confirmation from environment variable.
	if cmd.Flags().Lookup("yes") == nil {
		_, skipConfirmation = os.LookupEnv("ANEMOS_APPLY_YES")
	}

	applySetName := args[0]

	deleteContext := &deleteContext{
		program:          program,
		applySetName:     applySetName,
		namespace:        namespace,
		skipConfirmation: skipConfirmation,
		timeout:          timeout,
	}

	return deleteManifests(deleteContext)
}

func deleteManifests(context *deleteContext) error {
	kubernetesClient, err := client.NewKubernetesClient()
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	return kubernetesClient.Delete(context.applySetName, context.namespace, context.skipConfirmation)
}
