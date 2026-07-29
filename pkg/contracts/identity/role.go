package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	identityenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/identity"
)

// Role is the projection of a role definition stored in mgmt_roles.
//
// The role enum value (Key) is the wire-stable identifier used by JWT
// claims and middleware checks; Permissions is the list of permission
// strings granted to anyone holding this role.
//
// IsSystem marks the six built-in workforce roles that the platform ships
// with (superAdmin, admin, sales, warehouse, warehouseOperator, marketing).
// System roles cannot be deleted but their permissions can
// be tweaked by a superAdmin.
type Role struct {
	Key                         identityenum.UserRole `json:"key"`
	Label                       string                `json:"label"`
	Description                 string                `json:"description,omitempty"`
	Permissions                 []string              `json:"permissions"`
	IsSystem                    bool                  `json:"is_system"`
	OwnerID                     string                `json:"owner_id,omitempty"`
	LeastPrivilegeJustification string                `json:"least_privilege_justification,omitempty"`
	AccessReviewedAt            *time.Time            `json:"access_reviewed_at,omitempty"`

	common.AuditFields
}
