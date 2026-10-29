package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/listing"

// SaleEvidence carries exactly one evidence variant for a Pricing quote line.
// Producers and consumers enforce exclusivity, channel suitability and owner
// revalidation. No variant grants a price or permission. Offline POS sales use
// the separate offline snapshot rather than pretending to have an online quote.
// An Orders-owned digital gift purchase is a separate face-value instrument,
// not a Pricing quote or a stock sale. Its synthetic snapshot has no variant;
// Orders must establish its purchase identity and digital classification.
type SaleEvidence struct {
	Stocked            *listing.SaleEligibilitySnapshot `json:"stocked,omitempty"`
	WholesaleBackorder *listing.BackorderEvidence       `json:"wholesale_backorder,omitempty"`
	POSMerchandise     *listing.POSMerchandiseEvidence  `json:"pos_merchandise,omitempty"`
}
