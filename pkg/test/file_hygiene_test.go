package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestProductionEnumFilesContainExactlyOneClosedStringEnum ensures a
// finite enum has a dedicated leaf package and source file. Open typed codes
// such as CountryCode intentionally remain model values and are not enums.
func TestProductionEnumFilesContainExactlyOneClosedStringEnum(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	var violations []string

	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") || filepath.Base(path) == "doc.go" {
			return nil
		}
		file, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		normalized := relativePkgPath(t, pkgRoot, path)
		stringTypes := modelBoundaryStringTypes(file)
		enumTypes := modelBoundaryClosedStringEnumTypes(file, stringTypes)
		isEnumFile := modelBoundaryIsLeafEnumFile(normalized)

		if isEnumFile {
			if len(stringTypes) != 1 || len(enumTypes) != 1 {
				violations = append(violations, normalized+": enum source must define exactly one closed string enum")
			}
			return nil
		}
		if len(enumTypes) > 0 {
			violations = append(violations, normalized+": closed enums must live in a leaf _enums package: "+strings.Join(sortedStringSet(enumTypes), ", "))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan enum file hygiene: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("enum file hygiene violations:\n%s", strings.Join(violations, "\n"))
	}
}

// TestProductionModelFilesContainOneExportedStruct keeps every production
// model in a dedicated source file. Tests may group helpers and fixtures, but
// production records must not accumulate unrelated exported shapes.
func TestProductionModelFilesContainOneExportedStruct(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	var violations []string

	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") || filepath.Base(path) == "doc.go" {
			return nil
		}
		normalized := relativePkgPath(t, pkgRoot, path)
		if modelBoundaryIsLeafEnumFile(normalized) {
			return nil
		}
		file, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		structNames := exportedStructNames(file)
		if len(structNames) <= 1 {
			return nil
		}
		violations = append(violations, normalized+": multiple exported structs in one production model file: "+strings.Join(structNames, ", "))
		return nil
	})
	if err != nil {
		t.Fatalf("scan model file hygiene: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("model file hygiene violations:\n%s", strings.Join(violations, "\n"))
	}
}

func exportedStructNames(file *ast.File) []string {
	var names []string
	for _, declaration := range file.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.TYPE {
			continue
		}
		for _, specification := range general.Specs {
			typeSpec, ok := specification.(*ast.TypeSpec)
			if !ok || !typeSpec.Name.IsExported() {
				continue
			}
			if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
				names = append(names, typeSpec.Name.Name)
			}
		}
	}
	sort.Strings(names)
	return names
}

func sortedStringSet(values map[string]struct{}) []string {
	keys := make([]string, 0, len(values))
	for value := range values {
		keys = append(keys, value)
	}
	sort.Strings(keys)
	return keys
}
