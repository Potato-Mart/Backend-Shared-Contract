package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/pubsub/pricing/promotion_enums"
)

func TestPromotionPublicationContextsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "promotion_enums.PromotionPublicationContext",
			valid: []stringEnum{
				promotion_enums.PromotionPublicationContextStandalone,
				promotion_enums.PromotionPublicationContextCampaignPublishTogether,
			},
			invalid: promotion_enums.PromotionPublicationContext("__invalid__"),
		},
	})
}
