package payment_enums

// BusinessNumberScheme names the public register a merchant's business number
// belongs to, so a number is never interpreted against the wrong register.
type BusinessNumberScheme string

const (
	// BusinessNumberSchemeABN is the Australian Business Number register.
	BusinessNumberSchemeABN BusinessNumberScheme = "abn"
	// BusinessNumberSchemeACN is the Australian Company Number register.
	BusinessNumberSchemeACN BusinessNumberScheme = "acn"
	// BusinessNumberSchemeNZBN is the New Zealand Business Number register.
	BusinessNumberSchemeNZBN BusinessNumberScheme = "nzbn"
	// BusinessNumberSchemeUEN is the Singapore Unique Entity Number
	// register.
	BusinessNumberSchemeUEN BusinessNumberScheme = "uen"
	// BusinessNumberSchemeVAT is a value-added-tax registration number.
	BusinessNumberSchemeVAT BusinessNumberScheme = "vat"
	// BusinessNumberSchemeEIN is the United States Employer Identification
	// Number register.
	BusinessNumberSchemeEIN BusinessNumberScheme = "ein"
	// BusinessNumberSchemeOther is any other national register, named in
	// the profile's own scheme label.
	BusinessNumberSchemeOther BusinessNumberScheme = "other"
)

// IsValid reports whether s is a known BusinessNumberScheme.
func (s BusinessNumberScheme) IsValid() bool {
	switch s {
	case BusinessNumberSchemeABN, BusinessNumberSchemeACN, BusinessNumberSchemeNZBN,
		BusinessNumberSchemeUEN, BusinessNumberSchemeVAT, BusinessNumberSchemeEIN,
		BusinessNumberSchemeOther:
		return true
	}
	return false
}

func (s BusinessNumberScheme) String() string { return string(s) }

// TaxRegistrationStatus records whether the merchant is registered for the
// market's consumption tax. A document may only be rendered as a tax invoice
// when the issuer is registered.
type TaxRegistrationStatus string

const (
	TaxRegistrationStatusRegistered    TaxRegistrationStatus = "registered"
	TaxRegistrationStatusNotRegistered TaxRegistrationStatus = "not_registered"
	TaxRegistrationStatusUnknown       TaxRegistrationStatus = "unknown"
)

// IsValid reports whether s is a known TaxRegistrationStatus.
func (s TaxRegistrationStatus) IsValid() bool {
	switch s {
	case TaxRegistrationStatusRegistered, TaxRegistrationStatusNotRegistered, TaxRegistrationStatusUnknown:
		return true
	}
	return false
}

func (s TaxRegistrationStatus) String() string { return string(s) }

// MerchantProfileStatus is the lifecycle state of an effective-dated merchant
// legal profile.
type MerchantProfileStatus string

const (
	MerchantProfileStatusDraft    MerchantProfileStatus = "draft"
	MerchantProfileStatusActive   MerchantProfileStatus = "active"
	MerchantProfileStatusInactive MerchantProfileStatus = "inactive"
)

// IsValid reports whether s is a known MerchantProfileStatus.
func (s MerchantProfileStatus) IsValid() bool {
	switch s {
	case MerchantProfileStatusDraft, MerchantProfileStatusActive, MerchantProfileStatusInactive:
		return true
	}
	return false
}

func (s MerchantProfileStatus) String() string { return string(s) }

// DocumentKind names the customer document a transaction produced. A tax
// invoice is rendered only when the document qualifies; a receipt is always
// issued for a completed sale.
type DocumentKind string

const (
	DocumentKindReceipt        DocumentKind = "receipt"
	DocumentKindTaxInvoice     DocumentKind = "tax_invoice"
	DocumentKindInvoice        DocumentKind = "invoice"
	DocumentKindAdjustmentNote DocumentKind = "adjustment_note"
)

// IsValid reports whether k is a known DocumentKind.
func (k DocumentKind) IsValid() bool {
	switch k {
	case DocumentKindReceipt, DocumentKindTaxInvoice, DocumentKindInvoice, DocumentKindAdjustmentNote:
		return true
	}
	return false
}

func (k DocumentKind) String() string { return string(k) }
