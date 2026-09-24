package courier

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/courier/courier_enums"
)

// DeliveryCompany is the Supply-owned, admin-safe delivery company catalogue
// record. Code identifies an instance; Integration retains the service-owned
// api/manual mode. Adapter is the separate open server-side adapter identifier.
// Adding a record does not install an adapter.
// Revision identifies the entire configuration, including areas and schedules.
type DeliveryCompany struct {
	Code                string `json:"code"`
	Name                string `json:"name"`
	Integration         string `json:"integration"`
	Adapter             string `json:"adapter,omitempty"`
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
	// Legacy configured windows remain in the JSON shape for compatibility.
	// New availability and routing decisions must use provider-backed slots.
	Schedules []DeliveryServiceWindow `json:"schedules,omitempty"`

	audit.AuditFields
}
