package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ohayocorp/anemos/pkg/components"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/js"
	"github.com/ohayocorp/anemos/pkg/k8s"
	"github.com/ohayocorp/anemos/pkg/util"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func getBuildCommand(program *AnemosProgram) *cobra.Command {
	command := &cobra.Command{
		Use:   "build [js_file|ts_file]",
		Short: "Builds a project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return build(cmd, args, program)
		},
		Args: cobra.MinimumNArgs(1),
	}

	command.Flags().Bool("apply", false, "Apply the generated manifests to the cluster.")
	command.Flags().Bool("yes", false, "Skip confirmation prompt and apply changes directly")
	command.Flags().Bool("force-conflicts", false, "Forcefully apply changes even if there are conflicts")
	command.Flags().StringArrayP("document-groups", "d", nil, "Document groups to apply, other groups will be skipped")

	return command
}

func build(cmd *cobra.Command, args []string, program *AnemosProgram) error {
	apply := cmdutil.GetFlagBool(cmd, "apply")
	skipConfirmation := cmdutil.GetFlagBool(cmd, "yes")
	forceConflicts := cmdutil.GetFlagBool(cmd, "force-conflicts")
	documentGroups := cmdutil.GetFlagStringArray(cmd, "document-groups")

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

	mainScriptPath := jsFile

	if strings.HasSuffix(jsFile, ".ts") {
		tsFile := jsFile
		tsFileDir := filepath.Dir(tsFile)

		err = compileTypeScript(program, tsFile)
		if err != nil {
			return err
		}

		compiledJsFile := fmt.Sprintf("%s.js", strings.TrimSuffix(filepath.Base(tsFile), filepath.Ext(tsFile)))
		jsFile = filepath.Join(tsFileDir, "dist", compiledJsFile)
	}

	scriptContents, err := os.ReadFile(jsFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %s, %w", jsFile, err)
	}

	runtime, err := InitializeNewRuntime(program)
	if err != nil {
		return err
	}

	runtime.BuilderDefaultsContext.Set("apply", apply)
	runtime.BuilderDefaultsContext.Set("skipConfirmation", skipConfirmation)
	runtime.BuilderDefaultsContext.Set("forceConflicts", forceConflicts)
	runtime.BuilderDefaultsContext.Set("documentGroups", documentGroups)

	script := &js.JsScript{
		Contents:       string(scriptContents),
		FilePath:       jsFile,
		MainScriptPath: mainScriptPath,
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
	// Delete the existing type declarations directory if it exists.
	typeDeclarationsDir := filepath.Join(directory, ".anemos", "types")

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

	return writeDeclarations(program, filepath.Join(directory, ".anemos", "types"))
}

func compileTypeScript(program *AnemosProgram, tsFile string) error {
	tempDir, err := os.MkdirTemp("", "anemos-tsc-")
	if err != nil {
		return fmt.Errorf("failed to create temporary directory for tsc: %w", err)
	}
	defer os.RemoveAll(tempDir)

	err = writeTypeDeclarations(program, tempDir)
	if err != nil {
		return fmt.Errorf("failed to write type declarations: %w", err)
	}

	tsFileSlash := filepath.ToSlash(tsFile)
	tsFileDir := filepath.ToSlash(filepath.Dir(tsFileSlash))

	tsconfig := util.ParseTemplate(`
		{
		  "compilerOptions": {
		    "target": "ES2019",
		    "lib": [
		        "ES2019"
		    ],
		    "moduleResolution": "nodenext",
		    "module": "NodeNext",
		    "strict": true,
		    "declaration": true,
		    "inlineSourceMap": true,
		    "outDir": "{{ .tsFileDir }}/dist",
		    "rootDirs": [
		      ".",
		      "{{ .tsFileDir }}"
		    ],
		    "typeRoots": [
		      ".anemos/types"
		    ],
		    "baseUrl": ".",
		    "paths": {
		      "@ohayocorp/anemos": [
		        ".anemos/types/index.d.ts"
		      ],
		      "@ohayocorp/anemos/*": [
		        ".anemos/types/*"
		      ]
		    }
		  },
		  "include": [
		    "{{ .tsFileSlash }}"
		  ]
		}`,
		map[string]string{
			"tsFileDir":   tsFileDir,
			"tsFileSlash": tsFileSlash,
		},
	)

	tsconfigPath := filepath.Join(tempDir, "tsconfig.json")
	err = os.WriteFile(tsconfigPath, []byte(tsconfig), 0644)
	if err != nil {
		return fmt.Errorf("failed to write tsconfig.json: %w", err)
	}

	err = js.RunTsc(tempDir)
	if err != nil {
		return err
	}

	return nil
}
