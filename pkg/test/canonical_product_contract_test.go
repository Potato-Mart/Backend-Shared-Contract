package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"
)

func TestV26CanonicalProductUsesOnlyReusableComponents(t *testing.T) {
	canonicalProductAssertExactFields(t, reflect.TypeOf(product.Product{}), map[string]canonicalProductField{
		"SKUCode": {
			json:   "sku_code",
			typeOf: reflect.TypeOf(""),
		},
		"Content": {
			json:   "content",
			typeOf: reflect.TypeOf(product.ProductContent{}),
		},
		"Classification": {
			json:   "classification",
			typeOf: reflect.TypeOf(product.ProductClassification{}),
		},
		"Packaging": {
			json:   "packaging",
			typeOf: reflect.TypeOf(product.ProductPackaging{}),
		},
		"Commerce": {
			json:   "commerce",
			typeOf: reflect.TypeOf(product.ProductCommerce{}),
		},
		"Metrics": {
			json:   "metrics",
			typeOf: reflect.TypeOf(product.ProductMetrics{}),
		},
		"Supply": {
			json:   "supply,omitempty",
			typeOf: reflect.TypeOf((*classification.ProductSupply)(nil)),
		},
		"Administration": {
			json:   "administration,omitempty",
			typeOf: reflect.TypeOf((*product.ProductAdministration)(nil)),
		},
	})

	content, ok := reflect.TypeOf(product.ProductContent{}).FieldByName("Images")
	if !ok {
		t.Fatal("product.ProductContent must expose Images")
	}
	if want := reflect.TypeOf((*product.Images)(nil)); content.Type != want {
		t.Errorf("product.ProductContent.Images type = %s, want %s", content.Type, want)
	}
	if got, want := content.Tag.Get("json"), "images,omitempty"; got != want {
		t.Errorf("product.ProductContent.Images JSON tag = %q, want %q", got, want)
	}
}

func TestV26ProductImagesHaveOnlyObjectMediaCoverGalleryAndDetails(t *testing.T) {
	canonicalProductAssertExactFields(t, reflect.TypeOf(product.Images{}), map[string]canonicalProductField{
		"Cover": {
			json:   "cover,omitempty",
			typeOf: reflect.TypeOf((*security.ObjectMedia)(nil)),
		},
		"Gallery": {
			json:   "gallery,omitempty",
			typeOf: reflect.TypeOf([]security.ObjectMedia{}),
		},
		"Details": {
			json:   "details,omitempty",
			typeOf: reflect.TypeOf([]security.ObjectMedia{}),
		},
	})
}

func TestV26ProductAdministrationIsOptionalAndOwnsHistoryAndAudit(t *testing.T) {
	administrationField, ok := reflect.TypeOf(product.Product{}).FieldByName("Administration")
	if !ok {
		t.Fatal("product.Product must expose optional Administration")
	}
	if want := reflect.TypeOf((*product.ProductAdministration)(nil)); administrationField.Type != want {
		t.Errorf("product.Product.Administration type = %s, want %s", administrationField.Type, want)
	}
	if got, want := administrationField.Tag.Get("json"), "administration,omitempty"; got != want {
		t.Errorf("product.Product.Administration JSON tag = %q, want %q", got, want)
	}

	administration := reflect.TypeOf(product.ProductAdministration{})
	canonicalProductAssertExactFields(t, administration, map[string]canonicalProductField{
		"History": {
			json:   "history,omitempty",
			typeOf: reflect.TypeOf([]security.HistoryEntry{}),
		},
		"AuditFields": {
			json:   "",
			typeOf: reflect.TypeOf(audit.AuditFields{}),
		},
	})
	auditField, ok := administration.FieldByName("AuditFields")
	if !ok || !auditField.Anonymous {
		t.Error("product.ProductAdministration must embed audit.AuditFields")
	}
}

func TestV26ProductCommerceExcludesOperationalInventoryAndAuditEvidence(t *testing.T) {
	canonicalProductRequireJSONFields(t, reflect.TypeOf(product.ProductCommerce{}), map[string]string{
		"Status":        "status,omitempty",
		"Selling":       "selling,omitempty",
		"FirstListedAt": "first_listed_at,omitempty",
		"Packages":      "packages,omitempty",
	})
	canonicalProductRequireJSONFields(t, reflect.TypeOf(product.ProductPackageCommerce{}), map[string]string{
		"PackageOptionID": "package_option_id",
		"PackagePrice":    "package_price",
		"TaxAmount":       "tax_amount",
		"StockState":      "stock_state,omitempty",
		"AsOf":            "as_of",
	})

	for _, model := range []reflect.Type{
		reflect.TypeOf(product.ProductCommerce{}),
		reflect.TypeOf(product.ProductPackageCommerce{}),
	} {
		canonicalProductAssertNoOperationalCommerceEvidence(t, model)
	}
}

