package promotion_enums

// PromotionPublicationContext records the trusted publication path that
// produced a promotion.changed event. It is provenance, not campaign ownership.
type PromotionPublicationContext string

const (
	PromotionPublicationContextStandalone              PromotionPublicationContext = "standalone"
	PromotionPublicationContextCampaignPublishTogether PromotionPublicationContext = "campaign_publish_together"
)

// IsValid reports whether context is supported by promotion.changed v3.
func (context PromotionPublicationContext) IsValid() bool {
	switch context {
	case PromotionPublicationContextStandalone, PromotionPublicationContextCampaignPublishTogether:
		return true
	}
	return false
}

// String returns the wire value for context.
func (context PromotionPublicationContext) String() string { return string(context) }
