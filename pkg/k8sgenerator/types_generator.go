package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-openapi/spec"
	"github.com/ohayocorp/anemos/pkg/core"
)

const packageBase = "github.com/ohayocorp/anemos/pkg/k8s"

//go:generate go run .

var (
	//go:embed k8s-openapi-spec-1.34.json
	openAPISpec1_34 []byte
	typeMappings    map[string]*typeInfo

	// Output directories for generated Go and TypeScript files.
	outputDir   = filepath.Join("..", "k8s")
	jsOutputDir = filepath.Join("..", "jsdeclarations", "k8s")

	// Mapping of OpenAPI primitive types to Go primitive types.
	literalMappings = map[string]string{
		"string":  "string",
		"integer": "int",
		"number":  "float64",
		"boolean": "bool",
	}

	// Mapping of OpenAPI primitive types to TypeScript primitive types.
	literalMappingsJs = map[string]string{
		"string":  "string",
		"integer": "number",
		"number":  "number",
		"boolean": "boolean",
	}
)

// typeInfo holds metadata about a Kubernetes type being generated.
// It includes information about the type's schema, package location,
// and how it should be rendered in Go and TypeScript.
type typeInfo struct {
	Schema       *spec.Schema
	PackageName  string
	PackagePath  string
	PackageAlias string
	Name         string
	LiteralType  *string // Non-nil for types that map to primitives
	IsExcluded   bool
}

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
	fieldNames := []string{}
	fields := []string{}
	jsFields := []string{}
	jsRegistrationFields := []string{}

	for propertyName := range typeInfo.Schema.Properties {
		fieldNames = append(fieldNames, propertyName)
	}
	sort.Strings(fieldNames)

	for _, propertyName := range fieldNames {
		description := typeInfo.Schema.Properties[propertyName].Description
		// Skip system-populated fields since they are not user-modifiable.
		if strings.Contains(description, "Populated by the system") {
			continue
		}

		isOptional := !slices.Contains(typeInfo.Schema.Required, propertyName)
		optionalTag := ""
		if isOptional {
			optionalTag = "?"
		}

		fieldName, fieldType := typeInfo.getFieldNameAndType(propertyName, isOptional)
		jsFieldType := typeInfo.getFieldTypeJs(propertyName)
		jsonTag := typeInfo.getJsonTag(propertyName, isOptional)
		yamlTag := typeInfo.getYamlTag(propertyName, isOptional)
		comment := getDescription(description)

		goFieldBuilder := strings.Builder{}
		goFieldBuilder.WriteString(fmt.Sprintf("%s\n", comment))
		goFieldBuilder.WriteString(fmt.Sprintf("%s %s `%s %s`", fieldName, fieldType, jsonTag, yamlTag))

		jsFieldBuilder := strings.Builder{}
		jsFieldBuilder.WriteString(fmt.Sprintf("%s\n", toJsComment(comment)))
		jsFieldBuilder.WriteString(fmt.Sprintf("%s%s: %s", propertyName, optionalTag, jsFieldType))

		fields = append(fields, core.IndentTab(goFieldBuilder.String(), 1))
		jsFields = append(jsFields, core.Indent(jsFieldBuilder.String(), 4))
		jsRegistrationFields = append(jsRegistrationFields, fmt.Sprintf(`js.Field("%s"),`, fieldName))
	}

	err := typeInfo.writeGoContents(fields, jsRegistrationFields)
	if err != nil {
		return err
	}

	err = typeInfo.writeJsContents(jsFields)
	if err != nil {
		return err
	}

	return nil
}

