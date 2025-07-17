package core

import (
	"fmt"
)

func ExampleSerializeToYaml() {
	document := NewEmptyDocument("test.yaml")
	foo := NewEmptyMapping()

	document.GetRoot().SetMapping("foo", foo)
	foo.SetValue("bar", "baz")

	yaml := SerializeToYaml(document)

	fmt.Println(yaml)

	// Output:
	// foo:
	//   bar: baz
}

func ExampleParseDocument() {
	text := MultilineString(`
		apiVersion: v1
		kind: ConfigMap
		metadata:
		  name: example
		  namespace: default
		  labels:
		    app.kubernetes.io/name: example
		data:
		  multiline: |
		    lorem
		    ipsum`)

	document := ParseDocument("test.yaml", text)

	fmt.Println(*document.GetLabel("app.kubernetes.io/name"))
	fmt.Println(*document.GetRoot().GetValueChain("data", "multiline"))

	// Output:
	// example
	// lorem
	// ipsum
}

func ExampleParseTemplateAsDocument() {
	document := ParseTemplateAsDocument(
		"test.yaml",
		`
		apiVersion: v1
		kind: ConfigMap
		metadata:
		  name: {{ .Name }}
		  namespace: {{ .Namespace }}
		`,
		map[string]interface{}{
			"Name":      "example",
			"Namespace": "default",
		})

	fmt.Println(*document.GetName())
	fmt.Println(*document.GetNamespace())

	// Output:
	// example
	// default
}
