package membership

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/membership/membership_enums"
)

// SubscriptionPlan defines a recurring purchase option available through the
// membership domain. It remains separate from points accounting.
type SubscriptionPlan struct {
	ID              string      `json:"id"`
	SKUID           string      `json:"sku_id"`
	UnitPrice       money.Money `json:"unit_price"`
	FrequencyDays   int         `json:"frequency_days"`
	FrequencyLabel  string      `json:"frequency_label"`
	DiscountPercent float64     `json:"discount_percent"`
	MinCycles       int         `json:"min_cycles"`
	IsActive        bool        `json:"is_active"`
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

	audit.AuditFields
}
