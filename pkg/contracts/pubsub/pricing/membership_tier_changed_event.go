package pricing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/membership/membership_enums"
)

// MembershipTierChangedEvent is membership.tier_changed event_version v1,
// published by Pricing on storefront-events after a committed change to a
// customer's effective membership tier. Customers keeps a projection of the
// current tier per customer. AggregateID is the CustomerNumber.
// Pricing owns qualification, tier definitions and validation; this fact
// carries no spend, money, points, benefits, actor or contact data.
type MembershipTierChangedEvent struct {
	// CustomerNumber is the retail customer's canonical number. Required.
	CustomerNumber string `json:"customer_number"`
	// CountryCode identifies the country membership tier group the tiers
	// belong to. Required.
	CountryCode geography.CountryCode `json:"country_code"`
	// TierKey is the effective tier after the change. Null means the customer
	// currently has no effective tier in the group.
	TierKey *string `json:"tier_key"`
	// PreviousTierKey is the effective tier before the change, null when the
	// customer had none.
	PreviousTierKey *string `json:"previous_tier_key"`
	// Revision increases monotonically per customer and starts at 1. Consumers
	// ignore a delivery whose Revision is not greater than the stored one.
	Revision  int64                                       `json:"revision"`
	Reason    membership_enums.MembershipTierChangeReason `json:"reason"`
	ChangedAt time.Time                                   `json:"changed_at"`
}
