package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ohayocorp/anemos/pkg/components"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/k8s"
	"github.com/spf13/cobra"
)

func getBuildCommand(program *AnemosProgram) *cobra.Command {
	var tscDirs []string
	var apply bool
	var skipConfirmation bool

	command := &cobra.Command{
		Use:   "build [js_file]",
		Short: "Builds a project.",
		RunE: func(ctx *cobra.Command, args []string) error {
			return build(program, args, tscDirs, apply, skipConfirmation)
		},
		Args: cobra.MinimumNArgs(1),
	}

	command.Flags().StringSliceVar(&tscDirs, "tsc", nil, "Directories to compile with tsc.")
	command.Flags().BoolVar(&apply, "apply", false, "Apply the generated manifests to the cluster.")
	command.Flags().BoolVarP(&skipConfirmation, "yes", "y", false, "Skip confirmation prompt and apply changes directly")

	return command
}

func build(program *AnemosProgram, args []string, tscDirs []string, apply bool, skipConfirmation bool) error {
	var jsFile string
	if len(args) > 0 {
		jsFile = args[0]
		args = args[1:]
	} else {
		return fmt.Errorf("no JS file provided")
	}

	jsFile, err := js.ResolvePath(jsFile, true)
	if err != nil {
		return err
	}

	err = writeTypeDeclarations(program, filepath.Dir(jsFile))
	if err != nil {
		return fmt.Errorf("failed to write type declarations: %w", err)
	}

	for _, tscDir := range tscDirs {
		tscDir, err := filepath.Abs(tscDir)
		if err != nil {
			return fmt.Errorf("failed to get absolute path: %s, %w", tscDir, err)
		}

		err = js.RunTsc(tscDir)
		if err != nil {
			return err
		}
	}

	runtime, err := InitializeNewRuntime(program)
	if err != nil {
		return err
	}

	if apply {
		runtime.Flags[core.JsRuntimeMetadataBuilderApply] = "true"
	}

	if skipConfirmation {
		runtime.Flags[core.JsRuntimeMetadataBuilderSkipConfirmation] = "true"
	}

	scriptContents, err := os.ReadFile(jsFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %s, %w", jsFile, err)
	}

	script := &js.JsScript{
		Contents: string(scriptContents),
		FilePath: jsFile,
	}

	return runtime.Run(script, args)
}

func InitializeNewRuntime(program *AnemosProgram) (*js.JsRuntime, error) {
	runtime := js.NewJsRuntime()

	k8s.RegisterK8S(runtime)
	core.RegisterCore(runtime)
	components.RegisterComponents(runtime)

	if program.RegisterRuntimeCallback != nil {
		if err := program.RegisterRuntimeCallback(runtime); err != nil {
			return nil, fmt.Errorf("failed to call register runtime callback: %w", err)
		}
	}

	err := runtime.InitializeNativeLibraries()
	if err != nil {
		return nil, err
	}

	if program.InitializeRuntimeCallback != nil {
		if err := program.InitializeRuntimeCallback(runtime); err != nil {
			return nil, fmt.Errorf("failed to initialize runtime: %w", err)
		}
	}

	return runtime, nil
}

func writeTypeDeclarations(program *AnemosProgram, directory string) error {
	// Find the package.json file in the directory of the jsFile and its parent directories
	dirAbs, err := filepath.Abs(directory)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %s, %w", directory, err)
	}

	var packageJsonPath string
	dir := dirAbs

	for {
		packageJsonPath = filepath.Join(dir, "package.json")
		if _, err := os.Stat(packageJsonPath); err == nil {
			break
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Reached the root directory without finding package.json
			return nil
		}

		dir = parentDir
	}

	packageJson, err := os.ReadFile(packageJsonPath)
	if err != nil {
		return fmt.Errorf("failed to read package.json: %s, %w", packageJsonPath, err)
	}

	if !containsAnemosDependency(packageJson) {
		return nil
	}

	// Delete the existing type declarations directory if it exists.
	typeDeclarationsDir := filepath.Join(dir, ".anemos", "types")

	entries, err := os.ReadDir(typeDeclarationsDir)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read type declarations directory: %s, %w", typeDeclarationsDir, err)
	}

	for _, entry := range entries {
		path := filepath.Join(typeDeclarationsDir, entry.Name())

		if err := os.RemoveAll(path); err != nil {
			return fmt.Errorf("failed to remove existing type declarations: %s, %w", path, err)
		}
	}

	return writeDeclarations(program, filepath.Join(dir, ".anemos", "types"))
}

func containsAnemosDependency(packageJson []byte) bool {
	// Parse the package.json and check for anemos package in dependencies or devDependencies
	parsedJson := make(map[string]interface{})
	if err := json.Unmarshal(packageJson, &parsedJson); err != nil {
		return false
	}

	if dependencies, ok := parsedJson["dependencies"].(map[string]interface{}); ok {
		if _, ok := dependencies[js.PackageName]; ok {
			return true
		}
	}

	if devDependencies, ok := parsedJson["devDependencies"].(map[string]interface{}); ok {
		if _, ok := devDependencies[js.PackageName]; ok {
			return true
		}
	}

	return false
}
