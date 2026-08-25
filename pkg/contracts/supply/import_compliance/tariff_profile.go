package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// TariffProfile is an approved reusable classification for one SKU and
// jurisdiction. Assessment records retain their own frozen classification.
type TariffProfile struct {
	ID       string           `json:"id"`
	Revision RevisionMetadata `json:"revision"`
	// MarketCode and CountryCode are the denormalized market and country the
	// record belongs to, carried so a geographically scoped staff query is
	// a plain indexed match.
	MarketCode           string                               `json:"market_code,omitempty"`
	CountryCode          geography.CountryCode                `json:"country_code,omitempty"`
	SKUCode              string                               `json:"sku_code"`
	Jurisdiction         import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Classification       TariffClassification                 `json:"classification"`
	EffectiveFrom        temporal.Date                        `json:"effective_from,omitempty"`
	EffectiveTo          temporal.Date                        `json:"effective_to,omitempty"`
	TrademarkEvidenceIDs []string                             `json:"trademark_evidence_ids,omitempty"`
	Evidence             []EvidenceReference                  `json:"evidence,omitempty"`

	audit.AuditFields
}
