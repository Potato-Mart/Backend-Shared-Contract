package ads

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdAttributedOrder is one paid order credited to an ad, with the ad it came
// from and how long the person took to buy.
//
// It carries no contact details and no click identifiers: no email, phone,
// address, click id or browser cookie value ever appears here. CustomerNumber
// is the only customer identity, and is empty for a guest order. DisplayName
// is the already-masked name, present only when the caller is also allowed to
// see customer names; it is never a substitute for the customer number.
// Revenue is the attributed revenue for the order, net of tax and discount and
// excluding shipping, as a list by currency so that markets are never summed.
// ClickToOrderSeconds is the gap between ClickedAt and PaidAt, carried so a
// reader does not have to subtract two instants itself.
type AdAttributedOrder struct {
	OrderNumber         string                     `json:"order_number,omitempty"`
	PaidAt              *time.Time                 `json:"paid_at,omitempty"`
	CustomerNumber      string                     `json:"customer_number,omitempty"`
	DisplayName         string                     `json:"display_name,omitempty"`
	IsNewCustomer       bool                       `json:"is_new_customer,omitempty"`
	Revenue             []money.Money              `json:"revenue,omitempty"`
	Platform            ads_enums.AdPlatform       `json:"platform,omitempty"`
	SourcePlatform      ads_enums.AdSourcePlatform `json:"source_platform,omitempty"`
	Campaign            *AdEntityRef               `json:"campaign,omitempty"`
	AdSet               *AdEntityRef               `json:"ad_set,omitempty"`
	Ad                  *AdEntityRef               `json:"ad,omitempty"`
	ClickedAt           *time.Time                 `json:"clicked_at,omitempty"`
	ClickToOrderSeconds int64                      `json:"click_to_order_seconds,omitempty"`
}
