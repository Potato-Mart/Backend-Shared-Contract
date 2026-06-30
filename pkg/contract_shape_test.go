package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode"
)

var forbiddenExportedTypeSuffixes = []string{
	"Request",
	"Response",
	"Input",
	"Payload",
	"DTO",
	"Dto",
}

var allowedSharedWireDTOs = map[string]map[string]struct{}{
	"contracts/wholesale/application.go": {
		"WholesaleApplicationRequest":          {},
		"WholesaleApplicationResponse":         {},
		"WholesaleApplicationStatusResponse":   {},
		"WholesaleApplicationDecisionRequest":  {},
		"WholesaleApplicationDecisionResponse": {},
	},
}

func TestSharedContractDoesNotExportWireDTOs(t *testing.T) {
	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok || !isExported(ts.Name.Name) {
					continue
				}
				for _, suffix := range forbiddenExportedTypeSuffixes {
					if strings.HasSuffix(ts.Name.Name, suffix) {
						if isAllowedSharedWireDTO(path, ts.Name.Name) {
							continue
						}
						pos := fset.Position(ts.Pos())
						t.Fatalf("%s exports %s; request/response/input/payload DTOs belong in the owning backend", pos, ts.Name.Name)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan contract package types: %v", err)
	}
}

func TestSharedContractHasNoDatabaseSpecificTerms(t *testing.T) {
	terms := blockedStorageTerms()
	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		lower := strings.ToLower(string(body))
		for _, term := range terms {
			if strings.Contains(lower, term) {
				t.Fatalf("%s contains database-specific term %q", path, term)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan repository text: %v", err)
	}
}

func TestSharedContractStructTagsAreDatabaseNeutral(t *testing.T) {
	terms := blockedStorageTerms()
	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}
		ast.Inspect(file, func(n ast.Node) bool {
			st, ok := n.(*ast.StructType)
			if !ok || st.Fields == nil {
				return true
			}
			for _, field := range st.Fields.List {
				if field.Tag == nil {
					continue
				}
				tag := strings.ToLower(field.Tag.Value)
				for _, term := range terms {
					if strings.Contains(tag, term) {
						pos := fset.Position(field.Tag.Pos())
						t.Fatalf("%s uses database-specific struct tag %q", pos, term)
					}
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan struct tags: %v", err)
	}
}

func blockedStorageTerms() []string {
	return []string{
		string([]byte{98, 115, 111, 110}),
		string([]byte{109, 111, 110, 103, 111}),
		string([]byte{109, 111, 110, 103, 111, 100, 98}),
	}
}

func isAllowedSharedWireDTO(path, name string) bool {
	types := allowedSharedWireDTOs[filepath.ToSlash(path)]
	if types == nil {
		return false
	}
	_, ok := types[name]
	return ok
}

func isExported(name string) bool {
	for _, r := range name {
		return unicode.IsUpper(r)
	}
	return false
}
