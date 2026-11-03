package campaign_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/campaign"
	ce "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/campaign/campaign_enums"
	"reflect"
	"testing"

	audienceenum "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/marketing/audience/audience_enums"
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
	if _, exists := fields["card_media"]; exists {
		t.Fatal("absent card_media must be omitted")
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

func TestCampaignCardMediaIsCampaignLevelAndOptional(t *testing.T) {
	c := campaign.Campaign{
		CampaignCode: "CAM-1",
		Media:        &security.ObjectMedia{Code: "cover-16x9", URL: "/media/cover"},
		CardMedia:    &security.ObjectMedia{Code: "card-9x16", URL: "/media/card"},
		Placements:   []campaign.Placement{{Position: ce.CampaignPlacementHomeBanner, Platforms: []ce.CampaignPlatform{ce.CampaignPlatformMobile}}},
		Offers:       []campaign.OfferLink{},
		Audience:     []ce.CampaignCustomerType{ce.CampaignCustomerTypeRetail},
	}
	raw, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	if string(fields["card_media"]) != `{"code":"card-9x16","url":"/media/card"}` || string(fields["media"]) != `{"code":"cover-16x9","url":"/media/cover"}` {
		t.Fatalf("card_media or media wire shape changed: %s", raw)
	}
	var decoded campaign.Campaign
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c, decoded) {
		t.Fatalf("round trip changed model: %+v", decoded)
	}
	if _, onPlacement := reflect.TypeOf(campaign.Placement{}).FieldByName("CardMedia"); onPlacement {
		t.Fatal("card media is campaign-level, not a placement field")
	}

	// JSON written before card_media existed still decodes, leaving it absent
	// so apps fall back to media.
	var legacy campaign.Campaign
	if err := json.Unmarshal([]byte(`{"campaign_code":"CAM-0","media":{"code":"cover-16x9"},"placements":[],"offers":[],"published":false,"audience":[],"publication_status":"draft","effective_status":"","revision":0,"publication_revision":0}`), &legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.CardMedia != nil || legacy.Media == nil {
		t.Fatal("legacy campaign must keep media and leave card_media absent")
	}
}
