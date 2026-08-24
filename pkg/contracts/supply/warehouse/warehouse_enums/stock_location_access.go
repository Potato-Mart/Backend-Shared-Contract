package warehouse_enums

// StockLocationAccess describes who may access a stock location.
type StockLocationAccess string

const (
	StockLocationAccessCustomerAccessible StockLocationAccess = "CUSTOMER_ACCESSIBLE"
	StockLocationAccessStaffOnly          StockLocationAccess = "STAFF_ONLY"
)

func (a StockLocationAccess) IsValid() bool {
	switch a {
	case StockLocationAccessCustomerAccessible, StockLocationAccessStaffOnly:
		return true
	}
	return false
}

func (a StockLocationAccess) String() string { return string(a) }
