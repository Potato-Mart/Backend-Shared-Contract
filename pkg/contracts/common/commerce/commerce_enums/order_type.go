package commerce_enums

// OrderType identifies the channel through which an order is placed.
// POS is a channel, not a buyer type.
type OrderType string

const (
	OrderTypeOnline OrderType = "online"
	OrderTypePOS    OrderType = "pos"
	OrderTypeB2B    OrderType = "b2b"
	OrderTypeRelay  OrderType = "relay"
	OrderTypeManual OrderType = "manual"
	OrderTypeImport OrderType = "import"
)

// IsValid reports whether p is a known OrderType.
func (p OrderType) IsValid() bool {
	switch p {
	case OrderTypeOnline, OrderTypePOS, OrderTypeB2B,
		OrderTypeRelay, OrderTypeManual, OrderTypeImport:
		return true
	}
	return false
}

// String returns the wire value for p.
func (p OrderType) String() string { return string(p) }
