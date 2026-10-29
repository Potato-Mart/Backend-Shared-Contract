package campaign_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/marketing/campaign"
	ce "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/marketing/campaign/campaign_enums"
	"reflect"
	"testing"

	audienceenum "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/marketing/audience/audience_enums"
)

func TestCampaignBoundaryAndRoundTrip(t *testing.T) {
	c := campaign.Campaign{CampaignCode: "CAM-1", Audience: []ce.CampaignCustomerType{ce.CampaignCustomerTypeRetail, ce.CampaignCustomerTypeWholesale}, Placements: []campaign.Placement{{Position: ce.CampaignPlacementHomeBanner, Platforms: []ce.CampaignPlatform{ce.CampaignPlatformWeb}, Priority: 2}, {Position: ce.CampaignPlacementAccountBanner, Platforms: []ce.CampaignPlatform{ce.CampaignPlatformMobile}, Priority: 9}}, Offers: []campaign.OfferLink{{Kind: "promotion", ID: "PROMO-1", DisplayOrder: 2, Revision: "r7"}}, Status: ce.CampaignStatusDraft}
	raw, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	var decoded campaign.Campaign
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}
	var fields map[string]json.RawMessage
	_ = json.Unmarshal(raw, &fields)
	for _, retired := range []string{"series_code", "placement", "priority", "severity", "message_style", "cta", "cta_text", "cta_href", "notification_topic_code", "benefit_refs", "is_active"} {
		if _, exists := fields[retired]; exists {
			t.Fatalf("retired field %s remains", retired)
		}
	}
	if ce.CampaignCustomerType("guest").IsValid() {
		t.Fatal("guest is not a campaign audience")
	}
	if !audienceenum.CustomerTypeGuest.IsValid() {
		t.Fatal("offer guest semantics must stay separate")
	}
	if reflect.TypeOf(audienceenum.CustomerTypeGuest) == reflect.TypeOf(ce.CampaignCustomerTypeRetail) {
		t.Fatal("campaign and commercial audience share a type")
	}
	for _, position := range []ce.CampaignPlacement{"announcement_bar", "home_banner", "account_banner", "home_modal"} {
		if !position.IsValid() {
			t.Fatal(position)
		}
	}
	for _, state := range []ce.CampaignStatus{"draft", "published", "paused", "archived"} {
		if !state.IsValid() {
			t.Fatal(state)
		}
	}
}
