package membership

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

// QualifyingSpendLedgerEntry is immutable spend evidence for one retail
// customer. Positive amounts qualify spend and negative amounts reverse it.
type QualifyingSpendLedgerEntry struct {
	ID                 string                `json:"id"`
	CustomerNumber     string                `json:"customer_number"`
	Amount             common.Money          `json:"amount"`
	Reason             QualifyingSpendReason `json:"reason"`
	RelatedOrderNumber string                `json:"related_order_number,omitempty"`
	RelatedRefundID    string                `json:"related_refund_id,omitempty"`
	IdempotencyKey     string                `json:"idempotency_key"`
	OccurredAt         time.Time             `json:"occurred_at"`
	CreatedAt          time.Time             `json:"created_at"`
}
