package ads

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdBidding is how one platform object is asked to bid, together with the
// target that strategy needs.
//
// At most one target is set, and which one is decided by Strategy: a cost
// target for a cost-per-action goal, a bid cap for manual bidding, and a
// return target for a value goal. TargetReturnBasisPoints is a ratio in basis
// points, so 400 % is 40000; it is an integer because a return target must not
// drift through floating point. Money targets are in the ad account currency.
type AdBidding struct {
	Strategy                ads_enums.AdBidStrategy `json:"strategy,omitempty"`
	TargetCostPerAction     *money.Money            `json:"target_cost_per_action,omitempty"`
	BidCap                  *money.Money            `json:"bid_cap,omitempty"`
	TargetReturnBasisPoints int32                   `json:"target_return_basis_points,omitempty"`
}
