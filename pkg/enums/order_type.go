package enums

type OrderType string

const (
	OrderTypeOnline OrderType = "ONLINE"
	OrderTypePOS    OrderType = "POS"
	OrderTypeB2B    OrderType = "B2B"
	OrderTypeRelay  OrderType = "RELAY"
	OrderTypeManual OrderType = "MANUAL"
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

func (p OrderType) String() string { return string(p) }
