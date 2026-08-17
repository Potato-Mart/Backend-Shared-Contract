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

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/measurement"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product/product_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

// v27FrozenSKUCodeTypes is the frozen allowlist of transaction-evidence types
// that may keep a frozen sku_code alongside their sku_id link. Every other
// model links to a SKU by sku_id only.
var v27FrozenSKUCodeTypes = map[string]struct{}{
	"orders/order.CartItem":               {},
	"orders/order.OrderItem":              {},
	"orders/order.OrderLineSummary":       {},
	"orders/pos.ReceiptLine":              {},
	"supply/purchase.OrderItem":           {},
	"supply/purchase.ReceiptItem":         {},
	"supply/purchase.SupplierInvoiceLine": {},
}

func TestV27GlobalProductCarriesOnlyContentClassificationAndAdministration(t *testing.T) {
	canonicalProductAssertExactFields(t, reflect.TypeOf(product.Product{}), map[string]canonicalProductField{
		"ID": {
			json:   "id",
			typeOf: reflect.TypeOf(""),
		},
		"Status": {
			json:   "status",
			typeOf: reflect.TypeOf(product_enums.ProductStatus("")),
		},
		"Content": {
			json:   "content",
			typeOf: reflect.TypeOf(product.ProductContent{}),
		},
		"Classification": {
			json:   "classification",
			typeOf: reflect.TypeOf(product.ProductClassification{}),
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

func TestV27SKUIsTheSellableIdentityWithExactFields(t *testing.T) {
	canonicalProductAssertExactFields(t, reflect.TypeOf(product.SKU{}), map[string]canonicalProductField{
		"ID":        {json: "id", typeOf: reflect.TypeOf("")},
		"ProductID": {json: "product_id", typeOf: reflect.TypeOf("")},
		"Code":      {json: "code", typeOf: reflect.TypeOf("")},
		"PackageOptions": {
			json:   "package_options",
			typeOf: reflect.TypeOf([]product.ProductPackageOption{}),
		},
		"BarcodeAssignments": {
			json:   "barcode_assignments,omitempty",
			typeOf: reflect.TypeOf([]product.ProductBarcodeAssignment{}),
		},
		"NetContent": {
			json:   "net_content,omitempty",
			typeOf: reflect.TypeOf((*measurement.NetContent)(nil)),
		},
		"StorageType": {
			json:   "storage_type,omitempty",
			typeOf: reflect.TypeOf(warehouse_enums.StorageType("")),
		},
		"Status": {
			json:   "status",
			typeOf: reflect.TypeOf(product_enums.SKUStatus("")),
		},
		"AuditFields": {
			json:   "",
			typeOf: reflect.TypeOf(audit.AuditFields{}),
		},
	})
	auditField, ok := reflect.TypeOf(product.SKU{}).FieldByName("AuditFields")
	if !ok || !auditField.Anonymous {
		t.Error("product.SKU must embed audit.AuditFields")
	}
}

func TestV27ProductAndSKUCarryNoPriceTaxOrMarketFacts(t *testing.T) {
	forbiddenFragments := []string{
		"price", "tax", "market", "currency", "amount_minor", "discount",
		"selling", "commerce", "channel", "audience", "visibility",
		"depot", "available", "reserved", "stock",
	}
	for _, model := range []reflect.Type{
		reflect.TypeOf(product.Product{}),
		reflect.TypeOf(product.ProductContent{}),
		reflect.TypeOf(product.ProductClassification{}),
		reflect.TypeOf(product.SKU{}),
	} {
		for index := 0; index < model.NumField(); index++ {
			field := model.Field(index)
			for _, value := range []string{field.Name, field.Tag.Get("json")} {
				lowerValue := strings.ToLower(value)
				for _, forbidden := range forbiddenFragments {
					if strings.Contains(lowerValue, forbidden) {
						t.Errorf("%s exposes forbidden market/commercial fact through %s", model, field.Name)
						break
					}
				}
			}
		}
	}
}

func TestV27ProductImagesHaveOnlyObjectMediaCoverGalleryAndDetails(t *testing.T) {
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

func TestV27ProductAdministrationIsOptionalAndOwnsHistoryAndAudit(t *testing.T) {
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

func TestV27CanonicalProductRetiresLegacyCatalogueProjections(t *testing.T) {
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
		"ProductCommerce",
		"ProductPackageCommerce",
		"ProductMetrics",
		"ProductPackaging",
		"Selling",
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
		"contracts/supply/classification/sku.go",
	} {
		path := filepath.Join(sharedContractPkgRoot(t), filepath.FromSlash(retiredPath))
		if _, err := os.Stat(path); err == nil {
			t.Errorf("retired canonical-product source file remains: %s", retiredPath)
		} else if !os.IsNotExist(err) {
			t.Errorf("inspect retired canonical-product path %s: %v", retiredPath, err)
		}
	}
}

func TestV27OnlyProductPackageDeclaresGlobalCatalogueModels(t *testing.T) {
	const productImportPath = "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product"

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
				if typeSpecification.Name.Name != "Product" && typeSpecification.Name.Name != "SKU" {
					continue
				}
				if directory != "contracts/supply/product" {
					t.Errorf("%s declares %s outside supply/product; product.Product and product.SKU are the only global catalogue models", fset.Position(typeSpecification.Pos()), typeSpecification.Name.Name)
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
			t.Errorf("%s dot-imports supply/product; cross-domain models must use SKUID", path)
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
			if canonicalProductTypeUsesGlobalCatalogueModel(field.Type, aliases) {
				t.Errorf("%s embeds product.Product or product.SKU outside supply/product; use sku_id plus transaction-owned frozen facts", fset.Position(field.Pos()))
			}
			return true
		})
	})
}

func TestV27CrossDomainSKULinksUseScalarSKUID(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	canonicalProductWalkProductionGoFiles(t, func(path string, relativePath string, fset *token.FileSet, file *ast.File) {
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
				typeKey := v27ProductionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				_, frozenAllowed := v27FrozenSKUCodeTypes[typeKey]
				for _, field := range structure.Fields.List {
					jsonKey, present := v27JSONFieldName(t, path, field)
					for _, name := range field.Names {
						switch name.Name {
						case "SKUID":
							identifier, stringScalar := field.Type.(*ast.Ident)
							if !stringScalar || identifier.Name != "string" {
								t.Errorf("%s uses non-scalar SKUID on %s", fset.Position(field.Pos()), typeSpecification.Name.Name)
							}
							if !present || jsonKey != "sku_id" {
								t.Errorf("%s SKUID on %s must use JSON key sku_id", fset.Position(field.Pos()), typeSpecification.Name.Name)
							}
						case "SKUCode":
							if !frozenAllowed {
								t.Errorf("%s declares frozen SKUCode on unlisted type %s; only transaction evidence may freeze a sku_code", fset.Position(field.Pos()), typeKey)
							}
						}
					}
					if !present {
						continue
					}
					if jsonKey == "sku_code" && !frozenAllowed {
						t.Errorf("%s uses frozen JSON key sku_code on unlisted type %s", fset.Position(field.Pos()), typeKey)
					}
					if jsonKey == "product_sku_code" || jsonKey == "product_sku_codes" {
						t.Errorf("%s retains retired cross-domain JSON key %q on %s", fset.Position(field.Pos()), jsonKey, typeKey)
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

func canonicalProductTypeUsesGlobalCatalogueModel(expression ast.Expr, aliases map[string]struct{}) bool {
	usesGlobalModel := false
	ast.Inspect(expression, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if !ok || (selector.Sel.Name != "Product" && selector.Sel.Name != "SKU") {
			return true
		}
		qualifier, ok := selector.X.(*ast.Ident)
		if !ok {
			return true
		}
		if _, imported := aliases[qualifier.Name]; imported {
			usesGlobalModel = true
		}
		return true
	})
	return usesGlobalModel
}

func canonicalProductStringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}
