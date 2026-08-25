package enums_test

import (
	"testing"

	insights_marketing_enums "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/insights/marketing/marketing_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/marketing/message/message_enums"
)

func TestMarketingMessageAndInsightEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "message.MarketingMessageStatus", valid: []stringEnum{message_enums.MarketingMessageStatusDraft, message_enums.MarketingMessageStatusScheduled, message_enums.MarketingMessageStatusSending, message_enums.MarketingMessageStatusSent, message_enums.MarketingMessageStatusPartial, message_enums.MarketingMessageStatusFailed, message_enums.MarketingMessageStatusCancelled}, invalid: message_enums.MarketingMessageStatus("__invalid__")},
		{name: "message.MarketingTemplateStatus", valid: []stringEnum{message_enums.MarketingTemplateStatusActive, message_enums.MarketingTemplateStatusArchived}, invalid: message_enums.MarketingTemplateStatus("__invalid__")},
		{name: "insights.CampaignPredictionStatus", valid: []stringEnum{insights_marketing_enums.CampaignPredictionStatusNotApplicable, insights_marketing_enums.CampaignPredictionStatusReady, insights_marketing_enums.CampaignPredictionStatusWarning}, invalid: insights_marketing_enums.CampaignPredictionStatus("__invalid__")},
		{name: "insights.CampaignPredictionSource", valid: []stringEnum{insights_marketing_enums.CampaignPredictionSourceSameSeries, insights_marketing_enums.CampaignPredictionSourceSimilarEvent, insights_marketing_enums.CampaignPredictionSourceLast14DaysDoubled}, invalid: insights_marketing_enums.CampaignPredictionSource("__invalid__")},
	})
}
