package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v5/pkg/enums"
)

// UserProfile is the public projection of a user account. Secret fields such
// as password hashes and refresh token material never appear here —
// they live only inside the service that manages identity.
type UserProfile struct {
	ID                      string                       `json:"id"`
	Email                   string                       `json:"email"`
	DisplayName             string                       `json:"display_name,omitempty"`
	Active                  bool                         `json:"active"`
	UserRole                enums.UserRole               `json:"user_role"`
	MFAEnabled              bool                         `json:"mfa_enabled,omitempty"`
	NotificationPreferences *UserNotificationPreferences `json:"notification_preferences,omitempty"`
	LastLoginAt             *time.Time                   `json:"last_login_at,omitempty"`
	PasswordChangedAt       *time.Time                   `json:"password_changed_at,omitempty"`
	AccessReviewedAt        *time.Time                   `json:"access_reviewed_at,omitempty"`

	common.AuditFields
}
