package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

// OrganisationAccess links a user account/persona to a wholesale organisation
// and carries the organisation-scoped role key.
type OrganisationAccess struct {
	ID                        string                         `json:"id"`
	WholesaleOrganisationCode string                         `json:"wholesale_organisation_code"`
	UserID                    string                         `json:"user_id"`
	AccountID                 string                         `json:"account_id"`
	RoleKey                   string                         `json:"role_key"`
	Status                    enums.OrganisationAccessStatus `json:"status"`
	JobTitle                  string                         `json:"job_title,omitempty"`
	IsPrimary                 bool                           `json:"is_primary,omitempty"`
	Invitation                *common.LifecycleAction        `json:"invitation,omitempty"`
	JoinedAt                  *time.Time                     `json:"joined_at,omitempty"`
	Suspension                *common.LifecycleAction        `json:"suspension,omitempty"`
	Revocation                *common.LifecycleAction        `json:"revocation,omitempty"`
	ExpiresAt                 *time.Time                     `json:"expires_at,omitempty"`
	Metadata                  common.Metadata                `json:"metadata,omitempty"`

	common.AuditFields
}

// OrganisationAccessSummary is the compact access projection carried by
// access/session and organisation contact responses.
type OrganisationAccessSummary struct {
	ID                        string                         `json:"id"`
	WholesaleOrganisationCode string                         `json:"wholesale_organisation_code"`
	UserID                    string                         `json:"user_id"`
	AccountID                 string                         `json:"account_id,omitempty"`
	RoleKey                   string                         `json:"role_key"`
	Status                    enums.OrganisationAccessStatus `json:"status"`
	JobTitle                  string                         `json:"job_title,omitempty"`
	IsPrimary                 bool                           `json:"is_primary,omitempty"`
}
