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
