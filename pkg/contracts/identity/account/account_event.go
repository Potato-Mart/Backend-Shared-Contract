package account

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security/security_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/identity/account/account_enums"
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

// AuthIdentityLinkedEvent is emitted when a non-secret auth identity is linked
// to a canonical user.
type AuthIdentityLinkedEvent struct {
	AuthIdentityID string                             `json:"auth_identity_id"`
	UserID         string                             `json:"user_id"`
	Provider       account_enums.AuthIdentityProvider `json:"provider"`
	IdentityDomain security_enums.IdentityDomain      `json:"identity_domain"`
	LinkedBy       string                             `json:"linked_by,omitempty"`
	LinkedAt       time.Time                          `json:"linked_at"`
	RequestID      string                             `json:"request_id,omitempty"`
}