// writeGoContents generates a Go source file for a Kubernetes type.
// This includes the struct definition, constructors, and JS registration code.
func (typeInfo *typeInfo) writeGoContents(fields []string, jsFields []string) error {
	fileName := fmt.Sprintf("%s.go", strings.ToLower(typeInfo.Name))
	filePath := filepath.Join(outputDir, typeInfo.PackagePath, fileName)

	file, err := NewGeneratedFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	template := core.MultilineString(`
		// Code generated by types_generator.go; DO NOT EDIT.

		package {{ .PackageName }}

		import (
			"reflect"

			"github.com/ohayocorp/anemos/pkg/js"

			{{ .Imports }}
		)

		{{ .Description }}
		type {{ .TypeName }} struct {
			{{ .Fields }}
		}

		func New{{ .TypeName }}() *{{ .TypeName }} {
			return &{{ .TypeName }}{}
		}

		func New{{ .TypeName }}WithSpec(spec *{{ .TypeName }}) *{{ .TypeName }} {
			{{ .TypeMetaSetter }}
			return spec
		}

		func Register{{ .TypeName }}(jsRuntime *js.JsRuntime) {
			jsRuntime.Type(reflect.TypeFor[{{ .TypeName }}]()).JsNamespace("k8s.{{ .JsNamespace }}").Fields(
				{{ .JsFields }}
			).Constructors(
				js.Constructor(reflect.ValueOf(New{{ .TypeName }})),
				js.Constructor(reflect.ValueOf(New{{ .TypeName }}WithSpec)),
			)
		}
		`)

	contents := core.ParseTemplate(template, map[string]any{
		"PackageName":    typeInfo.PackageName,
		"Imports":        core.IndentTab(strings.Join(typeInfo.getImports(), "\n"), 1),
		"Description":    getDescription(typeInfo.Schema.Description),
		"TypeName":       typeInfo.Name,
		"JsNamespace":    strings.ReplaceAll(typeInfo.PackagePath, "/", "."),
		"TypeMetaSetter": typeInfo.getTypeMetaSetter(),
		"Fields":         strings.Join(fields, "\n\n\t"),
		"JsFields":       strings.Join(jsFields, "\n\t\t"),
	})

	file.Write(contents)
	return nil
}

// writeJsContents generates a TypeScript declaration file for a Kubernetes type.
func (typeInfo *typeInfo) writeJsContents(fields []string) error {
	fileName := fmt.Sprintf("%s.d.ts", typeInfo.Name)
	filePath := filepath.Join(jsOutputDir, typeInfo.PackagePath, fileName)

	file, err := NewGeneratedFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	omittedTypeName := typeInfo.getOmittedTypeName()
	if omittedTypeName == "" {
		omittedTypeName = typeInfo.Name
	}

	template := core.MultilineString(`
		// Auto generated code; DO NOT EDIT.

		{{ .Imports }}

		{{ .Description }}
		export declare class {{ .TypeName }} {
		    constructor();
		    constructor(spec: {{ .OmittedTypeName }});

			{{ .Fields }}
		}
		`)

	contents := core.ParseTemplate(template, map[string]any{
		"Imports":         strings.Join(typeInfo.getJsImports(), "\n"),
		"Description":     toJsComment(getDescription(typeInfo.Schema.Description)),
		"TypeName":        typeInfo.Name,
		"OmittedTypeName": omittedTypeName,
		"Fields":          strings.Join(fields, "\n\n\t"),
	})

	file.Write(contents)

	return nil
}

// generateJsRegistrations creates a single Go file that registers all Kubernetes types
// with the JavaScript runtime.
func generateJsRegistrations() error {
	registerFile, err := NewGeneratedFile(filepath.Join(outputDir, "js_registrations.go"))
	if err != nil {
		return fmt.Errorf("failed to create js_registrations file: %w", err)
	}
	defer registerFile.Close()

	registrationFuncs := make([]string, 0, len(typeMappings))
	imports := mapset.NewSet[string]()

	for _, identifier := range core.SortedKeys(typeMappings) {
		typeInfo := typeMappings[identifier]
		if typeInfo.IsExcluded {
			continue
		}

		registration := fmt.Sprintf("%s.Register%s(jsRuntime)", typeInfo.PackageAlias, typeInfo.Name)
		importLine := fmt.Sprintf(`%s "%s"`, typeInfo.PackageAlias, path.Join(packageBase, typeInfo.PackagePath))

		registrationFuncs = append(registrationFuncs, registration)
		imports.Add(importLine)
	}

	importsSlice := imports.ToSlice()
	sort.Strings(importsSlice)

	jsRegistrations := core.ParseTemplate(`
		// Code generated by types_generator.go; DO NOT EDIT.

		package k8s

		import (
			"github.com/ohayocorp/anemos/pkg/js"

			{{ .Imports }}
		)

		func RegisterK8S(jsRuntime *js.JsRuntime) {
			{{ .RegistrationFuncs }}
		}
		`,
		map[string]any{
			"Imports":           core.IndentTab(strings.Join(importsSlice, "\n"), 1),
			"RegistrationFuncs": core.IndentTab(strings.Join(registrationFuncs, "\n"), 1),
		})

	registerFile.Write(jsRegistrations)

	return nil
}

