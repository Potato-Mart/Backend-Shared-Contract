package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/membership/membership_enums"
)

// MemberSubscription is an active recurring purchase attached to a customer.
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
	MarketCode      string                                    `json:"market_code,omitempty"`
	CountryCode     geography.CountryCode                     `json:"country_code,omitempty"`

	audit.AuditFields
}
