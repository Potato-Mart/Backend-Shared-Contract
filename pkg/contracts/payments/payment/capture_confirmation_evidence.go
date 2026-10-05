package payment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/payments/payment/payment_enums"
	"time"
)

// CaptureConfirmationEvidence is immutable, authenticated evidence that the
// captured amount had succeeded no later than UpperBoundAt. A later upper bound
// is inconclusive about a deadline, never proof of late capture.
// Payments authenticates and stores the source proof privately. Timestamp
// precision must be accounted for when deriving a conservative upper bound.
// EvidenceReference is an opaque provider event ID or immutable observation ID,
// never raw provider data, a credential, URL or customer contact information.
type CaptureConfirmationEvidence struct {
	UpperBoundAt      time.Time                               `json:"upper_bound_at"`
	Source            payment_enums.CaptureConfirmationSource `json:"source"`
	Provider          string                                  `json:"provider"`
	EvidenceReference string                                  `json:"evidence_reference"`
}
