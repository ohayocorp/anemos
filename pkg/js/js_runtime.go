package js

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/grafana/sobek"
	"github.com/ohayocorp/anemos/pkg"
	"github.com/ohayocorp/anemos/pkg/util"
	"github.com/ohayocorp/sobek_nodejs/console"
	"github.com/ohayocorp/sobek_nodejs/process"
	"github.com/ohayocorp/sobek_nodejs/require"
)

const PackageName = "@ohayocorp/anemos"

type JsRuntime struct {
	MainScriptPath         string
	Registry               *require.Registry
	Runtime                *sobek.Runtime
	Flags                  map[string]string
	EmbeddedModules        []*EmbeddedModule
	variableRegistrations  []*VariableRegistration
	functionRegistrations  []*FunctionRegistration
	typeRegistrations      map[reflect.Type]*TypeRegistration
	templates              map[reflect.Type]*DynamicObjectTemplate
	typeConversions        map[reflect.Type][]*TypeConversion
	functions              []*DynamicFunction
	disabledObjectMappings mapset.Set[reflect.Type]
}

type JsScript struct {
	Contents string
	FilePath string
}

type EmbeddedModule struct {
	ModulePath string
	Files      fs.FS
}

func (jsRuntime *JsRuntime) CheckInsideTheMainScriptDirectory(filePath string) error {
	// Main script path has already been resolved, so we can use it directly.
	mainScriptDirectory := filepath.Dir(jsRuntime.MainScriptPath)

	filePath, err := ResolvePath(filePath, true)
	if err != nil {
		return err
	}

	relPath, err := filepath.Rel(mainScriptDirectory, filePath)
	if err != nil || strings.HasPrefix(relPath, "..") || filepath.IsAbs(relPath) {
		return fmt.Errorf("file path %s is not inside the main script directory %s", filePath, mainScriptDirectory)
	}

	return nil
}

func ResolvePath(path string, mayNotExist bool) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	path, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return "", fmt.Errorf("failed to resolve file path, %w", err)
	}

	// Resolve symlinks in the target path.
	resolvedPath, err := filepath.EvalSymlinks(path)
	if err == nil {
		path = resolvedPath
	} else {
		if mayNotExist && os.IsNotExist(err) {
			// If the file doesn't exist, continue with the original path. Throw
			// for other errors.
			return path, nil
		}

		return "", fmt.Errorf("failed to resolve file path symlinks, %w", err)
	}

	return path, nil
}

func pathResolver(jsRuntime *JsRuntime, base, path string) string {
	for _, module := range jsRuntime.EmbeddedModules {
		if strings.HasPrefix(path, module.ModulePath) {
			// Module will be loaded from embedded modules, base path is ignored.
			return path
		}
	}

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}

	fileName := filepath.Base(path)
	if fileName == "package.json" {
		// If the path is a package.json file, search all directories traversing up
		// the file system until we find a directory with a package.json file.
		parent := base
		previousParent := ""

		for {
			// Break if we passed the main script directory.
			err := jsRuntime.CheckInsideTheMainScriptDirectory(parent)
			if err != nil {
				break
			}

			if parent == previousParent {
				break
			}

			pkgPath := filepath.Join(parent, "package.json")
			if _, err := os.Stat(pkgPath); err == nil {
				return pkgPath
			}

			previousParent = parent
			parent = filepath.Dir(parent)
		}
	}

	return require.DefaultPathResolver(base, path)
}

