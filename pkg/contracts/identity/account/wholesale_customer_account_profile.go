package account

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"

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