// generateIndexTsFiles creates index.d.ts files for TypeScript modules.
// These files re-export types from individual definition files.
func generateIndexTsFiles() error {
	// Store the entries under each path. This includes every directory recursively because
	// each directory will need an index.d.ts file. E.g:
	//   pathParts["foo/bar/baz"] = {"type1"}
	//   pathParts["foo/bar"] = {"baz", "type2", "type3"}
	//   pathParts["foo/other"] = {"type4"}
	//   pathParts["foo"] = {"bar", "other"}
	//   pathParts["another"] = {"type5"}
	pathParts := map[string]mapset.Set[string]{}

	// Types themselves will be exported as-is such as: export * from './type'
	// Other non-type exports will be re-exported as namespaces such as: export * as foo from './foo'
	isType := map[string]bool{}

	// Root paths will be exported from the root index file, so keep track of them.
	rootsSet := mapset.NewSet[string]()

	for _, identifier := range core.SortedKeys(typeMappings) {
		typeInfo := typeMappings[identifier]
		if typeInfo.IsExcluded {
			continue
		}

		path := fmt.Sprintf("%s/%s", typeInfo.PackagePath, typeInfo.Name)
		isType[path] = true

		for {
			parent := filepath.Dir(path)
			self := filepath.Base(path)

			// Check if we have reached the root directory. In that case
			// add this path to the roots set.
			if parent == "." {
				rootsSet.Add(path)
				break
			}

			path = parent
			key := filepath.ToSlash(parent)

			set := pathParts[key]
			if set == nil {
				set = mapset.NewSet[string]()
				pathParts[key] = set
			}

			set.Add(self)
		}
	}

	// Generate index.d.ts files for each directory
	for _, path := range core.SortedKeys(pathParts) {
		parts := pathParts[path].ToSlice()
		sort.Strings(parts)

		indexFile, err := NewGeneratedFile(filepath.Join(jsOutputDir, path, "index.d.ts"))
		if err != nil {
			return fmt.Errorf("failed to create index.d.ts file: %w", err)
		}
		defer indexFile.Close()

		indexFile.Write("// Auto generated code; DO NOT EDIT.\n\n")

		for _, part := range parts {
			if isType[fmt.Sprintf("%s/%s", path, part)] {
				// Export types as-is.
				indexFile.Write(fmt.Sprintf("export * from './%s';\n", part))
			} else {
				// Export non-types as aliases.
				indexFile.Write(fmt.Sprintf("export * as %s from './%s';\n", part, part))
			}
		}
	}

	// Generate root index.d.ts file
	roots := rootsSet.ToSlice()
	sort.Strings(roots)

	indexFile, err := NewGeneratedFile(filepath.Join(jsOutputDir, "index.d.ts"))
	if err != nil {
		return fmt.Errorf("failed to create index.d.ts file: %w", err)
	}
	defer indexFile.Close()

	indexFile.Write("// Auto generated code; DO NOT EDIT.\n\n")

	// Export the root directories.
	for _, root := range roots {
		indexFile.Write(fmt.Sprintf("export * as %s from './%s';\n", root, root))
	}

	return nil
}

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

// initializeTypeInfo builds the global type mapping from OpenAPI schema definitions.
// This creates a comprehensive map of all Kubernetes types with their package paths,
// names, and schema information needed for code generation.
func initializeTypeInfo(swagger *spec.Swagger) {
	typeMappings = make(map[string]*typeInfo)

	// First pass: create type info objects for all definitions.
	for name, schema := range swagger.Definitions {
		packagePath, packageAlias, typeName := getPackagePathPackageAliasAndTypeName(name)
		packageName := path.Base(packagePath)

		typeInfo := &typeInfo{
			Schema:       &schema,
			PackageName:  packageName,
			PackagePath:  packagePath,
			PackageAlias: packageAlias,
			Name:         typeName,
		}

		typeInfo.IsExcluded = isExcludedType(typeInfo)
		typeMappings[fmt.Sprintf("%s/%s", typeInfo.PackagePath, typeInfo.Name)] = typeInfo
	}

	// Second pass: identify types that are just aliases for primitive types like strings or numbers.
	for identifier, info := range typeMappings {
		if len(info.Schema.Type) > 1 {
			panic(fmt.Sprintf("Type %s has multiple types: %v", identifier, info.Schema.Type))
		}

		if len(info.Schema.Type) == 0 {
			info.LiteralType = core.Pointer("any")
			continue
		}

		schemaType := info.Schema.Type[0]
		if schemaType == "object" {
			continue
		}

		if literalType, ok := literalMappings[schemaType]; ok {
			info.LiteralType = core.Pointer(literalType)
			continue
		}
	}
}

