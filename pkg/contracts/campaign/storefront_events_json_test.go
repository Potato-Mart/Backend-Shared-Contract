package campaign_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/promotion"
	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/campaign"
	eventsenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/events"
)

func TestCustomerSafeStorefrontEventsJSON(t *testing.T) {
	now := time.Date(2026, 7, 29, 1, 2, 3, 0, time.UTC)
	values := []any{
		promotion.PromotionChangedEvent{
			PromotionID: "promotion_1", CampaignID: "campaign_1", CampaignKey: "winter-sale",
			Published: true, Revision: 3, ChangedAt: now,
		},
		campaign.CampaignChangedEvent{
			CampaignID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
			Status: campaignenum.CampaignStatusActive, IsActive: true,
			ActivationRevision: 2, ContentRevision: 5, ChangedAt: now,
		},
	}
	for _, value := range values {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal storefront event: %v", err)
		}
		for _, forbidden := range []string{"discount_value", "recipient", "title", "message", "provider"} {
			if strings.Contains(string(payload), forbidden) {
				t.Fatalf("storefront event leaked %q: %s", forbidden, payload)
			}
		}
	}

	if eventsenum.EventTopicStorefrontEvents.String() != "storefront-events" ||
		eventsenum.EventTypePromotionChanged.String() != "promotion.changed" ||
		eventsenum.EventTypeCampaignChanged.String() != "campaign.changed" {
		t.Fatal("storefront event topic/type values changed")
	}
}

func TestCampaignLinkRevisionCTAAndMediaJSON(t *testing.T) {
	value := campaign.Campaign{
		ID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
		Title: "Winter sale", CTAHref: "potatomart://product/SKU-1",
		CTA: &campaign.CTADestination{
			Type: campaignenum.CampaignCTADestinationProduct, ProductSKUCode: "SKU-1",
		},
		MediaID: "media_1", MediaURL: "/v1/storefront/campaigns/campaign_1/media",
		Placement: campaignenum.CampaignPlacementHomeHero,
		Severity:  campaignenum.CampaignSeverityInfo,
		Status:    campaignenum.CampaignStatusActive,
		Revision:  5, ActivationRevision: 2,
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal v21 campaign: %v", err)
	}
	for _, field := range []string{`"promotion_id":"promotion_1"`, `"cta_href":"potatomart://product/SKU-1"`, `"cta":{"type":"product","product_sku_code":"SKU-1"}`, `"media_id":"media_1"`, `"revision":5`, `"activation_revision":2`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("v21 campaign missing %s: %s", field, payload)
		}
	}
}
