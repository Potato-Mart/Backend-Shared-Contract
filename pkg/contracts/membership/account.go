package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/membership"
)

// MembershipWalletSummary is a projected wallet balance. PointLedgerEntry is
// the source of truth.
type MembershipWalletSummary struct {
	TotalPoints     int        `json:"total_points"`
	ReservedPoints  int        `json:"reserved_points"`
	AvailablePoints int        `json:"available_points"`
	ExpiringPoints  int        `json:"expiring_points"`
	NextExpiryAt    *time.Time `json:"next_expiry_at,omitempty"`
	CalculatedAt    time.Time  `json:"calculated_at"`
}

// MembershipAccount is the programme account for a retail customer. ID is the
// retail customer's customer number.
type MembershipAccount struct {
	ID          string                                 `json:"id"`
	TierKey     string                                 `json:"tier_key,omitempty"`
	Status      membershipenum.MembershipAccountStatus `json:"status"`
	Wallet      MembershipWalletSummary                `json:"wallet"`
	EnrolledAt  time.Time                              `json:"enrolled_at"`
	SuspendedAt *time.Time                             `json:"suspended_at,omitempty"`
	ClosedAt    *time.Time                             `json:"closed_at,omitempty"`
	Metadata    common.Metadata                        `json:"metadata,omitempty"`

	common.AuditFields
}
