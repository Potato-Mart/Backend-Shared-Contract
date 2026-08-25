package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// TrademarkEvidence is a cited search result or manually verified record. It
// does not itself make a legal-clearance claim.
type TrademarkEvidence struct {
	ID                 string                               `json:"id"`
	Revision           RevisionMetadata                     `json:"revision"`
	SKUCode            string                               `json:"sku_code,omitempty"`
	Jurisdiction       import_compliance_enums.Jurisdiction `json:"jurisdiction"`
	Mark               string                               `json:"mark"`
	Status             string                               `json:"status,omitempty"`
	Owner              string                               `json:"owner,omitempty"`
	RegistrationNumber string                               `json:"registration_number,omitempty"`
	Classes            []string                             `json:"classes,omitempty"`
	Source             string                               `json:"source"`
	CheckedAt          *time.Time                           `json:"checked_at,omitempty"`
	Evidence           []EvidenceReference                  `json:"evidence,omitempty"`

	audit.AuditFields
}
