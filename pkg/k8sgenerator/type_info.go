package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/go-openapi/spec"
)

// typeInfo holds metadata about a Kubernetes type being generated.
// It includes information about the type's schema, package location,
// and how it should be rendered in Go and TypeScript.
type typeInfo struct {
	Schema              *spec.Schema
	Identifier          string
	PackageName         string
	PackagePath         string
	PackageAlias        string
	Name                string
	NativeType          *nativeTypeInfo // Non-nil for types that map to primitives or types that map to native Go types
	IsExcluded          bool
	GenerateAliasOnRoot bool
	IsDocument          bool
	IsWorkload          bool
}

type nativeTypeInfo struct {
	Name        string
	IsPrimitive bool
}

// initializeTypeInfo builds the global type mapping from OpenAPI schema definitions.
// This creates a comprehensive map of all Kubernetes types with their package paths,
// names, and schema information needed for code generation.
func initializeTypeInfo(swagger *spec.Swagger) {
	typeMappings = make(map[string]*typeInfo)

	// First pass: create type info objects for all definitions.
	for name, schema := range swagger.Definitions {
		packagePath, packageAlias, typeName := getPackagePathPackageAliasAndTypeName(name)
		packageName := path.Base(packagePath)
		identifier := fmt.Sprintf("%s/%s", packagePath, typeName)

		typeInfo := &typeInfo{
			Schema:       &schema,
			Identifier:   identifier,
			PackageName:  packageName,
			PackagePath:  packagePath,
			PackageAlias: packageAlias,
			Name:         typeName,
		}

		typeInfo.IsExcluded = isExcludedType(typeInfo)
		typeMappings[identifier] = typeInfo
	}

	// Second pass: identify types that are just aliases for primitive types like strings or numbers.
	for identifier, info := range typeMappings {
		if len(info.Schema.Type) > 1 {
			panic(fmt.Sprintf("Type %s has multiple types: %v", identifier, info.Schema.Type))
		}

		if len(info.Schema.Type) == 0 {
			info.NativeType = &nativeTypeInfo{
				Name:        "any",
				IsPrimitive: true,
			}

			continue
		}

		schemaType := info.Schema.Type[0]
		if schemaType == "object" {
			continue
		}

		if literalType, ok := literalMappings[schemaType]; ok {
			info.NativeType = &nativeTypeInfo{
				Name:        literalType,
				IsPrimitive: true,
			}

			continue
		}
	}

	// Third pass: customize types based on predefined rules.
	for identifier, info := range typeMappings {
		customization := typeCustomizations[identifier]
		if customization.Path != nil {
			info.PackagePath = *customization.Path
		}

		if customization.NativeType != nil {
			info.NativeType = customization.NativeType
			info.Name = customization.NativeType.Name
		}

		info.GenerateAliasOnRoot = customization.GenerateAliasOnRoot
		info.IsDocument = customization.IsDocument
		info.IsWorkload = customization.IsWorkload
	}
}

// getPackagePathPackageAliasAndTypeName extracts package path, alias, and type name
// from a Kubernetes definition identifier. This handles the conversion from
// the dot-separated notation in OpenAPI specs to filesystem paths.
func getPackagePathPackageAliasAndTypeName(identifier string) (string, string, string) {
	// Example ref format: "#/definitions/io.k8s.api.apps.v1.Deployment".
	identifier = strings.TrimPrefix(identifier, "/definitions/")
	identifier = strings.TrimPrefix(identifier, "io.k8s.api.")
	identifier = strings.TrimPrefix(identifier, "io.k8s.apiextensions-apiserver.pkg.apis.")
	identifier = strings.TrimPrefix(identifier, "io.k8s.kube-aggregator.pkg.apis.")
	identifier = strings.ReplaceAll(identifier, "io.k8s.apimachinery.pkg.apis", "apimachinery")
	identifier = strings.ReplaceAll(identifier, "io.k8s.apimachinery.pkg", "apimachinery")

	tokens := strings.Split(identifier, ".")
	// Last segment after dot is the type name.
	typeName := tokens[len(tokens)-1]
	// Remaining segments are the package path.
	packageTokens := tokens[:len(tokens)-1]

	packageAlias := strings.Join(packageTokens, "")
	path := strings.Join(packageTokens, "/")

	return path, packageAlias, typeName
}

// isExcludedType determines if a type should be excluded from generation.
// Some types are internal implementation details or not useful for the API.
func isExcludedType(typeInfo *typeInfo) bool {
	// Some types are inside the packages we are interested in, but has no value for our use case.
	if strings.HasPrefix(typeInfo.PackagePath, "io/k8s") {
		return true
	}

	if strings.HasPrefix(typeInfo.PackagePath, "apiserverinternal") {
		return true
	}

	identifier := fmt.Sprintf("%s/%s", typeInfo.PackagePath, typeInfo.Name)

	return excludedTypes.Contains(identifier)
}
