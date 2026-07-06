package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/common"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/customer"
)

// RetailCustomerActivity is a single entry in the retail customer interaction timeline.
// It records orders, calls, complaints, tier changes, refunds, and any
// other touchpoints for a customer profile.
type RetailCustomerActivity struct {
	ID                   string                            `json:"id"`
	RetailCustomerNumber string                            `json:"retail_customer_number"`
	ActivityType         customerenum.CustomerActivityType `json:"activity_type"`
	Title                string                            `json:"title"`
	Body                 string                            `json:"body,omitempty"`
	Amount               *common.Money                     `json:"amount,omitempty"`
	Metadata             common.Metadata                   `json:"metadata,omitempty"`
	OccurredAt           time.Time                         `json:"occurred_at"`
	CreatedBy            string                            `json:"created_by,omitempty"`
	CreatedAt            time.Time                         `json:"created_at"`
}

// RetailCustomerExternalIdentity is a single cross-channel identifier
// belonging to a retail customer profile.
type RetailCustomerExternalIdentity struct {
	ID                   string                            `json:"id"`
	RetailCustomerNumber string                            `json:"retail_customer_number"`
	AccountID            string                            `json:"account_id,omitempty"`
	AuthIdentityID       string                            `json:"auth_identity_id,omitempty"`
	Kind                 customerenum.CustomerIdentityKind `json:"kind"`
	Value                string                            `json:"value"`
	Label                string                            `json:"label,omitempty"`
	Verified             bool                              `json:"verified"`
	CreatedAt            time.Time                         `json:"created_at"`

	common.DataProtectionFields
}
