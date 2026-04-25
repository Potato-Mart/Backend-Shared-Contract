package enums

// PaymentMethod tracks the money payment methods.

type PaymentMethod string

const (
	PaymentMethodCard         PaymentMethod = "CARD"
	PaymentMethodCash         PaymentMethod = "CASH"
	PaymentMethodQR           PaymentMethod = "QR"
	PaymentMethodBankTransfer PaymentMethod = "BANK_TRANSFER"
	PaymentMethodLinePay      PaymentMethod = "LINE_PAY"
	PaymentMethodECPay        PaymentMethod = "EC_PAY"
	PaymentMethodManual       PaymentMethod = "MANUAL"

	// Manual pay will add period
)

// IsValid reports whether p is a known PaymentMethod.
func (p PaymentMethod) IsValid() bool {
	switch p {
	case PaymentMethodCard, PaymentMethodCash, PaymentMethodQR, PaymentMethodBankTransfer,
		PaymentMethodLinePay, PaymentMethodECPay, PaymentMethodManual:
		return true
	}
	return false
}

func (p PaymentMethod) String() string { return string(p) }
