package membership

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/membership/membership_enums"
)

// QualifyingSpendLedgerEntry is immutable spend evidence for one retail
// customer. Positive amounts qualify spend and negative amounts reverse it.
type QualifyingSpendLedgerEntry struct {
	ID                 string                                 `json:"id"`
	CustomerNumber     string                                 `json:"customer_number"`
	Amount             money.Money                            `json:"amount"`
	Reason             membership_enums.QualifyingSpendReason `json:"reason"`
	RelatedOrderNumber string                                 `json:"related_order_number,omitempty"`
	RelatedRefundID    string                                 `json:"related_refund_id,omitempty"`
	OccurredAt         time.Time                              `json:"occurred_at"`
	CreatedAt          time.Time                              `json:"created_at"`

	security.DataProtectionFields
}
