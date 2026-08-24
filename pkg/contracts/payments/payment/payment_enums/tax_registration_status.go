package payment_enums

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
