package access

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/customers/wholesale/wholesale_enums"
	"time"
)

// OrganisationAccessChangedEvent is emitted when wholesale organisation access
// changes status or role.
type OrganisationAccessChangedEvent struct {
	OrganisationAccessID      string                                   `json:"organisation_access_id"`
	WholesaleOrganisationCode string                                   `json:"wholesale_organisation_code"`
	UserID                    string                                   `json:"user_id"`
	AccountID                 string                                   `json:"account_id,omitempty"`
	RoleKey                   string                                   `json:"role_key,omitempty"`
	PreviousStatus            wholesale_enums.OrganisationAccessStatus `json:"previous_status,omitempty"`
	Status                    wholesale_enums.OrganisationAccessStatus `json:"status"`
	ChangedBy                 string                                   `json:"changed_by,omitempty"`
	ChangedAt                 time.Time                                `json:"changed_at"`
	Reason                    string                                   `json:"reason,omitempty"`
	RequestID                 string                                   `json:"request_id,omitempty"`
}
