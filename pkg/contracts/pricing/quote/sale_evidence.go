package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/supply/catalogue/listing"

// SaleEvidence carries exactly one Supply-owned evidence variant for a priced
// line. Producers and consumers enforce exclusivity, channel suitability and
// owner revalidation. No variant by itself grants a price or buyer permission.
type SaleEvidence struct {
	Stocked            *listing.SaleEligibilitySnapshot `json:"stocked,omitempty"`
	WholesaleBackorder *listing.BackorderEvidence       `json:"wholesale_backorder,omitempty"`
	POSMerchandise     *listing.POSMerchandiseEvidence  `json:"pos_merchandise,omitempty"`
}
