package subscription

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v4/pkg/enums"
)

// SubscriptionPlan defines a recurring purchase option for a product.
// A product may have multiple plans (e.g. weekly, bi-weekly, monthly).
type SubscriptionPlan struct {
	ID              string           `json:"id"`
	Product         product.Snapshot `json:"product"`    // snapshot for display without join
	UnitPrice       common.Money     `json:"unit_price"` // price snapshot at plan creation
	FrequencyDays   int              `json:"frequency_days"`
	FrequencyLabel  string           `json:"frequency_label"` // e.g. "weekly" / "monthly"
	DiscountPercent float64          `json:"discount_percent"`
	MinCycles       int              `json:"min_cycles"` // 1 = cancel any time
	IsActive        bool             `json:"is_active"`
	SortOrder       int              `json:"sort_order"`

	common.AuditFields
}

// CustomerSubscription is an active subscription between a customer and a plan.
type CustomerSubscription struct {
	ID                string                   `json:"id"`
	CustomerProfileID string                   `json:"customer_profile_id"`
	PlanID            string                   `json:"plan_id"`
	Qty               int                      `json:"qty"`
	Status            enums.SubscriptionStatus `json:"status"`
	StartedAt         time.Time                `json:"started_at"`
	NextOrderAt       time.Time                `json:"next_order_at"`
	LastOrderAt       *time.Time               `json:"last_order_at,omitempty"`
	PausedAt          *time.Time               `json:"paused_at,omitempty"`
	CancelledAt       *time.Time               `json:"cancelled_at,omitempty"`
	CyclesCompleted   int                      `json:"cycles_completed"`
	Note              string                   `json:"note,omitempty"`

	common.AuditFields
}
