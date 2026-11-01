package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging/packaging_enums"
)

// NonStockSaleEvidence is Supply's listing-only evidence for a SKU that is
// sold without physical inventory, such as the placeholder open-price SKUs
// OPEN0001 to OPEN0020. Such a line never reserves, stages or settles stock,
// and the evidence asserts no inventory, reservation, buyer eligibility or
// price. OpenPrice marks a placeholder whose price is entered at sale time.
// Supply must revalidate the token before an order relies on this evidence.
type NonStockSaleEvidence struct {
	SKUCode         string                              `json:"sku_code"`
	MarketCode      string                              `json:"market_code"`
	DepotCode       string                              `json:"depot_code,omitempty"`
	OpenPrice       bool                                `json:"open_price"`
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
