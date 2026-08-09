package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/product/product_enums"
)

func TestProductEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "productenum.BarcodeFormat", valid: []stringEnum{product_enums.BarcodeFormatEAN8, product_enums.BarcodeFormatEAN13, product_enums.BarcodeFormatUPCA, product_enums.BarcodeFormatUPCE, product_enums.BarcodeFormatCode128, product_enums.BarcodeFormatQRCode}, invalid: product_enums.BarcodeFormat("__invalid__")},
		{name: "securityenum.MediaStatus", valid: []stringEnum{security_enums.MediaStatusPending, security_enums.MediaStatusActive, security_enums.MediaStatusDeleted}, invalid: security_enums.MediaStatus("__invalid__")},
		{name: "productenum.PriceAudience", valid: []stringEnum{product_enums.PriceAudienceRetail, product_enums.PriceAudienceWholesale}, invalid: product_enums.PriceAudience("__invalid__")},
		{name: "productenum.PriceVisibility", valid: []stringEnum{product_enums.PriceVisibilityPublic, product_enums.PriceVisibilityLoginRequired, product_enums.PriceVisibilityWholesaleApprovedOnly, product_enums.PriceVisibilityHidden}, invalid: product_enums.PriceVisibility("__invalid__")},
		{name: "productenum.StorefrontStockState", valid: []stringEnum{product_enums.StorefrontStockStateUnknown, product_enums.StorefrontStockStateInStock, product_enums.StorefrontStockStateOutOfStock}, invalid: product_enums.StorefrontStockState("__invalid__")},
		{name: "productenum.ProductStatus", valid: []stringEnum{product_enums.ProductStatusDraft, product_enums.ProductStatusActive, product_enums.ProductStatusArchived, product_enums.ProductStatusDiscontinued}, invalid: product_enums.ProductStatus("__invalid__")},
	})
}
