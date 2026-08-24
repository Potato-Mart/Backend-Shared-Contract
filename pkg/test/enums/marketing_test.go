package enums_test

import (
	"testing"

	legacy_marketing_enums "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/insights/marketing/marketing_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/marketing/marketing_enums"
)

func TestMarketingEnumsValidateLockedV30Values(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "marketing.CampaignStatus", valid: []stringEnum{marketing_enums.CampaignStatusDraft, marketing_enums.CampaignStatusScheduled, marketing_enums.CampaignStatusActive, marketing_enums.CampaignStatusCompleted, marketing_enums.CampaignStatusArchived}, invalid: marketing_enums.CampaignStatus("__invalid__")},
		{name: "marketing.CampaignType", valid: []stringEnum{marketing_enums.CampaignTypeCoupon, marketing_enums.CampaignTypePromotion, marketing_enums.CampaignTypeMixed}, invalid: marketing_enums.CampaignType("__invalid__")},
		{name: "marketing.CampaignPlacement", valid: []stringEnum{marketing_enums.CampaignPlacementTopBanner, marketing_enums.CampaignPlacementHomeHero, marketing_enums.CampaignPlacementModal, marketing_enums.CampaignPlacementCheckoutNotice, marketing_enums.CampaignPlacementProductNotice}, invalid: marketing_enums.CampaignPlacement("__invalid__")},
		{name: "marketing.CampaignCustomerType", valid: []stringEnum{marketing_enums.CampaignCustomerTypeGuest, marketing_enums.CampaignCustomerTypeRetail, marketing_enums.CampaignCustomerTypeWholesale, marketing_enums.CampaignCustomerTypeAll}, invalid: marketing_enums.CampaignCustomerType("__invalid__")},
		{name: "marketing.CampaignPlatform", valid: []stringEnum{marketing_enums.CampaignPlatformWeb, marketing_enums.CampaignPlatformMobile, marketing_enums.CampaignPlatformAll}, invalid: marketing_enums.CampaignPlatform("__invalid__")},
		{name: "marketing.CouponType", valid: []stringEnum{marketing_enums.CouponTypePercentage, marketing_enums.CouponTypeFixedAmount, marketing_enums.CouponTypeFreeShipping}, invalid: marketing_enums.CouponType("__invalid__")},
		{name: "marketing.CouponScopeType", valid: []stringEnum{marketing_enums.CouponScopeTypeSKUCode, marketing_enums.CouponScopeTypeCollection, marketing_enums.CouponScopeTypeCategoryTag, marketing_enums.CouponScopeTypeProducts, marketing_enums.CouponScopeTypeMembershipTier, marketing_enums.CouponScopeTypeNewRegistration}, invalid: marketing_enums.CouponScopeType("__invalid__")},
		{name: "marketing.CouponStatus", valid: []stringEnum{marketing_enums.CouponStatusDraft, marketing_enums.CouponStatusScheduled, marketing_enums.CouponStatusActive, marketing_enums.CouponStatusCompleted, marketing_enums.CouponStatusArchived}, invalid: marketing_enums.CouponStatus("__invalid__")},
		{name: "marketing.PromotionType", valid: []stringEnum{marketing_enums.PromotionTypeProductDiscount, marketing_enums.PromotionTypeVolumeDiscount, marketing_enums.PromotionTypeGroupOrderDiscount, marketing_enums.PromotionTypeAddOnBundle, marketing_enums.PromotionTypeScopeBundle, marketing_enums.PromotionTypeBOGO, marketing_enums.PromotionTypeTieredPricing, marketing_enums.PromotionTypeMembershipPointMultiplier}, invalid: marketing_enums.PromotionType("__invalid__")},
		{name: "marketing.PromotionScopeType", valid: []stringEnum{marketing_enums.PromotionScopeTypeSKUCode, marketing_enums.PromotionScopeTypeCollection, marketing_enums.PromotionScopeTypeCategoryTag, marketing_enums.PromotionScopeTypeProducts, marketing_enums.PromotionScopeTypeMembershipTier}, invalid: marketing_enums.PromotionScopeType("__invalid__")},
		{name: "marketing.PromotionStatus", valid: []stringEnum{marketing_enums.PromotionStatusDraft, marketing_enums.PromotionStatusScheduled, marketing_enums.PromotionStatusActive, marketing_enums.PromotionStatusCompleted, marketing_enums.PromotionStatusArchived}, invalid: marketing_enums.PromotionStatus("__invalid__")},
		{name: "marketing.DiscountType", valid: []stringEnum{marketing_enums.DiscountTypePercentage, marketing_enums.DiscountTypeFixedAmount, marketing_enums.DiscountTypeFixedPackagePrice, marketing_enums.DiscountTypeFreeShipping}, invalid: marketing_enums.DiscountType("__invalid__")},
		{name: "marketing.MarketingChannel", valid: []stringEnum{marketing_enums.MarketingChannelEmail, marketing_enums.MarketingChannelSMS, marketing_enums.MarketingChannelLine, marketing_enums.MarketingChannelExport}, invalid: marketing_enums.MarketingChannel("__invalid__")},
		{name: "marketing.MarketingMessageStatus", valid: []stringEnum{marketing_enums.MarketingMessageStatusDraft, marketing_enums.MarketingMessageStatusScheduled, marketing_enums.MarketingMessageStatusSending, marketing_enums.MarketingMessageStatusSent, marketing_enums.MarketingMessageStatusPartial, marketing_enums.MarketingMessageStatusFailed, marketing_enums.MarketingMessageStatusCancelled, marketing_enums.MarketingMessageStatusExported}, invalid: marketing_enums.MarketingMessageStatus("__invalid__")},
		{name: "marketing.TemplateStatus", valid: []stringEnum{marketing_enums.TemplateStatusActive, marketing_enums.TemplateStatusArchived}, invalid: marketing_enums.TemplateStatus("__invalid__")},
	})
}

