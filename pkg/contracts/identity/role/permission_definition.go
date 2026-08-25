package role

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/identity/role/role_enums"
)

// PermissionDefinition is the display and risk metadata for one workforce
// permission. Identity owns the catalogue contents and role policy.
type PermissionDefinition struct {
	Key            role_enums.PermissionKey            `json:"key"`
	Label          string                              `json:"label"`
	Description    string                              `json:"description,omitempty"`
	Module         string                              `json:"module"`
	RiskLevel      security_enums.SecurityRiskLevel    `json:"risk_level"`
	RequiresMFA    bool                                `json:"requires_mfa"`
	Classification role_enums.PermissionClassification `json:"classification"`
}
