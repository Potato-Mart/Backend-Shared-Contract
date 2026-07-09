package identity

import "github.com/Potato-Mart/Backend-Shared-Contract/v15/pkg/common"

// AdminAccountProfile contains optional workforce profile data for an
// adminUser account.
type AdminAccountProfile struct {
	ID             string   `json:"id,omitempty"`
	UserID         string   `json:"user_id"`
	AccountID      string   `json:"account_id"`
	EmployeeID     string   `json:"employee_id,omitempty"`
	Department     string   `json:"department,omitempty"`
	JobTitle       string   `json:"job_title,omitempty"`
	SupportRegion  string   `json:"support_region,omitempty"`
	SupportRegions []string `json:"support_regions,omitempty"`
	ManagerUserID  string   `json:"manager_user_id,omitempty"`

	common.AuditFields
}

// RetailCustomerAccountProfile contains lightweight references to existing
// retail customer, membership, and marketing records for a retailCustomer account.
type RetailCustomerAccountProfile struct {
	ID                  string `json:"id,omitempty"`
	UserID              string `json:"user_id"`
	AccountID           string `json:"account_id"`
	CustomerNumber      string `json:"customer_number,omitempty"`
	MembershipAccountID string `json:"membership_account_id,omitempty"`
	MembershipTierKey   string `json:"membership_tier_key,omitempty"`
	MarketingConsentRef string `json:"marketing_consent_ref,omitempty"`
	ReferralCode        string `json:"referral_code,omitempty"`

	common.AuditFields
}

// WholesaleCustomerAccountProfile contains lightweight wholesale organisation
// references for a wholesaleCustomer organisation-principal account.
type WholesaleCustomerAccountProfile struct {
	ID                               string `json:"id,omitempty"`
	UserID                           string `json:"user_id"`
	AccountID                        string `json:"account_id"`
	DefaultWholesaleOrganisationCode string `json:"default_wholesale_organisation_code,omitempty"`
	DefaultOrganisationAccessID      string `json:"default_organisation_access_id,omitempty"`
	MembershipAccountID              string `json:"membership_account_id,omitempty"`
	JobTitle                         string `json:"job_title,omitempty"`

	common.AuditFields
}
