package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/identity/account/account_enums"
	"time"
)

// UserAccountStatusChangedEvent is emitted when an account/persona lifecycle
// status changes.
type UserAccountStatusChangedEvent struct {
	UserID         string                      `json:"user_id"`
	AccountID      string                      `json:"account_id"`
	AccountType    account_enums.AccountType   `json:"account_type"`
	PreviousStatus account_enums.AccountStatus `json:"previous_status,omitempty"`
	Status         account_enums.AccountStatus `json:"status"`
	ChangedBy      string                      `json:"changed_by,omitempty"`
	ChangedAt      time.Time                   `json:"changed_at"`
	Reason         string                      `json:"reason,omitempty"`
	RequestID      string                      `json:"request_id,omitempty"`
}
