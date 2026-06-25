package enums

// PriceVisibility controls who may see a product's price or listing. It is
// a sellability rule consumed by storefront/portal backends; the contract
// only defines the vocabulary, never the enforcement logic.
type PriceVisibility string

const (
	// PriceVisibilityPublic is visible to everyone, including guests.
	PriceVisibilityPublic PriceVisibility = "public"
	// PriceVisibilityLoginRequired is visible only to authenticated users.
	PriceVisibilityLoginRequired PriceVisibility = "login_required"
	// PriceVisibilityWholesaleApprovedOnly is visible only to approved
	// wholesale organisations.
	PriceVisibilityWholesaleApprovedOnly PriceVisibility = "wholesale_approved_only"
	// PriceVisibilityHidden is never shown.
	PriceVisibilityHidden PriceVisibility = "hidden"
)

// IsValid reports whether v is a known PriceVisibility.
func (v PriceVisibility) IsValid() bool {
	switch v {
	case PriceVisibilityPublic, PriceVisibilityLoginRequired,
		PriceVisibilityWholesaleApprovedOnly, PriceVisibilityHidden:
		return true
	}
	return false
}

func (v PriceVisibility) String() string { return string(v) }
