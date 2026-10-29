package order

import "github.com/Potato-Mart/Backend-Shared-Contract/v50/pkg/contracts/common/money"

// OfflineCashTenderEvidence records the actual cash taken and change given at
// the till. Payments owns reconciliation with the order and register session.
type OfflineCashTenderEvidence struct {
	AmountDue      money.Money `json:"amount_due"`
	AmountReceived money.Money `json:"amount_received"`
	Change         money.Money `json:"change"`
}
