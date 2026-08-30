package geography

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/geography/geography_enums"

// GeographicScope is either explicitly global or an inclusive list of
// country, subdivision, depot-region, or depot targets.
type GeographicScope struct {
	Mode    geography_enums.GeographicScopeMode `json:"mode"`
	Targets []GeographicTarget                  `json:"targets,omitempty"`
}
