package common

// OrganisationDetail contains the shared public and administrative details for
// a company or organisation. It is intended for suppliers, wholesale
// organisations, invoices, fulfilment paperwork, and other contracts that need
// consistent organisation information.
//
// It embeds PartyRef for the organisation's identity and primary contact
// (id / name / phone / email) and adds richer registration, tax, address,
// branding, and operating metadata.
type OrganisationDetail struct {
	PartyRef `bson:",inline"` // organisation id / name / phone / email

	TradingName string `json:"trading_name,omitempty"`
	LegalName   string `json:"legal_name,omitempty"`
	Description string `json:"description,omitempty"`

	// Registration and tax details.
	ABN                 string `json:"abn,omitempty"`
	ACN                 string `json:"acn,omitempty"`
	TaxRegistrationID   string `json:"tax_registration_id,omitempty"`
	TaxRegistered       bool   `json:"tax_registered,omitempty"`
	BusinessNumberLabel string `json:"business_number_label,omitempty"`

	// Additional contact details (primary phone/email come from PartyRef).
	Website       string `json:"website,omitempty"`
	ContactPerson string `json:"contact_person,omitempty"`

	// Addresses.
	RegisteredAddress *ContactAddress `json:"registered_address,omitempty"`
	BillingAddress    *ContactAddress `json:"billing_address,omitempty"`
	ShippingAddress   *ContactAddress `json:"shipping_address,omitempty"`

	// Branding and public references.
	LogoURL   string `json:"logo_url,omitempty"`
	BrandName string `json:"brand_name,omitempty"`

	// Additional business metadata.
	Industry     string   `json:"industry,omitempty"`
	Timezone     string   `json:"timezone,omitempty"`
	CurrencyCode string   `json:"currency_code,omitempty"`
	Metadata     Metadata `json:"metadata,omitempty"`
}
