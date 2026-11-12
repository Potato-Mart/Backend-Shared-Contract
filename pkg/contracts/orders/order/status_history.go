package order

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/orders/order/order_enums"
)

// StatusHistory is one entry on an order's status timeline.
//
// Kind is empty for an ordinary workflow transition and set for a staff
// override or payment correction. ReasonCode is an open, service-owned code
// for why the entry was written (for example PAYMENT_CORRECTION); Note keeps
// the staff-entered reason text. Changes lists the field-level before and
// after values the entry applied.
type StatusHistory struct {
	FromValue  string                        `json:"from_value,omitempty"`
	ToValue    string                        `json:"to_value"`
	Note       string                        `json:"note,omitempty"`
	Kind       order_enums.StatusHistoryKind `json:"kind,omitempty"`
	ReasonCode string                        `json:"reason_code,omitempty"`
	Changes    []security.HistoryChange      `json:"changes,omitempty"`
	ActorID    string                        `json:"actor_id,omitempty"`
	ActorEmail string                        `json:"actor_email,omitempty"`
	CreatedAt  time.Time                     `json:"created_at"`
}
