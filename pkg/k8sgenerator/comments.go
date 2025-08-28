package main

import (
	"fmt"
	"strings"
)

// getDescription formats an OpenAPI description as Go comments.
// It capitalizes the first sentence and prefixes each line with "//".
func getDescription(description string) string {
	if description == "" {
		return ""
	}

	sb := strings.Builder{}

	for i, line := range strings.Split(description, "\n") {
		if line == "" {
			continue
		}

		if i == 0 {
			line = strings.ToUpper(line[0:1]) + line[1:]
		}

		sb.WriteString(fmt.Sprintf("// %s\n", line))
	}

	return strings.TrimSpace(sb.String())
}

// toJsComment converts Go comment format to JSDoc format for TypeScript declarations.
// It handles multi-line comments and escapes JSDoc comment terminators.
func toJsComment(comment string) string {
	builder := strings.Builder{}
	builder.WriteString("/**\n")

	lines := strings.Split(comment, "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}

		line = strings.TrimPrefix(line, "// ")
		// Prevent closing comment in the middle of a line:
		// https://github.com/jsdoc/jsdoc/issues/821#issuecomment-385324492
		line = strings.ReplaceAll(line, "*/", "*&#8205;/")

		builder.WriteString(" * ")
		builder.WriteString(line)
		builder.WriteString("\n")

		// Putting empty line results in better formatting in VSCode.
		if i != len(lines)-1 {
			builder.WriteString("\n")
		}
	}

	builder.WriteString(" */")

	return builder.String()
}