// getPackagePathPackageAliasAndTypeName extracts package path, alias, and type name
// from a Kubernetes definition identifier. This handles the conversion from
// the dot-separated notation in OpenAPI specs to filesystem paths.
func getPackagePathPackageAliasAndTypeName(identifier string) (string, string, string) {
	// Example ref format: "#/definitions/io.k8s.api.apps.v1.Deployment".
	identifier = strings.TrimPrefix(identifier, "/definitions/")
	identifier = strings.TrimPrefix(identifier, "io.k8s.api.")
	identifier = strings.ReplaceAll(identifier, "io.k8s.apimachinery.pkg.apis", "apimachinery")

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

	switch fmt.Sprintf("%s/%s", typeInfo.PackagePath, typeInfo.Name) {
	case "apimachinery/meta/v1/WatchEvent":
		return true
	}

	return false
}

// getFieldNameAndType determines the Go field name and type for a property.
// It handles proper casing for field names and pointer types for optional fields.
func (typeInfo *typeInfo) getFieldNameAndType(propertyName string, isOptional bool) (string, string) {
	propertySchema := typeInfo.Schema.Properties[propertyName]

	fieldName := fmt.Sprintf("%s%s", strings.ToUpper(string(propertyName[0])), propertyName[1:])
	fieldType, fieldTypeInfo := typeInfo.schemaToType(&propertySchema, true)

	// Make optional primitive fields pointers for proper omitempty behavior.
	if isOptional && fieldType != "any" && !strings.HasPrefix(fieldType, "map[") && (fieldTypeInfo == nil || fieldTypeInfo.LiteralType != nil) {
		fieldType = fmt.Sprintf("*%s", fieldType)
	}

	return fieldName, fieldType
}

// getFieldTypeJs determines the TypeScript type for a property.
func (typeInfo *typeInfo) getFieldTypeJs(propertyName string) string {
	propertySchema := typeInfo.Schema.Properties[propertyName]

	fieldType, fieldTypeInfo := typeInfo.schemaToType(&propertySchema, false)
	if fieldTypeInfo != nil && fieldTypeInfo.PackagePath != typeInfo.PackagePath {
		return fieldTypeInfo.Name
	}

	if fieldTypeInfo != nil && fieldTypeInfo.LiteralType != nil {
		return *fieldTypeInfo.LiteralType
	}

	return fieldType
}

// getImports generates a list of import statements needed for the Go file.
// It only includes imports for types referenced in the struct fields.
func (typeInfo *typeInfo) getImports() []string {
	imports := map[string]string{}

	for propertyName := range typeInfo.Schema.Properties {
		propertySchema := typeInfo.Schema.Properties[propertyName]

		_, fieldTypeInfo := typeInfo.schemaToType(&propertySchema, true)
		if fieldTypeInfo == nil {
			continue
		}

		if fieldTypeInfo.IsExcluded {
			continue
		}

		if fieldTypeInfo.PackagePath == typeInfo.PackagePath {
			continue
		}

		imports[fieldTypeInfo.PackagePath] = fieldTypeInfo.PackageAlias
	}

	importPaths := core.SortedKeys(imports)
	importsList := make([]string, 0, len(importPaths))

	for _, path := range importPaths {
		alias := imports[path]
		importsList = append(importsList, fmt.Sprintf(`%s "%s/%s"`, alias, packageBase, path))
	}

	return importsList
}

// getJsImports generates TypeScript import statements for the .d.ts file.
// It determines the correct relative paths and handles imports from the same package.
func (typeInfo *typeInfo) getJsImports() []string {
	imports := map[string]mapset.Set[string]{}

	for propertyName := range typeInfo.Schema.Properties {
		propertySchema := typeInfo.Schema.Properties[propertyName]

		_, fieldTypeInfo := typeInfo.schemaToType(&propertySchema, true)
		if fieldTypeInfo == nil {
			continue
		}

		if fieldTypeInfo.IsExcluded {
			continue
		}

		packagePath := fieldTypeInfo.PackagePath

		if fieldTypeInfo.PackagePath == typeInfo.PackagePath {
			packagePath = fmt.Sprintf("%s/%s", typeInfo.PackagePath, fieldTypeInfo.Name)
		}

		set := imports[packagePath]
		if set == nil {
			set = mapset.NewSet[string]()
			imports[packagePath] = set
		}

		set.Add(fieldTypeInfo.Name)
	}

	importPaths := core.SortedKeys(imports)
	importsList := make([]string, 0, len(importPaths))

	for _, importPath := range importPaths {
		types := imports[importPath].ToSlice()
		sort.Strings(types)

		relativePath, err := filepath.Rel(typeInfo.PackagePath, importPath)
		if err != nil {
			panic(fmt.Errorf("failed to get relative path for %s: %w", importPath, err))
		}

		relativePath = filepath.ToSlash(relativePath)
		if !strings.HasPrefix(relativePath, ".") {
			relativePath = "./" + relativePath
		}

		importsList = append(importsList, fmt.Sprintf(`import { %s } from "%s"`, strings.Join(types, ", "), relativePath))
	}

	return importsList
}

