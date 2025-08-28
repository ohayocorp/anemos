package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/spec"
)

// getSwaggerSpec parses the embedded Kubernetes OpenAPI spec into a swagger object.
func getSwaggerSpec() (*spec.Swagger, error) {
	swagger := &spec.Swagger{}
	if err := json.Unmarshal(openAPISpec1_34, swagger); err != nil {
		return nil, fmt.Errorf("failed to parse OpenAPI spec: %w", err)
	}

	fmt.Printf("Successfully parsed OpenAPI spec version: %s\n", swagger.Info.Version)
	fmt.Printf("Found %d definitions in the spec.\n", len(swagger.Definitions))

	return swagger, nil
}

func (typeInfo *typeInfo) schemaToType(schema *spec.Schema, isGo bool) (string, *typeInfo) {
	if len(schema.Type) > 1 {
		panic(fmt.Sprintf("Type %s has multiple types: %v", schema.ID, schema.Type))
	}

	if len(schema.Type) == 1 {
		schemaType := schema.Type[0]

		if literalType, ok := literalMappings[schemaType]; ok {
			return literalType, nil
		}

		if schemaType == "array" {
			itemsSchema := schema.Items.Schema
			itemType, fieldTypeInfo := typeInfo.schemaToType(itemsSchema, isGo)

			if isGo {
				return fmt.Sprintf("[]%s", itemType), fieldTypeInfo
			}

			return fmt.Sprintf("Array<%s>", itemType), fieldTypeInfo
		}

		if schemaType != "object" {
			panic(fmt.Sprintf("Unsupported schema type: %s for %s", schemaType, schema.ID))
		}
	}

	if schema.AdditionalProperties != nil {
		mapValueType, mapValueTypeInfo := typeInfo.schemaToType(schema.AdditionalProperties.Schema, isGo)

		if isGo {
			return fmt.Sprintf("map[string]%s", mapValueType), mapValueTypeInfo
		} else {
			return fmt.Sprintf("Record<string, %s>", mapValueType), mapValueTypeInfo
		}
	}

	ref := schema.Ref.GetURL()
	if ref == nil || ref.Fragment == "" {
		return "any", nil
	}

	packagePath, _, typeName := getPackagePathPackageAliasAndTypeName(ref.Fragment)
	identifier := fmt.Sprintf("%s/%s", packagePath, typeName)

	fieldTypeInfo, ok := typeMappings[identifier]
	if !ok {
		panic(fmt.Errorf("unknown schema reference: %s", identifier))
	}

	if fieldTypeInfo.IsExcluded {
		return "any", nil
	}

	if fieldTypeInfo.NativeType != nil {
		return fieldTypeInfo.NativeType.Name, fieldTypeInfo
	}

	if isGo {
		if fieldTypeInfo.PackagePath == typeInfo.PackagePath {
			typeName = fmt.Sprintf("*%s", fieldTypeInfo.Name)
		} else {
			// Different package, use alias. Packages will always be imported with their alias.
			typeName = fmt.Sprintf("*%s.%s", fieldTypeInfo.PackageAlias, typeName)
		}
	}

	return typeName, fieldTypeInfo
}
