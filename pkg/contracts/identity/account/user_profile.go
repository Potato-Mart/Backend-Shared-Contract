package account

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/audit"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/identity/access"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/identity/account/account_enums"
)

// UserProfile is the public projection of a canonical user. Secret fields such
// as password hashes and refresh token material never appear here —
// they live only inside the service that manages identity.
type UserProfile struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	// Phone is the Identity-owned canonical E.164 value. An empty value means
	// that no phone has been registered; normalization and verification policy
	// remain owned by Identity.
	Phone              string                    `json:"phone,omitempty"`
	DisplayName        string                    `json:"display_name,omitempty"`
	Avatar             *security.ObjectMedia     `json:"avatar,omitempty"`
	Active             bool                      `json:"active"`
	Accounts           []UserAccountSummary      `json:"accounts,omitempty"`
	PrimaryAccountID   string                    `json:"primary_account_id,omitempty"`
	PrimaryAccountType account_enums.AccountType `json:"primary_account_type,omitempty"`
	MFAEnabled         bool                      `json:"mfa_enabled,omitempty"`
	EmailVerified      bool                      `json:"email_verified"`
	// PhoneVerifiedAt is nil until Phone has been verified. Identity clears it
	// when the canonical phone value changes.
	PhoneVerifiedAt *time.Time `json:"phone_verified_at,omitempty"`
	// GeoScope is the workforce principal's geographic grant. It is absent
	// on customer profiles and on workforce profiles whose scope has not
	// been assigned yet; consumers fail closed on an absent scope rather
	// than widening it.
	GeoScope *access.StaffGeoScope `json:"geo_scope,omitempty"`

	// User Operations
	UserDevice        UserDevice `json:"user_device,omitempty"`
	LastLoginAt       *time.Time `json:"last_login_at,omitempty"`
	PasswordChangedAt *time.Time `json:"password_changed_at,omitempty"`
	AccessReviewedAt  *time.Time `json:"access_reviewed_at,omitempty"`

	audit.AuditFields
	security.DataProtectionFields
}
