package js

import (
	"errors"
	"reflect"
	"slices"

	"github.com/grafana/sobek"
)

type DynamicObject struct {
	jsRuntime       *JsRuntime
	template        *DynamicObjectTemplate
	backingObject   reflect.Value
	jsPropertyStore map[string]sobek.Value
}

func (d *DynamicObject) Get(originalKey string) sobek.Value {
	template := d.template
	backingObject := d.backingObject

	if backingObject.Kind() == reflect.Ptr {
		backingObject = backingObject.Elem()
	}

	mappedKeys, ok := template.jsToGoNameMappings[originalKey]
	if !ok {
		mappedKeys = []string{originalKey}
	}

	marshalErrors := make([]error, 0)
	hasNullResult := false

	for _, key := range mappedKeys {
		if _, ok := backingObject.Type().FieldByName(key); ok {
			fieldValue := backingObject.FieldByName(key)
			result, err := d.jsRuntime.MarshalToJs(fieldValue)
			if err != nil {
				marshalErrors = append(marshalErrors, err)
				continue
			}

			if result == sobek.Null() {
				hasNullResult = true
				continue
			}

			return result
		}

		if d.jsPropertyStore != nil {
			if value, ok := d.jsPropertyStore[key]; ok {
				return value
			}
		}
	}

	if len(marshalErrors) == len(mappedKeys) {
		panic(d.jsRuntime.runtime.ToValue(errors.Join(marshalErrors...)))
	}

	if hasNullResult {
		return sobek.Null()
	}

	return template.prototype.Get(originalKey)
}

func (d *DynamicObject) Set(originalKey string, value sobek.Value) bool {
	template := d.template
	backingObject := d.backingObject

	mappedKeys, ok := template.jsToGoNameMappings[originalKey]
	if !ok {
		mappedKeys = []string{originalKey}
	}

	underlyingObject := backingObject.Elem()
	underlyingType := underlyingObject.Type()

	marshalErrors := make([]error, 0)

	for _, key := range mappedKeys {
		if _, ok := underlyingType.FieldByName(key); !ok {
			continue
		}

		field := underlyingObject.FieldByName(key)

		exported, err := d.jsRuntime.MarshalToGo(value, field.Type())
		if err != nil {
			marshalErrors = append(marshalErrors, err)
			continue
		}

		field.Set(exported)
		return true
	}

	if len(marshalErrors) == len(mappedKeys) {
		panic(d.jsRuntime.runtime.ToValue(errors.Join(marshalErrors...)))
	}

	if d.jsPropertyStore == nil {
		d.jsPropertyStore = make(map[string]sobek.Value)
	}

	d.jsPropertyStore[originalKey] = value
	return true
}

func (d *DynamicObject) Has(key string) bool {
	return slices.Contains(d.template.sortedKeys, key)
}

func (d *DynamicObject) Delete(key string) bool {
	return false
}

func (d *DynamicObject) Keys() []string {
	return d.template.sortedKeys
}
