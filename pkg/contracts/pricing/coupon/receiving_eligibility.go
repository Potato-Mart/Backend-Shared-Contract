package coupon

import "github.com/Potato-Mart/Backend-Shared-Contract/v47/pkg/contracts/common/geography"

type ReceivingEligibility struct {
	Membership      ReceivingTierPolicy          `json:"membership"`
	Profile         ProfileCompletionRestriction `json:"profile"`
	GeographicScope *geography.GeographicScope   `json:"geographic_scope,omitempty"`
}