func TestV26CanonicalProductRetiresLegacyCatalogueProjections(t *testing.T) {
	retiredProductTypes := canonicalProductStringSet(
		"Snapshot",
		"StorefrontProduct",
		"StorefrontCommercial",
		"DetailImage",
		"ProductPackageOptionSnapshot",
		"ProductBarcodeAssignmentSnapshot",
		"StorefrontMerchandising",
		"StorefrontDisplay",
		"StorefrontPromotionBadge",
		"StorefrontOrigin",
		"PreorderPolicy",
		"SoonExpiryMerchandisingPolicy",
		"StorefrontExpiryDisplay",
		"StorefrontPreorderDisplay",
	)
	retiredPOSETypes := canonicalProductStringSet("CatalogProduct")

	canonicalProductWalkProductionGoFiles(t, func(path string, relativePath string, fset *token.FileSet, file *ast.File) {
		directory := filepath.ToSlash(filepath.Dir(relativePath))
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if directory == "contracts/supply/product" {
					if _, retired := retiredProductTypes[typeSpecification.Name.Name]; retired {
						t.Errorf("%s retains retired canonical-product type %s", fset.Position(typeSpecification.Pos()), typeSpecification.Name.Name)
					}
				}
				if directory == "contracts/orders/pos" {
					if _, retired := retiredPOSETypes[typeSpecification.Name.Name]; retired {
						t.Errorf("%s retains retired POS catalogue type %s", fset.Position(typeSpecification.Pos()), typeSpecification.Name.Name)
					}
				}
			}
		}
	})

	for _, retiredPath := range []string{
		"contracts/supply/product/detail_image.go",
		"contracts/supply/product/merchandising.go",
		"contracts/supply/product/snapshot.go",
		"contracts/supply/product/storefront.go",
	} {
		path := filepath.Join(sharedContractPkgRoot(t), filepath.FromSlash(retiredPath))
		if _, err := os.Stat(path); err == nil {
			t.Errorf("retired canonical-product source file remains: %s", retiredPath)
		} else if !os.IsNotExist(err) {
			t.Errorf("inspect retired canonical-product path %s: %v", retiredPath, err)
		}
	}
}

func TestV26OnlyProductPackageEmbedsFullProduct(t *testing.T) {
	const productImportPath = "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product"

	canonicalProductWalkProductionGoFiles(t, func(path string, relativePath string, fset *token.FileSet, file *ast.File) {
		directory := filepath.ToSlash(filepath.Dir(relativePath))
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if !ok || typeSpecification.Name.Name != "Product" {
					continue
				}
				if directory != "contracts/supply/product" {
					t.Errorf("%s declares Product outside supply/product; product.Product is the only full product model", fset.Position(typeSpecification.Pos()))
				}
			}
		}

		if directory == "contracts/supply/product" {
			return
		}

		aliases := make(map[string]struct{})
		dotImported := false
		for _, imported := range file.Imports {
			importPath, err := strconv.Unquote(imported.Path.Value)
			if err != nil || importPath != productImportPath {
				continue
			}
			if imported.Name == nil {
				aliases["product"] = struct{}{}
				continue
			}
			switch imported.Name.Name {
			case "_":
			case ".":
				dotImported = true
			default:
				aliases[imported.Name.Name] = struct{}{}
			}
		}
		if dotImported {
			t.Errorf("%s dot-imports supply/product; cross-domain models must use ProductSKUCode", path)
			return
		}
		if len(aliases) == 0 {
			return
		}

		ast.Inspect(file, func(node ast.Node) bool {
			field, ok := node.(*ast.Field)
			if !ok {
				return true
			}
			if canonicalProductTypeUsesFullProduct(field.Type, aliases) {
				t.Errorf("%s embeds product.Product outside supply/product; use ProductSKUCode plus transaction-owned frozen facts", fset.Position(field.Pos()))
			}
			return true
		})
	})
}

