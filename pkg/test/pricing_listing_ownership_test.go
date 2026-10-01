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

// Commercial listing records, enums and their event each have one definition.
// Supply stock evidence remains independently owned; no aliases or duplicate
// compatibility representations may restore the removed listing surface.
func TestPricingListingHasOneCanonicalOwner(t *testing.T) {
	want := map[string]string{
		"MarketListing":              "contracts/pricing/listing",
		"SaleRestriction":            "contracts/pricing/listing",
		"MarketListingStatus":        "contracts/pricing/listing/listing_enums",
		"SaleRestrictionKind":        "contracts/pricing/listing/listing_enums",
		"CatalogListingChangedEvent": "contracts/pubsub/pricing",
		"SaleEligibilitySnapshot":    "contracts/supply/catalogue/listing",
		"DamageSaleApproval":         "contracts/supply/catalogue/listing",
	}
	counts := make(map[string]int)
	root := sharedContractPkgRoot(t)
	err := filepath.WalkDir(filepath.Join(root, "contracts"), func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			return err
		}
		owner := filepath.ToSlash(filepath.Dir(filepath.FromSlash(relativePkgPath(t, root, path))))
		for _, decl := range file.Decls {
			general, ok := decl.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, spec := range general.Specs {
				model := spec.(*ast.TypeSpec)
				if expected, tracked := want[model.Name.Name]; tracked {
					counts[model.Name.Name]++
					if owner != expected {
						t.Errorf("%s belongs to %s, found %s", model.Name.Name, expected, owner)
					}
					if model.Name.Name == "MarketListing" || model.Name.Name == "CatalogListingChangedEvent" {
						if record, ok := model.Type.(*ast.StructType); ok {
							for _, field := range record.Fields.List {
								for _, name := range field.Names {
									if name.Name == "DisplayName" || name.Name == "UnitPricingRequired" {
										t.Errorf("%s restores removed listing field %s", model.Name.Name, name.Name)
									}
								}
								if field.Tag != nil && (strings.Contains(field.Tag.Value, "display_name") || strings.Contains(field.Tag.Value, "unit_pricing_required")) {
									t.Errorf("%s restores removed listing JSON key %s", model.Name.Name, field.Tag.Value)
								}
							}
						}
					}
				}
			}
		}
		for _, imported := range file.Imports {
			if strings.Contains(imported.Path.Value, "/supply/catalogue/listing/listing_enums") {
				t.Errorf("%s imports removed Supply listing enums", path)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for name := range want {
		if counts[name] != 1 {
			t.Errorf("%s has %d definitions, want exactly one", name, counts[name])
		}
	}
}
