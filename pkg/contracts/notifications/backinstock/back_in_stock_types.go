package backinstock

// BackInStockChannel is a customer-selected delivery route for a restock alert.
type BackInStockChannel string

const (
	BackInStockChannelEmail BackInStockChannel = "email"
	BackInStockChannelSMS   BackInStockChannel = "sms"
)

func (c BackInStockChannel) String() string { return string(c) }

func (c BackInStockChannel) IsValid() bool {
	switch c {
	case BackInStockChannelEmail, BackInStockChannelSMS:
		return true
	default:
		return false
	}
}

// BackInStockStatus is the lifecycle state for a one-shot restock alert.
type BackInStockStatus string

const (
	BackInStockStatusPending   BackInStockStatus = "pending"
	BackInStockStatusNotified  BackInStockStatus = "notified"
	BackInStockStatusCancelled BackInStockStatus = "cancelled"
)

func (s BackInStockStatus) String() string { return string(s) }

func (s BackInStockStatus) IsValid() bool {
	switch s {
	case BackInStockStatusPending, BackInStockStatusNotified, BackInStockStatusCancelled:
		return true
	default:
		return false
	}
}

// BackInStockCustomerType records which storefront/account family requested an alert.
type BackInStockCustomerType string

const (
	BackInStockCustomerTypeRetail    BackInStockCustomerType = "retail"
	BackInStockCustomerTypeWholesale BackInStockCustomerType = "wholesale"
)

func (t BackInStockCustomerType) String() string { return string(t) }

func (t BackInStockCustomerType) IsValid() bool {
	switch t {
	case BackInStockCustomerTypeRetail, BackInStockCustomerTypeWholesale:
		return true
	default:
		return false
	}
}
