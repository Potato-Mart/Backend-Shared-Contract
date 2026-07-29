package wholesale

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/common"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/wholesale"
)

type GroupOrderManagerApplicantSnapshot struct {
	UserID               string `json:"user_id"`
	RetailAccountID      string `json:"retail_account_id"`
	RetailCustomerNumber string `json:"retail_customer_number"`
	Name                 string `json:"name,omitempty"`
	Email                string `json:"email,omitempty"`
}

type GroupOrderManagerApplication struct {
	ID                        string                                           `json:"id,omitempty"`
	ApplicationNumber         string                                           `json:"application_number"`
	Applicant                 GroupOrderManagerApplicantSnapshot               `json:"applicant"`
	ProposedGroupName         string                                           `json:"proposed_group_name"`
	ApplicantNote             string                                           `json:"applicant_note,omitempty"`
	Status                    wholesaleenum.GroupOrderManagerApplicationStatus `json:"status"`
	WholesaleAccountID        string                                           `json:"wholesale_account_id,omitempty"`
	WholesaleOrganisationCode string                                           `json:"wholesale_organisation_code,omitempty"`
	OrganisationAccessID      string                                           `json:"organisation_access_id,omitempty"`
	SubmittedAt               time.Time                                        `json:"submitted_at"`
	ReviewedAt                *time.Time                                       `json:"reviewed_at,omitempty"`
	ReviewedBy                string                                           `json:"reviewed_by,omitempty"`
	DecisionReason            string                                           `json:"decision_reason,omitempty"`

	common.AuditFields
}
