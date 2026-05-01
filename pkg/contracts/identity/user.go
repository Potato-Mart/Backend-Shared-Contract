package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v3/pkg/enums"
)

// UserProfile is the public projection of a user account. Secret fields such
// as password hashes and refresh token material never appear here —
// they live only inside the service that manages identity.
type UserProfile struct {
	ID          string         `json:"id"`
	Email       string         `json:"email"`
	DisplayName string         `json:"display_name,omitempty"`
	Active      bool           `json:"active"`
	UserRole    enums.UserRole `json:"user_role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
