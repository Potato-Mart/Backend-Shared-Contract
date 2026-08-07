package customer

import (
	"time"
)

// CustomerNotificationDelivery records one channel's durable delivery state.
type CustomerNotificationDelivery struct {
	Channel       CustomerNotificationChannel        `json:"channel"`
	Status        CustomerNotificationDeliveryStatus `json:"status"`
	AttemptCount  int                                `json:"attempt_count"`
	LastAttemptAt *time.Time                         `json:"last_attempt_at,omitempty"`
	DeliveredAt   *time.Time                         `json:"delivered_at,omitempty"`
	ErrorCode     string                             `json:"error_code,omitempty"`
	ErrorMessage  string                             `json:"error_message,omitempty"`
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
	ID             string                         `json:"id"`
	EventID        string                         `json:"event_id"`
	Topic          CustomerNotificationTopic      `json:"topic"`
	Title          string                         `json:"title"`
	Message        string                         `json:"message"`
	ActionURL      string                         `json:"action_url,omitempty"`
	OrderNumber    string                         `json:"order_number,omitempty"`
	ProductSKUCode string                         `json:"product_sku_code,omitempty"`
	ProductName    string                         `json:"product_name,omitempty"`
	Campaign       *CampaignReference             `json:"campaign,omitempty"`
	Deliveries     []CustomerNotificationDelivery `json:"deliveries,omitempty"`
	CreatedAt      time.Time                      `json:"created_at"`
	Status         CustomerNotificationStatus     `json:"status"`
	ReadAt         *time.Time                     `json:"read_at,omitempty"`
	DismissedAt    *time.Time                     `json:"dismissed_at,omitempty"`
	ExpiresAt      time.Time                      `json:"expires_at"`
}
