package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
)

// Depot is a fulfilment site within a depot region. Code is its canonical
// cross-service reference and Timezone is an IANA time-zone identifier.
type Depot struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	RegionCode string `json:"region_code"`
	// CountryCode is the denormalized country the depot trades in, carried
	// so a country-scoped staff query is a plain indexed match. Depots are
	// the only site identity in the platform: some depots trade as stores,
	// and there is no separate store record or store code.
	CountryCode geography.CountryCode `json:"country_code,omitempty"`
	Address     geography.Address     `json:"address"`
	Timezone    string                `json:"timezone"`
	Phone       string                `json:"phone,omitempty"`
	IsActive    bool                  `json:"is_active"`
	LayoutID    string                `json:"layout_id,omitempty"`

	audit.AuditFields
}
