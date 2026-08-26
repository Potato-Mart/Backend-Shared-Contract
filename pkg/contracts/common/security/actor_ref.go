package security

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security/security_enums"

// ActorRef identifies the authenticated principal that performed an action.
// It is shared by audit, access, and security records, which customers,
// workforce users, and background services all reach.
//
// ActorDomain names the trust domain that acted, so a customer-caused change
// is distinguishable from a service-caused one on the same timeline. An absent
// domain is no evidence of an actor; a consumer attributing an action fails
// closed rather than assuming one.
//
// ActorRole is an open code because the role vocabulary follows ActorDomain: a
// workforce actor carries one of the closed role_enums.UserRole keys and a
// wholesale buyer carries an organisation buyer role key. Typing this field to
// either catalogue would make the other unrecordable.
type ActorRef struct {
	ActorID     string                        `json:"actor_id,omitempty"`
	ActorEmail  string                        `json:"actor_email,omitempty"`
	ActorDomain security_enums.IdentityDomain `json:"actor_domain,omitempty"`
	ActorRole   string                        `json:"actor_role,omitempty"`
}
