package pos

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/pos/pos_enums"
	"time"
)

// CashMovement is one cash-drawer movement recorded during a session.
type CashMovement struct {
	ID         string                     `json:"id"`
	SessionID  string                     `json:"session_id"`
	RegisterID string                     `json:"register_id"`
	DepotCode  string                     `json:"depot_code,omitempty"`
	Kind       pos_enums.CashMovementKind `json:"kind"`
	Amount     money.Money                `json:"amount"`
	Reason     string                     `json:"reason,omitempty"`
	RecordedBy string                     `json:"recorded_by,omitempty"`
	OccurredAt time.Time                  `json:"occurred_at"`
}
