package retail

import (
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/orders/shipping"
)

// RetailCustomer is the grouped business profile for a retailCustomer
// account/persona. Portal admission is controlled by identity.AccountType and
// identity.PortalAccess, not by this profile.
type RetailCustomer struct {
	ID                    string                            `json:"id"`
	CustomerNumber        string                            `json:"customer_number,omitempty"`
	UserID                string                            `json:"user_id,omitempty"`
	AccountID             string                            `json:"account_id,omitempty"`
	PrimaryAuthIdentityID string                            `json:"primary_auth_identity_id,omitempty"`
	AuthIdentityIDs       []string                          `json:"auth_identity_ids,omitempty"`
	BasicInfo             RetailCustomerBasicInfo           `json:"basic_info"`
	Lifecycle             RetailCustomerLifecycle           `json:"lifecycle"`
	Management            RetailCustomerManagementProfile   `json:"management"`
	ReceiptPreferences    *RetailCustomerReceiptPreferences `json:"receipt_preferences,omitempty"`
	Commerce              RetailCustomerCommerceProfile     `json:"commerce"`
	Referral              *RetailCustomerReferralProfile    `json:"referral,omitempty"`
	ProfileCompletion     *RetailCustomerProfileCompletion  `json:"profile_completion,omitempty"`
	DefaultShipping       *party.ContactAddress             `json:"default_shipping,omitempty"`
	DefaultBilling        *party.ContactAddress             `json:"default_billing,omitempty"`
	BillingSameAsDelivery bool                              `json:"billing_same_as_delivery"`
	PreferredDeliverySlot *shipping.PreferredDeliverySlot   `json:"preferred_delivery_slot,omitempty"`
	ShippingAddresses     []party.ContactAddress            `json:"shipping_addresses,omitempty"`
	History               []security.HistoryEntry           `json:"history,omitempty"`

	audit.AuditFields
	security.DataProtectionFields
}
