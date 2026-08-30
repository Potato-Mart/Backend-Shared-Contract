package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/wholesale/wholesale_enums"
)

// WholesaleOrganisationSummary is the compact organisation projection carried
// by access/session and membership responses.
type WholesaleOrganisationSummary struct {
	party.PartyRef

	PrincipalUserID             string                                        `json:"principal_user_id,omitempty"`
	PrincipalAccountID          string                                        `json:"principal_account_id,omitempty"`
	PrimaryAuthIdentityID       string                                        `json:"primary_auth_identity_id,omitempty"`
	PrimaryOrganisationAccessID string                                        `json:"primary_organisation_access_id,omitempty"`
	TradingName                 string                                        `json:"trading_name,omitempty"`
	LegalName                   string                                        `json:"legal_name,omitempty"`
	Category                    wholesale_enums.WholesaleOrganisationCategory `json:"category,omitempty"`
	Status                      wholesale_enums.WholesaleOrganisationStatus   `json:"status"`
}
