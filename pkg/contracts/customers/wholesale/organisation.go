package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/party"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/customers/wholesale/wholesale_enums"
)

// WholesaleOrganisation represents an approved or prospective B2B organisation
// and is the canonical wholesale customer/business account. People who act for
// the organisation are represented by OrganisationAccess records.
type WholesaleOrganisation struct {
	party.OrganisationDetail

	PrincipalUserID             string                                      `json:"principal_user_id,omitempty"`
	PrincipalAccountID          string                                      `json:"principal_account_id,omitempty"`
	PrimaryAuthIdentityID       string                                      `json:"primary_auth_identity_id,omitempty"`
	AuthIdentityIDs             []string                                    `json:"auth_identity_ids,omitempty"`
	PrimaryOrganisationAccessID string                                      `json:"primary_organisation_access_id,omitempty"`
	Status                      wholesale_enums.WholesaleOrganisationStatus `json:"status"`
	Approval                    *audit.LifecycleAction                      `json:"approval,omitempty"`
	Suspension                  *audit.LifecycleAction                      `json:"suspension,omitempty"`
	Rejection                   *audit.LifecycleAction                      `json:"rejection,omitempty"`
	Closure                     *audit.LifecycleAction                      `json:"closure,omitempty"`

	audit.AuditFields
}

// WholesaleOrganisationSummary is the compact organisation projection carried
// by access/session and membership responses.
type WholesaleOrganisationSummary struct {
	party.PartyRef

	PrincipalUserID             string                                      `json:"principal_user_id,omitempty"`
	PrincipalAccountID          string                                      `json:"principal_account_id,omitempty"`
	PrimaryAuthIdentityID       string                                      `json:"primary_auth_identity_id,omitempty"`
	PrimaryOrganisationAccessID string                                      `json:"primary_organisation_access_id,omitempty"`
	TradingName                 string                                      `json:"trading_name,omitempty"`
	LegalName                   string                                      `json:"legal_name,omitempty"`
	Status                      wholesale_enums.WholesaleOrganisationStatus `json:"status"`
}