func SourceLoader(jsRuntime *JsRuntime, path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		// Download the file from the URL.
		response, err := http.Get(path)
		if err != nil {
			return nil, fmt.Errorf("failed to download file from %s: %w", path, err)
		}
		defer response.Body.Close()

		return io.ReadAll(response.Body)
	}

	for _, module := range jsRuntime.EmbeddedModules {
		pathClean := filepath.Clean(path)
		modulePathClean := filepath.Clean(module.ModulePath)

		relativePath, err := filepath.Rel(modulePathClean, pathClean)
		if err != nil {
			continue
		}

		relativePath = filepath.ToSlash(relativePath)

		if relativePath == "." || relativePath == ".." || strings.HasPrefix(relativePath, "../") || strings.HasPrefix(relativePath, "./") {
			continue
		}

		// Remove the module identifier from the path and get the file from embedded filesystem.
		result, err := module.Files.Open(relativePath)
		defer func() {
			if result != nil {
				result.Close()
			}
		}()

		// Goja expects ModuleFileDoesNotExistError if the file doesn't exist to continue searching for the module.
		if errors.Is(err, fs.ErrNotExist) {
			return nil, require.ModuleFileDoesNotExistError
		}

		if err != nil {
			return nil, err
		}

		if stat, err := result.Stat(); err == nil {
			// Check if the file is a regular file.
			if !stat.Mode().IsRegular() {
				return nil, require.ModuleFileDoesNotExistError
			}
		}

		// Read the file contents from the embedded filesystem.
		data, err := io.ReadAll(result)
		if err != nil {
			return nil, fmt.Errorf("failed to read embedded file %s: %w", relativePath, err)
		}

		return data, nil
	}

	if runtime.GOOS == "windows" {
		match, _ := regexp.Match(`^([a-zA-Z]):///`, []byte(path))
		if match {
			path = regexp.MustCompile(`^([a-zA-Z]):///`).ReplaceAllString(path, `$1:/`)
			path = strings.ReplaceAll(path, "/", "\\")
		}
	}

	return require.DefaultSourceLoader(path)
}

func NewJsRuntime() *JsRuntime {
	runtime := sobek.New()

	jsRuntime := &JsRuntime{
		Runtime:                runtime,
		Flags:                  make(map[string]string),
		typeRegistrations:      make(map[reflect.Type]*TypeRegistration),
		typeConversions:        make(map[reflect.Type][]*TypeConversion),
		templates:              make(map[reflect.Type]*DynamicObjectTemplate),
		disabledObjectMappings: mapset.NewSet[reflect.Type](),
	}

	registry := &require.Registry{}

	require.WithLoader(func(path string) ([]byte, error) {
		return SourceLoader(jsRuntime, path)
	})(registry)

	require.WithPathResolver(func(base, path string) string {
		return pathResolver(jsRuntime, base, path)
	})(registry)

	registry.Enable(runtime)
	console.Enable(runtime)
	process.Enable(runtime)
	runtime.Set("exports", runtime.NewObject())

	jsRuntime.Registry = registry

	return jsRuntime
}

func (jsRuntime *JsRuntime) GetStackTrace() []sobek.StackFrame {
	return jsRuntime.Runtime.CaptureCallStack(0, nil)
}

func (jsRuntime *JsRuntime) GetEnv(key string) *string {
	value, err := jsRuntime.Runtime.RunString(fmt.Sprintf("process.env.%s", key))
	if err != nil {
		Throw(fmt.Errorf("failed to get environment variable %s: %w", key, err))
	}

	if value == nil || value == sobek.Undefined() || value == sobek.Null() {
		return nil
	}

	valueString := value.String()
	return &valueString
}

func (jsRuntime *JsRuntime) Run(script *JsScript, args []string) error {
	if script == nil || script.Contents == "" {
		return fmt.Errorf("no script provided to run")
	}

	jsRuntime.MainScriptPath = script.FilePath
	defer func() {
		jsRuntime.MainScriptPath = ""
	}()

	jsArgs, err := jsRuntime.MarshalToJs(reflect.ValueOf(args))
	if err != nil {
		return fmt.Errorf("failed to marshal args: %w", err)
	}

	err = jsRuntime.Runtime.Set("args", jsArgs)
	if err != nil {
		return fmt.Errorf("failed to set args: %w", err)
	}

	_, err = jsRuntime.Runtime.RunString("require('process').argv = args; delete args;")
	if err != nil {
		return fmt.Errorf("failed to set process.argv: %w", err)
	}

	_, err = jsRuntime.Runtime.RunScript(script.FilePath, string(script.Contents))
	if err != nil {
		if ex, ok := err.(*sobek.Exception); ok {
			return fmt.Errorf("failed to run script %s:\n%s", script.FilePath, ex.String())
		} else {
			return fmt.Errorf("failed to run script %s:\n%w", script.FilePath, err)
		}
	}

	return nil
}

