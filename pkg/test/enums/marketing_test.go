package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/insights/marketing/marketing_enums"
)

func TestMarketingEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "marketingenum.MarketingCampaignStatus", valid: []stringEnum{marketing_enums.MarketingCampaignStatusDraft, marketing_enums.MarketingCampaignStatusSending, marketing_enums.MarketingCampaignStatusSent, marketing_enums.MarketingCampaignStatusPartial, marketing_enums.MarketingCampaignStatusFailed, marketing_enums.MarketingCampaignStatusCancelled, marketing_enums.MarketingCampaignStatusExported}, invalid: marketing_enums.MarketingCampaignStatus("__invalid__")},
		{name: "marketingenum.MarketingChannel", valid: []stringEnum{marketing_enums.MarketingChannelEmail, marketing_enums.MarketingChannelSMS, marketing_enums.MarketingChannelLine, marketing_enums.MarketingChannelExport}, invalid: marketing_enums.MarketingChannel("__invalid__")},
		{name: "marketingenum.MarketingRecipientStatus", valid: []stringEnum{marketing_enums.MarketingRecipientStatusPending, marketing_enums.MarketingRecipientStatusSent, marketing_enums.MarketingRecipientStatusDelivered, marketing_enums.MarketingRecipientStatusBounced, marketing_enums.MarketingRecipientStatusOpened, marketing_enums.MarketingRecipientStatusClicked, marketing_enums.MarketingRecipientStatusFailed, marketing_enums.MarketingRecipientStatusUnsubscribed}, invalid: marketing_enums.MarketingRecipientStatus("__invalid__")},
	})
}
