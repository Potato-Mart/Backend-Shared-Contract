package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

// WholesaleCustomer is the grouped business profile for a wholesaleCustomer
// account/persona. Organisation approval and access lifecycle live on
// WholesaleOrganisation and OrganisationAccess.
type WholesaleCustomer struct {
	ID                    string                             `json:"id"`
	CustomerNumber        string                             `json:"customer_number,omitempty"`
	UserID                string                             `json:"user_id,omitempty"`
	AccountID             string                             `json:"account_id,omitempty"`
	PrimaryAuthIdentityID string                             `json:"primary_auth_identity_id,omitempty"`
	AuthIdentityIDs       []string                           `json:"auth_identity_ids,omitempty"`
	OrganisationCode      string                             `json:"organisation_code"`
	OrganisationAccessID  string                             `json:"organisation_access_id,omitempty"`
	BasicInfo             WholesaleCustomerBasicInfo         `json:"basic_info"`
	Commercial            WholesaleCustomerCommercialProfile `json:"commercial"`
	AccountProfile        WholesaleCustomerAccountProfile    `json:"account_profile"`
	Terms                 WholesaleTerms                     `json:"terms"`
	// RecentOrders is a bounded, sync-populated display strip (hard-capped,
	// typically 5) of the organisation's latest orders for the account page.
	// Denormalised projection only; canonical orders live in Commerce.
	RecentOrders           []sales.OrderSummary    `json:"recent_orders,omitempty"`
	DefaultBillingAddress  *common.ContactAddress  `json:"default_billing_address,omitempty"`
	DefaultShippingAddress *common.ContactAddress  `json:"default_shipping_address,omitempty"`
	AdditionalShipTo       []common.ContactAddress `json:"additional_ship_to,omitempty"`
	History                []shared.HistoryEntry   `json:"history,omitempty"`

	common.AuditFields
	common.DataProtectionFields
}

// WholesaleCustomerSummary is a compact wholesale customer projection for
// lists, search results, and organisation-member references.
type WholesaleCustomerSummary struct {
	ID                   string               `json:"id"`
	AccountID            string               `json:"account_id,omitempty"`
	UserID               string               `json:"user_id,omitempty"`
	OrganisationCode     string               `json:"organisation_code"`
	OrganisationAccessID string               `json:"organisation_access_id,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	Email                string               `json:"email,omitempty"`
	Phone                string               `json:"phone,omitempty"`
	RoleKey              string               `json:"role_key,omitempty"`
	Status               enums.CustomerStatus `json:"status"`
	TierKey              string               `json:"tier_key,omitempty"`
}

// WholesaleCustomerBasicInfo groups the person/contact details for a wholesale
// customer contact.
type WholesaleCustomerBasicInfo struct {
	Name              common.PersonName               `json:"name"`
	Contacts          common.ContactChannels          `json:"contacts"`
	JobTitle          string                          `json:"job_title,omitempty"`
	Department        string                          `json:"department,omitempty"`
	AcquisitionSource enums.CustomerAcquisitionSource `json:"acquisition_source,omitempty"`
}

// WholesaleCustomerCommercialProfile groups staff-managed commercial details
// for the wholesale customer contact.
type WholesaleCustomerCommercialProfile struct {
	SalesRep     string          `json:"sales_rep,omitempty"`
	CRMTier      string          `json:"crm_tier,omitempty"`
	PaymentTerms string          `json:"payment_terms,omitempty"`
	TaxID        string          `json:"tax_id,omitempty"`
	Notes        string          `json:"notes,omitempty"`
	Tags         []string        `json:"tags,omitempty"`
	Metadata     common.Metadata `json:"metadata,omitempty"`
}

// WholesaleCustomerAccountProfile groups account/persona and organisation access
// references for wholesale access.
type WholesaleCustomerAccountProfile struct {
	Status        enums.CustomerStatus `json:"status"`
	RoleKey       string               `json:"role_key,omitempty"`
	IsPrimary     bool                 `json:"is_primary,omitempty"`
	JoinedAt      *time.Time           `json:"joined_at,omitempty"`
	LastOrderedAt *time.Time           `json:"last_ordered_at,omitempty"`
}

// WholesaleTerms groups B2B price-tier configuration and freight terms for a
// wholesale customer profile.
type WholesaleTerms struct {
	TierKey                string          `json:"tier_key,omitempty"`
	PriceTier              int             `json:"price_tier,omitempty"`
	PriceTierSea           *int            `json:"price_tier_sea,omitempty"`
	PriceTierAir           *int            `json:"price_tier_air,omitempty"`
	RebateRate             *float64        `json:"rebate_rate,omitempty"`
	ShippingFee            *common.Money   `json:"shipping_fee,omitempty"`
	FreightRule            common.Metadata `json:"freight_rule,omitempty"`
	AccountCheckoutEnabled bool            `json:"account_checkout_enabled,omitempty"`
	DueDays                int             `json:"due_days,omitempty"`
	PurchaseOrderRequired  bool            `json:"purchase_order_required,omitempty"`
	PaymentInstructions    string          `json:"payment_instructions,omitempty"`
	CreditLimit            *common.Money   `json:"credit_limit,omitempty"`
}
