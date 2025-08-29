package util

import (
	"fmt"
	"testing"
)

func ExampleIndent() {
	indented := Indent(`line1
  line2
line3`, 2)

	fmt.Println(indented)

	// Output:
	// line1
	//     line2
	//   line3
}

func ExampleMultilineString() {
	indented := MultilineString(`
	line1
	  line2
	line3`)

	fmt.Println(indented)

	// Output:
	// line1
	//   line2
	// line3
}

type dedentTestCase struct {
	text, expected string
}

var dedentTestCases = []dedentTestCase{
	{"1\n2\n3", "1\n2\n3"},                       // No indentation
	{"1\n\n2", "1\n\n2"},                         // No indentation, with blank line
	{"1\n  2", "1\n  2"},                         // Some indentation, but no common whitespace prefix
	{"1\n\n  2\n", "1\n\n  2\n"},                 // Some indentation, but no common whitespace prefix with blank line
	{"  1\n  2\n  3", "1\n2\n3"},                 // All lines indented by two spaces
	{"  1\n\n  2\n  3\n", "1\n\n2\n3\n"},         // All lines indented by two spaces, with blank line
	{"  1\n  \n  2\n  3", "1\n\n2\n3"},           // All lines indented by two spaces, with whitespace-only line
	{"  1\n    2\n      3\n", "1\n  2\n    3\n"}, // Lines indented unevenly
	{"  1\n    2\n\n   3\n", "1\n  2\n\n 3\n"},   // Uneven indentation with a blank line
	{"  1\n    2\n \n   3\n", "1\n  2\n\n 3\n"},  // Uneven indentation with a whitespace-only line
	{"\t1\n\t2", "1\n2"},                         // Tabs are dedented
	{"\t  1\n\t  2", "1\n2"},                     // Tabs and spaces mixed
	{"  \t  1\n  \t  2", "1\n2"},                 // Tabs and spaces mixed
	{"  \t1\n  \t  2", "1\n  2"},                 // Tabs and spaces mixed
	{"  1  2\n  3  4", "1  2\n3  4"},             // Spaces are preserved
	{"  1\t2\n  3\t4", "1\t2\n3\t4"},             // Tabs are preserved
	{"1  2\n3  4", "1  2\n3  4"},                 // Spaces are preserved
	{"1\t2\n3\t4", "1\t2\n3\t4"},                 // Tabs are preserved
	{"  1\n\t2", "  1\n\t2"},                     // Tabs and spaces are not equivalent
	{"    1\n\t2", "    1\n\t2"},                 // Tabs and spaces are not equivalent
	{"        1\n\t2", "        1\n\t2"},         // Tabs and spaces are not equivalent
}

func TestDedent(t *testing.T) {
	for _, test := range dedentTestCases {
		if Dedent(test.text) != test.expected {
			t.Errorf("\nexpected %q\ngot %q", test.expected, Dedent(test.text))
		}
	}
}

func ExampleDedent() {
	s := `
		Lorem ipsum dolor sit amet,
		consectetur adipiscing elit.
		Curabitur justo tellus, facilisis nec efficitur dictum,
		fermentum vitae ligula. Sed eu convallis sapien.`
	fmt.Println(Dedent(s))
	fmt.Println("-------------")
	fmt.Println(s)
	// Output:
	// Lorem ipsum dolor sit amet,
	// consectetur adipiscing elit.
	// Curabitur justo tellus, facilisis nec efficitur dictum,
	// fermentum vitae ligula. Sed eu convallis sapien.
	// -------------
	//
	//		Lorem ipsum dolor sit amet,
	//		consectetur adipiscing elit.
	//		Curabitur justo tellus, facilisis nec efficitur dictum,
	//		fermentum vitae ligula. Sed eu convallis sapien.
}
