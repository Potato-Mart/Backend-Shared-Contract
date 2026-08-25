package packaging

import measurement "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/measurement"

// PhysicalPackage is a package-level shipping measurement snapshot.
type PhysicalPackage struct {
	Dimensions *measurement.Dimensions `json:"dimensions,omitempty"`
	Weight     *measurement.Weight     `json:"weight,omitempty"`
	Quantity   int                     `json:"quantity,omitempty"`
}
