package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/packaging/packaging_enums"
)

// BackorderEvidence is Supply's listing-only evidence for a wholesale order.
// It asserts neither physical inventory, a reservation nor an arrival time.
// Supply must revalidate the token before an order relies on this evidence.
type BackorderEvidence struct {
	Allowed         bool                                `json:"allowed"`
	SKUCode         string                              `json:"sku_code"`
	MarketCode      string                              `json:"market_code"`
	PackageOption   packaging.PackageOptionRef          `json:"package_option"`
	HandlingUnit    packaging_enums.PackageHandlingUnit `json:"handling_unit"`
	PackageCount    int64                               `json:"package_count"`
	UnitsPerPackage int64                               `json:"units_per_package"`
	BaseUnits       int64                               `json:"base_units"`
	ListingRevision int64                               `json:"listing_revision"`
	EvidenceToken   string                              `json:"evidence_token"`
	ValidUntil      time.Time                           `json:"valid_until"`
	CapturedAt      time.Time                           `json:"captured_at"`
}
