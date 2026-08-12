package event_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography/geography_enums"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/customers/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/customers/campaign/campaign_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pubsub/event"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pubsub/event/event_enums"
)

func TestCustomerSafeStorefrontEventsJSON(t *testing.T) {
	now := time.Date(2026, 7, 29, 1, 2, 3, 0, time.UTC)
	values := []any{
		event.PromotionChangedEvent{
			PromotionID: "promotion_1", CampaignID: "campaign_1", CampaignKey: "winter-sale",
			ScopeMode: geography_enums.GeographicScopeModeGlobal, ScopeRevision: 3,
			Published: true, Revision: 3, RefetchRequired: true, ChangedAt: now,
		},
		event.CampaignChangedEvent{
			CampaignID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
			Status: campaign_enums.CampaignStatusActive, IsActive: true,
			ScopeMode: geography_enums.GeographicScopeModeGlobal, ScopeRevision: 5,
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

	if event_enums.EventTopicStorefrontEvents.String() != "storefront-events" ||
		event_enums.EventTypePromotionChanged.String() != "promotion.changed" ||
		event_enums.EventTypeCampaignChanged.String() != "campaign.changed" {
		t.Fatal("storefront event topic/type values changed")
	}
}

func TestCampaignLinkRevisionCTAAndMediaJSON(t *testing.T) {
	value := campaign.Campaign{
		ID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
		Title: "Winter sale", CTAHref: "potatomart://product/SKU-1",
		CTA: &campaign.CTADestination{
			Type: campaign_enums.CampaignCTADestinationProduct, ProductSKUCode: "SKU-1",
		},
		Media:            &security.ObjectMedia{ID: "media_1", URL: "/v1/storefront/campaigns/campaign_1/media"},
		Placement:        campaign_enums.CampaignPlacementHomeHero,
		Severity:         campaign_enums.CampaignSeverityInfo,
		Status:           campaign_enums.CampaignStatusActive,
		ScheduleTimezone: "Etc/UTC",
		GeographicScope:  geography.GeographicScope{Mode: geography_enums.GeographicScopeModeGlobal},
		Revision:         5, ActivationRevision: 2,
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal campaign: %v", err)
	}
	for _, field := range []string{`"promotion_id":"promotion_1"`, `"cta_href":"potatomart://product/SKU-1"`, `"cta":{"type":"product","product_sku_code":"SKU-1"}`, `"media":{"id":"media_1","url":"/v1/storefront/campaigns/campaign_1/media"}`, `"schedule_timezone":"Etc/UTC"`, `"geographic_scope":{"mode":"GLOBAL"}`, `"revision":5`, `"activation_revision":2`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("campaign missing %s: %s", field, payload)
		}
	}
	if strings.Contains(string(payload), `"region"`) {
		t.Fatalf("campaign retained removed audience region: %s", payload)
	}
	for _, legacy := range []string{`"media_id"`, `"media_url"`} {
		if strings.Contains(string(payload), legacy) {
			t.Fatalf("campaign retained legacy %s: %s", legacy, payload)
		}
	}
}
