package warehouse

import "github.com/Potato-Mart/Backend-Shared-Contract/v22/pkg/common"

// Depot is a fulfilment site within a depot region. Code is its canonical
// cross-service reference and Timezone is an IANA time-zone identifier.
type Depot struct {
	ID         string         `json:"id"`
	Code       string         `json:"code"`
	Name       string         `json:"name"`
	RegionCode string         `json:"region_code"`
	Address    common.Address `json:"address"`
	Timezone   string         `json:"timezone"`
	Phone      string         `json:"phone,omitempty"`
	IsActive   bool           `json:"is_active"`
	LayoutID   string         `json:"layout_id,omitempty"`

	common.AuditFields
}
