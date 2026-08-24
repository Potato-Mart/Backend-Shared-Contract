package payment_enums

// InvoiceResendStatus tracks a customer-requested invoice email redelivery.
type InvoiceResendStatus string

const (
	InvoiceResendStatusQueued InvoiceResendStatus = "queued"
	InvoiceResendStatusSent   InvoiceResendStatus = "sent"
	InvoiceResendStatusFailed InvoiceResendStatus = "failed"
)

func (s InvoiceResendStatus) IsValid() bool {
	switch s {
	case InvoiceResendStatusQueued, InvoiceResendStatusSent, InvoiceResendStatusFailed:
		return true
	default:
		return false
	}
}

func (s InvoiceResendStatus) String() string { return string(s) }
