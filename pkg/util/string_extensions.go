package util

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Base64Decode(text string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Indents the each line with given number of spaces except the first line.
func Indent(text string, numberOfSpaces int) string {
	return indent(text, numberOfSpaces, " ")
}

// Indents the each line with given number of tabs except the first line.
func IndentTab(text string, numberOfTabs int) string {
	return indent(text, numberOfTabs, "\t")
}

func indent(text string, repeat int, indentation string) string {
	var builder strings.Builder
	fullIndentation := strings.Repeat(indentation, repeat)

	currentPos := 0
	processedFirstLine := false

	for currentPos < len(text) {
		if processedFirstLine {
			builder.WriteString(fullIndentation)
		}

		nextLF := strings.Index(text[currentPos:], "\n")

		if nextLF == -1 {
			builder.WriteString(text[currentPos:])
			currentPos = len(text)
		} else {
			// The line includes the newline character itself.
			lineEnd := currentPos + nextLF + 1
			builder.WriteString(text[currentPos:lineEnd])

			currentPos = lineEnd
			processedFirstLine = true
		}
	}

	return builder.String()
}

func Dedent(text string) string {
	getLeadingWhitespace := func(line string) string {
		for index, char := range line {
			if char != ' ' && char != '\t' {
				return line[:index]
			}
		}

		return line
	}

	scanner := bufio.NewScanner(strings.NewReader(text))

	// Find the longest common whitespace prefix.
	var longestPrefix string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		leadingWhitespace := getLeadingWhitespace(line)

		// If the leading whitespace is empty, then there is no common prefix. Return the text as is.
		if leadingWhitespace == "" {
			return text
		}

		// If the line only contains whitespace, then skip it.
		if leadingWhitespace == line {
			continue
		}

		// If the longest prefix is empty, then set it to the leading whitespace of the first line.
		if longestPrefix == "" {
			longestPrefix = leadingWhitespace
			continue
		}

		// If the current leading whitespace is longer than the longest prefix and starts with the longest prefix
		// then the longest prefix doesn't change since it is still the common prefix.
		if strings.HasPrefix(leadingWhitespace, longestPrefix) {
			continue
		}

		// Here the leading whitespace doesn't start with the longest prefix. If the longest prefix instead starts with the
		// leading whitespace, then the longest prefix will be the current leading whitespace which will be shorter.
		if strings.HasPrefix(longestPrefix, leadingWhitespace) {
			longestPrefix = leadingWhitespace
			continue
		}

		// Longest prefix and the current leading whitespace don't have a common prefix. Return the text as is.
		return text
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("failed to scan text for dedent: %w", err))
	}

	// This function is a copy of bufio.ScanLines with the difference that it doesn't remove the '\r' character.
	// It also sets eof variable to prevent extra newlines at the end of the text.
	isEOF := false
	scanLinesWithCR := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			isEOF = true
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	}

	scanner = bufio.NewScanner(strings.NewReader(text))
	scanner.Split(scanLinesWithCR)
	builder := strings.Builder{}

	appendNewLine := func() {
		if !isEOF {
			builder.WriteString("\n")
		}
	}

	// Remove the longest common whitespace prefix from each line.
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			appendNewLine()
			continue
		}

		whiteSpace := getLeadingWhitespace(line)

		// Remove the whitespace if it is the only content in the line.
		if whiteSpace == line {
			appendNewLine()
			continue
		}
		if line == whiteSpace+"\r" {
			builder.WriteString("\r")
			appendNewLine()
			continue
		}

		line = strings.TrimPrefix(line, longestPrefix)
		builder.WriteString(line)

		appendNewLine()
	}

	return builder.String()
}

// Dedents the data using [Dedent] so that the multiline strings with indentation are handled properly.
// Trims the spaces after dedent.
func MultilineString(text string) string {
	return strings.TrimSpace(Dedent(text))
}

var kubernetesIdentifierRegexp = regexp.MustCompile("[^a-zA-Z0-9-]")

// Converts a string to a valid Kubernetes identifier by replacing invalid characters.
func ToKubernetesIdentifier(name string) string {
	name = strings.ToLower(name)
	name = kubernetesIdentifierRegexp.ReplaceAllString(name, "-")

	// Maximum length of a label value is 63 characters.
	if len(name) > 63 {
		name = name[:63]
	}

	name = strings.Trim(name, "-")

	return name
}