func TestLegacyMarketingEnumsRemainCoveredDuringV30Foundation(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "insights.marketing.MarketingChannel", valid: []stringEnum{legacy_marketing_enums.MarketingChannelEmail, legacy_marketing_enums.MarketingChannelSMS, legacy_marketing_enums.MarketingChannelLine, legacy_marketing_enums.MarketingChannelExport}, invalid: legacy_marketing_enums.MarketingChannel("__invalid__")},
		{name: "insights.marketing.MarketingCampaignStatus", valid: []stringEnum{legacy_marketing_enums.MarketingCampaignStatusDraft, legacy_marketing_enums.MarketingCampaignStatusSending, legacy_marketing_enums.MarketingCampaignStatusSent, legacy_marketing_enums.MarketingCampaignStatusPartial, legacy_marketing_enums.MarketingCampaignStatusFailed, legacy_marketing_enums.MarketingCampaignStatusCancelled, legacy_marketing_enums.MarketingCampaignStatusExported}, invalid: legacy_marketing_enums.MarketingCampaignStatus("__invalid__")},
		{name: "insights.marketing.MarketingRecipientStatus", valid: []stringEnum{legacy_marketing_enums.MarketingRecipientStatusPending, legacy_marketing_enums.MarketingRecipientStatusSent, legacy_marketing_enums.MarketingRecipientStatusDelivered, legacy_marketing_enums.MarketingRecipientStatusBounced, legacy_marketing_enums.MarketingRecipientStatusOpened, legacy_marketing_enums.MarketingRecipientStatusClicked, legacy_marketing_enums.MarketingRecipientStatusFailed, legacy_marketing_enums.MarketingRecipientStatusUnsubscribed}, invalid: legacy_marketing_enums.MarketingRecipientStatus("__invalid__")},
	})
}
