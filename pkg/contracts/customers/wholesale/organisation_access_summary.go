package wholesale

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/wholesale/wholesale_enums"
)

// OrganisationAccessSummary is the compact access projection carried by
// access/session and organisation contact responses.
type OrganisationAccessSummary struct {
	ID                        string                                   `json:"id"`
	WholesaleOrganisationCode string                                   `json:"wholesale_organisation_code"`
	UserID                    string                                   `json:"user_id"`
	AccountID                 string                                   `json:"account_id,omitempty"`
	RoleKey                   wholesale_enums.WholesaleBuyerRole       `json:"role_key"`
	Status                    wholesale_enums.OrganisationAccessStatus `json:"status"`
	Name                      party.PersonName                         `json:"name,omitempty"`
	Contacts                  party.ContactChannels                    `json:"contacts,omitempty"`
	JobTitle                  string                                   `json:"job_title,omitempty"`
	Department                string                                   `json:"department,omitempty"`
	IsPrimary                 bool                                     `json:"is_primary,omitempty"`
}
