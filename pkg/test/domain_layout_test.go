package pkg_test

import (
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestV27DomainPackageLayout verifies the v27 contract's requested domain paths.
func TestV27DomainPackageLayout(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	legacyImportCompliance := "import" + "compliance"
	requiredFiles := map[string]string{
		"contracts/supply/classification/sku.go":                                          "classification",
		"contracts/supply/classification/category_tag.go":                                 "classification",
		"contracts/supply/classification/collection.go":                                   "classification",
		"contracts/supply/classification/brand.go":                                        "classification",
		"contracts/supply/classification/supplier.go":                                     "classification",
		"contracts/supply/classification/favourite_list.go":                               "classification",
		"contracts/supply/classification/classification_enums/favourite_list_types.go":    "classification_enums",
		"contracts/supply/import_compliance/declaration.go":                               "import_compliance",
		"contracts/supply/import_compliance/import_compliance_enums/import_compliance.go": "import_compliance_enums",
		"contracts/supply/operations/availability.go":                                     "operations",
		"contracts/supply/operations/packing.go":                                          "operations",
		"contracts/supply/operations/picking.go":                                          "operations",
		"contracts/supply/operations/reservation.go":                                      "operations",
		"contracts/supply/operations/inbound.go":                                          "operations",
		"contracts/supply/operations/stock_movement.go":                                   "operations",
		"contracts/supply/operations/inventory_category_tag_evidence.go":                  "operations",
		"contracts/orders/shipping/shipment.go":                                           "shipping",
		"contracts/pricing/promotion/promotion.go":                                        "promotion",
		"contracts/pricing/promotion/scope.go":                                            "promotion",
		"contracts/pricing/promotion/relation.go":                                         "promotion",
		"contracts/pricing/promotion/application.go":                                      "promotion",
		"contracts/pricing/promotion/package_pricing.go":                                  "promotion",
		"contracts/pricing/promotion/promotion_enums/promotion_status.go":                 "promotion_enums",
		"contracts/pricing/promotion/promotion_enums/promotion_match_mode.go":             "promotion_enums",
		"contracts/pricing/wallet/coupon.go":                                              "wallet",
	}

	for relativePath, wantPackage := range requiredFiles {
		path := filepath.Join(pkgRoot, filepath.FromSlash(relativePath))
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.PackageClauseOnly)
		if err != nil {
			t.Errorf("required v27 model file %s is missing or invalid: %v", relativePath, err)
			continue
		}
		if got := file.Name.Name; got != wantPackage {
			t.Errorf("%s package = %s, want %s", relativePath, got, wantPackage)
		}
	}

	for _, relativePath := range []string{
		"contracts/supply/category",
		"contracts/supply/favourite",
		"contracts/supply/" + legacyImportCompliance,
		"contracts/supply/product/brand.go",
		"contracts/supply/product/category_tag.go",
		"contracts/supply/product/collection.go",
		"contracts/supply/product/supply.go",
		"contracts/supply/purchase/supplier.go",
		"contracts/supply/warehouse/packing.go",
		"contracts/supply/warehouse/picking.go",
		"contracts/supply/warehouse/reservation.go",
		"contracts/supply/warehouse/inbound.go",
		"contracts/supply/warehouse/stock_movement.go",
		"contracts/supply/warehouse/shipment.go",
		"contracts/pricing/promotion/coupon.go",
		"contracts/pricing/promotion/promotion_enums/coupon_type.go",
		"contracts/pricing/promotion/effective_promotion.go",
		"contracts/pricing/promotion/group_order_discount.go",
		"contracts/pricing/promotion/receipt_offer.go",
		"contracts/pricing/promotion/shared.go",
		"contracts/pricing/promotion/storefront_promotion.go",
		"contracts/supply/product/availability.go",
		"contracts/supply/product/detail_image.go",
		"contracts/supply/product/offer.go",
		"contracts/orders/order/volume_discount.go",
		"test/enums/" + legacyImportCompliance + "_test.go",
	} {
		path := filepath.Join(pkgRoot, filepath.FromSlash(relativePath))
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("retired v27 path must not exist: %s", relativePath)
		}
	}
}

func TestV27SourcesRejectLegacyImportComplianceIdentifier(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	legacyIdentifier := "import" + "compliance"
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(contents), legacyIdentifier) {
			t.Errorf("%s retains legacy %s identifier", relativePkgPath(t, pkgRoot, path), legacyIdentifier)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("scan v27 source layout: %v", err)
	}
}
