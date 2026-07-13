package notification

import (
	"time"

	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/enums/notification"
)

// CustomerNotificationDelivery records one channel's durable delivery state.
type CustomerNotificationDelivery struct {
	Channel       notificationenum.CustomerNotificationChannel        `json:"channel"`
	Status        notificationenum.CustomerNotificationDeliveryStatus `json:"status"`
	AttemptCount  int                                                 `json:"attempt_count"`
	LastAttemptAt *time.Time                                          `json:"last_attempt_at,omitempty"`
	DeliveredAt   *time.Time                                          `json:"delivered_at,omitempty"`
	ErrorCode     string                                              `json:"error_code,omitempty"`
	ErrorMessage  string                                              `json:"error_message,omitempty"`
}

// CustomerNotification is the customer-safe portal projection. Recipient
// addresses and provider details deliberately do not appear on this contract.
type CustomerNotification struct {
	ID             string                                     `json:"id"`
	EventID        string                                     `json:"event_id"`
	Topic          notificationenum.CustomerNotificationTopic `json:"topic"`
	Title          string                                     `json:"title"`
	Message        string                                     `json:"message"`
	ActionURL      string                                     `json:"action_url,omitempty"`
	OrderNumber    string                                     `json:"order_number,omitempty"`
	ProductSKUCode string                                     `json:"product_sku_code,omitempty"`
	ProductName    string                                     `json:"product_name,omitempty"`
	Deliveries     []CustomerNotificationDelivery             `json:"deliveries,omitempty"`
	CreatedAt      time.Time                                  `json:"created_at"`
	ReadAt         *time.Time                                 `json:"read_at,omitempty"`
}
