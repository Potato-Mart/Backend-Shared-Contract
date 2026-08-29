package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

type RFIExternalEvent struct {
	ID                string                                     `json:"id"`
	State             import_compliance_enums.RFISubmissionState `json:"state"`
	ExternalReference string                                     `json:"external_reference,omitempty"`
	OccurredAt        time.Time                                  `json:"occurred_at"`
	RecordedBy        string                                     `json:"recorded_by,omitempty"`
	Note              string                                     `json:"note,omitempty"`
	Evidence          []EvidenceReference                        `json:"evidence,omitempty"`
}
