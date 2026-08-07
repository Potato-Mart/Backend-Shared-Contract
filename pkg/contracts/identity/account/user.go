package account

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/identity/account/account_enums"
)

// UserProfile is the public projection of a canonical user. Secret fields such
// as password hashes and refresh token material never appear here —
// they live only inside the service that manages identity.
type UserProfile struct {
	ID                      string                       `json:"id"`
	Email                   string                       `json:"email"`
	DisplayName             string                       `json:"display_name,omitempty"`
	AvatarMediaID           string                       `json:"avatar_media_id,omitempty"`
	AvatarURL               string                       `json:"avatar_url,omitempty"`
	Active                  bool                         `json:"active"`
	Accounts                []UserAccountSummary         `json:"accounts,omitempty"`
	PrimaryAccountID        string                       `json:"primary_account_id,omitempty"`
	PrimaryAccountType      account_enums.AccountType    `json:"primary_account_type,omitempty"`
	MFAEnabled              bool                         `json:"mfa_enabled,omitempty"`
	EmailVerified           bool                         `json:"email_verified"`
	NotificationPreferences *UserNotificationPreferences `json:"notification_preferences,omitempty"`

	// User Operations
	UserDevice        UserDevice `json:"user_device,omitempty"`
	LastLoginAt       *time.Time `json:"last_login_at,omitempty"`
	PasswordChangedAt *time.Time `json:"password_changed_at,omitempty"`
	AccessReviewedAt  *time.Time `json:"access_reviewed_at,omitempty"`

	audit.AuditFields
}
