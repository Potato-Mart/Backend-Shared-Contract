package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/metadata"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/wallet"
)

// MembershipAccount is the programme account for a retail customer. ID is the
// retail customer's customer number.
type MembershipAccount struct {
	ID          string                                   `json:"id"`
	TierKey     string                                   `json:"tier_key,omitempty"`
	Status      membership_enums.MembershipAccountStatus `json:"status"`
	Wallet      wallet.PointsSummary                     `json:"wallet"`
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
