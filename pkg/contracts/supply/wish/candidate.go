package wish

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/wish/wish_enums"
)

// WishCandidate is an admin-authored, customer-safe ballot choice. Name,
// Description, and ImageURLs contain only approved storefront content.
type WishCandidate struct {
	ID           string                              `json:"id"`
	Name         []localization.LocalizedName        `json:"name"`
	Description  []localization.LocalizedDescription `json:"description,omitempty"`
	ImageURLs    []string                            `json:"image_urls,omitempty"`
	State        wish_enums.WishCandidateState       `json:"state"`
	CreatedSKUID string                              `json:"created_sku_id,omitempty"`
	PublishedAt  *time.Time                          `json:"published_at,omitempty"`
	FulfilledAt  *time.Time                          `json:"fulfilled_at,omitempty"`
	CreatedAt    time.Time                           `json:"created_at"`
	UpdatedAt    time.Time                           `json:"updated_at"`
}
