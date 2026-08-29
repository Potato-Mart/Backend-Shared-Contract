package wish

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/wish/wish_enums"
)

// WishCandidate is an admin-authored, customer-safe ballot choice. Name,
// Description, and ImageURLs contain only approved storefront content.
type WishCandidate struct {
	ID             string                              `json:"id"`
	Name           []localization.LocalizedName        `json:"name"`
	Description    []localization.LocalizedDescription `json:"description,omitempty"`
	ImageURLs      []string                            `json:"image_urls,omitempty"`
	State          wish_enums.WishCandidateState       `json:"state"`
	CreatedSKUCode string                              `json:"created_sku_code,omitempty"`
	PublishedAt    *time.Time                          `json:"published_at,omitempty"`
	FulfilledAt    *time.Time                          `json:"fulfilled_at,omitempty"`

	audit.AuditFields
}
