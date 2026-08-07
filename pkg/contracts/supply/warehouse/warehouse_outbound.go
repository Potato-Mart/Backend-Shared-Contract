package warehouse

type OutboundShipmentStatus string

const (
	OutboundShipmentStatusPacked     OutboundShipmentStatus = "packed"
	OutboundShipmentStatusDispatched OutboundShipmentStatus = "dispatched"
	OutboundShipmentStatusDelivered  OutboundShipmentStatus = "delivered"
)

// IsValid reports whether s is a known OutboundShipmentStatus.
func (s OutboundShipmentStatus) IsValid() bool {
	switch s {
	case OutboundShipmentStatusPacked, OutboundShipmentStatusDispatched, OutboundShipmentStatusDelivered:
		return true
	}
	return false
}

func (s OutboundShipmentStatus) String() string { return string(s) }
