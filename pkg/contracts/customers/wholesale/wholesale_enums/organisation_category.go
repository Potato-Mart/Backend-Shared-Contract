package wholesale_enums

// WholesaleOrganisationCategory is the trade category of an approved wholesale
// organisation. Pricing binds a category price book to it during resolution,
// between the exact-organisation override and the general wholesale channel
// book.
type WholesaleOrganisationCategory string

const (
	// WholesaleOrganisationCategoryRegionalAgent is a regional distribution
	// agent reselling into its own territory.
	WholesaleOrganisationCategoryRegionalAgent WholesaleOrganisationCategory = "regional_agent"
	// WholesaleOrganisationCategorySupermarket is a supermarket or grocery
	// retailer buying for resale.
	WholesaleOrganisationCategorySupermarket WholesaleOrganisationCategory = "supermarket"
	// WholesaleOrganisationCategoryRestaurant is a restaurant or food
	// service business buying for consumption on premises.
	WholesaleOrganisationCategoryRestaurant WholesaleOrganisationCategory = "restaurant"
)

// IsValid reports whether c is a known WholesaleOrganisationCategory.
func (c WholesaleOrganisationCategory) IsValid() bool {
	switch c {
	case WholesaleOrganisationCategoryRegionalAgent, WholesaleOrganisationCategorySupermarket,
		WholesaleOrganisationCategoryRestaurant:
		return true
	}
	return false
}

func (c WholesaleOrganisationCategory) String() string { return string(c) }
