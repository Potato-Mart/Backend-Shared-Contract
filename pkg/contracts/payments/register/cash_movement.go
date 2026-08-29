package register

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	pos_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/register/register_enums"
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
