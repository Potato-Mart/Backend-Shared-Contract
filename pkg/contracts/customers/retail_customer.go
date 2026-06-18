package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// RetailCustomer is the grouped business profile for a generalCustomer
// account/persona. Portal admission is controlled by identity.AccountType and
// identity.PortalAccess, not by this profile.
type RetailCustomer struct {
	ID                string                          `json:"id"`
	Identity          common.IdentityLink             `json:"identity"`
	BasicInfo         RetailCustomerBasicInfo         `json:"basic_info"`
	Lifecycle         RetailCustomerLifecycle         `json:"lifecycle"`
	Management        RetailCustomerManagementProfile `json:"management"`
	Loyalty           RetailCustomerLoyaltyProfile    `json:"loyalty"`
	Marketing         RetailCustomerMarketingProfile  `json:"marketing"`
	Commerce          RetailCustomerCommerceProfile   `json:"commerce"`
	Analytics         *RetailCustomerAnalyticsProfile `json:"analytics,omitempty"`
	Referral          *RetailCustomerReferralProfile  `json:"referral,omitempty"`
	DefaultShipping   *common.ContactAddress          `json:"default_shipping,omitempty"`
	ShippingAddresses []common.ContactAddress         `json:"shipping_addresses,omitempty"`
	History           []shared.HistoryEntry           `json:"history,omitempty"`

	common.AuditFields          `bson:",inline"`
	common.DataProtectionFields `bson:",inline"`
}

// RetailCustomerSummary is a compact retail customer projection for lists,
// search results, and relationship references.
type RetailCustomerSummary struct {
	ID             string               `json:"id"`
	AccountID      string               `json:"account_id,omitempty"`
	UserID         string               `json:"user_id,omitempty"`
	CustomerNumber string               `json:"customer_number,omitempty"`
	DisplayName    string               `json:"display_name,omitempty"`
	Email          string               `json:"email,omitempty"`
	Phone          string               `json:"phone,omitempty"`
	Status         enums.CustomerStatus `json:"status"`
	LoyaltyTierKey string               `json:"loyalty_tier_key,omitempty"`
	Tags           []string             `json:"tags,omitempty"`
	Metadata       common.Metadata      `json:"metadata,omitempty"`
}

// RetailCustomerBasicInfo groups stable identity, name, contact, and
// acquisition fields for a retail customer profile.
type RetailCustomerBasicInfo struct {
	CustomerNumber    string                          `json:"customer_number,omitempty"`
	Name              common.PersonName               `json:"name"`
	Contacts          common.ContactChannels          `json:"contacts"`
	DateOfBirth       *time.Time                      `json:"date_of_birth,omitempty"`
	AcquisitionSource enums.CustomerAcquisitionSource `json:"acquisition_source,omitempty"`
}

// RetailCustomerLifecycle groups the retail customer profile lifecycle.
type RetailCustomerLifecycle struct {
	Status        enums.CustomerStatus `json:"status"`
	RegisteredAt  *time.Time           `json:"registered_at,omitempty"`
	ActivatedAt   *time.Time           `json:"activated_at,omitempty"`
	BlockedAt     *time.Time           `json:"blocked_at,omitempty"`
	BlockedReason string               `json:"blocked_reason,omitempty"`
	ClosedAt      *time.Time           `json:"closed_at,omitempty"`
	ClosedReason  string               `json:"closed_reason,omitempty"`
}
