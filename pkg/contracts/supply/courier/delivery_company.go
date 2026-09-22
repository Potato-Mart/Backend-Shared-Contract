package courier

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/courier/courier_enums"
)

// DeliveryCompany is the Supply-owned, admin-safe delivery company catalogue
// record. Code identifies an instance; Integration identifies its server-side
// adapter. Both are open codes. Adding a record does not install an adapter.
// Revision identifies the entire configuration, including areas and schedules.
type DeliveryCompany struct {
	Code                string `json:"code"`
	Name                string `json:"name"`
	Integration         string `json:"integration"`
	Enabled             bool   `json:"enabled"`
	DefaultInstructions string `json:"default_instructions,omitempty"`
	// DispatchCapable is derived by Supply from adapter support and readiness;
	// it is not an editable permission to dispatch.
	DispatchCapable bool                             `json:"dispatch_capable"`
	Revision        int64                            `json:"revision"`
	CountryCodes    []geography.CountryCode          `json:"country_codes"`
	Capabilities    DeliveryCapabilities             `json:"capabilities"`
	SlotSource      courier_enums.DeliverySlotSource `json:"slot_source"`
	Connection      *DeliveryConnection              `json:"connection,omitempty"`
	ServiceAreas    []DeliveryServiceArea            `json:"service_areas"`
	Schedules       []DeliveryServiceWindow          `json:"schedules,omitempty"`

	audit.AuditFields
}
