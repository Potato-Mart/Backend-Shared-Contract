package pricing_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/marketing/campaign/campaign_enums"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pubsub/pricing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pubsub/pricing/promotion_enums"
	event_enums "github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pubsub/routing"
)

func TestCustomerSafeStorefrontEventsJSON(t *testing.T) {
	now := time.Date(2026, 7, 29, 1, 2, 3, 0, time.UTC)
	values := []any{
		event.PromotionChangedEvent{
			PromotionID: "promotion_1", CampaignCode: "winter-sale",
			ScopeMode: geography_enums.GeographicScopeModeGlobal, ScopeRevision: 3,
			Published: true, Revision: 3, RefetchRequired: true, ChangedAt: now,
		},
		event.CampaignChangedEvent{
			CampaignCode: "winter-sale",
			Status:       campaign_enums.CampaignStatusPublished, Published: true,
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
		event_enums.EventTypeCouponChanged.String() != "coupon.changed" ||
		event_enums.EventTypeCampaignChanged.String() != "campaign.changed" {
		t.Fatal("storefront event topic/type values changed")
	}
}

func TestPromotionChangedPublicationContextVariantsJSON(t *testing.T) {
	now := time.Date(2026, 9, 24, 2, 3, 4, 0, time.UTC)
	legacyV2, err := json.Marshal(event.PromotionChangedEvent{
		PromotionID: "promotion_1", Published: true, Revision: 4,
		RefetchRequired: true, ChangedAt: now,
	})
	if err != nil {
		t.Fatalf("marshal legacy promotion.changed v2 payload: %v", err)
	}
	if strings.Contains(string(legacyV2), `"publication_context"`) {
		t.Fatalf("legacy promotion.changed v2 payload must omit publication_context: %s", legacyV2)
	}

	context := promotion_enums.PromotionPublicationContextCampaignPublishTogether
	v3Payload, err := json.Marshal(event.PromotionChangedEvent{
		PromotionID: "promotion_1", Published: true, Revision: 5,
		PublicationContext: context, RefetchRequired: true, ChangedAt: now,
	})
	if err != nil {
		t.Fatalf("marshal promotion.changed v3 payload: %v", err)
	}
	if !strings.Contains(string(v3Payload), `"publication_context":"campaign_publish_together"`) {
		t.Fatalf("promotion.changed v3 payload missing publication context: %s", v3Payload)
	}
	var decoded event.PromotionChangedEvent
	if err := json.Unmarshal(v3Payload, &decoded); err != nil {
		t.Fatalf("unmarshal promotion.changed v3 payload: %v", err)
	}
	if decoded.PublicationContext != context || !decoded.PublicationContext.IsValid() {
		t.Fatalf("promotion.changed publication context = %q, want %q", decoded.PublicationContext, context)
	}
	if promotion_enums.PromotionPublicationContext("unknown").IsValid() {
		t.Fatal("unknown promotion publication context must be invalid")
	}
}

func TestCouponChangedEventJSONIsIdentityAndRefetchOnly(t *testing.T) {
	now := time.Date(2026, 9, 24, 2, 3, 4, 0, time.UTC)
	value := event.CouponChangedEvent{
		CouponID: "opaque-coupon-id", Revision: 7, RefetchRequired: true, ChangedAt: now,
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal coupon.changed payload: %v", err)
	}
	for _, field := range []string{`"coupon_id":"opaque-coupon-id"`, `"revision":7`, `"refetch_required":true`, `"changed_at":`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("coupon.changed payload missing %s: %s", field, payload)
		}
	}
	for _, forbidden := range []string{"coupon_code", "terms", "scope", "customer", "recipient"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("coupon.changed payload leaked %q: %s", forbidden, payload)
		}
	}
	var decoded event.CouponChangedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal coupon.changed payload: %v", err)
	}
	if decoded.CouponID != value.CouponID || decoded.Revision != value.Revision || !decoded.RefetchRequired || !decoded.ChangedAt.Equal(now) {
		t.Fatalf("coupon.changed round trip = %+v, want %+v", decoded, value)
	}
}

func TestCampaignChangedHasNoRetiredIdentity(t *testing.T) {
	raw, err := json.Marshal(event.CampaignChangedEvent{CampaignCode: "CAM-1", Published: true, ActivationRevision: 3, Status: campaign_enums.CampaignStatusPublished})
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"series_code", "promotion_id", "is_active", "activation_revision"} {
		if strings.Contains(string(raw), `"`+field+`"`) {
			t.Fatal(field)
		}
	}
	for _, field := range []string{`"published":true`, `"publication_revision":3`, `"publication_status":"published"`} {
		if !strings.Contains(string(raw), field) {
			t.Fatal(field)
		}
	}
}
