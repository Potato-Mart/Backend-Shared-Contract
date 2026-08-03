package campaign_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/campaign"
	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/contracts/promotion"
	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/campaign"
	eventsenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/events"
	geographyenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/geography"
)

func TestCustomerSafeStorefrontEventsJSON(t *testing.T) {
	now := time.Date(2026, 7, 29, 1, 2, 3, 0, time.UTC)
	values := []any{
		promotion.PromotionChangedEvent{
			PromotionID: "promotion_1", CampaignID: "campaign_1", CampaignKey: "winter-sale",
			ScopeMode: geographyenum.GeographicScopeModeGlobal, ScopeRevision: 3,
			Published: true, Revision: 3, RefetchRequired: true, ChangedAt: now,
		},
		campaign.CampaignChangedEvent{
			CampaignID: "campaign_1", CampaignKey: "winter-sale", PromotionID: "promotion_1",
			Status: campaignenum.CampaignStatusActive, IsActive: true,
			ScopeMode: geographyenum.GeographicScopeModeGlobal, ScopeRevision: 5,
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
		Placement:        campaignenum.CampaignPlacementHomeHero,
		Severity:         campaignenum.CampaignSeverityInfo,
		Status:           campaignenum.CampaignStatusActive,
		ScheduleTimezone: "Etc/UTC",
		GeographicScope:  common.GeographicScope{Mode: geographyenum.GeographicScopeModeGlobal},
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
