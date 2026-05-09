package enums

// TerminalTxType is the kind of transaction submitted to the payment
// terminal. The values map 1:1 to MX51 SCI request blocks
// (purchase_details, refund_details, cashout_details,
// purchase_with_cashout_details, moto_details) plus settlement and
// settlement-enquiry, which share the same polling envelope.
type TerminalTxType string

const (
	TerminalTxTypePurchase            TerminalTxType = "purchase"
	TerminalTxTypeRefund              TerminalTxType = "refund"
	TerminalTxTypeCashout             TerminalTxType = "cashout"
	TerminalTxTypePurchaseWithCashout TerminalTxType = "purchase_with_cashout"
	TerminalTxTypeMOTO                TerminalTxType = "moto"
	TerminalTxTypeSettlement          TerminalTxType = "settlement"
	TerminalTxTypeSettlementEnquiry   TerminalTxType = "settlement_enquiry"
)

// IsValid reports whether t is a known TerminalTxType.
func (t TerminalTxType) IsValid() bool {
	switch t {
	case TerminalTxTypePurchase, TerminalTxTypeRefund, TerminalTxTypeCashout,
		TerminalTxTypePurchaseWithCashout, TerminalTxTypeMOTO,
		TerminalTxTypeSettlement, TerminalTxTypeSettlementEnquiry:
		return true
	}
	return false
}

func (t TerminalTxType) String() string { return string(t) }