func (jsRuntime *JsRuntime) getTemplate(objectType reflect.Type) *DynamicObjectTemplate {
	return jsRuntime.templates[objectType]
}

func (jsRuntime *JsRuntime) createTemplate(objectType reflect.Type) *DynamicObjectTemplate {
	if objectType.Kind() != reflect.Struct {
		panic(fmt.Sprintf("objectType %v must be a struct type", objectType))
	}

	if template, ok := jsRuntime.templates[objectType]; ok {
		return template
	}

	template := NewDynamicObjectTemplate(jsRuntime, objectType)
	jsRuntime.templates[objectType] = template

	return template
}

func (jsRuntime *JsRuntime) InitializeNativeLibraries() error {
	jsRuntime.registerTypes()
	jsRuntime.registerFunctions()

	moduleNames := mapset.NewSet[string]()
	variables := map[string][]*VariableRegistration{}
	templates := map[string][]*DynamicObjectTemplate{}
	functions := map[string][]*DynamicFunction{}

	for _, variable := range jsRuntime.variableRegistrations {
		variables[variable.jsModule] = append(variables[variable.jsModule], variable)
		moduleNames.Add(variable.jsModule)
	}

	for _, template := range jsRuntime.templates {
		templates[template.jsModule] = append(templates[template.jsModule], template)
		moduleNames.Add(template.jsModule)
	}

	for _, function := range jsRuntime.functions {
		functions[function.jsModule] = append(functions[function.jsModule], function)
		moduleNames.Add(function.jsModule)
	}

	// Ensure intermediate module names are registered so that require("@ohayocorp/anemos/k8s/core")
	// works even if only deeper modules like "@ohayocorp/anemos/k8s/core/v1" exist.
	initialModules := moduleNames.ToSlice()
	for _, m := range initialModules {
		if m == "" {
			continue
		}

		parts := strings.Split(m, "/")
		// Add all parent prefixes like "k8s" and "k8s/core".
		for i := 1; i < len(parts); i++ {
			parent := strings.Join(parts[:i], "/")
			moduleNames.Add(parent)
		}
	}

	moduleNamesSlice := moduleNames.ToSlice()
	sort.Strings(moduleNamesSlice)

	for _, moduleVariables := range variables {
		sort.SliceStable(moduleVariables, func(i, j int) bool {
			return moduleVariables[i].jsName < moduleVariables[j].jsName
		})
	}

	for _, moduleTemplates := range templates {
		sort.SliceStable(moduleTemplates, func(i, j int) bool {
			return moduleTemplates[i].jsName < moduleTemplates[j].jsName
		})
	}

	for _, moduleFunctions := range functions {
		sort.SliceStable(moduleFunctions, func(i, j int) bool {
			return moduleFunctions[i].jsName < moduleFunctions[j].jsName
		})
	}

	// Register the main module so that require("@ohayocorp/anemos") works.
	require.RegisterNativeModule(PackageName, func(runtime *sobek.Runtime, jsModule *sobek.Object) {
	})

	for _, module := range moduleNamesSlice {
		moduleName := PackageName
		if module != "" {
			moduleName = fmt.Sprintf("%s/%s", moduleName, module)
		}

		require.RegisterNativeModule(moduleName, func(runtime *sobek.Runtime, jsModule *sobek.Object) {
			exports := jsModule.Get("exports").(*sobek.Object)

			for _, variable := range variables[module] {
				value := variable.value

				object, err := jsRuntime.MarshalToJs(value)
				if err != nil {
					panic(err)
				}

				value = reflect.ValueOf(object)

				err = exports.Set(variable.jsName, value.Interface())
				if err != nil {
					panic(fmt.Errorf("failed to set value name: %s on module: %s, %w", variable.jsName, variable.jsModule, err))
				}
			}

			for _, template := range templates[module] {
				template.Initialize(exports)

				class := exports.Get(template.jsName)

				// Set the prototype of the class to the prototype of the template. This way users can
				// add functions to the prototypes of the built-in classes.
				if classObject, ok := class.(*sobek.Object); ok {
					classObject.Set("prototype", template.prototype)
				} else {
					// Class doesn't have a constructor, create a new object and set the prototype.
					object := runtime.NewObject()
					object.Set("prototype", template.prototype)

					err := exports.Set(template.jsName, object)
					if err != nil {
						panic(err)
					}
				}
			}

			jsRuntime.initializeFunctions(exports, functions[module], nil)
		})
	}

	// Export the child modules on the parent modules.
	for _, module := range moduleNamesSlice {
		if module == "" {
			continue
		}

		moduleName := fmt.Sprintf("%s/%s", PackageName, module)
		tokens := strings.Split(moduleName, "/")

		parentModule := strings.Join(tokens[:len(tokens)-1], "/")
		parent := require.Require(jsRuntime.Runtime, parentModule)

		err := parent.ToObject(jsRuntime.Runtime).Set(tokens[len(tokens)-1], require.Require(jsRuntime.Runtime, moduleName))
		if err != nil {
			return fmt.Errorf("failed to set submodule %s on parent module %s: %w", tokens[len(tokens)-1], parentModule, err)
		}
	}

	jsRuntime.initializeStringExtensions()

	return jsRuntime.initializeLib()
}

