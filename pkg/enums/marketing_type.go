package enums

// MarketingChannel is the delivery channel for a marketing campaign.
type MarketingChannel string

const (
	MarketingChannelEmail  MarketingChannel = "EMAIL"
	MarketingChannelSMS    MarketingChannel = "SMS"
	MarketingChannelLine   MarketingChannel = "LINE"
	MarketingChannelExport MarketingChannel = "EXPORT"
)

func (m MarketingChannel) IsValid() bool {
	switch m {
	case MarketingChannelEmail, MarketingChannelSMS, MarketingChannelLine, MarketingChannelExport:
		return true
	}
	return false
}

func (m MarketingChannel) String() string { return string(m) }

// MarketingCampaignStatus is the lifecycle of a marketing campaign send.
type MarketingCampaignStatus string

const (
	MarketingCampaignStatusDraft     MarketingCampaignStatus = "DRAFT"
	MarketingCampaignStatusSending   MarketingCampaignStatus = "SENDING"
	MarketingCampaignStatusSent      MarketingCampaignStatus = "SENT"
	MarketingCampaignStatusPartial   MarketingCampaignStatus = "PARTIAL"
	MarketingCampaignStatusFailed    MarketingCampaignStatus = "FAILED"
	MarketingCampaignStatusCancelled MarketingCampaignStatus = "CANCELLED"
	MarketingCampaignStatusExported  MarketingCampaignStatus = "EXPORTED"
)

func (m MarketingCampaignStatus) IsValid() bool {
	switch m {
	case MarketingCampaignStatusDraft, MarketingCampaignStatusSending, MarketingCampaignStatusSent,
		MarketingCampaignStatusPartial, MarketingCampaignStatusFailed,
		MarketingCampaignStatusCancelled, MarketingCampaignStatusExported:
		return true
	}
	return false
}

func (m MarketingCampaignStatus) String() string { return string(m) }

// MarketingRecipientStatus tracks the delivery outcome for a single recipient.
type MarketingRecipientStatus string

const (
	MarketingRecipientStatusPending      MarketingRecipientStatus = "PENDING"
	MarketingRecipientStatusSent         MarketingRecipientStatus = "SENT"
	MarketingRecipientStatusDelivered    MarketingRecipientStatus = "DELIVERED"
	MarketingRecipientStatusBounced      MarketingRecipientStatus = "BOUNCED"
	MarketingRecipientStatusOpened       MarketingRecipientStatus = "OPENED"
	MarketingRecipientStatusClicked      MarketingRecipientStatus = "CLICKED"
	MarketingRecipientStatusFailed       MarketingRecipientStatus = "FAILED"
	MarketingRecipientStatusUnsubscribed MarketingRecipientStatus = "UNSUBSCRIBED"
)

func (m MarketingRecipientStatus) IsValid() bool {
	switch m {
	case MarketingRecipientStatusPending, MarketingRecipientStatusSent,
		MarketingRecipientStatusDelivered, MarketingRecipientStatusBounced,
		MarketingRecipientStatusOpened, MarketingRecipientStatusClicked,
		MarketingRecipientStatusFailed, MarketingRecipientStatusUnsubscribed:
		return true
	}
	return false
}

func (m MarketingRecipientStatus) String() string { return string(m) }
