package warehouseenum

type OutboundShipmentStatus string

const (
	OutboundShipmentStatusPacked     OutboundShipmentStatus = "packed"
	OutboundShipmentStatusDispatched OutboundShipmentStatus = "dispatched"
)

// IsValid reports whether s is a known OutboundShipmentStatus.
func (s OutboundShipmentStatus) IsValid() bool {
	switch s {
	case OutboundShipmentStatusPacked, OutboundShipmentStatusDispatched:
		return true
	}
	return false
}

func (s OutboundShipmentStatus) String() string { return string(s) }
