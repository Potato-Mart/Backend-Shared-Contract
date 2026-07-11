package notification

import (
	"time"

	notificationenum "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/enums/notification"
)

const (
	// PathCustomerNotifications lists notifications for the authenticated
	// customer. Management derives ownership from the access token.
	PathCustomerNotifications = "/v1/notifications"
	// PathCustomerNotificationRead marks one owned notification read. Consumers
	// replace :id with the notification id.
	PathCustomerNotificationRead = "/v1/notifications/:id/read"
	// PathInternalPreorderAvailability is the service-authenticated command
	// Commerce uses after it has reserved a paid preorder line.
	PathInternalPreorderAvailability = "/v1/internal/notifications/preorder-availability"
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

// PreorderAvailabilityCommand carries stable business identifiers only.
// Management resolves the verified email recipient and renders server-owned
// copy; the caller cannot choose an address, subject, or body.
type PreorderAvailabilityCommand struct {
	EventID              string     `json:"event_id"`
	RetailCustomerNumber string     `json:"retail_customer_number"`
	OrderNumber          string     `json:"order_number"`
	PreorderNumber       string     `json:"preorder_number,omitempty"`
	ProductSKUCode       string     `json:"product_sku_code"`
	ProductName          string     `json:"product_name,omitempty"`
	Quantity             int        `json:"quantity"`
	AvailableAt          time.Time  `json:"available_at"`
	ExpectedAvailableAt  *time.Time `json:"expected_available_at,omitempty"`
	Locale               string     `json:"locale,omitempty"`
}

// NotificationDeliveryReceipt is returned for both first delivery and
// idempotent retries of the same stable event.
type NotificationDeliveryReceipt struct {
	EventID        string                                     `json:"event_id"`
	NotificationID string                                     `json:"notification_id"`
	Topic          notificationenum.CustomerNotificationTopic `json:"topic"`
	Deliveries     []CustomerNotificationDelivery             `json:"deliveries"`
	Replayed       bool                                       `json:"replayed,omitempty"`
	CreatedAt      time.Time                                  `json:"created_at"`
}