func (jsRuntime *JsRuntime) initializeLib() error {
	jsRuntime.EmbeddedModules = append(jsRuntime.EmbeddedModules, &EmbeddedModule{
		ModulePath: PackageName,
		Files:      pkg.LibJavaScript,
	})

	libIndexJs, err := fs.ReadFile(pkg.LibJavaScript, "index.js")
	if err != nil {
		return fmt.Errorf("failed to read lib index.js: %w", err)
	}

	script := util.ParseTemplate(`
		__anemos_existing_exports = exports;
        exports = require("@ohayocorp/anemos");

		{{ .IndexJs }}

		exports = __anemos_existing_exports;
		delete __anemos_existing_exports;
		`,
		map[string]string{
			"IndexJs": string(libIndexJs),
		})

	_, err = jsRuntime.Runtime.RunString(script)
	if err != nil {
		if sobekEx, ok := err.(*sobek.Exception); ok {
			return fmt.Errorf("failed to run lib index.js: %s", sobekEx.String())
		}

		return fmt.Errorf("failed to run lib index.js: %w", err)
	}

	return nil
}

func (jsRuntime *JsRuntime) initializeStringExtensions() {
	runtime := jsRuntime.Runtime

	runtime.Get("String").ToObject(runtime).Get("prototype").ToObject(runtime).Set("indent", func(call sobek.FunctionCall) sobek.Value {
		this := call.This.String()
		numberOfSpaces := call.Argument(0).ToInteger()

		return runtime.ToValue(util.Indent(this, int(numberOfSpaces)))
	})

	runtime.Get("String").ToObject(runtime).Get("prototype").ToObject(runtime).Set("dedent", func(call sobek.FunctionCall) sobek.Value {
		return runtime.ToValue(util.Dedent(call.This.String()))
	})

	runtime.Get("String").ToObject(runtime).Get("prototype").ToObject(runtime).Set("toKubernetesIdentifier", func(call sobek.FunctionCall) sobek.Value {
		return runtime.ToValue(util.ToKubernetesIdentifier(call.This.String()))
	})

	runtime.Get("String").ToObject(runtime).Get("prototype").ToObject(runtime).Set("base64Encode", func(call sobek.FunctionCall) sobek.Value {
		return runtime.ToValue(util.Base64Encode(call.This.String()))
	})

	runtime.Get("String").ToObject(runtime).Get("prototype").ToObject(runtime).Set("base64Decode", func(call sobek.FunctionCall) sobek.Value {
		result, err := util.Base64Decode(call.This.String())
		if err != nil {
			return runtime.ToValue(err)
		}

		return runtime.ToValue(result)
	})
}
