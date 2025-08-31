package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/util"
)

const packageBase = "github.com/ohayocorp/anemos/pkg/k8s"

//go:generate go run .

var (
	//go:embed k8s-openapi-spec-1.34.json
	openAPISpec1_34 []byte
	typeMappings    map[string]*typeInfo

	// Output directories for generated Go and TypeScript files.
	outputDir   = filepath.Join("..", "k8s")
	jsOutputDir = filepath.Join("..", "jslib", "native", "k8s")

	// Mapping of OpenAPI primitive types to TypeScript primitive types.
	literalMappings = map[string]string{
		"string":  "string",
		"integer": "number",
		"number":  "number",
		"boolean": "boolean",
	}
)

func main() {
	if err := generateTypes(); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating types: %v\n", err)
		os.Exit(1)
	}
}

// generateTypes orchestrates the entire type generation process:
// 1. Initializes output directories.
// 2. Parses Kubernetes OpenAPI spec.
// 3. Builds type information for all definitions.
// 4. Generates Go and TypeScript files for each type.
// 5. Creates registration code for JS interop.
func generateTypes() error {
	err := initializeOutputDirs()
	if err != nil {
		return err
	}

	swagger, err := getSwaggerSpec()
	if err != nil {
		return err
	}

	initializeTypeInfo(swagger)

	for _, identifier := range core.SortedKeys(typeMappings) {
		typeInfo := typeMappings[identifier]
		if typeInfo.IsExcluded {
			fmt.Printf("Skipping excluded type: %s\n", identifier)
			continue
		}

		if typeInfo.NativeType != nil {
			fmt.Printf("Skipping native type: %s\n", identifier)
			continue
		}

		err = generateType(typeInfo)
		if err != nil {
			return fmt.Errorf("failed to generate struct for %s: %w", typeInfo.Name, err)
		}

		fmt.Printf("Successfully generated Go types for %s\n", identifier)
	}

	if err := generateJsRegistrations(); err != nil {
		return fmt.Errorf("failed to generate js registrations: %w", err)
	}

	if err := generateIndexTsFiles(); err != nil {
		return fmt.Errorf("failed to generate index.d.ts files: %w", err)
	}

	return nil
}

// generateType creates Go and TypeScript representations for a Kubernetes type.
// It extracts field information from the OpenAPI schema, applying proper naming
// conventions, type mappings, and tags for serialization.
func generateType(typeInfo *typeInfo) error {
	propertyNames := []string{}
	jsFields := []string{}

	for propertyName := range typeInfo.Schema.Properties {
		propertyNames = append(propertyNames, propertyName)
	}
	sort.Strings(propertyNames)

	customization := typeCustomizations[typeInfo.Identifier]

	for _, propertyName := range propertyNames {
		description := typeInfo.Schema.Properties[propertyName].Description

		if slices.Contains(customization.ExcludedFields, propertyName) {
			continue
		}

		isOptional := !slices.Contains(typeInfo.Schema.Required, propertyName)
		optionalTag := ""
		if isOptional {
			optionalTag = "?"
		}

		jsFieldType := typeInfo.getFieldTypeJs(propertyName)
		comment := getDescription(description)
		jsFieldName := propertyName
		if strings.Contains(jsFieldName, "-") {
			jsFieldName = fmt.Sprintf(`"%s"`, jsFieldName)
		}

		jsFieldBuilder := strings.Builder{}
		jsFieldBuilder.WriteString(fmt.Sprintf("%s\n", toJsComment(comment)))
		jsFieldBuilder.WriteString(fmt.Sprintf("%s%s: %s", jsFieldName, optionalTag, jsFieldType))

		jsFields = append(jsFields, util.Indent(jsFieldBuilder.String(), 4))
	}

	err := typeInfo.writeGoContents()
	if err != nil {
		return err
	}

	err = typeInfo.writeJsContents(jsFields)
	if err != nil {
		return err
	}

	return nil
}

// initializeOutputDirs prepares the output directories by removing existing content
// and creating fresh directories for the generated code.
func initializeOutputDirs() error {
	initialize := func(directory string) error {
		// Remove old output directory if it exists.
		err := os.RemoveAll(directory)
		if err != nil {
			return fmt.Errorf("failed to remove old output directory: %w", err)
		}

		// Create the output directory.
		if err := os.MkdirAll(directory, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		return nil
	}

	err := initialize(outputDir)
	if err != nil {
		return err
	}

	return initialize(jsOutputDir)
}
