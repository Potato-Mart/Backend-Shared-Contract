package marketing_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/commerce/commerce_enums"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/marketing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/marketing/marketing_enums"
)

func TestCampaignUsesPublicBenefitReferencesAndSafeMedia(t *testing.T) {
	now := time.Date(2026, 8, 13, 10, 0, 0, 0, time.UTC)
	campaign := marketing.Campaign{
		CampaignCode: "spring-launch",
		Title:        []localization.LocalizedName{{Language: "en", Name: "Spring launch"}},
		Cover:        &security.ObjectMedia{Code: "media_campaign_1", URL: "https://cdn.example.test/campaign.png"},
		CampaignDetail: marketing.CampaignDetail{
			Message:      []localization.LocalizedText{{Language: "en", Text: "Fresh offers are here"}},
			CampaignType: marketing_enums.CampaignTypeMixed,
			CouponDetails: []marketing.BenefitRef{{
				Code: "WELCOME10", Name: []localization.LocalizedName{{Language: "en", Name: "Welcome 10"}}, Path: "/promotions/spring-launch",
			}},
			PromotionDetails: []marketing.BenefitRef{{
				Code: "SPRING-BUNDLE", Name: []localization.LocalizedName{{Language: "en", Name: "Spring bundle"}}, Path: "/promotions/spring-launch",
			}},
		},
		CampaignPosition: marketing.CampaignPosition{
			Placement: marketing_enums.CampaignPlacementHomeHero,
			GeographicScope: geography.GeographicScope{
				Mode: geography_enums.GeographicScopeModeGlobal,
			},
			ScheduleTimezone: "Australia/Sydney",
		},
		CampaignStatus: marketing.CampaignStatus{
			Status: marketing_enums.CampaignStatusScheduled, Dismissible: true, StartsAt: &now,
		},
		Audience: marketing.Audience{
			CustomerType: marketing_enums.CampaignCustomerTypeAll,
			Platform:     marketing_enums.CampaignPlatformAll,
		},
		AuditFields: audit.AuditFields{CreatedAt: now, UpdatedAt: now},
	}

	payload, err := json.Marshal(campaign)
	if err != nil {
		t.Fatalf("marshal campaign: %v", err)
	}
	for _, expected := range []string{
		`"campaign_code":"spring-launch"`,
		`"campaign_type":"mixed"`,
		`"coupon_details":[{"code":"WELCOME10"`,
		`"promotion_details":[{"code":"SPRING-BUNDLE"`,
		`"path":"/promotions/spring-launch"`,
		`"cover":{"code":"media_campaign_1","url":"https://cdn.example.test/campaign.png"}`,
		`"customer_type":"all"`,
		`"platform":"all"`,
	} {
		if !strings.Contains(string(payload), expected) {
			t.Fatalf("campaign JSON missing %s: %s", expected, payload)
		}
	}
	if got := strings.Count(string(payload), `"path":"/promotions/spring-launch"`); got != 2 {
		t.Fatalf("campaign benefit paths = %d, want 2 Campaign landing paths: %s", got, payload)
	}
	for _, forbidden := range []string{"bucket", "storage_path", "coupon_id", "promotion_id", "discount_value", "/v1/coupons/", "/v1/promotions/"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("campaign JSON leaked %q: %s", forbidden, payload)
		}
	}

	var decoded marketing.Campaign
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal campaign: %v", err)
	}
	if decoded.CampaignDetail.CouponDetails[0].Code != "WELCOME10" || decoded.Cover == nil || decoded.Cover.Code != "media_campaign_1" {
		t.Fatalf("campaign did not round-trip: %+v", decoded)
	}
}

