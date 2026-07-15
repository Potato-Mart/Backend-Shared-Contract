package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	membershipenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/membership"
)

// MembershipOwnerRef identifies the business entity that owns the membership
// account and points wallet.
type MembershipOwnerRef struct {
	OwnerType membershipenum.MembershipOwnerType `json:"owner_type"`
	OwnerID   string                             `json:"owner_id"`
}

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

// MembershipAccount is the global programme account for a retail customer or
// wholesale organisation.
type MembershipAccount struct {
	ID          string                                 `json:"id"`
	Owner       MembershipOwnerRef                     `json:"owner"`
	TierKey     string                                 `json:"tier_key,omitempty"`
	Status      membershipenum.MembershipAccountStatus `json:"status"`
	Wallet      MembershipWalletSummary                `json:"wallet"`
	EnrolledAt  time.Time                              `json:"enrolled_at"`
	SuspendedAt *time.Time                             `json:"suspended_at,omitempty"`
	ClosedAt    *time.Time                             `json:"closed_at,omitempty"`
	Metadata    common.Metadata                        `json:"metadata,omitempty"`

	common.AuditFields
}

// MembershipAccountSummary is safe to embed in customer, organisation, and
// session-facing projections.
type MembershipAccountSummary struct {
	ID              string                                 `json:"id"`
	Owner           MembershipOwnerRef                     `json:"owner"`
	TierKey         string                                 `json:"tier_key,omitempty"`
	Status          membershipenum.MembershipAccountStatus `json:"status"`
	AvailablePoints int                                    `json:"available_points"`
	TotalPoints     int                                    `json:"total_points"`
}
