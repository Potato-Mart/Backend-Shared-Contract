package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

func TestV19ProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := map[string]struct{}{
		"MembershipOwnerRef":              {},
		"MembershipOwnerType":             {},
		"EligibleOwnerTypes":              {},
		"PackingSessionStatusSyncPending": {},
		"SortOrder":                       {},
	}
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if parseErr != nil {
			return parseErr
		}
		for _, group := range file.Comments {
			if strings.Contains(group.Text(), "Deprecated:") || strings.Contains(group.Text(), "@deprecated") {
				t.Errorf("%s contains a deprecated production declaration", path)
			}
		}
		ast.Inspect(file, func(node ast.Node) bool {
			if identifier, ok := node.(*ast.Ident); ok {
				if _, removed := removedIdentifiers[identifier.Name]; removed {
					t.Errorf("%s contains removed v19 identifier %s", path, identifier.Name)
				}
			}
			field, ok := node.(*ast.Field)
			if !ok {
				return true
			}
			if field.Tag != nil && strings.Contains(field.Tag.Value, "sort_order") {
				t.Errorf("%s contains removed serialized field sort_order", path)
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
