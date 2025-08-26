package js

import (
	"errors"
	"reflect"
	"slices"
	"sort"

	"github.com/grafana/sobek"
)

type DynamicObjectCustomGetterSetter interface {
	GetKeys(jsRuntime *JsRuntime) []string
	Get(jsRuntime *JsRuntime, key string) (any, bool)
	Set(jsRuntime *JsRuntime, key string, value sobek.Value) bool
}

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

	jsObject := getSobekObject(d.backingObject)
	if jsObject != nil {
		return jsObject.Get(originalKey)
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
		panic(d.jsRuntime.Runtime.ToValue(errors.Join(marshalErrors...)))
	}

	if hasNullResult {
		return sobek.Null()
	}

	if dynamicGetterSetter, ok := d.backingObject.Interface().(DynamicObjectCustomGetterSetter); ok {
		if value, ok := dynamicGetterSetter.Get(d.jsRuntime, originalKey); ok {
			result, err := d.jsRuntime.MarshalToJs(reflect.ValueOf(value))
			if err != nil {
				panic(d.jsRuntime.Runtime.ToValue(err))
			}

			return result
		}
	}

	prototypeValue := template.prototype.Get(originalKey)
	if prototypeValue != nil && prototypeValue != sobek.Undefined() {
		return prototypeValue
	}

	return nil
}

func (d *DynamicObject) Set(originalKey string, value sobek.Value) bool {
	template := d.template
	backingObject := d.backingObject

	jsObject := getSobekObject(d.backingObject)
	if jsObject != nil {
		err := jsObject.Set(originalKey, value)
		return err == nil
	}

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
		panic(d.jsRuntime.Runtime.ToValue(errors.Join(marshalErrors...)))
	}

	if dynamicGetterSetter, ok := backingObject.Interface().(DynamicObjectCustomGetterSetter); ok {
		if ok := dynamicGetterSetter.Set(d.jsRuntime, originalKey, value); ok {
			return true
		}
	}

	if d.jsPropertyStore == nil {
		d.jsPropertyStore = make(map[string]sobek.Value)
	}

	d.jsPropertyStore[originalKey] = value
	return true
}

func (d *DynamicObject) Has(key string) bool {
	keys := d.Keys()
	return slices.Contains(keys, key)
}

func (d *DynamicObject) Delete(key string) bool {
	jsObject := getSobekObject(d.backingObject)
	if jsObject != nil {
		err := jsObject.Delete(key)
		return err == nil
	}

	if d.jsPropertyStore != nil {
		if _, ok := d.jsPropertyStore[key]; ok {
			delete(d.jsPropertyStore, key)
			return true
		}
	}

	return false
}

func (d *DynamicObject) Keys() []string {
	jsObject := getSobekObject(d.backingObject)
	if jsObject != nil {
		return jsObject.GetOwnPropertyNames()
	}

	keys := make([]string, 0, len(d.template.jsToGoNameMappings)+len(d.jsPropertyStore))

	for key := range d.template.jsToGoNameMappings {
		if !d.template.exportedFields.Contains(key) {
			continue
		}

		includeKey := true

		if d.template.keysWithOmitEmpty.Contains(key) {
			value := d.Get(key)
			includeKey = value != sobek.Undefined() && value != sobek.Null()
		}

		if includeKey {
			keys = append(keys, key)
		}
	}

	for key := range d.jsPropertyStore {
		keys = append(keys, key)
	}

	if dynamicGetterSetter, ok := d.backingObject.Interface().(DynamicObjectCustomGetterSetter); ok {
		keys = append(keys, dynamicGetterSetter.GetKeys(d.jsRuntime)...)
	}

	sort.Strings(keys)

	return keys
}
