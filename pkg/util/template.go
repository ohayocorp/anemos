package util

import (
	"bytes"
	"fmt"
	"text/template"
)

// Parses a template as a string by calling [MultilineString] on template text beforehand.
// Panics if template is invalid.
func ParseTemplate(templateText string, data any) string {
	templateText = MultilineString(templateText)
	template := template.Must(template.New("template").Parse(templateText))

	var buffer bytes.Buffer

	if err := template.Execute(&buffer, data); err != nil {
		panic(fmt.Errorf("can't parse template, %v\n%s", err, templateText))
	}

	return buffer.String()
}
