package retail

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/customers/retail/retail_enums"
)

// RetailCustomerSummary is a compact retail customer projection for lists,
// search results, and relationship references.
type RetailCustomerSummary struct {
	ID                 string                            `json:"id"`
	AccountID          string                            `json:"account_id,omitempty"`
	UserID             string                            `json:"user_id,omitempty"`
	CustomerNumber     string                            `json:"customer_number,omitempty"`
	DisplayName        string                            `json:"display_name,omitempty"`
	Email              string                            `json:"email,omitempty"`
	Phone              string                            `json:"phone,omitempty"`
	Status             retail_enums.CustomerStatus       `json:"status"`
	Tags               []string                          `json:"tags,omitempty"`
	ReceiptPreferences *RetailCustomerReceiptPreferences `json:"receipt_preferences,omitempty"`
	Metadata           metadata.Metadata                 `json:"metadata,omitempty"`
}
