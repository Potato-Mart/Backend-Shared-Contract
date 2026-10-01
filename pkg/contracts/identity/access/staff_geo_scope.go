package access

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v41/pkg/contracts/common/geography/geography_enums"
)

// StaffGeoScope records workforce geographic selections. Services validate
// paths and intersect them with role-derived grants. Selecting GLOBAL never
// elevates a role; existing global super-admin authority remains role-owned.
// Customer principals never carry workforce grants.
type StaffGeoScope struct {
	Mode    geography_enums.GeographicScopeMode `json:"mode"`
	Targets []geography.GeographicPath          `json:"targets,omitempty"`
}
