package favourite_enums

// FavouriteListOwnerType identifies the principal that owns a saved list.
type FavouriteListOwnerType string

const (
	FavouriteListOwnerTypeRetailUser            FavouriteListOwnerType = "retail_user"
	FavouriteListOwnerTypeWholesaleOrganisation FavouriteListOwnerType = "wholesale_organisation"
)

func (t FavouriteListOwnerType) String() string { return string(t) }

func (t FavouriteListOwnerType) IsValid() bool {
	switch t {
	case FavouriteListOwnerTypeRetailUser, FavouriteListOwnerTypeWholesaleOrganisation:
		return true
	default:
		return false
	}
}
