package favourite

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/favourite/favourite_enums"
)

// FavouriteListOwner identifies either a retail user or a wholesale
// organisation. Only the identifier appropriate to Type is populated.
type FavouriteListOwner struct {
	Type             favourite_enums.FavouriteListOwnerType `json:"type"`
	UserID           string                                 `json:"user_id,omitempty"`
	OrganisationCode string                                 `json:"organisation_code,omitempty"`
}
