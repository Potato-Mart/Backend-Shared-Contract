package access

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/account/account_enums"
	"time"
)

// PortalAccessRevokedEvent is emitted when an account loses access to a
// portal.
type PortalAccessRevokedEvent struct {
	PortalAccessID string                    `json:"portal_access_id"`
	UserID         string                    `json:"user_id"`
	AccountID      string                    `json:"account_id"`
	AccountType    account_enums.AccountType `json:"account_type"`
	Portal         identity_enums.Portal     `json:"portal"`
	RevokedBy      string                    `json:"revoked_by,omitempty"`
	RevokedAt      time.Time                 `json:"revoked_at"`
	Reason         string                    `json:"reason,omitempty"`
	RequestID      string                    `json:"request_id,omitempty"`
}
