package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/wholesale"
)

// WholesaleOrganisation represents an approved or prospective B2B organisation
// and is the canonical wholesale customer/business account. People who act for
// the organisation are represented by OrganisationAccess records.
type WholesaleOrganisation struct {
	common.OrganisationDetail

	PrincipalUserID             string                                    `json:"principal_user_id,omitempty"`
	PrincipalAccountID          string                                    `json:"principal_account_id,omitempty"`
	PrimaryAuthIdentityID       string                                    `json:"primary_auth_identity_id,omitempty"`
	AuthIdentityIDs             []string                                  `json:"auth_identity_ids,omitempty"`
	PrimaryOrganisationAccessID string                                    `json:"primary_organisation_access_id,omitempty"`
	MembershipAccountID         string                                    `json:"membership_account_id,omitempty"`
	Status                      wholesaleenum.WholesaleOrganisationStatus `json:"status"`
	TierKey                     string                                    `json:"tier_key,omitempty"`
	PriceTier                   int                                       `json:"price_tier,omitempty"`
	Approval                    *common.LifecycleAction                   `json:"approval,omitempty"`
	Suspension                  *common.LifecycleAction                   `json:"suspension,omitempty"`
	Rejection                   *common.LifecycleAction                   `json:"rejection,omitempty"`
	Closure                     *common.LifecycleAction                   `json:"closure,omitempty"`

	common.AuditFields
}

// WholesaleOrganisationSummary is the compact organisation projection carried
// by access/session and membership responses.
type WholesaleOrganisationSummary struct {
	common.PartyRef

	PrincipalUserID             string                                    `json:"principal_user_id,omitempty"`
	PrincipalAccountID          string                                    `json:"principal_account_id,omitempty"`
	PrimaryAuthIdentityID       string                                    `json:"primary_auth_identity_id,omitempty"`
	PrimaryOrganisationAccessID string                                    `json:"primary_organisation_access_id,omitempty"`
	MembershipAccountID         string                                    `json:"membership_account_id,omitempty"`
	TradingName                 string                                    `json:"trading_name,omitempty"`
	LegalName                   string                                    `json:"legal_name,omitempty"`
	Status                      wholesaleenum.WholesaleOrganisationStatus `json:"status"`
	TierKey                     string                                    `json:"tier_key,omitempty"`
	PriceTier                   int                                       `json:"price_tier,omitempty"`
}
