package classification

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/classification/classification_enums"
)

// FavouriteListOwner identifies either a retail user or a wholesale
// organisation. Only the identifier appropriate to Type is populated.
type FavouriteListOwner struct {
	Type             classification_enums.FavouriteListOwnerType `json:"type"`
	UserID           string                                      `json:"user_id,omitempty"`
	OrganisationCode string                                      `json:"organisation_code,omitempty"`
}

// FavouriteListProduct records product membership without cart quantities.
type FavouriteListProduct struct {
	SKUID   string    `json:"sku_id"`
	AddedAt time.Time `json:"added_at"`
}

// FavouriteList is one named, persisted customer product collection.
// DefaultNameSlot is populated only while a system-generated List-N name is in
// use; renaming the list clears it so that slot can be allocated again.
type FavouriteList struct {
	ID              string                 `json:"id"`
	Owner           FavouriteListOwner     `json:"owner"`
	Name            string                 `json:"name"`
	DefaultNameSlot int                    `json:"default_name_slot,omitempty"`
	Products        []FavouriteListProduct `json:"products,omitempty"`

	audit.AuditFields
}
