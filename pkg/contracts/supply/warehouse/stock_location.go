package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geometry"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// StockLocation is a depot-qualified physical inventory location.
type StockLocation struct {
	ID                    string                                      `json:"id"`
	DepotCode             string                                      `json:"depot_code"`
	LocationCode          string                                      `json:"location_code"`
	Name                  string                                      `json:"name,omitempty"`
	StorageType           warehouse_enums.LocationStorageType         `json:"storage_type"`
	Purpose               warehouse_enums.StockLocationPurpose        `json:"purpose"`
	HandlingMode          warehouse_enums.StockLocationHandlingMode   `json:"handling_mode"`
	Access                warehouse_enums.StockLocationAccess         `json:"access"`
	CollectionMode        warehouse_enums.StockLocationCollectionMode `json:"collection_mode"`
	CollectionEligibility []StockLocationCollectionEligibility        `json:"collection_eligibility,omitempty"`
	IsSystemManaged       bool                                        `json:"is_system_managed"`
	IsActive              bool                                        `json:"is_active"`
	LayoutNodeID          string                                      `json:"layout_node_id,omitempty"`
	Transform             *geometry.Transform                         `json:"transform,omitempty"`
	Size                  *geometry.Size3D                            `json:"size,omitempty"`
	Shape                 warehouse_enums.ShapeType                   `json:"shape,omitempty"`
	Color                 string                                      `json:"color,omitempty"`
	// Category is a short free-text area label for the location, for example
	// "飲料區". It is a display grouping only.
	Category string `json:"category,omitempty"`
	// CapacityCartons is the whole-case capacity of the location. Nil means
	// unlimited; a set value is at least 1.
	CapacityCartons *int `json:"capacity_cartons,omitempty"`

	audit.AuditFields
}
