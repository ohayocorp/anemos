package js

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/grafana/sobek"
	"github.com/ohayocorp/sobek_nodejs/console"
	"github.com/ohayocorp/sobek_nodejs/process"
	"github.com/ohayocorp/sobek_nodejs/require"
)

const PackageName = "@ohayocorp/anemos"

type JsRuntime struct {
	MainScriptPath         string
	runtime                *sobek.Runtime
	variableRegistrations  []*VariableRegistration
	functionRegistrations  []*FunctionRegistration
	typeRegistrations      map[reflect.Type]*TypeRegistration
	templates              map[reflect.Type]*DynamicObjectTemplate
	typeConversions        map[reflect.Type][]*TypeConversion
	functions              []*DynamicFunction
	disabledObjectMappings mapset.Set[reflect.Type]
}

func (jsRuntime *JsRuntime) CheckInsideTheMainScriptDirectory(filePath string) error {
	// Main script path has already been resolved, so we can use it directly.
	mainScriptDirectory := filepath.Dir(jsRuntime.MainScriptPath)

	filePath, err := resolvePath(filePath, true)
	if err != nil {
		return err
	}

	relPath, err := filepath.Rel(mainScriptDirectory, filePath)
	if err != nil || strings.HasPrefix(relPath, "..") || filepath.IsAbs(relPath) {
		return fmt.Errorf("file path %s is not inside the main script directory %s", filePath, mainScriptDirectory)
	}

	return nil
}

