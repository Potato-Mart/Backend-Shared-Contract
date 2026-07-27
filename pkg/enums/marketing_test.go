package enums_test

import (
	"testing"

	marketingenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/marketing"
)

func TestMarketingEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "marketingenum.MarketingCampaignStatus", valid: []stringEnum{marketingenum.MarketingCampaignStatusDraft, marketingenum.MarketingCampaignStatusSending, marketingenum.MarketingCampaignStatusSent, marketingenum.MarketingCampaignStatusPartial, marketingenum.MarketingCampaignStatusFailed, marketingenum.MarketingCampaignStatusCancelled, marketingenum.MarketingCampaignStatusExported}, invalid: marketingenum.MarketingCampaignStatus("__invalid__")},
		{name: "marketingenum.MarketingChannel", valid: []stringEnum{marketingenum.MarketingChannelEmail, marketingenum.MarketingChannelSMS, marketingenum.MarketingChannelLine, marketingenum.MarketingChannelExport}, invalid: marketingenum.MarketingChannel("__invalid__")},
		{name: "marketingenum.MarketingRecipientStatus", valid: []stringEnum{marketingenum.MarketingRecipientStatusPending, marketingenum.MarketingRecipientStatusSent, marketingenum.MarketingRecipientStatusDelivered, marketingenum.MarketingRecipientStatusBounced, marketingenum.MarketingRecipientStatusOpened, marketingenum.MarketingRecipientStatusClicked, marketingenum.MarketingRecipientStatusFailed, marketingenum.MarketingRecipientStatusUnsubscribed}, invalid: marketingenum.MarketingRecipientStatus("__invalid__")},
	})
}
