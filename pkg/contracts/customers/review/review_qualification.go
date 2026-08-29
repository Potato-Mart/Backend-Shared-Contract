package review

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/review/review_enums"
)

// ReviewQualification stores service-verified authorization evidence. The
// review service owns authorization; clients must not infer eligibility from a
// boolean field in the contract.
type ReviewQualification struct {
	Kind       review_enums.ReviewQualificationKind `json:"kind"`
	Reference  string                               `json:"reference,omitempty"`
	VerifiedAt *time.Time                           `json:"verified_at,omitempty"`
}
