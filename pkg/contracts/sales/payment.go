package sales

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
)

type Payment struct {
	ID                   string              `json:"id"`
	OrderID              string              `json:"order_id"`
	Amount               float64             `json:"amount"`
	Currency             string              `json:"currency"`
	Method               enums.PaymentMethod `json:"method"`
	Gateway              string              `json:"gateway,omitempty"`
	GatewayTransactionID string              `json:"gateway_transaction_id,omitempty"`
	PaidAt               *time.Time          `json:"paid_at,omitempty"`
	RefundedAt           *time.Time          `json:"refunded_at,omitempty"`
	RefundAmount         float64             `json:"refund_amount,omitempty"`
	RefundReason         string              `json:"refund_reason,omitempty"`
	Metadata             common.Metadata     `json:"metadata,omitempty"`
	CreatedAt            time.Time           `json:"created_at"`
}

// Will add this later: Status enums.PaymentRecordStatus `json:"status"`
