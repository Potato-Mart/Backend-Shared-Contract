package common

// CompanyDetail contains the shared public and administrative details for a company.
// It is intended for invoices, purchase orders, fulfilment paperwork, supplier/customer
// profiles, and other contracts that need consistent organisation information.
//
// It embeds PartyRef for the company's identity and primary contact
// (id / name / phone / email) and adds the richer organisation fields, so it
// can be embedded wherever full company information is required.
type CompanyDetail struct {
	PartyRef `bson:",inline"` // company id / name / phone / email

	TradingName        string `json:"trading_name,omitempty"`
	LegalName          string `json:"legal_name,omitempty"`
	CompanyDescription string `json:"company_description,omitempty"`

	// Registration and tax details.
	CompanyABN          string `json:"company_abn,omitempty"`
	CompanyACN          string `json:"company_acn,omitempty"`
	TaxRegistrationID   string `json:"tax_registration_id,omitempty"`
	TaxRegistered       bool   `json:"tax_registered,omitempty"`
	BusinessNumberLabel string `json:"business_number_label,omitempty"`

	// Additional contact details (primary phone/email come from PartyRef).
	CompanyWebsite       string `json:"company_website,omitempty"`
	CompanyContactPerson string `json:"company_contact_person,omitempty"`

	// Addresses.
	CompanyAddress  ContactAddress  `json:"company_address"`
	BillingAddress  *ContactAddress `json:"billing_address,omitempty"`
	ShippingAddress *ContactAddress `json:"shipping_address,omitempty"`

	// Branding and public references.
	CompanyLogoURL string `json:"logo_url,omitempty"`
	BrandName      string `json:"brand_name,omitempty"`

	// Additional business metadata.
	Industry     string            `json:"industry,omitempty"`
	Timezone     string            `json:"timezone,omitempty"`
	CurrencyCode string            `json:"currency_code,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}
