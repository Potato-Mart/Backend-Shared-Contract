package warehouse

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/geometry"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/classification/classification_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/warehouse/warehouse_enums"
)

const (
	StockLocationCodeQualityHoldAmbient = "SYS-QH-AMBIENT"
	StockLocationCodeQualityHoldChilled = "SYS-QH-CHILLED"
	StockLocationCodeQualityHoldFrozen  = "SYS-QH-FROZEN"
	StockLocationCodeOnlineStageAmbient = "SYS-ONLINE-STAGE-AMBIENT"
	StockLocationCodeOnlineStageChilled = "SYS-ONLINE-STAGE-CHILLED"
	StockLocationCodeOnlineStageFrozen  = "SYS-ONLINE-STAGE-FROZEN"
)

// StockLocationRef qualifies a location code by its depot.
type StockLocationRef struct {
	DepotCode    string `json:"depot_code"`
	LocationCode string `json:"location_code"`
}

// StockLocation is a depot-qualified physical inventory location.
type StockLocation struct {
	ID                    string                                      `json:"id"`
	DepotCode             string                                      `json:"depot_code"`
	LocationCode          string                                      `json:"location_code"`
	Name                  string                                      `json:"name,omitempty"`
	StorageType           classification_enums.StorageType            `json:"storage_type"`
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
}

// StockLocationCollectionEligibility identifies a collection's primary or
// ordered overflow placement at a location.
type StockLocationCollectionEligibility struct {
	CollectionCode string                                      `json:"collection_code"`
	Role           warehouse_enums.StockLocationCollectionRole `json:"role"`
	Priority       int                                         `json:"priority"`
}

// StockLocationAssignment links one product SKU to one depot-qualified
// location independently of its current quantity.
type StockLocationAssignment struct {
	ID                       string `json:"id"`
	DepotCode                string `json:"depot_code"`
	LocationCode             string `json:"location_code"`
	SKUCode                  string `json:"sku_code"`
	ElectronicShelfLabelCode string `json:"electronic_shelf_label_code,omitempty"`
	IsActive                 bool   `json:"is_active"`

	audit.AuditFields
}

// StockLocationProductBalance is a revisioned quantity projection for one
// stock-location assignment.
type StockLocationProductBalance struct {
	AssignmentID       string                               `json:"assignment_id"`
	DepotCode          string                               `json:"depot_code"`
	LocationCode       string                               `json:"location_code"`
	SKUCode            string                               `json:"sku_code"`
	PackageComposition packaging.PackageCompositionSnapshot `json:"package_composition"`
	OnHandBaseUnits    int64                                `json:"on_hand_base_units"`
	ReservedBaseUnits  int64                                `json:"reserved_base_units"`
	AvailableBaseUnits int64                                `json:"available_base_units"`
	IsOutOfStock       bool                                 `json:"is_out_of_stock"`
	Revision           int64                                `json:"revision"`
	LastRestockedAt    *time.Time                           `json:"last_restocked_at,omitempty"`
	DepotTimezone      string                               `json:"depot_timezone"`
	AsOf               time.Time                            `json:"as_of"`
}
