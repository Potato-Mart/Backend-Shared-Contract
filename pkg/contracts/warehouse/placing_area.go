// Package warehouse holds contracts that describe the physical
// storage layout of a depot: placing areas (a labelled cargo
// location inside a depot where one or more SKUs live).
//
// Placing areas are referenced by Code (a short, human-readable
// identifier such as "FROZEN-A1-01") rather than by ID, because
// operators on the warehouse floor type the code directly.
package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/pkg/enums"
)

// PlacingArea is a labelled physical location inside a depot.
// Each Product references a PlacingArea by its Code.
//
// Storage on the area constrains which products can live there:
// only products whose SKU category storage matches PlacingArea.Storage
// may be assigned to this area.
type PlacingArea struct {
	ID          string            `json:"id"`
	Code        string            `json:"code"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	DepotCode   string            `json:"depot_code,omitempty"`
	Zone        string            `json:"zone,omitempty"`
	Storage     enums.StorageType `json:"storage"`
	Capacity    int               `json:"capacity,omitempty"`
	IsActive    bool              `json:"is_active"`
	SortOrder   int               `json:"sort_order"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
