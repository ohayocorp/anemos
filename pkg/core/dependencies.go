package core

import (
	"fmt"
	"slices"

	"github.com/ohayocorp/anemos/pkg/js"
)

// Dependencies specifies prerequisites and dependents of the given type. This information is used to
// create dependency graphs in order to sort dependent objects.
type Dependencies[T comparable] struct {
	Prerequisites []T
	Dependents    []T
}

func NewDependencies[T comparable]() *Dependencies[T] {
	return &Dependencies[T]{
		Prerequisites: []T{},
		Dependents:    []T{},
	}
}

// Makes this instance a prerequisite of the given element.
func (d *Dependencies[T]) RunBefore(element T) {
	var nilT T
	if element == nilT {
		js.Throw(fmt.Errorf("element cannot be nil"))
	}

	if !slices.Contains(d.Dependents, element) {
		d.Dependents = append(d.Dependents, element)
	}
}

// Makes the given element a prerequisite of this instance.
func (d *Dependencies[T]) RunAfter(element T) {
	var nilT T
	if element == nilT {
		js.Throw(fmt.Errorf("element cannot be nil"))
	}

	if !slices.Contains(d.Prerequisites, element) {
		d.Prerequisites = append(d.Prerequisites, element)
	}
}

// Appends dependency specifications of the given object into this object.
func (d *Dependencies[T]) Merge(other *Dependencies[T]) {
	if other == nil {
		js.Throw(fmt.Errorf("other dependencies cannot be nil"))
	}

	if other.Prerequisites != nil {
		d.Prerequisites = append(d.Prerequisites, other.Prerequisites...)
	}

	if other.Dependents != nil {
		d.Dependents = append(d.Dependents, other.Dependents...)
	}
}
