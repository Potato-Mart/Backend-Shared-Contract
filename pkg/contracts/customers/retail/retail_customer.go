package retail

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/shipping"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/customers/retail/retail_enums"
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
	Marketing             RetailCustomerMarketingProfile    `json:"marketing"`
	ReceiptPreferences    *RetailCustomerReceiptPreferences `json:"receipt_preferences,omitempty"`
	Commerce              RetailCustomerCommerceProfile     `json:"commerce"`
	Analytics             *RetailCustomerAnalyticsProfile   `json:"analytics,omitempty"`
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

// RetailCustomerSummary is a compact retail customer projection for lists,
// search results, and relationship references.
type RetailCustomerSummary struct {
	ID                 string                            `json:"id"`
	AccountID          string                            `json:"account_id,omitempty"`
	UserID             string                            `json:"user_id,omitempty"`
	CustomerNumber     string                            `json:"customer_number,omitempty"`
	DisplayName        string                            `json:"display_name,omitempty"`
	Email              string                            `json:"email,omitempty"`
	Phone              string                            `json:"phone,omitempty"`
	Status             retail_enums.CustomerStatus       `json:"status"`
	Tags               []string                          `json:"tags,omitempty"`
	ReceiptPreferences *RetailCustomerReceiptPreferences `json:"receipt_preferences,omitempty"`
	Metadata           metadata.Metadata                 `json:"metadata,omitempty"`
}

// RetailCustomerBasicInfo groups stable name, contact, and acquisition fields
// for a retail customer profile.
type RetailCustomerBasicInfo struct {
	Name                   party.PersonName                       `json:"name"`
	Contacts               party.ContactChannels                  `json:"contacts"`
	PreferredContactMethod retail_enums.PreferredContactMethod    `json:"preferred_contact_method,omitempty"`
	DateOfBirth            *time.Time                             `json:"date_of_birth,omitempty"`
	Gender                 retail_enums.CustomerGender            `json:"gender,omitempty"`
	AcquisitionSource      retail_enums.CustomerAcquisitionSource `json:"acquisition_source,omitempty"`
}

// RetailCustomerLifecycle groups the retail customer profile lifecycle.
type RetailCustomerLifecycle struct {
	Status        retail_enums.CustomerStatus `json:"status"`
	RegisteredAt  *time.Time                  `json:"registered_at,omitempty"`
	ActivatedAt   *time.Time                  `json:"activated_at,omitempty"`
	BlockedAt     *time.Time                  `json:"blocked_at,omitempty"`
	BlockedReason string                      `json:"blocked_reason,omitempty"`
	ClosedAt      *time.Time                  `json:"closed_at,omitempty"`
	ClosedReason  string                      `json:"closed_reason,omitempty"`
}
