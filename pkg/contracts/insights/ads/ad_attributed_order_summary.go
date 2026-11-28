package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// AdAttributedOrderSummary totals the paid orders credited to ads over the
// period being read, independently of how the rows are paged.
//
// Orders counts every attributed order, of which GuestOrders had no customer
// account and NewCustomers were a customer's first paid order. Revenue is a
// list by currency, net of tax and discount and excluding shipping.
// MedianClickToOrderSeconds is a median rather than a mean, so one slow
// shopper does not distort it; it is absent when there is nothing to measure.
type AdAttributedOrderSummary struct {
	Orders                    int64         `json:"orders,omitempty"`
	NewCustomers              int64         `json:"new_customers,omitempty"`
	GuestOrders               int64         `json:"guest_orders,omitempty"`
	Revenue                   []money.Money `json:"revenue,omitempty"`
	MedianClickToOrderSeconds int64         `json:"median_click_to_order_seconds,omitempty"`
}
