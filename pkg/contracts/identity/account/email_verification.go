package account

import "time"

// EmailVerificationAudit is the immutable record created when an administrator
// verifies a customer's email manually.
type EmailVerificationAudit struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Email       string    `json:"email"`
	ActorID     string    `json:"actor_id"`
	ActorEmail  string    `json:"actor_email,omitempty"`
	Reason      string    `json:"reason"`
	WasVerified bool      `json:"was_verified"`
	IsVerified  bool      `json:"is_verified"`
	VerifiedAt  time.Time `json:"verified_at"`
}
