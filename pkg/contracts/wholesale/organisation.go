package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

// WholesaleOrganisation represents an approved or prospective B2B organisation
// that wholesaleCustomer accounts can access through memberships.
type WholesaleOrganisation struct {
	common.OrganisationDetail `bson:",inline"`

	PrimaryWholesaleCustomerID string                            `json:"primary_wholesale_customer_id,omitempty"`
	Status                     enums.WholesaleOrganisationStatus `json:"status"`
	TierKey                    string                            `json:"tier_key,omitempty"`
	PriceTier                  int                               `json:"price_tier,omitempty"`
	Approval                   *common.LifecycleAction           `json:"approval,omitempty"`
	Suspension                 *common.LifecycleAction           `json:"suspension,omitempty"`
	Rejection                  *common.LifecycleAction           `json:"rejection,omitempty"`
	Closure                    *common.LifecycleAction           `json:"closure,omitempty"`

	common.AuditFields `bson:",inline"`
}

// WholesaleOrganisationSummary is the compact organisation projection carried
// by access/session and membership responses.
type WholesaleOrganisationSummary struct {
	common.PartyRef `bson:",inline"`

	PrimaryWholesaleCustomerID string                            `json:"primary_wholesale_customer_id,omitempty"`
	TradingName                string                            `json:"trading_name,omitempty"`
	LegalName                  string                            `json:"legal_name,omitempty"`
	Status                     enums.WholesaleOrganisationStatus `json:"status"`
	TierKey                    string                            `json:"tier_key,omitempty"`
	PriceTier                  int                               `json:"price_tier,omitempty"`
}
