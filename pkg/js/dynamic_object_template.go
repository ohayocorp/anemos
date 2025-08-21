package js

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"weak"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/grafana/sobek"
)

type IteratorResult struct {
	Value sobek.Value
	Done  bool
}

type Iterator interface {
	Next(jsRuntime *JsRuntime, index int) IteratorResult
}

// WeakMap copied from: https://github.com/golang/go/issues/43615#issuecomment-2985815833
type WeakMap[K comparable, V any] struct {
	internalMap sync.Map
}

type entry[K, V any] struct {
	key       K
	weakValue weak.Pointer[V]
	cleanup   runtime.Cleanup
}

func (weakMap *WeakMap[K, V]) Store(key K, p *V) {
	e := &entry[K, V]{
		key:       key,
		weakValue: weak.Make(p),
	}

	e.cleanup = runtime.AddCleanup(p, func(e *entry[K, V]) {
		weakMap.internalMap.CompareAndDelete(e.key, e)
	}, e)

	old, ok := weakMap.internalMap.Swap(key, e)
	if ok {
		old.(*entry[K, V]).cleanup.Stop()
	}
}

func (weakMap *WeakMap[K, V]) Load(key K) *V {
	value, ok := weakMap.internalMap.Load(key)
	if !ok {
		return nil
	}

	return value.(*entry[K, V]).weakValue.Value()
}

type DynamicObjectTemplate struct {
	jsRuntime          *JsRuntime
	objectType         reflect.Type
	functions          []*DynamicFunction
	goToJsNameMappings map[string]string
	jsToGoNameMappings map[string][]string
	exportedFields     mapset.Set[string]
	prototype          *sobek.Object
	jsNamespace        string
	jsName             string
	objectStore        WeakMap[any, sobek.Object]
	keysWithOmitEmpty  mapset.Set[string]
}

func NewDynamicObjectTemplate(jsRuntime *JsRuntime, objectType reflect.Type) *DynamicObjectTemplate {
	return &DynamicObjectTemplate{
		jsRuntime:          jsRuntime,
		objectType:         objectType,
		goToJsNameMappings: make(map[string]string),
		jsToGoNameMappings: make(map[string][]string),
		exportedFields:     mapset.NewSet[string](),
		prototype:          jsRuntime.Runtime.NewObject(),
		keysWithOmitEmpty:  mapset.NewSet[string](),
	}
}

func (template *DynamicObjectTemplate) Initialize(module *sobek.Object) {
	template.jsRuntime.initializeFunctions(module, template.functions, template.prototype)

	for key := range template.jsToGoNameMappings {
		goNames := template.jsToGoNameMappings[key]
		for _, goName := range goNames {
			field, ok := template.objectType.FieldByName(goName)
			if !ok {
				continue
			}

			jsonTag := field.Tag.Get("json")
			if strings.Contains(jsonTag, "omitempty") {
				template.keysWithOmitEmpty.Add(key)
			}
		}
	}
}

func (template *DynamicObjectTemplate) NewObject(backingObject reflect.Value) *sobek.Object {
	pointerType := reflect.PointerTo(template.objectType)

	if backingObject.Kind() != reflect.Ptr {
		panic(fmt.Errorf("backingObject must be a pointer"))
	}

	if backingObject.Type() != pointerType {
		panic(fmt.Errorf("backingObject must be of type %s", pointerType.String()))
	}

	backingObjectSelf := backingObject.Interface()
	existingDynamicObject := template.objectStore.Load(backingObjectSelf)
	if existingDynamicObject != nil {
		return existingDynamicObject
	}

	dynamicObject := &DynamicObject{
		jsRuntime:     template.jsRuntime,
		template:      template,
		backingObject: backingObject,
	}

	object := template.jsRuntime.Runtime.NewDynamicObject(dynamicObject)

	if iterable, ok := backingObject.Interface().(Iterator); ok {
		prototype := template.jsRuntime.Runtime.NewObject()
		prototype.SetPrototype(template.prototype)

		prototype.SetSymbol(sobek.SymIterator, func() *sobek.Object {
			index := 0

			iterator := template.jsRuntime.Runtime.NewObject()
			iterator.Set("next", func() *sobek.Object {
				result := iterable.Next(template.jsRuntime, index)
				index++

				objectResult := template.jsRuntime.Runtime.NewObject()
				objectResult.Set("value", result.Value)
				objectResult.Set("done", result.Done)

				return objectResult
			})

			return iterator
		})

		object.SetPrototype(prototype)
	} else {
		object.SetPrototype(template.prototype)
	}

	template.objectStore.Store(backingObjectSelf, object)

	return object
}
