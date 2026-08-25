package classification

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/party"
)

// Supplier is the full supplier record. A supplier is an organisation, so it
// carries the complete organisation profile via party.OrganisationDetail
// (which embeds PartyRef for id / name / phone / email, plus registration,
// tax, addresses, branding and other organisation fields).
type Supplier struct {
	party.OrganisationDetail
	GeographicLocation   *geography.Address           `json:"geographic_location,omitempty"`
	AvailableMarketCodes []string                     `json:"available_market_codes,omitempty"`
	AvailableProducts    []SupplierAvailableProduct   `json:"available_products,omitempty"`
	AvailablePromotions  []SupplierAvailablePromotion `json:"available_promotions,omitempty"`
	audit.AuditFields
}
