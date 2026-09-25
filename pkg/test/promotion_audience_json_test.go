package pkg_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/commerce/commerce_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/marketing/audience"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/pricing/coupon"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/pricing/promotion"
)

func TestPromotionAndCouponAudienceRoundTripWithoutChangingChannels(t *testing.T) {
	for _, tc := range []struct {
		name     string
		audience *audience.Audience
		wire     string
	}{
		{name: "legacy absent"},
		{name: "explicit unrestricted", audience: &audience.Audience{}, wire: `{}`},
		{name: "customer only", audience: &audience.Audience{CustomerType: campaign_enums.CampaignCustomerTypeRetail}, wire: `{"customer_type":"retail"}`},
		{name: "platform only", audience: &audience.Audience{Platform: campaign_enums.CampaignPlatformMobile}, wire: `{"platform":"mobile"}`},
		{name: "both dimensions", audience: &audience.Audience{CustomerType: campaign_enums.CampaignCustomerTypeWholesale, Platform: campaign_enums.CampaignPlatformWeb}, wire: `{"customer_type":"wholesale","platform":"web"}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			controls := promotion.PromotionControls{
				Channels: []commerce_enums.OrderType{commerce_enums.OrderTypeOnline},
				Audience: tc.audience,
			}
			for name, model := range map[string]any{
				"promotion": promotion.Promotion{Controls: controls},
				"coupon":    coupon.Coupon{Controls: controls},
			} {
				t.Run(name, func(t *testing.T) {
					data, err := json.Marshal(model)
					if err != nil {
						t.Fatal(err)
					}
					var wire struct {
						Controls map[string]json.RawMessage `json:"controls"`
					}
					if err := json.Unmarshal(data, &wire); err != nil {
						t.Fatal(err)
					}
					if got := string(wire.Controls["audience"]); got != tc.wire {
						t.Fatalf("audience = %s, want %s", got, tc.wire)
					}
					if got := string(wire.Controls["channels"]); got != `["online"]` {
						t.Fatalf("audience changed order channels: %s", got)
					}
					decoded := reflect.New(reflect.TypeOf(model))
					if err := json.Unmarshal(data, decoded.Interface()); err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(decoded.Elem().Interface(), model) {
						t.Fatalf("offer audience did not round-trip: %s", data)
					}
				})
			}
		})
	}
}
