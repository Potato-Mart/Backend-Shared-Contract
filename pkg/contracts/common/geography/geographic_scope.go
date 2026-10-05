package geography

import "github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/geography/geography_enums"

// GeographicScope is global with no targets or targeted with one or more
// independent hierarchy paths. Owning market/currency remain separate facts.
type GeographicScope struct {
	Mode    geography_enums.GeographicScopeMode `json:"mode"`
	Targets []GeographicPath                    `json:"targets,omitempty"`
}
