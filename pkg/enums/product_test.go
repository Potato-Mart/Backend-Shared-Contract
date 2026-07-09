package enums_test

import (
	"testing"

	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/product"
)

func TestProductEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "productenum.MediaStatus", valid: []stringEnum{productenum.MediaStatusPending, productenum.MediaStatusActive, productenum.MediaStatusDeleted}, invalid: productenum.MediaStatus("__invalid__")},
		{name: "productenum.PriceAudience", valid: []stringEnum{productenum.PriceAudienceRetail, productenum.PriceAudienceWholesale}, invalid: productenum.PriceAudience("__invalid__")},
		{name: "productenum.PriceVisibility", valid: []stringEnum{productenum.PriceVisibilityPublic, productenum.PriceVisibilityLoginRequired, productenum.PriceVisibilityWholesaleApprovedOnly, productenum.PriceVisibilityHidden}, invalid: productenum.PriceVisibility("__invalid__")},
		{name: "productenum.StorefrontPreorderStatus", valid: []stringEnum{productenum.StorefrontPreorderStatusUnavailable, productenum.StorefrontPreorderStatusUpcoming, productenum.StorefrontPreorderStatusOpen, productenum.StorefrontPreorderStatusClosed, productenum.StorefrontPreorderStatusSoldOut}, invalid: productenum.StorefrontPreorderStatus("__invalid__")},
		{name: "productenum.StorefrontExpiryStatus", valid: []stringEnum{productenum.StorefrontExpiryStatusNotApplicable, productenum.StorefrontExpiryStatusSoonExpiry, productenum.StorefrontExpiryStatusExpired}, invalid: productenum.StorefrontExpiryStatus("__invalid__")},
		{name: "productenum.ProductStatus", valid: []stringEnum{productenum.ProductStatusDraft, productenum.ProductStatusActive, productenum.ProductStatusArchived, productenum.ProductStatusDiscontinued}, invalid: productenum.ProductStatus("__invalid__")},
		{name: "productenum.SalesPerformance", valid: []stringEnum{productenum.SalesPerformanceHot, productenum.SalesPerformanceNormal, productenum.SalesPerformanceSlow}, invalid: productenum.SalesPerformance("__invalid__")},
	})
}
