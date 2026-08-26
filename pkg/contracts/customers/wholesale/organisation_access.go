package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/customers/wholesale/wholesale_enums"
)

// OrganisationAccess links a user account/persona to a wholesale organisation
// and carries the organisation-scoped role key plus the person's team profile.
// At most one active access per organisation should be marked primary.
type OrganisationAccess struct {
	ID                        string                                   `json:"id"`
	WholesaleOrganisationCode string                                   `json:"wholesale_organisation_code"`
	UserID                    string                                   `json:"user_id"`
	AccountID                 string                                   `json:"account_id"`
	RoleKey                   wholesale_enums.WholesaleBuyerRole       `json:"role_key"`
	Status                    wholesale_enums.OrganisationAccessStatus `json:"status"`
	Name                      party.PersonName                         `json:"name,omitempty"`
	Contacts                  party.ContactChannels                    `json:"contacts,omitempty"`
	JobTitle                  string                                   `json:"job_title,omitempty"`
	Department                string                                   `json:"department,omitempty"`
	IsPrimary                 bool                                     `json:"is_primary,omitempty"`
	Invitation                *audit.LifecycleAction                   `json:"invitation,omitempty"`
	JoinedAt                  *time.Time                               `json:"joined_at,omitempty"`
	Suspension                *audit.LifecycleAction                   `json:"suspension,omitempty"`
	Revocation                *audit.LifecycleAction                   `json:"revocation,omitempty"`
	ExpiresAt                 *time.Time                               `json:"expires_at,omitempty"`
	Metadata                  metadata.Metadata                        `json:"metadata,omitempty"`

	audit.AuditFields
}
