package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/product"
)

// StorefrontMerchandising groups admin-configurable product merchandising policy
// used by storefront read models. It is not a customer-facing projection by
// itself; services should map it to StorefrontDisplay before returning public
// product responses.
type StorefrontMerchandising struct {
	Preorder   *PreorderPolicy                `json:"preorder,omitempty"`
	SoonExpiry *SoonExpiryMerchandisingPolicy `json:"soon_expiry,omitempty"`
}

// PreorderPolicy describes whether a product can accept preorder interest and
// what limits or dates storefronts may show after backend validation.
type PreorderPolicy struct {
	Enabled                bool                          `json:"enabled"`
	StartsAt               *time.Time                    `json:"starts_at,omitempty"`
	EndsAt                 *time.Time                    `json:"ends_at,omitempty"`
	ExpectedAvailableAt    *time.Time                    `json:"expected_available_at,omitempty"`
	MaxQuantityPerRequest  int                           `json:"max_quantity_per_request,omitempty"`
	MaxQuantityPerCustomer int                           `json:"max_quantity_per_customer,omitempty"`
	Labels                 []common.LocalizedName        `json:"labels,omitempty"`
	Descriptions           []common.LocalizedDescription `json:"descriptions,omitempty"`
}

// SoonExpiryMerchandisingPolicy is the admin-managed policy for highlighting
// products whose sellable stock is near expiry.
type SoonExpiryMerchandisingPolicy struct {
	Enabled             bool                          `json:"enabled"`
	WindowDays          int                           `json:"window_days,omitempty"`
	StartsAt            *time.Time                    `json:"starts_at,omitempty"`
	EndsAt              *time.Time                    `json:"ends_at,omitempty"`
	ShowExactExpiryDate bool                          `json:"show_exact_expiry_date,omitempty"`
	Labels              []common.LocalizedName        `json:"labels,omitempty"`
	Descriptions        []common.LocalizedDescription `json:"descriptions,omitempty"`
}

// StorefrontDisplay is the backend-computed, customer-safe merchandising state
// that storefront product/listing projections can embed.
type StorefrontDisplay struct {
	Preorder *StorefrontPreorderDisplay `json:"preorder,omitempty"`
	Expiry   *StorefrontExpiryDisplay   `json:"expiry,omitempty"`
}

// StorefrontPreorderDisplay contains only fields safe for customer-facing
// product cards and detail pages.
type StorefrontPreorderDisplay struct {
	Available              bool                                 `json:"available"`
	Status                 productenum.StorefrontPreorderStatus `json:"status,omitempty"`
	StartsAt               *time.Time                           `json:"starts_at,omitempty"`
	EndsAt                 *time.Time                           `json:"ends_at,omitempty"`
	ExpectedAvailableAt    *time.Time                           `json:"expected_available_at,omitempty"`
	MaxQuantityPerRequest  int                                  `json:"max_quantity_per_request,omitempty"`
	MaxQuantityPerCustomer int                                  `json:"max_quantity_per_customer,omitempty"`
	Labels                 []common.LocalizedName               `json:"labels,omitempty"`
	Descriptions           []common.LocalizedDescription        `json:"descriptions,omitempty"`
}

// StorefrontExpiryDisplay contains backend-computed expiry merchandising fields
// for customer-facing product cards and detail pages.
type StorefrontExpiryDisplay struct {
	SoonExpiry          bool                               `json:"soon_expiry"`
	Status              productenum.StorefrontExpiryStatus `json:"status,omitempty"`
	ExpiresAt           *time.Time                         `json:"expires_at,omitempty"`
	DaysToExpiry        *int                               `json:"days_to_expiry,omitempty"`
	WindowDays          int                                `json:"window_days,omitempty"`
	ShowExactExpiryDate bool                               `json:"show_exact_expiry_date,omitempty"`
	Labels              []common.LocalizedName             `json:"labels,omitempty"`
	Descriptions        []common.LocalizedDescription      `json:"descriptions,omitempty"`
}
