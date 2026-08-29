package group

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	wholesale_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/group/group_enums"
)

type GroupOrderManagerApplication struct {
	ID                        string                                             `json:"id,omitempty"`
	ApplicationNumber         string                                             `json:"application_number"`
	Applicant                 GroupOrderManagerApplicantSnapshot                 `json:"applicant"`
	ProposedGroupName         string                                             `json:"proposed_group_name"`
	ApplicantNote             string                                             `json:"applicant_note,omitempty"`
	Status                    wholesale_enums.GroupOrderManagerApplicationStatus `json:"status"`
	WholesaleAccountID        string                                             `json:"wholesale_account_id,omitempty"`
	WholesaleOrganisationCode string                                             `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                                             `json:"organisation_access_id,omitempty"`
	SubmittedAt               time.Time                                          `json:"submitted_at"`
	ReviewedAt                *time.Time                                         `json:"reviewed_at,omitempty"`
	ReviewedBy                string                                             `json:"reviewed_by,omitempty"`
	DecisionReason            string                                             `json:"decision_reason,omitempty"`

	audit.AuditFields
}
