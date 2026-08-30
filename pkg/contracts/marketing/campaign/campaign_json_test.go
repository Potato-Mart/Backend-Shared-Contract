package campaign_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/marketing/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/marketing/campaign/campaign_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/benefit"
)

func TestCampaignUsesOpenNotificationTopicAndPricingBenefitReference(t *testing.T) {
	payload, err := json.Marshal(campaign.Campaign{
		CampaignCode: "spring-2026", MarketCode: "au", Title: []localization.LocalizedName{{Language: "en-AU", Name: "Spring"}},
		Placement: campaign_enums.CampaignPlacementHomeHero, Severity: campaign_enums.CampaignSeverityInfo,
		Status: campaign_enums.CampaignStatusActive, NotificationTopicCode: "seasonal_offer",
		BenefitRefs: []benefit.BenefitRef{{Kind: "promotion", Code: "spring-price"}},
	})
	if err != nil {
		t.Fatalf("marshal campaign: %v", err)
	}
	for _, expected := range []string{`"campaign_code":"spring-2026"`, `"notification_topic_code":"seasonal_offer"`, `"kind":"promotion"`} {
		if !strings.Contains(string(payload), expected) {
			t.Fatalf("campaign missing %s: %s", expected, payload)
		}
	}
	for _, forbidden := range []string{`messaging_category`, `promotion_id`} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("campaign retained retired field %s: %s", forbidden, payload)
		}
	}
}
