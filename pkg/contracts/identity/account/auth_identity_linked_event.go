package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/identity/account/account_enums"
	"time"
)

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
