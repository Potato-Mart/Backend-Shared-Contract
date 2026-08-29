package account

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"

// AdminAccountProfile contains optional workforce profile data for an
// adminUser account.
type AdminAccountProfile struct {
	ID             string   `json:"id,omitempty"`
	UserID         string   `json:"user_id"`
	AccountID      string   `json:"account_id"`
	EmployeeID     string   `json:"employee_id,omitempty"`
	Department     string   `json:"department,omitempty"`
	JobTitle       string   `json:"job_title,omitempty"`
	SupportRegion  string   `json:"support_region,omitempty"`
	SupportRegions []string `json:"support_regions,omitempty"`
	ManagerUserID  string   `json:"manager_user_id,omitempty"`

	audit.AuditFields
}
