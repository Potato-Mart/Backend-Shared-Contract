package authorisation

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
)

// Role is the projection of a role definition stored in mgmt_roles.
//
// Key is an open wire-stable RoleCode used by JWT claims and middleware
// checks. Built-in values retain the closed role_enums.UserRole vocabulary;
// owner-defined roles may use other non-empty keys. Permissions is the list
// of permission keys granted to anyone holding this role.
//
// IsSystem marks the six built-in workforce roles that the platform ships
// with (superAdmin, countryAdmin, depotManager, marketing, warehouseManager,
// warehouseOperator). System roles cannot be deleted but their permissions
// can be tweaked by a superAdmin.
//
// Rank is the role's hierarchy position; a lower number is more privileged.
// Built-in roles hold 1 (superAdmin) through 6 (warehouseOperator). An
// owner-defined custom role carries the rank its owner assigned (2 or
// higher), which Backend-Identity stores and enforces. Rank is omitted only
// for a legacy custom role written without one.
type Role struct {
	Key                         RoleCode        `json:"key"`
	Label                       string          `json:"label"`
	Description                 string          `json:"description,omitempty"`
	Permissions                 []PermissionKey `json:"permissions"`
	Rank                        int             `json:"rank,omitempty"`
	IsSystem                    bool            `json:"is_system"`
	OwnerID                     string          `json:"owner_id,omitempty"`
	LeastPrivilegeJustification string          `json:"least_privilege_justification,omitempty"`
	AccessReviewedAt            *time.Time      `json:"access_reviewed_at,omitempty"`

	audit.AuditFields
}
