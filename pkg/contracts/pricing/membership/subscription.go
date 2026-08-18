package membership

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/membership/membership_enums"
)

// SubscriptionPlan defines a recurring purchase option available through the
// membership domain. It remains separate from points accounting.
type SubscriptionPlan struct {
	ID              string      `json:"id"`
	SKUCode         string      `json:"sku_code"`
	UnitPrice       money.Money `json:"unit_price"`
	FrequencyDays   int         `json:"frequency_days"`
	FrequencyLabel  string      `json:"frequency_label"`
	DiscountPercent float64     `json:"discount_percent"`
	MinCycles       int         `json:"min_cycles"`
	IsActive        bool        `json:"is_active"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	audit.AuditFields
}

// MemberSubscription is an active recurring purchase attached to a retail
// customer.
type MemberSubscription struct {
	ID              string                                    `json:"id"`
	CustomerNumber  string                                    `json:"customer_number"`
	PlanID          string                                    `json:"plan_id"`
	Qty             int                                       `json:"qty"`
	Status          membership_enums.MemberSubscriptionStatus `json:"status"`
	StartedAt       time.Time                                 `json:"started_at"`
	NextOrderAt     time.Time                                 `json:"next_order_at"`
	LastOrderAt     *time.Time                                `json:"last_order_at,omitempty"`
	PausedAt        *time.Time                                `json:"paused_at,omitempty"`
	CancelledAt     *time.Time                                `json:"cancelled_at,omitempty"`
	CyclesCompleted int                                       `json:"cycles_completed"`
	Note            string                                    `json:"note,omitempty"`
	History         []security.HistoryEntry                   `json:"history,omitempty"`
	// MarketCode and CountryCode are the denormalized owning market and its
	// country, carried so a geographically scoped staff query is a plain
	// indexed match. They mirror the plan the subscription was taken out
	// against; without them a staff-facing subscription list cannot be
	// scoped at all, because the subscription reaches its geography only
	// through PlanID.
	MarketCode  string                `json:"market_code,omitempty"`
	CountryCode geography.CountryCode `json:"country_code,omitempty"`

	audit.AuditFields
}
