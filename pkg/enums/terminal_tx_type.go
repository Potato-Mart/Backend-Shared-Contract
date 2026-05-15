package enums

// TerminalTxType is the kind of transaction submitted to the payment
// terminal. The values map to the shared POS operations used by Adyen
// Terminal API: PaymentRequest for purchases, unreferenced refunds,
// cashbacks/cashouts and MOTO; ReversalRequest for referenced refunds;
// and ReconciliationRequest/GetTotals-style settlement operations.
type TerminalTxType string

const (
	TerminalTxTypePurchase            TerminalTxType = "purchase"
	TerminalTxTypeRefund              TerminalTxType = "refund"
	TerminalTxTypeReversal            TerminalTxType = "reversal"
	TerminalTxTypeCashout             TerminalTxType = "cashout"
	TerminalTxTypePurchaseWithCashout TerminalTxType = "purchase_with_cashout"
	TerminalTxTypeMOTO                TerminalTxType = "moto"
	TerminalTxTypeSettlement          TerminalTxType = "settlement"
	TerminalTxTypeSettlementEnquiry   TerminalTxType = "settlement_enquiry"
)

// IsValid reports whether t is a known TerminalTxType.
func (t TerminalTxType) IsValid() bool {
	switch t {
	case TerminalTxTypePurchase, TerminalTxTypeRefund, TerminalTxTypeReversal,
		TerminalTxTypeCashout, TerminalTxTypePurchaseWithCashout,
		TerminalTxTypeMOTO, TerminalTxTypeSettlement, TerminalTxTypeSettlementEnquiry:
		return true
	}
	return false
}

func (t TerminalTxType) String() string { return string(t) }
