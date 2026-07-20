package wish

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
	wishenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/wish"
)

// WishCandidate is an admin-authored, customer-safe ballot choice. Name,
// Description, and ImageURLs contain only approved storefront content.
type WishCandidate struct {
	ID                    string                        `json:"id"`
	Name                  []common.LocalizedName        `json:"name"`
	Description           []common.LocalizedDescription `json:"description,omitempty"`
	ImageURLs             []string                      `json:"image_urls,omitempty"`
	State                 wishenum.WishCandidateState   `json:"state"`
	CreatedProductSKUCode string                        `json:"created_product_sku_code,omitempty"`
	PublishedAt           *time.Time                    `json:"published_at,omitempty"`
	FulfilledAt           *time.Time                    `json:"fulfilled_at,omitempty"`
	CreatedAt             time.Time                     `json:"created_at"`
	UpdatedAt             time.Time                     `json:"updated_at"`
}
