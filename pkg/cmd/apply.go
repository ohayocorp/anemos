package cmd

import (
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/ohayocorp/anemos/pkg/client"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/util"
	"github.com/ohayocorp/sobek_nodejs/require"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

//go:embed templates/applyJsFile.js
var applyJsFileScript string

var (
	SourceTypePackage SourceType = "package"
	SourceTypeURL     SourceType = "url"
)

type SourceType string

type applyContext struct {
	program          *AnemosProgram
	source           string
	namespace        string
	skipConfirmation bool
	forceConflicts   bool
	distribution     core.KubernetesDistribution
	environmentType  core.EnvironmentType
	documentGroups   []string
	options          map[string]any
}

func getApplyCommand(program *AnemosProgram) *cobra.Command {
	command := &cobra.Command{
		Use:   "apply [package|directory]",
		Short: "Apply manifests to the Kubernetes cluster",
		Long: util.Dedent(`
			Apply manifests to the Kubernetes cluster. Packages are downloaded using Bun, so you
			can use any package format that Bun supports.
			  - Apply a package from NPM: anemos apply my-package
			  - Apply a package from a tgz file: anemos apply my-package.tgz
			  - Apply a local JS file: anemos apply ./my-package.js
			  - Apply a JS file from a URL: anemos apply https://example.com/my-package.js
			`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runApplyCommand(cmd, args, program)
		},
		Args: cobra.ExactArgs(1),
	}

	command.Flags().StringP("namespace", "", "", "Namespace to apply the manifests to")
	command.Flags().BoolP("yes", "y", false, "Skip confirmation prompt and apply changes directly")
	command.Flags().Bool("force-conflicts", false, "Forcefully apply changes even if there are conflicts")
	command.Flags().StringArrayP("document-groups", "d", nil, "Document groups to apply, other groups will be skipped")
	command.Flags().StringP("options-file", "f", "", "Path to YAML file containing options to pass to the package")
	command.Flags().String("distribution", "", "Distribution of the target Kubernetes cluster, e.g., minikube, openshift, etc. If not set, it will be determined based on the cluster version.")
	command.Flags().String("environment-type", "", "Environment type such as dev, test or prod. If not set, it will be determined based on the cluster distribution.")

	return command
}

func runApplyCommand(cmd *cobra.Command, args []string, program *AnemosProgram) error {
	skipConfirmation := cmdutil.GetFlagBool(cmd, "yes")
	forceConflicts := cmdutil.GetFlagBool(cmd, "force-conflicts")
	namespace := cmdutil.GetFlagString(cmd, "namespace")
	optionsFile := cmdutil.GetFlagString(cmd, "options-file")
	documentGroups := cmdutil.GetFlagStringArray(cmd, "document-groups")
	distribution := cmdutil.GetFlagString(cmd, "distribution")
	environmentType := cmdutil.GetFlagString(cmd, "environment-type")

	// Check if we should skip confirmation from environment variable.
	if cmd.Flags().Lookup("yes") == nil {
		_, skipConfirmation = os.LookupEnv("ANEMOS_APPLY_YES")
	}

	var yamlOptions map[string]any
	var err error

	// Load options from file if provided
	if optionsFile != "" {
		yamlOptions, err = loadOptionsFromFile(optionsFile)
		if err != nil {
			return err
		}
	} else {
		yamlOptions = make(map[string]any)
	}

	source := args[0]

	applyContext := &applyContext{
		program:          program,
		source:           source,
		namespace:        namespace,
		skipConfirmation: skipConfirmation,
		forceConflicts:   forceConflicts,
		distribution:     core.KubernetesDistribution(distribution),
		environmentType:  core.EnvironmentType(environmentType),
		documentGroups:   documentGroups,
		options:          yamlOptions,
	}

	return applyManifests(applyContext)
}

func applyManifests(context *applyContext) error {
	// Determine if source is a package name, or URL.
	sourceType, err := determineSourceType(context.source)
	if err != nil {
		return err
	}

	switch sourceType {
	case SourceTypeURL:
		err = applyURL(context)
	case SourceTypePackage:
		err = applyPackage(context)
	default:
		return fmt.Errorf("unknown source type")
	}

	if err != nil {
		return err
	}

	return nil
}

func determineSourceType(source string) (SourceType, error) {
	// Check if source is a URL.
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return SourceTypeURL, nil
	}

	// If not a URL, assume it's a package that Bun can handle.
	return SourceTypePackage, nil
}

func applyURL(context *applyContext) error {
	parsedUrl, err := url.Parse(context.source)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %s, %w", context.source, err)
	}

	// Determine if this is a direct JavaScript file or a tar.gz package.
	isJavaScriptFile := strings.HasSuffix(strings.ToLower(parsedUrl.Path), ".js")

	if isJavaScriptFile {
		// Handle direct JavaScript file.
		contents, err := downloadContents(parsedUrl)
		if err != nil {
			return err
		}

		return applyJavaScriptFile(context, contents)
	}

	// For other types of URLs, we assume it's a package URL that Bun can handle.
	return applyPackage(context)
}

