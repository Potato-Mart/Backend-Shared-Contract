package geography

import "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography/geography_enums"

// GeographicTarget names one inclusive target in the geographic hierarchy.
type GeographicTarget struct {
	Kind geography_enums.GeographicTargetKind `json:"kind"`
	Code string                               `json:"code"`
}
