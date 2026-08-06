package product

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/product"
	securityenum "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/enums/security"
)

// StorefrontMerchandising groups admin-configurable product merchandising policy
// used by storefront read models. It is not a customer-facing projection by
// itself; services should map it to StorefrontDisplay before returning public
// product responses.
type StorefrontMerchandising struct {
	Preorder   *PreorderPolicy                `json:"preorder,omitempty"`
	SoonExpiry *SoonExpiryMerchandisingPolicy `json:"soon_expiry,omitempty"`
}

// SoonExpiryMerchandisingPolicy is the admin-managed policy for highlighting
// products whose currently sellable stock is inside the public expiry window.
// The actual lot eligibility and warning/critical thresholds remain backend
// inventory rules; this policy only controls display copy and date visibility.
type SoonExpiryMerchandisingPolicy struct {
	Enabled             bool                          `json:"enabled"`
	WindowDays          int                           `json:"window_days,omitempty"`
	StartsAt            *time.Time                    `json:"starts_at,omitempty"`
	EndsAt              *time.Time                    `json:"ends_at,omitempty"`
	ShowExactExpiryDate bool                          `json:"show_exact_expiry_date,omitempty"`
	Labels              []common.LocalizedName        `json:"labels,omitempty"`
	Descriptions        []common.LocalizedDescription `json:"descriptions,omitempty"`
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
	Expiry   *StorefrontExpiryDisplay   `json:"expiry,omitempty"`
}

// StorefrontExpiryStatus is the broad customer-facing expiry state. AlertLevel
// below carries the existing 30-day/7-day warning or critical classification.
type StorefrontExpiryStatus string

const (
	StorefrontExpiryStatusNotApplicable StorefrontExpiryStatus = "not_applicable"
	StorefrontExpiryStatusSoonExpiry    StorefrontExpiryStatus = "soon_expiry"
	StorefrontExpiryStatusExpired       StorefrontExpiryStatus = "expired"
)

// StorefrontExpiryDisplay contains only backend-computed expiry information
// safe for customer-facing product cards and detail pages. It never contains a
// lot ID, depot, supplier lot, or raw inventory quantity.
type StorefrontExpiryDisplay struct {
	SoonExpiry          bool                          `json:"soon_expiry"`
	Status              StorefrontExpiryStatus        `json:"status,omitempty"`
	AlertLevel          securityenum.AlertLevel       `json:"alert_level,omitempty"`
	ExpiresAt           *time.Time                    `json:"expires_at,omitempty"`
	DaysToExpiry        *int                          `json:"days_to_expiry,omitempty"`
	WindowDays          int                           `json:"window_days,omitempty"`
	ShowExactExpiryDate bool                          `json:"show_exact_expiry_date,omitempty"`
	Labels              []common.LocalizedName        `json:"labels,omitempty"`
	Descriptions        []common.LocalizedDescription `json:"descriptions,omitempty"`
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
