package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// UserProfile is the public projection of a canonical user. Secret fields such
// as password hashes and refresh token material never appear here —
// they live only inside the service that manages identity.
type UserProfile struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name,omitempty"`
	Active      bool   `json:"active"`
	// UserRole is a legacy/global RBAC projection kept for backward
	// compatibility. New code must not use it for portal admission; use
	// Accounts, AccountType, and PortalAccess instead.
	UserRole                enums.UserRole               `json:"user_role"`
	Accounts                []UserAccountSummary         `json:"accounts,omitempty"`
	PrimaryAccountID        string                       `json:"primary_account_id,omitempty"`
	PrimaryAccountType      enums.AccountType            `json:"primary_account_type,omitempty"`
	MFAEnabled bool `json:"mfa_enabled,omitempty"`
	// EmailVerified is the customer email-verification state. A nil pointer
	// means "not tracked" (legacy/admin accounts) and is treated as verified;
	// an explicit false means the account must verify its email before it can
	// place orders.
	EmailVerified           *bool                        `json:"email_verified,omitempty"`
	NotificationPreferences *UserNotificationPreferences `json:"notification_preferences,omitempty"`

	// User Operations
	UserDevice        UserDevice `json:"user_device,omitempty"`
	LastLoginAt       *time.Time `json:"last_login_at,omitempty"`
	PasswordChangedAt *time.Time `json:"password_changed_at,omitempty"`
	AccessReviewedAt  *time.Time `json:"access_reviewed_at,omitempty"`

	common.AuditFields `bson:",inline"`
}
