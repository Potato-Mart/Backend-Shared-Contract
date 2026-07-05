package identity

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

// Role is the projection of a role definition stored in mgmt_roles.
//
// The role enum value (Key) is the wire-stable identifier used by JWT
// claims and middleware checks; Permissions is the list of permission
// strings granted to anyone holding this role.
//
// IsSystem marks the seven built-in roles that the platform ships with
// (superAdmin, admin, sales, warehouse, warehouseOperator, marketing,
// customer). System roles cannot be deleted but their permissions can
// be tweaked by a superAdmin.
type Role struct {
	Key                         enums.UserRole `json:"key"`
	Label                       string         `json:"label"`
	Description                 string         `json:"description,omitempty"`
	Permissions                 []string       `json:"permissions"`
	IsSystem                    bool           `json:"is_system"`
	OwnerID                     string         `json:"owner_id,omitempty"`
	LeastPrivilegeJustification string         `json:"least_privilege_justification,omitempty"`
	AccessReviewedAt            *time.Time     `json:"access_reviewed_at,omitempty"`

	common.AuditFields
}

// Permission is the projection of a permission definition stored in
// mgmt_permissions. Permissions are addressed by a stable dotted key
// (e.g. "customer.write", "promotion.publish") and grouped by module
// for the admin UI.
type Permission struct {
	Key         string                  `json:"key"`
	Label       string                  `json:"label"`
	Description string                  `json:"description,omitempty"`
	Module      string                  `json:"module"`
	IsSystem    bool                    `json:"is_system"`
	RiskLevel   enums.SecurityRiskLevel `json:"risk_level,omitempty"`
	RequiresMFA bool                    `json:"requires_mfa,omitempty"`

	common.AuditFields
}
