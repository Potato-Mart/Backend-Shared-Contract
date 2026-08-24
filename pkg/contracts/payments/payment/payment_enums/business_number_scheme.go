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
