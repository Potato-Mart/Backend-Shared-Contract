package enums

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

func (p OrderType) String() string { return string(p) }
