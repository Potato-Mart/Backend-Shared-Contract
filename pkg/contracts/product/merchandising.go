package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
)

// StorefrontMerchandising groups admin-configurable product merchandising policy
// used by storefront read models. It is not a customer-facing projection by
// itself; services should map it to StorefrontDisplay before returning public
// product responses.
type StorefrontMerchandising struct {
	Preorder *PreorderPolicy `json:"preorder,omitempty"`
}

// PreorderPolicy describes whether a product can accept preorder interest and
// what limits or dates storefronts may show after backend validation.
type PreorderPolicy struct {
	Enabled                bool                          `json:"enabled"`
	StartsAt               *time.Time                    `json:"starts_at,omitempty"`
	EndsAt                 *time.Time                    `json:"ends_at,omitempty"`
	ExpectedAvailableAt    *time.Time                    `json:"expected_available_at,omitempty"`
	ScheduleTimezone       string                        `json:"schedule_timezone"`
	MaxQuantityPerOrder    int                           `json:"max_quantity_per_order,omitempty"`
	MaxQuantityPerCustomer int                           `json:"max_quantity_per_customer,omitempty"`
	Labels                 []common.LocalizedName        `json:"labels,omitempty"`
	Descriptions           []common.LocalizedDescription `json:"descriptions,omitempty"`
}

// StorefrontDisplay is the backend-computed, customer-safe merchandising state
// that storefront product/listing projections can embed.
type StorefrontDisplay struct {
	Preorder *StorefrontPreorderDisplay `json:"preorder,omitempty"`
}

// StorefrontPreorderDisplay contains only fields safe for customer-facing
// product cards and detail pages.
type StorefrontPreorderDisplay struct {
	Available              bool                                 `json:"available"`
	Status                 productenum.StorefrontPreorderStatus `json:"status,omitempty"`
	StartsAt               *time.Time                           `json:"starts_at,omitempty"`
	EndsAt                 *time.Time                           `json:"ends_at,omitempty"`
	ExpectedAvailableAt    *time.Time                           `json:"expected_available_at,omitempty"`
	ScheduleTimezone       string                               `json:"schedule_timezone"`
	MaxQuantityPerOrder    int                                  `json:"max_quantity_per_order,omitempty"`
	MaxQuantityPerCustomer int                                  `json:"max_quantity_per_customer,omitempty"`
	Labels                 []common.LocalizedName               `json:"labels,omitempty"`
	Descriptions           []common.LocalizedDescription        `json:"descriptions,omitempty"`
}