func TestV26CrossDomainProductSKUCodeLinksAreScalars(t *testing.T) {
	canonicalProductWalkProductionGoFiles(t, func(path string, relativePath string, fset *token.FileSet, file *ast.File) {
		if filepath.ToSlash(filepath.Dir(relativePath)) == "contracts/supply/product" {
			return
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if !ok {
					continue
				}
				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if name.Name != "ProductSKUCode" {
							continue
						}
						identifier, stringScalar := field.Type.(*ast.Ident)
						if !stringScalar || identifier.Name != "string" {
							t.Errorf("%s uses non-scalar ProductSKUCode on %s", fset.Position(field.Pos()), typeSpecification.Name.Name)
						}
						jsonKey, present := v26JSONFieldName(t, path, field)
						if !present || jsonKey != "product_sku_code" {
							t.Errorf("%s ProductSKUCode on %s must use JSON key product_sku_code", fset.Position(field.Pos()), typeSpecification.Name.Name)
						}
					}
				}
			}
		}
	})
}

type canonicalProductField struct {
	json   string
	typeOf reflect.Type
}

func canonicalProductAssertExactFields(t *testing.T, model reflect.Type, expected map[string]canonicalProductField) {
	t.Helper()
	if model.NumField() != len(expected) {
		t.Errorf("%s has %d fields, want exactly %d", model, model.NumField(), len(expected))
	}
	for name, want := range expected {
		field, ok := model.FieldByName(name)
		if !ok {
			t.Errorf("%s is missing required field %s", model, name)
			continue
		}
		if field.Type != want.typeOf {
			t.Errorf("%s.%s type = %s, want %s", model, name, field.Type, want.typeOf)
		}
		if got := field.Tag.Get("json"); got != want.json {
			t.Errorf("%s.%s JSON tag = %q, want %q", model, name, got, want.json)
		}
	}
	for index := 0; index < model.NumField(); index++ {
		field := model.Field(index)
		if _, expectedField := expected[field.Name]; !expectedField {
			t.Errorf("%s contains unexpected field %s", model, field.Name)
		}
	}
}

func canonicalProductRequireJSONFields(t *testing.T, model reflect.Type, expected map[string]string) {
	t.Helper()
	for name, wantTag := range expected {
		field, ok := model.FieldByName(name)
		if !ok {
			t.Errorf("%s is missing required field %s", model, name)
			continue
		}
		if got := field.Tag.Get("json"); got != wantTag {
			t.Errorf("%s.%s JSON tag = %q, want %q", model, name, got, wantTag)
		}
	}
}

func canonicalProductAssertNoOperationalCommerceEvidence(t *testing.T, model reflect.Type) {
	t.Helper()
	forbiddenFragments := []string{
		"depot", "location", "bucket", "lot", "unit_id", "unit_ids",
		"available", "reserved", "quantity", "count", "balance", "inventory",
		"audit", "history", "created", "updated", "deleted", "source",
		"condition", "disposition", "date_mark",
	}
	for index := 0; index < model.NumField(); index++ {
		field := model.Field(index)
		for _, value := range []string{field.Name, field.Tag.Get("json")} {
			lowerValue := strings.ToLower(value)
			for _, forbidden := range forbiddenFragments {
				if strings.Contains(lowerValue, forbidden) {
					t.Errorf("%s exposes forbidden operational evidence through %s", model, field.Name)
					break
				}
			}
		}
	}
}

func canonicalProductWalkProductionGoFiles(t *testing.T, inspect func(path string, relativePath string, fset *token.FileSet, file *ast.File)) {
	t.Helper()
	pkgRoot := sharedContractPkgRoot(t)
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
		inspect(path, relativePkgPath(t, pkgRoot, path), fset, file)
		return nil
	})
	if err != nil {
		t.Fatalf("scan canonical-product production sources: %v", err)
	}
}

func canonicalProductTypeUsesFullProduct(expression ast.Expr, aliases map[string]struct{}) bool {
	usesFullProduct := false
	ast.Inspect(expression, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != "Product" {
			return true
		}
		qualifier, ok := selector.X.(*ast.Ident)
		if !ok {
			return true
		}
		if _, imported := aliases[qualifier.Name]; imported {
			usesFullProduct = true
		}
		return true
	})
	return usesFullProduct
}

func canonicalProductStringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}
