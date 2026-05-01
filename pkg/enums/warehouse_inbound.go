package enums

type InboundReceiptStatus string

const (
	InboundReceiptStatusDraft     InboundReceiptStatus = "draft"
	InboundReceiptStatusConfirmed InboundReceiptStatus = "confirmed"
)

// IsValid reports whether s is a known InboundReceiptStatus.
func (s InboundReceiptStatus) IsValid() bool {
	switch s {
	case InboundReceiptStatusDraft, InboundReceiptStatusConfirmed:
		return true
	}
	return false
}

func (s InboundReceiptStatus) String() string { return string(s) }
