package js

import (
	"fmt"
	"reflect"

	"github.com/grafana/sobek"
)

func (jsRuntime *JsRuntime) MarshalToJs(object reflect.Value) (sobek.Value, error) {
	underlyingKind := object.Kind()
	underlyingObject := object

	if object.Kind() == reflect.Ptr || object.Kind() == reflect.Interface {
		underlyingObject = object.Elem()
		underlyingKind = underlyingObject.Kind()
	}

	switch underlyingKind {
	case reflect.Slice:
		if underlyingObject.IsNil() {
			return sobek.Null(), nil
		}

		dynamicArray := jsRuntime.NewDynamicArray(underlyingObject)
		return jsRuntime.runtime.ToValue(dynamicArray), nil
	case reflect.Bool:
		return jsRuntime.runtime.ToValue(underlyingObject.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return jsRuntime.runtime.ToValue(underlyingObject.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return jsRuntime.runtime.ToValue(underlyingObject.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return jsRuntime.runtime.ToValue(underlyingObject.Float()), nil
	case reflect.String:
		return jsRuntime.runtime.ToValue(underlyingObject.String()), nil
	}

	if object.Kind() == reflect.Interface {
		if object.IsNil() {
			return sobek.Null(), nil
		}

		if sobekValue, ok := object.Interface().(sobek.Value); ok {
			return sobekValue, nil
		}
	}

	if object.Kind() == reflect.Ptr {
		if object.IsNil() {
			return sobek.Null(), nil
		}

		template := jsRuntime.getTemplate(underlyingObject.Type())

		if template == nil {
			return nil, fmt.Errorf("no template found for object: %s", object.Type().String())
		}

		dynamicObject := template.NewObject(object)

		return dynamicObject, nil
	}

	return nil, fmt.Errorf("unsupported type %s", object.Type().String())
}
