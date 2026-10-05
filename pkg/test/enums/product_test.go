package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/security/security_enums"
	market_enums "github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/pricing/market/market_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/supply/catalogue/product/product_enums"
)

func TestProductEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "productenum.BarcodeFormat", valid: []stringEnum{product_enums.BarcodeFormatCode128}, invalid: product_enums.BarcodeFormat("__invalid__")},
		{name: "securityenum.MediaStatus", valid: []stringEnum{security_enums.MediaStatusPending, security_enums.MediaStatusActive, security_enums.MediaStatusDeleted}, invalid: security_enums.MediaStatus("__invalid__")},
		{name: "productenum.PriceAudience", valid: []stringEnum{market_enums.PriceAudienceRetail, market_enums.PriceAudienceWholesale}, invalid: market_enums.PriceAudience("__invalid__")},
		{name: "productenum.PriceVisibility", valid: []stringEnum{product_enums.PriceVisibilityPublic, product_enums.PriceVisibilityLoginRequired, product_enums.PriceVisibilityWholesaleApprovedOnly, product_enums.PriceVisibilityHidden}, invalid: product_enums.PriceVisibility("__invalid__")},
		{name: "productenum.StorefrontStockState", valid: []stringEnum{product_enums.StorefrontStockStateUnknown, product_enums.StorefrontStockStateInStock, product_enums.StorefrontStockStateOutOfStock}, invalid: product_enums.StorefrontStockState("__invalid__")},
		{name: "productenum.ProductStatus", valid: []stringEnum{product_enums.ProductStatusDraft, product_enums.ProductStatusActive, product_enums.ProductStatusArchived, product_enums.ProductStatusDiscontinued}, invalid: product_enums.ProductStatus("__invalid__")},
	})
}

func TestRetiredProductBarcodeFormatsAreInvalid(t *testing.T) {
	for _, value := range []string{"EAN_13", "EAN_8", "UPC_A", "UPC_E", "QR_CODE"} {
		if product_enums.BarcodeFormat(value).IsValid() {
			t.Errorf("retired barcode format %s remains valid", value)
		}
	}
}
