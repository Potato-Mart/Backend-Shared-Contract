package compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
)

// TariffAssessment is the revisioned, purchase-order-specific worksheet. An
// approved assessment may publish separate reusable TariffProfile revisions.
type TariffAssessment struct {
	ID            string                `json:"id"`
	Revision      RevisionMetadata      `json:"revision"`
	PurchaseOrder PurchaseOrderSnapshot `json:"purchase_order"`
	Rows          []TariffAssessmentRow `json:"rows"`
	Evidence      []EvidenceReference   `json:"evidence,omitempty"`
	Artifacts     []ArtifactReference   `json:"artifacts,omitempty"`

	audit.AuditFields
}
