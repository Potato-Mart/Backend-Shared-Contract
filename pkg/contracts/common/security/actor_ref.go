package security

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/role/role_enums"

// ActorRef identifies the authenticated principal that performed an action.
// It is shared by audit, access, and security records.
type ActorRef struct {
	ActorID    string              `json:"actor_id,omitempty"`
	ActorEmail string              `json:"actor_email,omitempty"`
	ActorRole  role_enums.UserRole `json:"actor_role,omitempty"`
}
