package ads_enums

// AdLandingKind says what an ad landing destination refers to. The ads module
// builds the per-locale landing URL from the kind and its reference, so the
// destination always stays on our own domain and carries the agreed tracking
// parameters.
type AdLandingKind string

const (
	// AdLandingKindCampaign refers to a storefront campaign by its code.
	AdLandingKindCampaign AdLandingKind = "campaign"
	// AdLandingKindProduct refers to one product.
	AdLandingKindProduct AdLandingKind = "product"
	// AdLandingKindSelection refers to a curated product selection.
	AdLandingKindSelection AdLandingKind = "selection"
	// AdLandingKindURL refers to a path on our own storefront domain. It is
	// never a shortened link and never already carries tracking parameters.
	AdLandingKindURL AdLandingKind = "url"
)

// IsValid reports whether k is a known AdLandingKind.
func (k AdLandingKind) IsValid() bool {
	switch k {
	case AdLandingKindCampaign, AdLandingKindProduct, AdLandingKindSelection, AdLandingKindURL:
		return true
	}
	return false
}

func (k AdLandingKind) String() string { return string(k) }
