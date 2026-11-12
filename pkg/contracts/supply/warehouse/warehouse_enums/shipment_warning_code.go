package warehouse_enums

// ShipmentWarningCode names a condition staff must see on an outbound
// shipment without it blocking dispatch. CarrierCancellationRequired means an
// earlier carrier booking could not be cancelled automatically after a carrier
// switch, so staff must cancel it with the carrier. OrderUnpaid means the
// order still has an unpaid balance.
type ShipmentWarningCode string

const (
	ShipmentWarningCodeCarrierCancellationRequired ShipmentWarningCode = "carrier_cancellation_required"
	ShipmentWarningCodeOrderUnpaid                 ShipmentWarningCode = "order_unpaid"
)

// IsValid reports whether c is a known ShipmentWarningCode value.
func (c ShipmentWarningCode) IsValid() bool {
	switch c {
	case ShipmentWarningCodeCarrierCancellationRequired, ShipmentWarningCodeOrderUnpaid:
		return true
	}
	return false
}

func (c ShipmentWarningCode) String() string { return string(c) }
