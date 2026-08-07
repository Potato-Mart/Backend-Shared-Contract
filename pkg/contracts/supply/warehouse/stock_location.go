package warehouse

import (
	"time"

	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
)

const (
	StockLocationCodeQualityHoldDry     = "SYS-QH-DRY"
	StockLocationCodeQualityHoldChilled = "SYS-QH-CHILLED"
	StockLocationCodeQualityHoldFrozen  = "SYS-QH-FROZEN"
	StockLocationCodeOnlineStageDry     = "SYS-ONLINE-STAGE-DRY"
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
	ID                    string                               `json:"id"`
	DepotCode             string                               `json:"depot_code"`
	LocationCode          string                               `json:"location_code"`
	Name                  string                               `json:"name,omitempty"`
	StorageType           StorageType                          `json:"storage_type"`
	Purpose               StockLocationPurpose                 `json:"purpose"`
	HandlingMode          StockLocationHandlingMode            `json:"handling_mode"`
	Access                StockLocationAccess                  `json:"access"`
	CollectionMode        StockLocationCollectionMode          `json:"collection_mode"`
	CollectionEligibility []StockLocationCollectionEligibility `json:"collection_eligibility,omitempty"`
	IsSystemManaged       bool                                 `json:"is_system_managed"`
	IsActive              bool                                 `json:"is_active"`
	LayoutNodeID          string                               `json:"layout_node_id,omitempty"`
	Transform             *common.Transform                    `json:"transform,omitempty"`
	Size                  *common.Size3D                       `json:"size,omitempty"`
	Shape                 ShapeType                            `json:"shape,omitempty"`
	Color                 string                               `json:"color,omitempty"`
}

// StockLocationCollectionEligibility identifies a collection's primary or
// ordered overflow placement at a location.
type StockLocationCollectionEligibility struct {
	CollectionID string                      `json:"collection_id"`
	Role         StockLocationCollectionRole `json:"role"`
	Priority     int                         `json:"priority"`
}

// StockLocationAssignment links one product SKU to one depot-qualified
// location independently of its current quantity.
type StockLocationAssignment struct {
	ID                       string `json:"id"`
	DepotCode                string `json:"depot_code"`
	LocationCode             string `json:"location_code"`
	ProductSKUCode           string `json:"product_sku_code"`
	ElectronicShelfLabelCode string `json:"electronic_shelf_label_code,omitempty"`
	IsActive                 bool   `json:"is_active"`

	common.AuditFields
}

// StockLocationProductBalance is a revisioned quantity projection for one
// stock-location assignment.
type StockLocationProductBalance struct {
	AssignmentID       string                            `json:"assignment_id"`
	DepotCode          string                            `json:"depot_code"`
	LocationCode       string                            `json:"location_code"`
	ProductSKUCode     string                            `json:"product_sku_code"`
	PackageComposition common.PackageCompositionSnapshot `json:"package_composition"`
	OnHandBaseUnits    int64                             `json:"on_hand_base_units"`
	ReservedBaseUnits  int64                             `json:"reserved_base_units"`
	AvailableBaseUnits int64                             `json:"available_base_units"`
	IsOutOfStock       bool                              `json:"is_out_of_stock"`
	Revision           int64                             `json:"revision"`
	LastRestockedAt    *time.Time                        `json:"last_restocked_at,omitempty"`
	DepotTimezone      string                            `json:"depot_timezone"`
	AsOf               time.Time                         `json:"as_of"`
}
