package shared

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/common"
)

// SecurityPolicySettings is the shared shape for backend-enforced security
// settings. Each service can persist or cache this policy and apply the
// controls that are relevant to its boundary.
type SecurityPolicySettings struct {
	ID                               string          `json:"id"`
	TenantID                         string          `json:"tenant_id,omitempty"`
	Environment                      string          `json:"environment,omitempty"` // "dev", "staging", "prod"
	MFARequiredForStaff              bool            `json:"mfa_required_for_staff"`
	MFARequiredForAdmins             bool            `json:"mfa_required_for_admins"`
	PrivilegedActionsRequireReauth   bool            `json:"privileged_actions_require_reauth"`
	PIIExportRequiresMFA             bool            `json:"pii_export_requires_mfa"`
	SessionIdleTimeoutMinutes        int             `json:"session_idle_timeout_minutes"`
	SessionAbsoluteTimeoutMinutes    int             `json:"session_absolute_timeout_minutes"`
	AccessTokenTTLSeconds            int             `json:"access_token_ttl_seconds"`
	RefreshTokenTTLSeconds           int             `json:"refresh_token_ttl_seconds"`
	RefreshTokenRotationEnabled      bool            `json:"refresh_token_rotation_enabled"`
	PasswordMinLength                int             `json:"password_min_length"`
	PasswordRequireUppercase         bool            `json:"password_require_uppercase"`
	PasswordRequireLowercase         bool            `json:"password_require_lowercase"`
	PasswordRequireNumber            bool            `json:"password_require_number"`
	PasswordRequireSymbol            bool            `json:"password_require_symbol"`
	FailedLoginLockoutThreshold      int             `json:"failed_login_lockout_threshold"`
	FailedLoginLockoutMinutes        int             `json:"failed_login_lockout_minutes"`
	AuditLogRetentionDays            int             `json:"audit_log_retention_days"`
	AccessLogRetentionDays           int             `json:"access_log_retention_days"`
	SecurityEventRetentionDays       int             `json:"security_event_retention_days"`
	EvidenceRetentionDays            int             `json:"evidence_retention_days"`
	DataRetentionDefaultDays         int             `json:"data_retention_default_days"`
	DataDeletionGraceDays            int             `json:"data_deletion_grace_days"`
	EncryptionAtRestRequired         bool            `json:"encryption_at_rest_required"`
	EncryptionInTransitRequired      bool            `json:"encryption_in_transit_required"`
	IntegrityHashRequiredForAuditLog bool            `json:"integrity_hash_required_for_audit_log"`
	LastReviewedAt                   *time.Time      `json:"last_reviewed_at,omitempty"`
	ReviewedBy                       string          `json:"reviewed_by,omitempty"`
	Metadata                         common.Metadata `json:"metadata,omitempty"`

	common.AuditFields
}