func resolvePath(path string, mayNotExist bool) (string, error) {
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

func sourceLoader(path string) ([]byte, error) {
	if runtime.GOOS == "windows" {
		match, _ := regexp.Match(`^([a-zA-Z]):///`, []byte(path))
		if match {
			path = regexp.MustCompile(`^([a-zA-Z]):///`).ReplaceAllString(path, `$1:/`)
			path = strings.ReplaceAll(path, "/", "\\")
		}
	}

	return require.DefaultSourceLoader(path)
}

func NewJsRuntime() (*JsRuntime, error) {
	runtime := sobek.New()

	jsRuntime := &JsRuntime{
		runtime:                runtime,
		typeRegistrations:      make(map[reflect.Type]*TypeRegistration),
		typeConversions:        make(map[reflect.Type][]*TypeConversion),
		templates:              make(map[reflect.Type]*DynamicObjectTemplate),
		disabledObjectMappings: mapset.NewSet[reflect.Type](),
	}

	registry := &require.Registry{}
	require.WithLoader(sourceLoader)(registry)
	require.WithPathResolver(func(base, path string) string {
		return pathResolver(jsRuntime, base, path)
	})(registry)

	registry.Enable(runtime)
	console.Enable(runtime)
	process.Enable(runtime)
	runtime.Set("exports", runtime.NewObject())

	err := jsRuntime.initialize()
	if err != nil {
		return nil, err
	}

	return jsRuntime, nil
}

func (jsRuntime *JsRuntime) GetStackTrace() []sobek.StackFrame {
	return jsRuntime.runtime.CaptureCallStack(0, nil)
}

func (jsRuntime *JsRuntime) GetEnv(key string) *string {
	value, err := jsRuntime.runtime.RunString(fmt.Sprintf("process.env.%s", key))
	if err != nil {
		Throw(fmt.Errorf("failed to get environment variable %s: %w", key, err))
	}

	if value == nil || value == sobek.Undefined() || value == sobek.Null() {
		return nil
	}

	valueString := value.String()
	return &valueString
}

func (jsRuntime *JsRuntime) ToSobekValue(object any) sobek.Value {
	return jsRuntime.runtime.ToValue(object)
}

func (jsRuntime *JsRuntime) ToSobekObject(value sobek.Value) *sobek.Object {
	return value.ToObject(jsRuntime.runtime)
}

func (jsRuntime *JsRuntime) Run(jsFile string, args []string) error {
	jsFile, err := filepath.Abs(jsFile)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for %s: %w", jsFile, err)
	}

	jsFile, err = filepath.EvalSymlinks(jsFile)
	if err != nil {
		return fmt.Errorf("failed to resolve symlinks for %s: %w", jsFile, err)
	}

	defer func() {
		jsRuntime.MainScriptPath = ""
	}()

	jsRuntime.MainScriptPath, err = resolvePath(jsFile, false)
	if err != nil {
		return err
	}

	jsArgs, err := jsRuntime.MarshalToJs(reflect.ValueOf(args))
	if err != nil {
		return fmt.Errorf("failed to marshal args: %w", err)
	}

	err = jsRuntime.runtime.Set("args", jsArgs)
	if err != nil {
		return fmt.Errorf("failed to set args: %w", err)
	}

	_, err = jsRuntime.runtime.RunString("require('process').argv = args; delete args;")
	if err != nil {
		return fmt.Errorf("failed to set process.argv: %w", err)
	}

	script, err := os.ReadFile(jsFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %s, %w", jsFile, err)
	}

	_, err = jsRuntime.runtime.RunScript(jsFile, string(script))
	if err != nil {
		if ex, ok := err.(*sobek.Exception); ok {
			return fmt.Errorf("failed to run script %s: %s", jsFile, ex.String())
		} else {
			return fmt.Errorf("failed to run script %s: %w", jsFile, err)
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

	template := &DynamicObjectTemplate{
		jsRuntime:          jsRuntime,
		objectType:         objectType,
		goToJsNameMappings: make(map[string]string),
		jsToGoNameMappings: make(map[string][]string),
		prototype:          jsRuntime.runtime.NewObject(),
	}

	jsRuntime.templates[objectType] = template

	return template
}

func (jsRuntime *JsRuntime) initialize() error {
	require.RegisterNativeModule(PackageName, func(runtime *sobek.Runtime, module *sobek.Object) {
		jsRuntime.registerTypes()
		jsRuntime.registerFunctions()

		rootObject := module.Get("exports").(*sobek.Object)

		for _, variable := range jsRuntime.variableRegistrations {
			value := variable.value

			object, err := jsRuntime.MarshalToJs(value)
			if err != nil {
				panic(err)
			}

			value = reflect.ValueOf(object)

			err = jsRuntime.addToNamespace(rootObject, variable.jsNamespace, variable.jsName, value)
			if err != nil {
				panic(err)
			}
		}

		for _, template := range jsRuntime.templates {
			template.Initialize(rootObject)

			namespace, err := jsRuntime.getNamespace(rootObject, template.jsNamespace)
			if err != nil {
				panic(err)
			}

			class := namespace.Get(template.jsName)

			// Set the prototype of the class to the prototype of the template. This way users can
			// add functions to the prototypes of the built-in classes.
			if classObject, ok := class.(*sobek.Object); ok {
				classObject.Set("prototype", template.prototype)
			} else {
				// Class doesn't have a constructor, create a new object and set the prototype.
				object := runtime.NewObject()
				object.SetPrototype(template.prototype)

				err := rootObject.Set(template.jsName, object)
				if err != nil {
					panic(err)
				}
			}
		}

		jsRuntime.initializeFunctions(rootObject, jsRuntime.functions, nil)
	})

	return nil
}

func (jsRuntime *JsRuntime) addToNamespace(rootObject *sobek.Object, namespace, name string, value reflect.Value) error {
	ns, err := jsRuntime.getNamespace(rootObject, namespace)
	if err != nil {
		return err
	}

	err = ns.Set(name, value.Interface())

	if err != nil {
		return fmt.Errorf("failed to set value name: %s, namespace: %s, %w", name, namespace, err)
	}

	return nil
}

func (jsRuntime *JsRuntime) getNamespace(rootObject *sobek.Object, namespace string) (*sobek.Object, error) {
	currentNamespace := rootObject

	tokens := strings.Split(namespace, ".")
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "" {
			continue
		}

		ns, ok := currentNamespace.Get(tokens[i]).(*sobek.Object)
		if !ok {
			ns = jsRuntime.runtime.NewObject()
			err := currentNamespace.Set(tokens[i], ns)
			if err != nil {
				return nil, fmt.Errorf("failed to set namespace: %s, token: %s, %w", namespace, tokens[i], err)
			}
		}

		currentNamespace = ns
	}

	return currentNamespace, nil
}
