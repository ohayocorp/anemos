package js

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var genericTypesRegex = regexp.MustCompile(`(.+)\[(.+)\]`)
var funcRegex = regexp.MustCompile(`func\((.*)\)`)

func typeToJsTypeName(t reflect.Type) string {
	return typeNameToJsTypeName(t.String())
}

func typeNameToJsTypeName(name string) string {
	genericMatches := genericTypesRegex.FindStringSubmatch(name)

	if genericMatches == nil {
		funcMatches := funcRegex.FindStringSubmatch(name)

		if funcMatches != nil {
			args := strings.Split(funcMatches[1], ", ")
			for i, arg := range args {
				args[i] = typeNameToJsTypeName(arg)
			}

			return fmt.Sprintf("func(%s)", strings.Join(args, ", "))
		}

		if strings.HasPrefix(name, "[]") {
			return fmt.Sprintf("%s[]", typeNameToJsTypeName(name[2:]))
		}

		tokens := strings.Split(name, ".")
		return tokens[len(tokens)-1]
	}

	if len(genericMatches) != 3 {
		panic(fmt.Errorf("invalid generic type: %s", name))
	}

	typeName := typeNameToJsTypeName(genericMatches[1])

	genericTypes := strings.Split(genericMatches[2], ",")
	for i, genericType := range genericTypes {
		genericTypes[i] = typeNameToJsTypeName(genericType)
	}

	return fmt.Sprintf("%s<%s>", typeName, strings.Join(genericTypes, ", "))
}

func toCamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToLower(s[0:1]) + s[1:]
}
