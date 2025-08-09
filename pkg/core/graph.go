package core

import (
	"bytes"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/ohayocorp/anemos/pkg/js"
)

type DependencyGraph[T comparable] struct {
	Elements           []T
	IdentifierGetter   func(element T) string
	DependenciesGetter func(T) *Dependencies[T]
}

func (d *DependencyGraph[T]) GetSortedElements() []T {
	d.sanitize()

	hash := func(element T) string {
		return d.IdentifierGetter(element)
	}

	hashes := make([]string, len(d.Elements))
	for i, element := range d.Elements {
		hashes[i] = hash(element)
	}

	g := graph.New(hash, graph.Directed(), graph.Acyclic(), graph.PreventCycles())

	addEdge := func(source, target T) error {
		err := g.AddEdge(hash(source), hash(target))
		if err == nil {
			return nil
		}

		if err == graph.ErrEdgeAlreadyExists {
			return nil
		}

		dot, dotErr := d.getDotRepresentation(g)
		if dotErr != nil {
			return dotErr
		}

		slog.Error("graph DOT representation: ${representation}", slog.String("representation", dot))
		return fmt.Errorf("can't add edge '%s' -> '%s' to graph, %w", hash(source), hash(target), err)
	}

	for _, element := range d.Elements {
		err := g.AddVertex(element)
		if err == graph.ErrVertexAlreadyExists {
			js.Throw(fmt.Errorf("duplicate element '%s' in the graph", d.IdentifierGetter(element)))
		}
	}

	for _, element := range d.Elements {
		dependencies := d.DependenciesGetter(element)

		for _, prerequisite := range dependencies.Prerequisites {
			if err := addEdge(prerequisite, element); err != nil {
				js.Throw(err)
			}
		}

		for _, dependent := range dependencies.Dependents {
			if err := addEdge(element, dependent); err != nil {
				js.Throw(err)
			}
		}
	}

	d.addAlphabeticalDependencies(g)

	// Keep the order of elements same as the given slice if dependencies don't
	// require a different order.
	sortedHashes, err := graph.StableTopologicalSort(g, func(x, y string) bool {
		xIndex := slices.Index(hashes, x)
		yIndex := slices.Index(hashes, y)
		return xIndex < yIndex
	})
	if err != nil {
		dot, dotErr := d.getDotRepresentation(g)
		if dotErr != nil {
			js.Throw(fmt.Errorf("can't get DOT representation of graph, %w", dotErr))
		}

		slog.Error("graph DOT representation: ${representation}", slog.String("representation", dot))
		js.Throw(fmt.Errorf("can't apply topological sort to graph, %w", err))
	}

	sortedElements := []T{}
	for _, h := range sortedHashes {
		element, err := g.Vertex(h)
		if err != nil {
			js.Throw(fmt.Errorf("can't get element from graph via identifier %s, %w", h, err))
		}

		sortedElements = append(sortedElements, element)
	}

	return sortedElements
}

func (d *DependencyGraph[T]) addAlphabeticalDependencies(g graph.Graph[string, T]) {
	hasDependency := func(element T) bool {
		return len(d.DependenciesGetter(element).Prerequisites) > 0 || len(d.DependenciesGetter(element).Dependents) > 0
	}

	// By default, elements are sorted alphabetically, and the topological sort will keep this order
	// unless dependencies require a different one. However, if an element has dependencies, elements
	// that come later alphabetically (and have no dependencies) might be scheduled before it (since
	// they are not blocked by any prerequisites).
	// To prevent this, we add extra edges between elements to enforce the alphabetical order,
	// ensuring that elements with dependencies are still run in the correct alphabetical sequence.
	for i := len(d.Elements) - 1; i >= 0; i-- {
		element := d.Elements[i]
		if !hasDependency(element) {
			continue
		}

		for j := 0; j < len(d.Elements); j++ {
			// Skip the same element.
			if i == j {
				continue
			}

			otherElement := d.Elements[j]

			var prerequisite, dependent T

			if i > j {
				// If the current element comes after the other element in the alphabetical order,
				// it should depend on the other element.
				prerequisite = otherElement
				dependent = element
			} else {
				// Otherwise, current element should depend on the other element.
				prerequisite = element
				dependent = otherElement
			}

			// Ignore the error since alphabetical edges are not critical.
			_ = g.AddEdge(d.IdentifierGetter(prerequisite), d.IdentifierGetter(dependent), func(ep *graph.EdgeProperties) {
				ep.Attributes["class"] = "alphabetical"
				ep.Attributes["color"] = "lightgreen"
			})
		}
	}
}

func (d *DependencyGraph[T]) getDotRepresentation(g graph.Graph[string, T]) (string, error) {
	buffer := bytes.Buffer{}
	if err := draw.DOT(g, &buffer); err != nil {
		return "", fmt.Errorf("can't get DOT representation of graph, %w", err)
	}

	result := buffer.String()
	result = strings.ReplaceAll(result, "\n\n", "\n")
	index := strings.Index(result, "{\n")

	if index == -1 {
		return "", fmt.Errorf("can't find '{' in DOT representation of graph")
	}

	attributes := MultilineString(`
		graph [ ratio = "auto", page = "100", compound = true, bgcolor = "#2e3e56", pad = 2, ranksep=1.0 ];
		node [ style = "filled", fillcolor = "#edad56", color = "#edad56", penwidth =3 ];
		edge [ color = "#fcfcfc", penwidth =2 ]
		`)

	attributes = strings.ReplaceAll(attributes, "\n", "\n\t")
	result = result[:index+1] + "\n\t" + attributes + result[index+1:]

	return result, nil
}

func (d *DependencyGraph[T]) sanitize() {
	if d.IdentifierGetter == nil {
		js.Throw(fmt.Errorf("IdentifierGetter can't be nil"))
	}

	if d.DependenciesGetter == nil {
		js.Throw(fmt.Errorf("DependenciesGetter can't be nil"))
	}
}
