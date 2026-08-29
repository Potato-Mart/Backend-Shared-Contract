package pkg_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// allowedStructuralDuplicateGroups records the deliberately distinct contract
// concepts that currently share a Go struct layout. Every listed group must
// remain a duplicate, and every discovered duplicate must be listed here, so
// copy-and-paste models and stale exceptions fail this repository gate.
var allowedStructuralDuplicateGroups = []string{
	"common/geometry.Rotation3|common/geometry.Vector3",
	"pricing/coupon.CouponContent|supply/catalogue/classification.SupplierAvailablePromotion",
	"supply/catalogue/classification.BrandRef|supply/catalogue/classification.CategoryTagRef|supply/catalogue/classification.CollectionRef|supply/catalogue/classification.ObjectMediaRef|supply/catalogue/classification.ProductSupplierRef",
}

func TestProductionStructLayoutsHaveNoUnreviewedDuplicates(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	contractsRoot := filepath.Join(pkgRoot, "contracts")
	byLayout := make(map[string][]string)

	err := filepath.WalkDir(contractsRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		imports := structuralDuplicateImportPaths(file)
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if !ok || !typeSpecification.Name.IsExported() {
					continue
				}
				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				layout, layoutErr := structuralDuplicateLayout(fset, imports, structure)
				if layoutErr != nil {
					return fmt.Errorf("describe %s: %w", path, layoutErr)
				}
				byLayout[layout] = append(byLayout[layout], productionTypeKey(pkgRoot, path, typeSpecification.Name.Name))
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	actual := make([]string, 0)
	for _, typeKeys := range byLayout {
		if len(typeKeys) < 2 {
			continue
		}
		sort.Strings(typeKeys)
		actual = append(actual, strings.Join(typeKeys, "|"))
	}
	sort.Strings(actual)

	want := append([]string(nil), allowedStructuralDuplicateGroups...)
	sort.Strings(want)
	if strings.Join(actual, "\n") != strings.Join(want, "\n") {
		t.Fatalf("production struct duplicate groups changed:\n got: %s\nwant: %s", strings.Join(actual, "\n"), strings.Join(want, "\n"))
	}
}

func structuralDuplicateImportPaths(file *ast.File) map[string]string {
	imports := make(map[string]string, len(file.Imports))
	for _, imported := range file.Imports {
		path := strings.Trim(imported.Path.Value, `\"`)
		name := filepath.Base(path)
		if imported.Name != nil {
			name = imported.Name.Name
		}
		if name != "_" && name != "." {
			imports[name] = path
		}
	}
	return imports
}

func structuralDuplicateLayout(fset *token.FileSet, imports map[string]string, structure *ast.StructType) (string, error) {
	fields := make([]string, 0, len(structure.Fields.List))
	for _, field := range structure.Fields.List {
		typeText, err := structuralDuplicateTypeText(fset, imports, field.Type)
		if err != nil {
			return "", err
		}
		tag := ""
		if field.Tag != nil {
			tag = field.Tag.Value
		}
		if len(field.Names) == 0 {
			fields = append(fields, "embedded:"+typeText+":"+tag)
			continue
		}
		for _, name := range field.Names {
			fields = append(fields, name.Name+":"+typeText+":"+tag)
		}
	}
	return strings.Join(fields, ";"), nil
}

func structuralDuplicateTypeText(fset *token.FileSet, imports map[string]string, expression ast.Expr) (string, error) {
	var rendered bytes.Buffer
	if err := printer.Fprint(&rendered, fset, expression); err != nil {
		return "", err
	}
	typeText := rendered.String()
	for name, path := range imports {
		typeText = strings.ReplaceAll(typeText, name+".", path+".")
	}
	return typeText, nil
}
