package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

// WholesaleOrganisation represents an approved or prospective B2B organisation
// that wholesaleCustomer accounts can access through OrganisationAccess grants.
type WholesaleOrganisation struct {
	common.OrganisationDetail

	PrimaryWholesaleCustomerID string                            `json:"primary_wholesale_customer_id,omitempty"`
	MembershipAccountID        string                            `json:"membership_account_id,omitempty"`
	Status                     enums.WholesaleOrganisationStatus `json:"status"`
	TierKey                    string                            `json:"tier_key,omitempty"`
	PriceTier                  int                               `json:"price_tier,omitempty"`
	Approval                   *common.LifecycleAction           `json:"approval,omitempty"`
	Suspension                 *common.LifecycleAction           `json:"suspension,omitempty"`
	Rejection                  *common.LifecycleAction           `json:"rejection,omitempty"`
	Closure                    *common.LifecycleAction           `json:"closure,omitempty"`

	common.AuditFields
}

// WholesaleOrganisationSummary is the compact organisation projection carried
// by access/session and membership responses.
type WholesaleOrganisationSummary struct {
	common.PartyRef

	PrimaryWholesaleCustomerID string                            `json:"primary_wholesale_customer_id,omitempty"`
	MembershipAccountID        string                            `json:"membership_account_id,omitempty"`
	TradingName                string                            `json:"trading_name,omitempty"`
	LegalName                  string                            `json:"legal_name,omitempty"`
	Status                     enums.WholesaleOrganisationStatus `json:"status"`
	TierKey                    string                            `json:"tier_key,omitempty"`
	PriceTier                  int                               `json:"price_tier,omitempty"`
}