func applyJavaScriptFile(context *applyContext, script string) error {
	jsRuntime, err := InitializeNewRuntime(context.program)
	if err != nil {
		return err
	}

	require.WithLoader(func(path string) ([]byte, error) {
		if path == filepath.Join("node_modules", "anemos-apply-package", "index.js") {
			// If the path is "node_modules/anemos-apply-package", we assume it's the main package file.
			return []byte(script), nil
		}

		return js.SourceLoader(jsRuntime, path)
	})(jsRuntime.Registry)

	setVariables(jsRuntime, context)

	return jsRuntime.Run(&js.JsScript{
		Contents: applyJsFileScript,
		FilePath: "index.js",
	}, nil)
}

func applyPackage(context *applyContext) error {
	// Check that if the source is a local JavaScript file.
	if strings.HasSuffix(context.source, ".js") {
		jsFile, err := filepath.Abs(context.source)
		if err != nil {
			return fmt.Errorf("failed to get absolute path for JavaScript file: %s, %w", context.source, err)
		}

		contents, err := os.ReadFile(jsFile)
		if err == nil {
			return applyJavaScriptFile(context, string(contents))
		}
	}

	return applyPackageIdentifier(context)
}

// applyPackageIdentifier applies a package by creating a temporary directory
// and running Bun to install the package.
func applyPackageIdentifier(context *applyContext) error {
	// Create a temporary directory for our work.
	tempDir, err := os.MkdirTemp("", "anemos-package-*")
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	stdout := strings.Builder{}
	stderr := strings.Builder{}

	// Add the package as given by the user.
	err = js.RunBunCommand(js.BunCommand{
		Description: "Adding dependencies",
		Args:        []string{"add", context.source},
		Cwd:         &tempDir,
		Stdout:      &stdout,
		Stderr:      &stderr,
	})

	if err != nil {
		return fmt.Errorf(
			"failed to install package: %s, error: %w\nbun stdout: %s\nbun stderr: %s",
			context.source,
			err,
			stdout.String(),
			stderr.String())
	}

	stdout = strings.Builder{}
	stderr = strings.Builder{}

	// Add the package as an alias to "anemos-apply-package".
	// This allows us to use the package in the JS script without needing to specify the full path.
	err = js.RunBunCommand(js.BunCommand{
		Description: "Adding dependencies",
		Args:        []string{"add", fmt.Sprintf("anemos-apply-package@npm:%s", context.source)},
		Cwd:         &tempDir,
		Stdout:      &stdout,
		Stderr:      &stderr,
	})

	if err != nil {
		return fmt.Errorf(
			"failed to make an alias to the package: %s, error: %w\nbun stdout: %s\nbun stderr: %s",
			context.source,
			err,
			stdout.String(),
			stderr.String())
	}

	slog.Info("Downloaded package ${package}", slog.String("package", context.source))

	// Create builder script that imports the package. File path is fixed to "index.js"
	// under the temporary directory so that modules can be resolved from node_modules.
	filePath := filepath.Join(tempDir, "index.js")

	jsRuntime, err := InitializeNewRuntime(context.program)
	if err != nil {
		return err
	}

	setVariables(jsRuntime, context)

	return jsRuntime.Run(&js.JsScript{
		Contents: applyJsFileScript,
		FilePath: filePath,
	}, nil)
}

// downloadContents downloads the content from a URL as a string.
func downloadContents(url *url.URL) (string, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return "", fmt.Errorf("failed to download content from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download content from %s: %s", url, resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(content), nil
}

func loadOptionsFromFile(filePath string) (map[string]interface{}, error) {
	if filePath == "" {
		return nil, nil
	}

	var data []byte
	var err error

	if filePath == "-" {
		// Read from stdin
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("failed to read from stdin: %w", err)
		}
	} else {
		data, err = os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read options file: %w", err)
		}
	}

	return core.ParseYaml[map[string]any](string(data))
}

func setVariables(jsRuntime *js.JsRuntime, context *applyContext) {
	clusterInfo, err := getClusterInfo(context)
	if err != nil {
		js.Throw(fmt.Errorf("failed to get cluster info: %w", err))
	}

	jsRuntime.Runtime.Set("options", context.options)
	jsRuntime.Runtime.Set("documentGroups", context.documentGroups)
	jsRuntime.Runtime.Set("namespace", context.namespace)
	jsRuntime.Runtime.Set("skipConfirmation", context.skipConfirmation)
	jsRuntime.Runtime.Set("forceConflicts", context.forceConflicts)
	jsRuntime.Runtime.Set("clusterInfo", clusterInfo)
	jsRuntime.Runtime.Set("environmentType", getEnvironmentType(clusterInfo, context))
}

func getClusterInfo(context *applyContext) (*client.ClusterInfo, error) {
	kubernetesClient, err := client.NewKubernetesClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	clusterInfo, err := kubernetesClient.GetClusterInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster info: %w", err)
	}

	if context.distribution != "" {
		clusterInfo.Distribution = core.KubernetesDistribution(context.distribution)
	}

	return clusterInfo, nil
}

func getEnvironmentType(clusterInfo *client.ClusterInfo, context *applyContext) core.EnvironmentType {
	if context.environmentType != "" {
		return context.environmentType
	}

	if clusterInfo.Distribution == core.KubernetesDistributionMinikube {
		return core.EnvironmentTypeDevelopment
	}

	// Default to production.
	return core.EnvironmentTypeProduction
}
