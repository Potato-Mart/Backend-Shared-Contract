package classification_enums

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

// FavouriteListErrorCode is a stable service error code shared by clients.
type FavouriteListErrorCode string

const (
	FavouriteListErrorCodeLimitReached        FavouriteListErrorCode = "FAVOURITE_LIST_LIMIT_REACHED"
	FavouriteListErrorCodeNameConflict        FavouriteListErrorCode = "FAVOURITE_LIST_NAME_CONFLICT"
	FavouriteListErrorCodeProductLimitReached FavouriteListErrorCode = "FAVOURITE_LIST_PRODUCT_LIMIT_REACHED"
)

func (c FavouriteListErrorCode) String() string { return string(c) }

func (c FavouriteListErrorCode) IsValid() bool {
	switch c {
	case FavouriteListErrorCodeLimitReached, FavouriteListErrorCodeNameConflict,
		FavouriteListErrorCodeProductLimitReached:
		return true
	default:
		return false
	}
}
