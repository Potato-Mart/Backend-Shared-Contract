package customers

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/shared"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/customer"
)

// RetailCustomer is the grouped business profile for a retailCustomer
// account/persona. Portal admission is controlled by identity.AccountType and
// identity.PortalAccess, not by this profile.
type RetailCustomer struct {
	ID                    string                           `json:"id"`
	CustomerNumber        string                           `json:"customer_number,omitempty"`
	UserID                string                           `json:"user_id,omitempty"`
	AccountID             string                           `json:"account_id,omitempty"`
	PrimaryAuthIdentityID string                           `json:"primary_auth_identity_id,omitempty"`
	AuthIdentityIDs       []string                         `json:"auth_identity_ids,omitempty"`
	BasicInfo             RetailCustomerBasicInfo          `json:"basic_info"`
	Lifecycle             RetailCustomerLifecycle          `json:"lifecycle"`
	Management            RetailCustomerManagementProfile  `json:"management"`
	Membership            RetailCustomerMembershipProfile  `json:"membership"`
	Marketing             RetailCustomerMarketingProfile   `json:"marketing"`
	Commerce              RetailCustomerCommerceProfile    `json:"commerce"`
	Analytics             *RetailCustomerAnalyticsProfile  `json:"analytics,omitempty"`
	Referral              *RetailCustomerReferralProfile   `json:"referral,omitempty"`
	ProfileCompletion     *RetailCustomerProfileCompletion `json:"profile_completion,omitempty"`
	DefaultShipping       *common.ContactAddress           `json:"default_shipping,omitempty"`
	DefaultBilling        *common.ContactAddress           `json:"default_billing,omitempty"`
	ShippingAddresses     []common.ContactAddress          `json:"shipping_addresses,omitempty"`
	History               []shared.HistoryEntry            `json:"history,omitempty"`

	common.AuditFields
	common.DataProtectionFields
}

// RetailCustomerSummary is a compact retail customer projection for lists,
// search results, and relationship references.
type RetailCustomerSummary struct {
	ID                  string                      `json:"id"`
	AccountID           string                      `json:"account_id,omitempty"`
	UserID              string                      `json:"user_id,omitempty"`
	CustomerNumber      string                      `json:"customer_number,omitempty"`
	DisplayName         string                      `json:"display_name,omitempty"`
	Email               string                      `json:"email,omitempty"`
	Phone               string                      `json:"phone,omitempty"`
	Status              customerenum.CustomerStatus `json:"status"`
	MembershipAccountID string                      `json:"membership_account_id,omitempty"`
	MembershipTierKey   string                      `json:"membership_tier_key,omitempty"`
	Tags                []string                    `json:"tags,omitempty"`
	Metadata            common.Metadata             `json:"metadata,omitempty"`
}

// RetailCustomerBasicInfo groups stable name, contact, and acquisition fields
// for a retail customer profile.
type RetailCustomerBasicInfo struct {
	Name              common.PersonName                      `json:"name"`
	Contacts          common.ContactChannels                 `json:"contacts"`
	DateOfBirth       *time.Time                             `json:"date_of_birth,omitempty"`
	Gender            customerenum.CustomerGender            `json:"gender,omitempty"`
	AcquisitionSource customerenum.CustomerAcquisitionSource `json:"acquisition_source,omitempty"`
}

// RetailCustomerLifecycle groups the retail customer profile lifecycle.
type RetailCustomerLifecycle struct {
	Status        customerenum.CustomerStatus `json:"status"`
	RegisteredAt  *time.Time                  `json:"registered_at,omitempty"`
	ActivatedAt   *time.Time                  `json:"activated_at,omitempty"`
	BlockedAt     *time.Time                  `json:"blocked_at,omitempty"`
	BlockedReason string                      `json:"blocked_reason,omitempty"`
	ClosedAt      *time.Time                  `json:"closed_at,omitempty"`
	ClosedReason  string                      `json:"closed_reason,omitempty"`
}
