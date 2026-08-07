package enums_test

import (
	"testing"

	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/security"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/product"
)

func TestProductEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "productenum.BarcodeFormat", valid: []stringEnum{productenum.BarcodeFormatEAN8, productenum.BarcodeFormatEAN13, productenum.BarcodeFormatUPCA, productenum.BarcodeFormatUPCE, productenum.BarcodeFormatCode128, productenum.BarcodeFormatQRCode}, invalid: productenum.BarcodeFormat("__invalid__")},
		{name: "securityenum.MediaStatus", valid: []stringEnum{securityenum.MediaStatusPending, securityenum.MediaStatusActive, securityenum.MediaStatusDeleted}, invalid: securityenum.MediaStatus("__invalid__")},
		{name: "productenum.PriceAudience", valid: []stringEnum{productenum.PriceAudienceRetail, productenum.PriceAudienceWholesale}, invalid: productenum.PriceAudience("__invalid__")},
		{name: "productenum.PriceVisibility", valid: []stringEnum{productenum.PriceVisibilityPublic, productenum.PriceVisibilityLoginRequired, productenum.PriceVisibilityWholesaleApprovedOnly, productenum.PriceVisibilityHidden}, invalid: productenum.PriceVisibility("__invalid__")},
		{name: "productenum.WholesalePriceMode", valid: []stringEnum{productenum.WholesalePriceModeFixed, productenum.WholesalePriceModeOnRequest}, invalid: productenum.WholesalePriceMode("__invalid__")},
		{name: "productenum.StorefrontPreorderStatus", valid: []stringEnum{productenum.StorefrontPreorderStatusUnavailable, productenum.StorefrontPreorderStatusUpcoming, productenum.StorefrontPreorderStatusOpen, productenum.StorefrontPreorderStatusClosed, productenum.StorefrontPreorderStatusSoldOut}, invalid: productenum.StorefrontPreorderStatus("__invalid__")},
		{name: "productenum.StorefrontStockState", valid: []stringEnum{productenum.StorefrontStockStateUnknown, productenum.StorefrontStockStateInStock, productenum.StorefrontStockStateOutOfStock}, invalid: productenum.StorefrontStockState("__invalid__")},
		{name: "productenum.ProductStatus", valid: []stringEnum{productenum.ProductStatusDraft, productenum.ProductStatusActive, productenum.ProductStatusArchived, productenum.ProductStatusDiscontinued}, invalid: productenum.ProductStatus("__invalid__")},
	})
	if got := productenum.WholesalePriceModeFixed.String(); got != "fixed" {
		t.Fatalf("fixed wholesale price mode wire value = %q", got)
	}
	if got := productenum.WholesalePriceModeOnRequest.String(); got != "on_request" {
		t.Fatalf("on-request wholesale price mode wire value = %q", got)
	}
}
