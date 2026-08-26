package retail

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/customers/retail/retail_enums"
)

// RetailCustomerLifecycle groups the retail customer profile lifecycle.
type RetailCustomerLifecycle struct {
	Status        retail_enums.CustomerStatus `json:"status"`
	RegisteredAt  *time.Time                  `json:"registered_at,omitempty"`
	ActivatedAt   *time.Time                  `json:"activated_at,omitempty"`
	BlockedAt     *time.Time                  `json:"blocked_at,omitempty"`
	BlockedReason string                      `json:"blocked_reason,omitempty"`
	ClosedAt      *time.Time                  `json:"closed_at,omitempty"`
	ClosedReason  string                      `json:"closed_reason,omitempty"`
}
