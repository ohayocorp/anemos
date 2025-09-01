package main

import (
	"fmt"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/util"
)

// writeJsContents generates a TypeScript declaration file for a Kubernetes type.
func (typeInfo *typeInfo) writeJsContents(fields []string) error {
	fileName := fmt.Sprintf("%s.d.ts", typeInfo.Name)
	filePath := filepath.Join(jsOutputDir, typeInfo.PackagePath, fileName)

	file, err := NewGeneratedFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	specTypeName := typeInfo.getSpecTypeName()
	if specTypeName == "" {
		specTypeName = typeInfo.Name
	}

	if !typeInfo.IsDocument {
		fields = append(fields, util.Indent(util.MultilineString(`
			/**
			 * This declaration allows setting and getting custom properties on the document without TypeScript
			 * compiler errors.
			 */
			[customProperties: string]: any;
			`), 4))
	}

	template := util.MultilineString(`
		// Auto generated code; DO NOT EDIT.
		{{ .Imports }}
		{{ .Description }}
		export declare class {{ .TypeName }}{{ .Extends }} {
		    constructor();
		    constructor(spec: {{ .SpecTypeName }});

			{{ .Fields }}
		}
		`)

	jsImports := typeInfo.getJsImports()

	extends := ""
	if typeInfo.IsDocument {
		extends = " extends Document"
		jsImports = append(jsImports, "import {Document} from '@ohayocorp/anemos';")
	} else {

	}

	imports := strings.Join(jsImports, "\n")
	if imports != "" {
		imports += "\n"
	}

	contents := util.ParseTemplate(template, map[string]any{
		"Imports":      imports,
		"Description":  toJsComment(getDescription(typeInfo.Schema.Description)),
		"TypeName":     typeInfo.Name,
		"Extends":      extends,
		"SpecTypeName": specTypeName,
		"Fields":       strings.Join(fields, "\n\n\t"),
	})

	file.Write(contents)

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
	rootAliases := []string{}

	for _, identifier := range core.SortedKeys(typeMappings) {
		typeInfo := typeMappings[identifier]
		if typeInfo.IsExcluded {
			continue
		}

		if typeInfo.NativeType != nil {
			continue
		}

		if typeInfo.GenerateAliasOnRoot {
			rootAliases = append(rootAliases, fmt.Sprintf("export {%s} from './%s';\n", typeInfo.Name, typeInfo.PackagePath))
		}

		p := fmt.Sprintf("%s/%s", typeInfo.PackagePath, typeInfo.Name)
		p = path.Clean(p)

		isType[p] = true

		for {
			parent := filepath.Dir(p)
			self := filepath.Base(p)

			// Check if we have reached the root directory. In that case
			// add this path to the roots set.
			if parent == "." {
				rootsSet.Add(p)
				break
			}

			p = parent
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
	sort.Strings(rootAliases)

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

	// Export the root aliases.
	for _, alias := range rootAliases {
		indexFile.Write(alias)
	}

	return nil
}

// getFieldTypeJs determines the TypeScript type for a property.
func (typeInfo *typeInfo) getFieldTypeJs(propertyName string) string {
	propertySchema := typeInfo.Schema.Properties[propertyName]

	fieldType, _ := typeInfo.schemaToType(&propertySchema)

	return fieldType
}

// getJsImports generates TypeScript import statements for the .d.ts file.
// It determines the correct relative paths and handles imports from the same package.
func (typeInfo *typeInfo) getJsImports() []string {
	imports := map[string]mapset.Set[string]{}

	for propertyName := range typeInfo.Schema.Properties {
		propertySchema := typeInfo.Schema.Properties[propertyName]

		_, fieldTypeInfo := typeInfo.schemaToType(&propertySchema)
		if fieldTypeInfo == nil {
			continue
		}

		if fieldTypeInfo.IsExcluded {
			continue
		}

		if fieldTypeInfo.NativeType != nil && fieldTypeInfo.NativeType.IsPrimitive {
			continue
		}

		packagePath := fieldTypeInfo.PackagePath

		if fieldTypeInfo.PackagePath == typeInfo.PackagePath {
			if fieldTypeInfo.Name == typeInfo.Name {
				continue
			}

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

		modulePath, err := filepath.Rel(typeInfo.PackagePath, importPath)
		if err != nil {
			panic(fmt.Errorf("failed to get relative path for %s: %w", importPath, err))
		}

		modulePath = filepath.ToSlash(modulePath)
		if !strings.HasPrefix(modulePath, "@") {
			modulePath = "./" + modulePath
		}

		importsList = append(importsList, fmt.Sprintf(`import { %s } from "%s"`, strings.Join(types, ", "), modulePath))
	}

	return importsList
}

// getSpecTypeName returns a TypeScript type expression that picks fields other thanapiVersion and kind,
// which are managed automatically for resource types.
func (typeInfo *typeInfo) getSpecTypeName() string {
	fields := []string{}
	for propertyName := range typeInfo.Schema.Properties {
		if propertyName == "apiVersion" || propertyName == "kind" {
			continue
		}

		tc, ok := typeCustomizations[typeInfo.Identifier]
		if ok && slices.Contains(tc.ExcludedFields, propertyName) {
			continue
		}

		fields = append(fields, fmt.Sprintf(`"%s"`, propertyName))
	}

	if len(fields) == 0 {
		return "{}"
	}

	sort.Strings(fields)

	return fmt.Sprintf("Pick<%s, %s>", typeInfo.Name, strings.Join(fields, " | "))
}
