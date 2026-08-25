package access

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/account/account_enums"
	"time"
)

// PortalAccessGrantedEvent is emitted when an account is granted access to a
// portal.
type PortalAccessGrantedEvent struct {
	PortalAccessID string                    `json:"portal_access_id"`
	UserID         string                    `json:"user_id"`
	AccountID      string                    `json:"account_id"`
	AccountType    account_enums.AccountType `json:"account_type"`
	Portal         identity_enums.Portal     `json:"portal"`
	GrantedBy      string                    `json:"granted_by,omitempty"`
	GrantedAt      time.Time                 `json:"granted_at"`
	ExpiresAt      *time.Time                `json:"expires_at,omitempty"`
	RequestID      string                    `json:"request_id,omitempty"`
}
