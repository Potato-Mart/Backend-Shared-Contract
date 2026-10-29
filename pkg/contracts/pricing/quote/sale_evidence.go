package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/supply/catalogue/listing"

// SaleEvidence carries exactly one evidence variant for a priced line.
// Producers and consumers enforce exclusivity, channel suitability and owner
// revalidation. The offline variant references captured book and journal facts
// rather than a fresh Supply token. No variant grants a price or permission.
type SaleEvidence struct {
	Stocked            *listing.SaleEligibilitySnapshot       `json:"stocked,omitempty"`
	WholesaleBackorder *listing.BackorderEvidence             `json:"wholesale_backorder,omitempty"`
	POSMerchandise     *listing.POSMerchandiseEvidence        `json:"pos_merchandise,omitempty"`
	OfflinePOS         *listing.OfflinePOSMerchandiseEvidence `json:"offline_pos,omitempty"`
}
