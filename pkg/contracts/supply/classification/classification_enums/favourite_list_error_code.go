package classification_enums

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
