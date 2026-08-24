package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/product/product_enums"
)

type canonicalProductField struct {
	json   string
	typeOf reflect.Type
}

func TestV30GlobalProductHasSKUCodeAndAuthoritativeStorage(t *testing.T) {
	assertExactFields(t, reflect.TypeOf(product.Product{}), map[string]canonicalProductField{
		"ID":             {json: "id", typeOf: reflect.TypeOf("")},
		"SKUCode":        {json: "sku_code", typeOf: reflect.TypeOf("")},
		"StorageType":    {json: "storage_type", typeOf: reflect.TypeOf(classification_enums.StorageType(""))},
		"Status":         {json: "status", typeOf: reflect.TypeOf(product_enums.ProductStatus(""))},
		"Content":        {json: "content", typeOf: reflect.TypeOf(product.ProductContent{})},
		"Classification": {json: "classification", typeOf: reflect.TypeOf(product.ProductClassification{})},
		"Supply":         {json: "supply,omitempty", typeOf: reflect.TypeOf((*classification.ProductSupply)(nil))},
		"Administration": {json: "administration,omitempty", typeOf: reflect.TypeOf((*product.ProductAdministration)(nil))},
	})
}

func TestV30SKUHasNoDatabaseOrProductIdentifier(t *testing.T) {
	assertExactFields(t, reflect.TypeOf(product.SKU{}), map[string]canonicalProductField{
		"SKUCode":            {json: "sku_code", typeOf: reflect.TypeOf("")},
		"PackageOptions":     {json: "package_options", typeOf: reflect.TypeOf([]product.ProductPackageOption{})},
		"BarcodeAssignments": {json: "barcode_assignments,omitempty", typeOf: reflect.TypeOf([]product.ProductBarcodeAssignment{})},
		"NetContent":         {json: "net_content,omitempty", typeOf: reflect.TypeOf((*measurement.NetContent)(nil))},
		"StorageType":        {json: "storage_type", typeOf: reflect.TypeOf(classification_enums.StorageType(""))},
		"Status":             {json: "status", typeOf: reflect.TypeOf(product_enums.SKUStatus(""))},
		"AuditFields":        {json: "", typeOf: reflect.TypeOf(audit.AuditFields{})},
	})
}

func TestV30ProductionModelsRejectLegacyCatalogueSymbols(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	legacyIdentifiers := map[string]struct{}{"SKUID": {}, "SKUIDs": {}, "ProductCategory": {}}
	legacyJSON := []string{"sku_id", "sku_ids", "product_category_code", "brand_ref", "package_option_id", "market_id", "price_book_id", "tax_category_id"}
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}
		ast.Inspect(file, func(node ast.Node) bool {
			switch value := node.(type) {
			case *ast.Ident:
				if _, legacy := legacyIdentifiers[value.Name]; legacy {
					t.Errorf("%s retains legacy identifier %s", fset.Position(value.Pos()), value.Name)
				}
			case *ast.Field:
				if value.Tag == nil {
					return true
				}
				for _, key := range legacyJSON {
					if strings.Contains(value.Tag.Value, `json:\"`+key) {
						t.Errorf("%s retains legacy JSON key %s", fset.Position(value.Tag.Pos()), key)
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

	for _, retired := range []string{
		"contracts/supply/classification/product_category.go",
		"contracts/supply/warehouse/warehouse_enums/warehouse_storage_type.go",
	} {
		if _, err := os.Stat(filepath.Join(pkgRoot, filepath.FromSlash(retired))); err == nil || !os.IsNotExist(err) {
			t.Errorf("retired source remains: %s", retired)
		}
	}
}

func assertExactFields(t *testing.T, model reflect.Type, expected map[string]canonicalProductField) {
	t.Helper()
	if model.NumField() != len(expected) {
		t.Errorf("%s has %d fields, want %d", model, model.NumField(), len(expected))
	}
	for name, want := range expected {
		field, ok := model.FieldByName(name)
		if !ok {
			t.Errorf("%s missing %s", model, name)
			continue
		}
		if field.Type != want.typeOf || field.Tag.Get("json") != want.json {
			t.Errorf("%s.%s = (%s,%q), want (%s,%q)", model, name, field.Type, field.Tag.Get("json"), want.typeOf, want.json)
		}
	}
}
