package customer

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/notifications/customer/customer_enums"
)

// CustomerNotificationDelivery records one channel's durable delivery state.
type CustomerNotificationDelivery struct {
	Channel       customer_enums.CustomerNotificationChannel        `json:"channel"`
	Status        customer_enums.CustomerNotificationDeliveryStatus `json:"status"`
	AttemptCount  int                                               `json:"attempt_count"`
	LastAttemptAt *time.Time                                        `json:"last_attempt_at,omitempty"`
	DeliveredAt   *time.Time                                        `json:"delivered_at,omitempty"`
	ErrorCode     string                                            `json:"error_code,omitempty"`
	ErrorMessage  string                                            `json:"error_message,omitempty"`
}

// CampaignReference gives a campaign notification enough typed identity to
// refetch authoritative content without embedding campaign copy or provider
// destinations in the durable notification row.
type CampaignReference struct {
	CampaignID         string `json:"campaign_id"`
	CampaignKey        string `json:"campaign_key"`
	PromotionID        string `json:"promotion_id,omitempty"`
	ActivationRevision int64  `json:"activation_revision"`
	ContentRevision    int64  `json:"content_revision"`
}

// CustomerNotification is the customer-safe portal projection. Recipient
// addresses and provider details deliberately do not appear on this contract.
type CustomerNotification struct {
	ID      string `json:"id"`
	EventID string `json:"event_id"`
	// MarketID and CountryCode are the denormalized market and country the
	// notification was raised for, carried so a geographically scoped staff
	// query is a plain indexed match.
	MarketID    string                                    `json:"market_id,omitempty"`
	CountryCode geography.CountryCode                     `json:"country_code,omitempty"`
	Topic       customer_enums.CustomerNotificationTopic  `json:"topic"`
	Title       string                                    `json:"title"`
	Message     string                                    `json:"message"`
	ActionURL   string                                    `json:"action_url,omitempty"`
	OrderNumber string                                    `json:"order_number,omitempty"`
	SKUID       string                                    `json:"sku_id,omitempty"`
	ProductName string                                    `json:"product_name,omitempty"`
	Campaign    *CampaignReference                        `json:"campaign,omitempty"`
	Deliveries  []CustomerNotificationDelivery            `json:"deliveries,omitempty"`
	CreatedAt   time.Time                                 `json:"created_at"`
	Status      customer_enums.CustomerNotificationStatus `json:"status"`
	ReadAt      *time.Time                                `json:"read_at,omitempty"`
	DismissedAt *time.Time                                `json:"dismissed_at,omitempty"`
	ExpiresAt   time.Time                                 `json:"expires_at"`
}
