package js

import (
	"fmt"
	"log/slog"
	"reflect"

	"github.com/grafana/sobek"
)

type DynamicArray struct {
	jsRuntime    *JsRuntime
	backingSlice reflect.Value
}

func (jsRuntime *JsRuntime) NewDynamicArray(backingSlice reflect.Value) *sobek.Object {
	if backingSlice.Kind() != reflect.Slice {
		panic(fmt.Errorf("backingSlice must be a slice"))
	}

	dynamicArray := &DynamicArray{
		jsRuntime:    jsRuntime,
		backingSlice: backingSlice,
	}

	return jsRuntime.Runtime.NewDynamicArray(dynamicArray)
}

func (d *DynamicArray) get(index int) reflect.Value {
	return d.backingSlice.Index(index)
}

func (d *DynamicArray) set(index int, object reflect.Value) {
	d.get(index).Set(object)
}

func (d *DynamicArray) elemtype() reflect.Type {
	return d.backingSlice.Type().Elem()
}

func (d *DynamicArray) Len() int {
	return d.backingSlice.Len()
}

func (d *DynamicArray) Get(index int) sobek.Value {
	if index < 0 && index >= d.Len() {
		panic(d.jsRuntime.Runtime.ToValue(fmt.Errorf("index out of range")))
	}

	result, err := d.jsRuntime.MarshalToJs(d.get(index))
	if err != nil {
		panic(d.jsRuntime.Runtime.ToValue(err))
	}

	return result
}

func (d *DynamicArray) Set(index int, value sobek.Value) bool {
	if index < 0 || index >= d.Len() {
		callStack := d.jsRuntime.Runtime.CaptureCallStack(-1, nil)
		if len(callStack) > 0 && callStack[0].FuncName() == "push" {
			// Sobek doesn't call SetLen when pushing to an array.
			// We need to call it manually to ensure the backing slice is resized.
			length := d.Len()
			if length == 0 {
				length = 1
			} else {
				length *= 2
			}

			d.SetLen(length)
			return d.Set(index, value)
		}

		panic(d.jsRuntime.Runtime.ToValue(fmt.Errorf(
			"index out of range slice type: %s index: %d len: %d",
			d.backingSlice.Type().String(),
			index,
			d.Len())))
	}

	if value == nil || value == sobek.Undefined() {
		zero := reflect.Zero(d.elemtype())
		d.set(index, zero)

		return true
	}

	exported, err := d.jsRuntime.MarshalToGo(value, d.elemtype())
	if err != nil {
		panic(d.jsRuntime.Runtime.ToValue(err))
	}

	if dynamicObject, ok := exported.Interface().(*DynamicObject); ok {
		d.set(index, dynamicObject.backingObject)
		return true
	}

	d.set(index, exported)
	return true
}

func (d *DynamicArray) SetLen(newLen int) bool {
	if newLen < 0 {
		slog.Error("new length ${newLen} must be greater than or equal to 0", slog.Int("newLen", newLen))
		return false
	}

	if newLen == d.Len() {
		return true
	}

	newSlice := reflect.MakeSlice(reflect.SliceOf(d.elemtype()), newLen, newLen)
	oldSlice := d.backingSlice

	reflect.Copy(newSlice, oldSlice)
	oldSlice.Set(newSlice)

	return true
}
