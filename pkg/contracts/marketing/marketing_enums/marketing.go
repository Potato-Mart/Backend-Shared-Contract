// Package marketing_enums defines the finite wire values used by the
// canonical marketing aggregates.
package marketing_enums

// CampaignStatus is the lifecycle of a campaign aggregate.
type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusActive    CampaignStatus = "active"
	CampaignStatusCompleted CampaignStatus = "completed"
	CampaignStatusArchived  CampaignStatus = "archived"
)

// IsValid reports whether s is a supported campaign status.
func (s CampaignStatus) IsValid() bool {
	switch s {
	case CampaignStatusDraft, CampaignStatusScheduled, CampaignStatusActive,
		CampaignStatusCompleted, CampaignStatusArchived:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s CampaignStatus) String() string { return string(s) }

// CampaignType identifies which benefit families a campaign presents.
type CampaignType string

const (
	CampaignTypeCoupon    CampaignType = "coupon"
	CampaignTypePromotion CampaignType = "promotion"
	CampaignTypeMixed     CampaignType = "mixed"
)

// IsValid reports whether c is a supported campaign type.
func (c CampaignType) IsValid() bool {
	switch c {
	case CampaignTypeCoupon, CampaignTypePromotion, CampaignTypeMixed:
		return true
	default:
		return false
	}
}

// String returns the wire value for c.
func (c CampaignType) String() string { return string(c) }

// CampaignPlacement identifies where a campaign renders in a storefront.
type CampaignPlacement string

const (
	CampaignPlacementTopBanner      CampaignPlacement = "top_banner"
	CampaignPlacementHomeHero       CampaignPlacement = "home_hero"
	CampaignPlacementModal          CampaignPlacement = "modal"
	CampaignPlacementCheckoutNotice CampaignPlacement = "checkout_notice"
	CampaignPlacementProductNotice  CampaignPlacement = "product_notice"
)

// IsValid reports whether p is a supported campaign placement.
func (p CampaignPlacement) IsValid() bool {
	switch p {
	case CampaignPlacementTopBanner, CampaignPlacementHomeHero, CampaignPlacementModal,
		CampaignPlacementCheckoutNotice, CampaignPlacementProductNotice:
		return true
	default:
		return false
	}
}

// String returns the wire value for p.
func (p CampaignPlacement) String() string { return string(p) }

// CampaignCustomerType identifies the customer population for a campaign.
type CampaignCustomerType string

const (
	CampaignCustomerTypeGuest     CampaignCustomerType = "guest"
	CampaignCustomerTypeRetail    CampaignCustomerType = "retail"
	CampaignCustomerTypeWholesale CampaignCustomerType = "wholesale"
	CampaignCustomerTypeAll       CampaignCustomerType = "all"
)

// IsValid reports whether c is a supported campaign customer type.
func (c CampaignCustomerType) IsValid() bool {
	switch c {
	case CampaignCustomerTypeGuest, CampaignCustomerTypeRetail,
		CampaignCustomerTypeWholesale, CampaignCustomerTypeAll:
		return true
	default:
		return false
	}
}

// String returns the wire value for c.
func (c CampaignCustomerType) String() string { return string(c) }

// CampaignPlatform identifies the client platform for a campaign.
type CampaignPlatform string

const (
	CampaignPlatformWeb    CampaignPlatform = "web"
	CampaignPlatformMobile CampaignPlatform = "mobile"
	CampaignPlatformAll    CampaignPlatform = "all"
)

// IsValid reports whether p is a supported campaign platform.
func (p CampaignPlatform) IsValid() bool {
	switch p {
	case CampaignPlatformWeb, CampaignPlatformMobile, CampaignPlatformAll:
		return true
	default:
		return false
	}
}

// String returns the wire value for p.
func (p CampaignPlatform) String() string { return string(p) }

// CouponType identifies the customer benefit granted by a coupon.
type CouponType string

const (
	CouponTypePercentage   CouponType = "percentage"
	CouponTypeFixedAmount  CouponType = "fixed_amount"
	CouponTypeFreeShipping CouponType = "free_shipping"
)

// IsValid reports whether c is a supported coupon type.
func (c CouponType) IsValid() bool {
	switch c {
	case CouponTypePercentage, CouponTypeFixedAmount, CouponTypeFreeShipping:
		return true
	default:
		return false
	}
}

// String returns the wire value for c.
func (c CouponType) String() string { return string(c) }

// CouponScopeType identifies the target family selected by a coupon.
type CouponScopeType string

const (
	CouponScopeTypeSKUCode         CouponScopeType = "sku_code"
	CouponScopeTypeCollection      CouponScopeType = "collection"
	CouponScopeTypeCategoryTag     CouponScopeType = "category_tag"
	CouponScopeTypeProducts        CouponScopeType = "products"
	CouponScopeTypeMembershipTier  CouponScopeType = "membership_tier"
	CouponScopeTypeNewRegistration CouponScopeType = "new_registration"
)

// IsValid reports whether s is a supported coupon scope type.
func (s CouponScopeType) IsValid() bool {
	switch s {
	case CouponScopeTypeSKUCode, CouponScopeTypeCollection, CouponScopeTypeCategoryTag,
		CouponScopeTypeProducts, CouponScopeTypeMembershipTier, CouponScopeTypeNewRegistration:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s CouponScopeType) String() string { return string(s) }

// CouponStatus is the lifecycle of a coupon aggregate.
type CouponStatus string

const (
	CouponStatusDraft     CouponStatus = "draft"
	CouponStatusScheduled CouponStatus = "scheduled"
	CouponStatusActive    CouponStatus = "active"
	CouponStatusCompleted CouponStatus = "completed"
	CouponStatusArchived  CouponStatus = "archived"
)

// IsValid reports whether s is a supported coupon status.
func (s CouponStatus) IsValid() bool {
	switch s {
	case CouponStatusDraft, CouponStatusScheduled, CouponStatusActive,
		CouponStatusCompleted, CouponStatusArchived:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s CouponStatus) String() string { return string(s) }

// PromotionType identifies the mechanic represented by a promotion.
type PromotionType string

const (
	PromotionTypeProductDiscount           PromotionType = "product_discount"
	PromotionTypeVolumeDiscount            PromotionType = "volume_discount"
	PromotionTypeGroupOrderDiscount        PromotionType = "group_order_discount"
	PromotionTypeAddOnBundle               PromotionType = "add_on_bundle"
	PromotionTypeScopeBundle               PromotionType = "scope_bundle"
	PromotionTypeBOGO                      PromotionType = "bogo"
	PromotionTypeTieredPricing             PromotionType = "tiered_pricing"
	PromotionTypeMembershipPointMultiplier PromotionType = "membership_point_multiplier"
)

// IsValid reports whether p is a supported promotion type.
func (p PromotionType) IsValid() bool {
	switch p {
	case PromotionTypeProductDiscount, PromotionTypeVolumeDiscount,
		PromotionTypeGroupOrderDiscount, PromotionTypeAddOnBundle,
		PromotionTypeScopeBundle, PromotionTypeBOGO, PromotionTypeTieredPricing,
		PromotionTypeMembershipPointMultiplier:
		return true
	default:
		return false
	}
}

// String returns the wire value for p.
func (p PromotionType) String() string { return string(p) }

// PromotionScopeType identifies the target family selected by a promotion.
type PromotionScopeType string

const (
	PromotionScopeTypeSKUCode        PromotionScopeType = "sku_code"
	PromotionScopeTypeCollection     PromotionScopeType = "collection"
	PromotionScopeTypeCategoryTag    PromotionScopeType = "category_tag"
	PromotionScopeTypeProducts       PromotionScopeType = "products"
	PromotionScopeTypeMembershipTier PromotionScopeType = "membership_tier"
)

// IsValid reports whether s is a supported promotion scope type.
func (s PromotionScopeType) IsValid() bool {
	switch s {
	case PromotionScopeTypeSKUCode, PromotionScopeTypeCollection, PromotionScopeTypeCategoryTag,
		PromotionScopeTypeProducts, PromotionScopeTypeMembershipTier:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s PromotionScopeType) String() string { return string(s) }

// PromotionStatus is the lifecycle of a promotion aggregate.
type PromotionStatus string

const (
	PromotionStatusDraft     PromotionStatus = "draft"
	PromotionStatusScheduled PromotionStatus = "scheduled"
	PromotionStatusActive    PromotionStatus = "active"
	PromotionStatusCompleted PromotionStatus = "completed"
	PromotionStatusArchived  PromotionStatus = "archived"
)

// IsValid reports whether s is a supported promotion status.
func (s PromotionStatus) IsValid() bool {
	switch s {
	case PromotionStatusDraft, PromotionStatusScheduled, PromotionStatusActive,
		PromotionStatusCompleted, PromotionStatusArchived:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s PromotionStatus) String() string { return string(s) }

// DiscountType defines how DiscountValue is interpreted.
type DiscountType string

const (
	DiscountTypePercentage        DiscountType = "percentage"
	DiscountTypeFixedAmount       DiscountType = "fixed_amount"
	DiscountTypeFixedPackagePrice DiscountType = "fixed_package_price"
	DiscountTypeFreeShipping      DiscountType = "free_shipping"
)

// IsValid reports whether d is a supported discount type.
func (d DiscountType) IsValid() bool {
	switch d {
	case DiscountTypePercentage, DiscountTypeFixedAmount,
		DiscountTypeFixedPackagePrice, DiscountTypeFreeShipping:
		return true
	default:
		return false
	}
}

// String returns the wire value for d.
func (d DiscountType) String() string { return string(d) }

// MarketingChannel identifies the delivery channel for a marketing message.
type MarketingChannel string

const (
	MarketingChannelEmail  MarketingChannel = "email"
	MarketingChannelSMS    MarketingChannel = "sms"
	MarketingChannelLine   MarketingChannel = "line"
	MarketingChannelExport MarketingChannel = "export"
)

// IsValid reports whether c is a supported marketing delivery channel.
func (c MarketingChannel) IsValid() bool {
	switch c {
	case MarketingChannelEmail, MarketingChannelSMS, MarketingChannelLine, MarketingChannelExport:
		return true
	default:
		return false
	}
}

// String returns the wire value for c.
func (c MarketingChannel) String() string { return string(c) }

// MarketingMessageStatus is the lifecycle of a marketing message aggregate.
type MarketingMessageStatus string

const (
	MarketingMessageStatusDraft     MarketingMessageStatus = "draft"
	MarketingMessageStatusScheduled MarketingMessageStatus = "scheduled"
	MarketingMessageStatusSending   MarketingMessageStatus = "sending"
	MarketingMessageStatusSent      MarketingMessageStatus = "sent"
	MarketingMessageStatusPartial   MarketingMessageStatus = "partial"
	MarketingMessageStatusFailed    MarketingMessageStatus = "failed"
	MarketingMessageStatusCancelled MarketingMessageStatus = "cancelled"
	MarketingMessageStatusExported  MarketingMessageStatus = "exported"
)

// IsValid reports whether s is a supported marketing message status.
func (s MarketingMessageStatus) IsValid() bool {
	switch s {
	case MarketingMessageStatusDraft, MarketingMessageStatusScheduled,
		MarketingMessageStatusSending, MarketingMessageStatusSent,
		MarketingMessageStatusPartial, MarketingMessageStatusFailed,
		MarketingMessageStatusCancelled, MarketingMessageStatusExported:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s MarketingMessageStatus) String() string { return string(s) }

// TemplateStatus is the lifecycle of an immutable template version.
type TemplateStatus string

const (
	TemplateStatusActive   TemplateStatus = "active"
	TemplateStatusArchived TemplateStatus = "archived"
)

// IsValid reports whether s is a supported template status.
func (s TemplateStatus) IsValid() bool {
	switch s {
	case TemplateStatusActive, TemplateStatusArchived:
		return true
	default:
		return false
	}
}

// String returns the wire value for s.
func (s TemplateStatus) String() string { return string(s) }
