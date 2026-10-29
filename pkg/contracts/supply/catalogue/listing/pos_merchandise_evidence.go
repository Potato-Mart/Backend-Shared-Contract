package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v51/pkg/contracts/common/packaging/packaging_enums"
)

// POSMerchandiseEvidence records Supply's stock-independent merchandise proof
// for a specific market, depot and package. It does not assert inventory,
// reservation, buyer eligibility or pricing; Supply must revalidate its token.
// It is online merchandise evidence, not an offline-sale journal or price book.
type POSMerchandiseEvidence struct {
	SKUCode             string                              `json:"sku_code"`
	MarketCode          string                              `json:"market_code"`
	DepotCode           string                              `json:"depot_code"`
	PackageOption       packaging.PackageOptionRef          `json:"package_option"`
	HandlingUnit        packaging_enums.PackageHandlingUnit `json:"handling_unit"`
	PackageCount        int64                               `json:"package_count"`
	UnitsPerPackage     int64                               `json:"units_per_package"`
	BaseUnits           int64                               `json:"base_units"`
	ListingRevision     int64                               `json:"listing_revision"`
	MarketRevision      int64                               `json:"market_revision"`
	TaxCategoryRevision int64                               `json:"tax_category_revision"`
	TaxCategoryCode     string                              `json:"tax_category_code"`
	TaxRuleID           string                              `json:"tax_rule_id"`
	ProductFactsDigest  string                              `json:"product_facts_digest"`
	TaxRuleRevision     int64                               `json:"tax_rule_revision"`
	Token               string                              `json:"token"`
	CapturedAt          time.Time                           `json:"captured_at"`
	ValidUntil          time.Time                           `json:"valid_until"`
}
