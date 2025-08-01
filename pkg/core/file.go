package core

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
)

func CheckPathInsideMainScriptDirectory(jsRuntime *js.JsRuntime, filePath string) {
	err := jsRuntime.CheckInsideTheMainScriptDirectory(filePath)
	if err != nil {
		js.Throw(err)
	}
}

func ReadAllText(jsRuntime *js.JsRuntime, filePath string) string {
	filePath = filepath.Clean(filePath)
	CheckPathInsideMainScriptDirectory(jsRuntime, filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		js.Throw(err)
	}

	return string(data)
}

func ReadAllBytes(jsRuntime *js.JsRuntime, filePath string) []byte {
	filePath = filepath.Clean(filePath)
	CheckPathInsideMainScriptDirectory(jsRuntime, filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		js.Throw(err)
	}

	return data
}

func WriteAllText(jsRuntime *js.JsRuntime, filePath string, content string) {
	filePath = filepath.Clean(filePath)
	CheckPathInsideMainScriptDirectory(jsRuntime, filePath)

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		js.Throw(err)
	}
}

func WriteAllBytes(jsRuntime *js.JsRuntime, filePath string, data []byte) {
	filePath = filepath.Clean(filePath)
	CheckPathInsideMainScriptDirectory(jsRuntime, filePath)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		js.Throw(err)
	}
}

func MainScriptPath(jsRuntime *js.JsRuntime) string {
	return jsRuntime.MainScriptPath
}

func MainScriptDirectory(jsRuntime *js.JsRuntime) string {
	return filepath.Dir(jsRuntime.MainScriptPath)
}

func CurrentScriptPath(jsRuntime *js.JsRuntime) string {
	frames := jsRuntime.GetStackTrace()
	for i := 0; i < len(frames); i++ {
		frame := frames[i]
		fileName := frame.SrcName()
		if fileName != "" && fileName != "<native>" {
			return filepath.Clean(fileName)
		}
	}

	return ""
}

func CurrentScriptDirectory(jsRuntime *js.JsRuntime) string {
	currentScriptPath := CurrentScriptPath(jsRuntime)
	return filepath.Dir(currentScriptPath)
}

func registerFile(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(ReadAllText)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(ReadAllBytes)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(WriteAllText)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(WriteAllBytes)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(MainScriptPath)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(MainScriptDirectory)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(MainScriptPath)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(MainScriptDirectory)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(CurrentScriptPath)).JsNamespace("file")
	jsRuntime.Function(reflect.ValueOf(CurrentScriptDirectory)).JsNamespace("file")
}
