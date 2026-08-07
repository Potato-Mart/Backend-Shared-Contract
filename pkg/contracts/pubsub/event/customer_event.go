package event

import (
	"time"

	retail "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/customers/retail"
)

// CustomerRegisteredEvent is emitted on the customer-events topic when a
// retail customer profile is created (self-registration or the identity
// ensure-profile path). AggregateID is the customer number. Deliberately
// carries no PII (no name/email/phone).
type CustomerRegisteredEvent struct {
	CustomerID        string                           `json:"customer_id"`
	CustomerNumber    string                           `json:"customer_number"`
	UserID            string                           `json:"user_id,omitempty"`
	AccountID         string                           `json:"account_id,omitempty"`
	AcquisitionSource retail.CustomerAcquisitionSource `json:"acquisition_source,omitempty"`
	RegisteredAt      time.Time                        `json:"registered_at"`
	RequestID         string                           `json:"request_id,omitempty"`
}

// CustomerProfileUpdatedEvent is emitted on the customer-events topic when a
// retail customer profile changes. UpdatedFields carries field names only,
// never values.
type CustomerProfileUpdatedEvent struct {
	CustomerNumber string    `json:"customer_number"`
	UserID         string    `json:"user_id,omitempty"`
	UpdatedFields  []string  `json:"updated_fields,omitempty"`
	UpdatedBy      string    `json:"updated_by,omitempty"`
	UpdatedAt      time.Time `json:"updated_at"`
	RequestID      string    `json:"request_id,omitempty"`
}

// CustomerConsentChangedEvent is emitted on the customer-events topic when
// marketing consent flags change.
type CustomerConsentChangedEvent struct {
	CustomerNumber string    `json:"customer_number"`
	UserID         string    `json:"user_id,omitempty"`
	EmailOptIn     bool      `json:"email_opt_in"`
	SMSOptIn       bool      `json:"sms_opt_in"`
	LineOptIn      bool      `json:"line_opt_in"`
	PushOptIn      bool      `json:"push_opt_in"`
	Source         string    `json:"source,omitempty"`
	ChangedAt      time.Time `json:"changed_at"`
	RequestID      string    `json:"request_id,omitempty"`
}