func TestCouponPromotionMessageAndTemplatesRoundTrip(t *testing.T) {
	now := time.Date(2026, 8, 13, 10, 0, 0, 0, time.UTC)
	usageLimit := 100
	perCustomerLimit := 2
	minimumUnits := int64(3)
	basisPoints := int64(1_000)
	bogoQuantity := 1
	amount := money.Money{AmountMinor: 500, Currency: "AUD"}
	scope := geography.GeographicScope{Mode: geography_enums.GeographicScopeModeGlobal}
	detail := marketing.ScopeDetail{
		DiscountType:       marketing_enums.DiscountTypePercentage,
		DiscountValue:      marketing.DiscountValue{BasisPoints: &basisPoints},
		MinimumOrderAmount: &amount,
		MinimumUnits:       &minimumUnits,
	}
	coupon := marketing.Coupon{
		CouponCode:  "WELCOME10",
		CouponName:  []localization.LocalizedName{{Language: "en", Name: "Welcome 10"}},
		CouponCover: &security.ObjectMedia{Code: "media_coupon_1"},
		CouponDetail: marketing.CouponDetail{
			Description: []localization.LocalizedDescription{{Language: "en", Description: "A welcome benefit"}},
			CouponType:  marketing_enums.CouponTypePercentage,
		},
		CouponScope: marketing.CouponScope{
			ScopeType:   marketing_enums.CouponScopeTypeProducts,
			Targets:     []marketing.ScopeTarget{{Code: "SKU-1", Name: []localization.LocalizedName{{Language: "en", Name: "Potatoes"}}}},
			ScopeDetail: detail,
		},
		CouponStatus:     marketing.CouponStatus{Status: marketing_enums.CouponStatusActive, StartsAt: &now},
		CouponPosition:   marketing.CouponPosition{GeographicScope: scope, ScheduleTimezone: "Australia/Sydney"},
		CouponConditions: marketing.CouponConditions{UsageLimit: &usageLimit, Stackable: false, Channels: []commerce_enums.OrderType{commerce_enums.OrderTypeOnline}},
		AuditFields:      audit.AuditFields{CreatedAt: now, UpdatedAt: now},
	}
	promotion := marketing.Promotion{
		PromotionCode: "SPRING-BUNDLE",
		PromotionName: []localization.LocalizedName{{Language: "en", Name: "Spring bundle"}},
		Cover:         &security.ObjectMedia{Code: "media_promotion_1"},
		PromotionDetail: marketing.PromotionDetail{
			Description:   []localization.LocalizedDescription{{Language: "en", Description: "Buy one, get one"}},
			PromotionType: marketing_enums.PromotionTypeBOGO,
		},
		PromotionScope: marketing.PromotionScope{
			ScopeType:   marketing_enums.PromotionScopeTypeSKUCode,
			Targets:     []marketing.ScopeTarget{{Code: "SKU-1", Name: []localization.LocalizedName{{Language: "en", Name: "Potatoes"}}}},
			ScopeDetail: detail,
		},
		PromotionStatus:   marketing.PromotionStatus{Status: marketing_enums.PromotionStatusActive, StartsAt: &now},
		PromotionPosition: marketing.PromotionPosition{GeographicScope: scope, ScheduleTimezone: "Australia/Sydney"},
		PromotionConditions: marketing.PromotionConditions{
			UsageLimit: &usageLimit, PerCustomerLimit: &perCustomerLimit, Stackable: true, Channels: []commerce_enums.OrderType{commerce_enums.OrderTypeOnline},
		},
		ScopeRelations: marketing.ScopeRelations{
			Targets:          []marketing.ScopeTarget{{Code: "SKU-2", Name: []localization.LocalizedName{{Language: "en", Name: "Free potatoes"}}}},
			RequiredProducts: []marketing.RequiredProduct{{Code: "SKU-1", Name: []localization.LocalizedName{{Language: "en", Name: "Potatoes"}}, Quantity: 1}},
			Tiers:            []marketing.PromotionTier{{MinimumUnits: 3, ScopeDetail: detail}},
			MixTargets:       boolPointer(true),
			BOGOGetQuantity:  &bogoQuantity,
		},
		AuditFields: audit.AuditFields{CreatedAt: now, UpdatedAt: now},
	}
	message := marketing.MarketingMessage{
		Code:             "msg-spring-launch",
		CampaignCode:     "spring-launch",
		CampaignName:     []localization.LocalizedName{{Language: "en", Name: "Spring launch"}},
		Channel:          marketing_enums.MarketingChannelEmail,
		Subject:          []localization.LocalizedText{{Language: "en", Text: "Spring is here"}},
		Message:          []localization.LocalizedText{{Language: "en", Text: "Fresh offers are ready"}},
		Images:           []security.ObjectMedia{{Code: "media_message_1", URL: "https://cdn.example.test/message.png"}},
		Status:           marketing_enums.MarketingMessageStatusScheduled,
		GeographicScope:  scope,
		ScheduleTimezone: "Australia/Sydney",
		ScheduledSendAt:  &now,
		AuditFields:      audit.AuditFields{CreatedAt: now, UpdatedAt: now},
	}

	values := []any{
		coupon,
		promotion,
		message,
		marketing.CampaignTemplate{TemplateCode: "campaign-spring", Version: 2, Status: marketing_enums.TemplateStatusActive, Payload: marketing.Campaign{CampaignCode: "spring-launch"}, AuditFields: audit.AuditFields{CreatedAt: now, UpdatedAt: now}},
		marketing.MarketingMessageTemplate{TemplateCode: "message-spring", Version: 3, Status: marketing_enums.TemplateStatusArchived, Payload: message, AuditFields: audit.AuditFields{CreatedAt: now, UpdatedAt: now}},
	}
	for _, value := range values {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal %T: %v", value, err)
		}
		for _, expected := range []string{`"geographic_scope":{"mode":"GLOBAL"}`, `"schedule_timezone":"Australia/Sydney"`} {
			if strings.Contains(string(payload), `"template_code"`) {
				continue
			}
			if !strings.Contains(string(payload), expected) {
				t.Fatalf("%T JSON missing %s: %s", value, expected, payload)
			}
		}
		for _, forbidden := range []string{"recipient", "customer_number", "bucket", "storage_path"} {
			if strings.Contains(string(payload), forbidden) {
				t.Fatalf("%T JSON leaked %q: %s", value, forbidden, payload)
			}
		}
	}
}

func boolPointer(value bool) *bool { return &value }

func TestMarketingAggregateEventsArePIIFree(t *testing.T) {
	now := time.Date(2026, 8, 13, 10, 0, 0, 0, time.UTC)
	values := []any{
		marketing.CampaignChangedEvent{CampaignCode: "spring-launch", Status: marketing_enums.CampaignStatusActive, ChangedAt: now},
		marketing.CouponChangedEvent{CouponCode: "WELCOME10", Status: marketing_enums.CouponStatusActive, ChangedAt: now},
		marketing.PromotionChangedEvent{PromotionCode: "SPRING-BUNDLE", Status: marketing_enums.PromotionStatusActive, ChangedAt: now},
		marketing.MarketingMessageChangedEvent{Code: "msg-spring-launch", CampaignCode: "spring-launch", Status: marketing_enums.MarketingMessageStatusScheduled, ChangedAt: now},
	}
	for _, value := range values {
		payload, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal %T: %v", value, err)
		}
		for _, forbidden := range []string{"name", "subject", "message", "recipient", "email", "phone", "provider"} {
			if strings.Contains(string(payload), forbidden) {
				t.Fatalf("%T event leaked %q: %s", value, forbidden, payload)
			}
		}
	}
}
