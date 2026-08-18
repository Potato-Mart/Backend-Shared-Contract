package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/metadata"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/membership/membership_enums"
)

// MembershipWalletSummary is a projected wallet balance. PointLedgerEntry is
// the source of truth.
type MembershipWalletSummary struct {
	TotalPoints     int        `json:"total_points"`
	ReservedPoints  int        `json:"reserved_points"`
	AvailablePoints int        `json:"available_points"`
	PointDebt       int        `json:"point_debt"`
	ExpiringPoints  int        `json:"expiring_points"`
	NextExpiryAt    *time.Time `json:"next_expiry_at,omitempty"`
	CalculatedAt    time.Time  `json:"calculated_at"`
}

// MembershipAccount is the programme account for a retail customer. ID is the
// retail customer's customer number.
type MembershipAccount struct {
	ID          string                                   `json:"id"`
	TierKey     string                                   `json:"tier_key,omitempty"`
	Status      membership_enums.MembershipAccountStatus `json:"status"`
	Wallet      MembershipWalletSummary                  `json:"wallet"`
	EnrolledAt  time.Time                                `json:"enrolled_at"`
	SuspendedAt *time.Time                               `json:"suspended_at,omitempty"`
	ClosedAt    *time.Time                               `json:"closed_at,omitempty"`
	Metadata    metadata.Metadata                        `json:"metadata,omitempty"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
