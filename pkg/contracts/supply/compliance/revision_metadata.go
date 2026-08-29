// Package compliance defines reusable import-compliance records and
// snapshots. Transport DTOs, validation, authorization, calculations, and
// lifecycle transitions belong to the owning backend.
package compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"

	compliance_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/compliance/compliance_enums"
)

// RevisionMetadata identifies one immutable-or-editable revision and records
// the review actions already applied to it. The owning backend defines which
// transitions are permitted.
type RevisionMetadata struct {
	Number             int64                        `json:"number"`
	BaseRevisionNumber *int64                       `json:"base_revision_number,omitempty"`
	State              compliance_enums.ReviewState `json:"state"`
	Submitted          *audit.LifecycleAction       `json:"submitted,omitempty"`
	Approved           *audit.LifecycleAction       `json:"approved,omitempty"`
	Rejected           *audit.LifecycleAction       `json:"rejected,omitempty"`
	Archived           *audit.LifecycleAction       `json:"archived,omitempty"`
}
