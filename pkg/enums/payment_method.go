package enums

// PaymentMethod tracks the money payment methods.

type PaymentMethod string

const (
	PaymentMethodCard         PaymentMethod = "card"
	PaymentMethodCash         PaymentMethod = "cash"
	PaymentMethodQR           PaymentMethod = "qr"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodLinePay      PaymentMethod = "line_pay"
	PaymentMethodECPay        PaymentMethod = "ecpay"
	PaymentMethodManual       PaymentMethod = "manual"

	// EFTPOS-terminal-backed methods. These are distinct from the
	// generic Card method so settlement reports and refund flows can
	// reconcile against the terminal provider (e.g. MX51 SCI splits
	// purchase / cashout / MOTO into separate transaction blocks).
	PaymentMethodEFTPOS  PaymentMethod = "eftpos"
	PaymentMethodMOTO    PaymentMethod = "moto"
	PaymentMethodCashout PaymentMethod = "cashout"

	// Manual pay will add period
)

// IsValid reports whether p is a known PaymentMethod.
func (p PaymentMethod) IsValid() bool {
	switch p {
	case PaymentMethodCard, PaymentMethodCash, PaymentMethodQR, PaymentMethodBankTransfer,
		PaymentMethodLinePay, PaymentMethodECPay, PaymentMethodManual,
		PaymentMethodEFTPOS, PaymentMethodMOTO, PaymentMethodCashout:
		return true
	}
	return false
}

func (p PaymentMethod) String() string { return string(p) }
