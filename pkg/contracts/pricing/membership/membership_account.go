package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/metadata"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/membership/membership_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/points"
)

// MembershipAccount is the programme account for a retail customer. ID is the
// retail customer's customer number.
type MembershipAccount struct {
	ID          string                                   `json:"id"`
	TierKey     string                                   `json:"tier_key,omitempty"`
	Status      membership_enums.MembershipAccountStatus `json:"status"`
	Wallet      points.PointsSummary                     `json:"wallet"`
	EnrolledAt  time.Time                                `json:"enrolled_at"`
	SuspendedAt *time.Time                               `json:"suspended_at,omitempty"`
	ClosedAt    *time.Time                               `json:"closed_at,omitempty"`
	Metadata    metadata.Metadata                        `json:"metadata,omitempty"`
	// MarketCodes selects membership availability and staff geographic access
	// within the account's country tier group. Pricing must validate a nonempty,
	// unique list of markets in CountryCode using the group's currency. Empty
	// or missing lists do not imply global availability. Pricing owns account
	// participation, validation, role permissions and access enforcement; the
	// list does not select an annual qualification timezone.
	MarketCodes []string              `json:"market_codes"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
	security.DataProtectionFields
}
