package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/account/account_enums"
	"time"
)

// UserAccountCreatedEvent is emitted when an account/persona is created for a
// canonical user.
type UserAccountCreatedEvent struct {
	UserID      string                    `json:"user_id"`
	AccountID   string                    `json:"account_id"`
	AccountType account_enums.AccountType `json:"account_type"`
	CreatedBy   string                    `json:"created_by,omitempty"`
	CreatedAt   time.Time                 `json:"created_at"`
	RequestID   string                    `json:"request_id,omitempty"`
}
