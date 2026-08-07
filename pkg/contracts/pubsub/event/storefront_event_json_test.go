package event_test

import (
	"encoding/json"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/customers/campaign"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/pubsub/event"
	"strings"
	"testing"
	"time"
)

func TestCustomerSafeStorefrontEventsJSON(t *testing.T) {
	now := time.Date(2026, 7, 29, 1, 2, 3, 0, time.UTC)
	values := []any{
		event.PromotionChangedEvent{
			PromotionID: "promotion_1", CampaignID: "campaign_1", CampaignKey: "winter-sale",
			ScopeMode: geography.GeographicScopeModeGlobal, ScopeRevision: 3,
			Published: true, Revision: 3, RefetchRequired: true, ChangedAt: now,
		},
		event.CampaignChangedEvent{
			CampaignID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
			Status: campaign.CampaignStatusActive, IsActive: true,
			ScopeMode: geography.GeographicScopeModeGlobal, ScopeRevision: 5,
			ActivationRevision: 2, ContentRevision: 5, RefetchRequired: true, ChangedAt: now,
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
		for _, field := range []string{`"scope_mode":"GLOBAL"`, `"scope_revision":`, `"refetch_required":true`} {
			if !strings.Contains(string(payload), field) {
				t.Fatalf("storefront event missing %s: %s", field, payload)
			}
		}
		if strings.Contains(string(payload), `"targets"`) {
			t.Fatalf("storefront event exposed geographic targets: %s", payload)
		}
	}

	if event.EventTopicStorefrontEvents.String() != "storefront-events" ||
		event.EventTypePromotionChanged.String() != "promotion.changed" ||
		event.EventTypeCampaignChanged.String() != "campaign.changed" {
		t.Fatal("storefront event topic/type values changed")
	}
}

func TestCampaignLinkRevisionCTAAndMediaJSON(t *testing.T) {
	value := campaign.Campaign{
		ID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
		Title: "Winter sale", CTAHref: "potatomart://product/SKU-1",
		CTA: &campaign.CTADestination{
			Type: campaign.CampaignCTADestinationProduct, ProductSKUCode: "SKU-1",
		},
		MediaID: "media_1", MediaURL: "/v1/storefront/campaigns/campaign_1/media",
		Placement:        campaign.CampaignPlacementHomeHero,
		Severity:         campaign.CampaignSeverityInfo,
		Status:           campaign.CampaignStatusActive,
		ScheduleTimezone: "Etc/UTC",
		GeographicScope:  geography.GeographicScope{Mode: geography.GeographicScopeModeGlobal},
		Revision:         5, ActivationRevision: 2,
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal campaign: %v", err)
	}
	for _, field := range []string{`"promotion_id":"promotion_1"`, `"cta_href":"potatomart://product/SKU-1"`, `"cta":{"type":"product","product_sku_code":"SKU-1"}`, `"media_id":"media_1"`, `"schedule_timezone":"Etc/UTC"`, `"geographic_scope":{"mode":"GLOBAL"}`, `"revision":5`, `"activation_revision":2`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("campaign missing %s: %s", field, payload)
		}
	}
	if strings.Contains(string(payload), `"region"`) {
		t.Fatalf("campaign retained removed audience region: %s", payload)
	}
}
