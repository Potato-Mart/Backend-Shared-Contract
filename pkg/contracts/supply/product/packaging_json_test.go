package product

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/product/product_enums"
)

func TestPackageAndBarcodeReferencesUseBusinessCodes(t *testing.T) {
	now := time.Date(2026, 8, 19, 0, 0, 0, 0, time.UTC)
	payload, err := json.Marshal(struct {
		Package ProductPackageOption     `json:"package"`
		Barcode ProductBarcodeAssignment `json:"barcode"`
	}{
		Package: ProductPackageOption{Code: "PKG-A00001-EACH", HandlingUnit: packaging_enums.PackageHandlingUnitEach, UnitsPerPackage: 1, EffectiveFrom: now},
		Barcode: ProductBarcodeAssignment{Code: "BAR-A00001", PackageOptionCode: "PKG-A00001-EACH", Value: "A00001", Format: product_enums.BarcodeFormatCode128, EffectiveFrom: now},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, retired := range []string{`"sku_id"`, `"package_option_id"`, `"manufacturer_id"`} {
		if strings.Contains(string(payload), retired) {
			t.Fatalf("package JSON retained %s: %s", retired, payload)
		}
	}
}
