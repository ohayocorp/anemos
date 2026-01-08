package core

import (
	"fmt"
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
	"k8s.io/apimachinery/pkg/api/resource"
)

func NewQuantity(value string) *resource.Quantity {
	quantity, err := resource.ParseQuantity(value)
	if err != nil {
		js.Throw(fmt.Errorf("invalid resource quantity %s: %w", value, err))
	}

	return &quantity
}

func AddQuantity(x *resource.Quantity, y *resource.Quantity) *resource.Quantity {
	if x == nil || y == nil {
		return nil
	}

	result := x.DeepCopy()
	result.Add(*y)

	return &result
}

func SubtractQuantity(x *resource.Quantity, y *resource.Quantity) *resource.Quantity {
	if x == nil || y == nil {
		return nil
	}

	result := x.DeepCopy()
	result.Sub(*y)

	return &result
}

func MultiplyQuantity(q *resource.Quantity, factor int64) *resource.Quantity {
	if q == nil {
		return nil
	}

	result := q.DeepCopy()
	result.Mul(factor)

	return &result
}

func CompareQuantity(x *resource.Quantity, y *resource.Quantity) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return x.Cmp(*y)
}

func EqualsQuantity(x *resource.Quantity, y *resource.Quantity) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil || y == nil {
		return false
	}

	return x.Equal(*y)
}

func CloneQuantity(x *resource.Quantity) *resource.Quantity {
	if x == nil {
		return nil
	}

	clone := x.DeepCopy()
	return &clone
}

func registerQuantity(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[resource.Quantity]()).JsModule(
		"quantity",
	).Methods(
		js.Method("String").JsName("toString"),
	).ExtensionMethods(
		js.ExtensionMethod(reflect.ValueOf(AddQuantity)).JsName("add"),
		js.ExtensionMethod(reflect.ValueOf(SubtractQuantity)).JsName("subtract"),
		js.ExtensionMethod(reflect.ValueOf(MultiplyQuantity)).JsName("multiply"),
		js.ExtensionMethod(reflect.ValueOf(CompareQuantity)).JsName("compare"),
		js.ExtensionMethod(reflect.ValueOf(EqualsQuantity)).JsName("equals"),
		js.ExtensionMethod(reflect.ValueOf(CloneQuantity)).JsName("clone"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewQuantity)),
	)
}
