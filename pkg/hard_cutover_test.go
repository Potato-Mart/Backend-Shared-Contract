package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestV21ProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := map[string]struct{}{
		"MembershipOwnerRef":              {},
		"MembershipOwnerType":             {},
		"EligibleOwnerTypes":              {},
		"PackingSessionStatusSyncPending": {},
		"SortOrder":                       {},
		"BrandKey":                        {},
		"BrandSummary":                    {},
		"ActiveProductCount":              {},
		"WholesaleProductCount":           {},
		"NotificationQuietHours":          {},
		"QuietHours":                      {},
		"FCMDestination":                  {},
		"FCMDestinations":                 {},
		"PushDestination":                 {},
		"PushDestinations":                {},
		"CouponPreview":                   {},
		"CouponPreviews":                  {},
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
					t.Errorf("%s contains removed v21 identifier %s", path, identifier.Name)
				}
			}
			field, ok := node.(*ast.Field)
			if !ok {
				return true
			}
			if field.Tag != nil {
				for _, removedTag := range []string{
					"sort_order", "quiet_hours", "fcm_destination", "fcm_destinations",
					"fcm_token", "coupon_preview", "coupon_previews",
				} {
					if strings.Contains(field.Tag.Value, removedTag) {
						t.Errorf("%s contains removed or service-local serialized field %s", path, removedTag)
					}
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestV21GoSourcesContainNoOlderContractImports(t *testing.T) {
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if parseErr != nil {
			return parseErr
		}
		for _, spec := range file.Imports {
			importPath, unquoteErr := strconv.Unquote(spec.Path.Value)
			if unquoteErr != nil {
				return unquoteErr
			}
			for _, oldMajor := range []string{"/v19/", "/v20/"} {
				if strings.Contains(importPath, "Backend-Shared-Contract"+oldMajor) {
					t.Errorf("%s imports older shared-contract major %s", path, importPath)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
