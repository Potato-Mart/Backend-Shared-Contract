package favourite

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
)

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
