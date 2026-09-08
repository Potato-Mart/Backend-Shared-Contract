package product_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/product/product_enums"
)

// Symbology and leading zeros are data, not values a consumer may infer from
// barcode length. Supply owns character, length, and checksum validation.
func TestBarcodeDisplayPreservesExplicitSymbologyAndLeadingZeros(t *testing.T) {
	cases := []struct {
		format product_enums.BarcodeFormat
		wire   string
		value  string
	}{
		{product_enums.BarcodeFormatCode128, "CODE_128", "00-LOT-A128"},
		{product_enums.BarcodeFormatEAN13, "EAN_13", "0012345678905"},
		{product_enums.BarcodeFormatUPCA, "UPC_A", "012345678905"},
	}
	for _, tc := range cases {
		t.Run(tc.wire, func(t *testing.T) {
			from := time.Date(2026, 9, 8, 0, 0, 0, 0, time.UTC)
			until := from.Add(24 * time.Hour)
			master := product.ProductBarcodeAssignment{
				Code: "BARCODE-A", SKUCode: "SKU-A", PackageOptionCode: "PACK-A",
				Value: tc.value, Format: tc.format, ManufacturerCode: "MAKER-A",
				IsPrimary: true, EffectiveFrom: from, EffectiveTo: &until,
			}
			encoded, err := json.Marshal(master)
			if err != nil {
				t.Fatal(err)
			}
			var decoded product.ProductBarcodeAssignment
			if err := json.Unmarshal(encoded, &decoded); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(decoded, master) {
				t.Fatalf("master barcode did not round trip: %s", encoded)
			}

			projection := product.SellingProductBarcode{
				PackageOptionCode: "PACK-A", Value: tc.value, Format: tc.format, IsPrimary: true,
			}
			encoded, err = json.Marshal(projection)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(encoded, &fields); err != nil {
				t.Fatal(err)
			}
			if len(fields) != 4 {
				t.Fatalf("display must contain only package/value/format/primary: %s", encoded)
			}
			for key, want := range map[string]string{
				"package_option_code": `"PACK-A"`, "value": `"` + tc.value + `"`,
				"format": `"` + tc.wire + `"`, "is_primary": "true",
			} {
				if string(fields[key]) != want {
					t.Errorf("%s = %s, want %s", key, fields[key], want)
				}
			}
			var decodedProjection product.SellingProductBarcode
			if err := json.Unmarshal(encoded, &decodedProjection); err != nil {
				t.Fatal(err)
			}
			if decodedProjection != projection {
				t.Fatalf("display barcode did not round trip: %s", encoded)
			}
		})
	}
}