// getTypeMetaSetter generates code to set apiVersion and kind fields for types
// that have x-kubernetes-group-version-kind extensions in their schema.
func (typeInfo *typeInfo) getTypeMetaSetter() string {
	version, kind := typeInfo.getVersionKind()
	if version == nil || kind == nil {
		return ""
	}

	setter := core.ParseTemplate(`
		version := "{{ .Version }}"
		kind := "{{ .Kind }}"

		spec.ApiVersion = &version
		spec.Kind = &kind
		`,
		map[string]any{
			"Version": version,
			"Kind":    kind,
		})

	return core.IndentTab(setter, 1)
}

// getOmittedTypeName returns a TypeScript type expression that omits apiVersion and kind
// fields, which are managed automatically for resource types.
func (typeInfo *typeInfo) getOmittedTypeName() string {
	version, kind := typeInfo.getVersionKind()
	if version == nil || kind == nil {
		return ""
	}

	return fmt.Sprintf(`Omit<%s, "apiVersion" | "kind">`, typeInfo.Name)
}

// getVersionKind extracts the API version and kind from a type's OpenAPI schema extensions.
func (typeInfo *typeInfo) getVersionKind() (*string, *string) {
	gvkExtension := typeInfo.Schema.Extensions["x-kubernetes-group-version-kind"]
	if gvkExtension == nil {
		return nil, nil
	}

	var version, kind *string

	if gvk, ok := gvkExtension.(map[string]any); ok {
		version = core.Pointer(gvk["version"].(string))
		kind = core.Pointer(gvk["kind"].(string))
	} else if gvks, ok := gvkExtension.([]any); ok && len(gvks) > 0 {
		gvk := gvks[0].(map[string]any)

		version = core.Pointer(gvk["version"].(string))
		kind = core.Pointer(gvk["kind"].(string))
	}

	if version == nil || kind == nil {
		panic(fmt.Errorf("invalid x-kubernetes-group-version-kind extension format"))
	}

	return version, kind
}

// schemaToType converts an OpenAPI schema to a Go or TypeScript type.
// This handles primitive types, arrays, and references to other types.
func (typeInfo *typeInfo) schemaToType(schema *spec.Schema, isGo bool) (string, *typeInfo) {
	if len(schema.Type) > 1 {
		panic(fmt.Sprintf("Type %s has multiple types: %v", schema.ID, schema.Type))
	}

	if len(schema.Type) == 1 {
		schemaType := schema.Type[0]

		literals := literalMappings
		if !isGo {
			literals = literalMappingsJs
		}

		if literalType, ok := literals[schemaType]; ok {
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

	if fieldTypeInfo.LiteralType != nil {
		return *fieldTypeInfo.LiteralType, nil
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
	builder.WriteString("/**")

	lines := strings.Split(comment, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		line = strings.TrimPrefix(line, "// ")
		// Prevent closing comment in the middle of a line:
		// https://github.com/jsdoc/jsdoc/issues/821#issuecomment-385324492
		line = strings.ReplaceAll(line, "*/", "*&#8205;/")

		builder.WriteString("\n * ")
		builder.WriteString(line)
		builder.WriteString("\n * ")
	}

	builder.WriteString("\n */")

	return builder.String()
}

// getJsonTag generates a JSON struct tag for a field.
// It adds the omitempty option for optional fields.
func (typeInfo *typeInfo) getJsonTag(propertyName string, isOptional bool) string {
	jsonTag := propertyName
	if isOptional {
		jsonTag += ",omitempty"
	}

	return fmt.Sprintf(`json:"%s"`, jsonTag)
}

// getYamlTag generates a YAML struct tag for a field.
// It adds the omitempty option for optional fields.
func (typeInfo *typeInfo) getYamlTag(propertyName string, isOptional bool) string {
	yamlTag := propertyName
	if isOptional {
		yamlTag += ",omitempty"
	}

	return fmt.Sprintf(`yaml:"%s"`, yamlTag)
}
