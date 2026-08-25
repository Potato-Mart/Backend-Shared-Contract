package order

import "time"

// PreorderItemSnapshot is server-stamped from Supply's active SKU policy.
// Clients cannot choose or alter preorder state.
type PreorderItemSnapshot struct {
	SKUCode                string     `json:"sku_code"`
	PolicyVersion          string     `json:"policy_version"`
	ExpectedAvailableAt    *time.Time `json:"expected_available_at,omitempty"`
	ScheduleTimezone       string     `json:"schedule_timezone"`
	MaxQuantityPerOrder    int        `json:"max_quantity_per_order,omitempty"`
	MaxQuantityPerCustomer int        `json:"max_quantity_per_customer,omitempty"`
	CapturedAt             time.Time  `json:"captured_at"`
}
