package warehouse_enums

type InboundReceiptStatus string

const (
	InboundReceiptStatusDraft     InboundReceiptStatus = "draft"
	InboundReceiptStatusConfirmed InboundReceiptStatus = "confirmed"
	// InboundReceiptStatusReversed marks an inbound receipt whose linked
	// purchase receipt was reversed. It is set only by the reversal command,
	// never by a status transition request.
	InboundReceiptStatusReversed InboundReceiptStatus = "reversed"
)

// IsValid reports whether s is a known InboundReceiptStatus.
func (s InboundReceiptStatus) IsValid() bool {
	switch s {
	case InboundReceiptStatusDraft, InboundReceiptStatusConfirmed, InboundReceiptStatusReversed:
		return true
	}
	return false
}

func (s InboundReceiptStatus) String() string { return string(s) }
