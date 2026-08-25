package account

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/account/account_enums"
)

// UserAccountSummary is the compact account/persona projection returned with
// user and access-resolution responses.
type UserAccountSummary struct {
	ID             string                      `json:"id"`
	UserID         string                      `json:"user_id"`
	AccountType    account_enums.AccountType   `json:"account_type"`
	Status         account_enums.AccountStatus `json:"status"`
	DisplayName    string                      `json:"display_name,omitempty"`
	PrimaryEmail   string                      `json:"primary_email,omitempty"`
	AllowedPortals []identity_enums.Portal     `json:"allowed_portals,omitempty"`
}
