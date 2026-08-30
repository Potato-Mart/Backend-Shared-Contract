package compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
)

// ManufacturerDeclaration is a revisioned declaration backed by an immutable
// purchase-order snapshot. Multiple declarations may reference the same order.
type ManufacturerDeclaration struct {
	ID       string           `json:"id"`
	Revision RevisionMetadata `json:"revision"`
	// MarketCode and CountryCode are the denormalized market and country the
	// record belongs to, carried so a geographically scoped staff query is
	// a plain indexed match.
	MarketCode           string                `json:"market_code,omitempty"`
	CountryCode          geography.CountryCode `json:"country_code,omitempty"`
	PurchaseOrder        PurchaseOrderSnapshot `json:"purchase_order"`
	DeclarationReference string                `json:"declaration_reference"`
	Shipment             DeclarationShipment   `json:"shipment"`
	Manufacturer         ManufacturerDetails   `json:"manufacturer"`
	Signatory            DeclarationSignatory  `json:"signatory"`
	Lines                []DeclarationLine     `json:"lines"`
	Evidence             []EvidenceReference   `json:"evidence,omitempty"`
	Artifacts            []ArtifactReference   `json:"artifacts,omitempty"`

	audit.AuditFields
}
