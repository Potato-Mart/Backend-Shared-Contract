package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/retail/retail_enums"
)

// CustomerRegisteredEvent is emitted on the customer-events topic when a
// retail customer profile is created (self-registration or the identity
// ensure-profile path). AggregateID is the customer number. Deliberately
// carries no PII (no name/email/phone).
type CustomerRegisteredEvent struct {
	CustomerID        string                                 `json:"customer_id"`
	CustomerNumber    string                                 `json:"customer_number"`
	UserID            string                                 `json:"user_id,omitempty"`
	AccountID         string                                 `json:"account_id,omitempty"`
	AcquisitionSource retail_enums.CustomerAcquisitionSource `json:"acquisition_source,omitempty"`
	RegisteredAt      time.Time                              `json:"registered_at"`
	RequestID         string                                 `json:"request_id,omitempty"`
}
