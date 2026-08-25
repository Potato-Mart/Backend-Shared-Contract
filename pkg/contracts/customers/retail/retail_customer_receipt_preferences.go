package retail

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/customers/retail/retail_enums"
)

// RetailCustomerReceiptPreferences groups the customer's receipt format
// election with provenance. An absent group means electronic-only (the
// default election); Formats never persists empty.
type RetailCustomerReceiptPreferences struct {
	Formats   []retail_enums.ReceiptFormat `json:"formats"`
	UpdatedAt *time.Time                   `json:"updated_at,omitempty"`
	Source    string                       `json:"source,omitempty"`
}
