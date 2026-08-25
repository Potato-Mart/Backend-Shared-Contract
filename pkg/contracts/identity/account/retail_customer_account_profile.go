package account

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"

// RetailCustomerAccountProfile contains lightweight references for a retail
// customer account. CustomerNumber is also the membership account key.
type RetailCustomerAccountProfile struct {
	ID             string `json:"id,omitempty"`
	UserID         string `json:"user_id"`
	AccountID      string `json:"account_id"`
	CustomerNumber string `json:"customer_number,omitempty"`
	ReferralCode   string `json:"referral_code,omitempty"`

	audit.AuditFields
}
