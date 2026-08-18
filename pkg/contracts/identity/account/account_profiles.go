package account

import "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"

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

	audit.AuditFields
}

// RetailCustomerAccountProfile contains lightweight references for a retail
// customer account. CustomerNumber is also the membership account key.
type RetailCustomerAccountProfile struct {
	ID                  string `json:"id,omitempty"`
	UserID              string `json:"user_id"`
	AccountID           string `json:"account_id"`
	CustomerNumber      string `json:"customer_number,omitempty"`
	MarketingConsentRef string `json:"marketing_consent_ref,omitempty"`
	ReferralCode        string `json:"referral_code,omitempty"`

	audit.AuditFields
}

// WholesaleCustomerAccountProfile contains lightweight wholesale organisation
// references for a wholesaleCustomer organisation-principal account.
type WholesaleCustomerAccountProfile struct {
	ID                               string `json:"id,omitempty"`
	UserID                           string `json:"user_id"`
	AccountID                        string `json:"account_id"`
	DefaultWholesaleOrganisationCode string `json:"default_wholesale_organisation_code,omitempty"`
	DefaultOrganisationAccessID      string `json:"default_organisation_access_id,omitempty"`
	JobTitle                         string `json:"job_title,omitempty"`

	audit.AuditFields
}
