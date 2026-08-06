package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestV22ProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := v22StringSet(
		// Earlier hard cut-overs that remain forbidden in v22.
		"MembershipOwnerRef",
		"MembershipOwnerType",
		"EligibleOwnerTypes",
		"PackingSessionStatusSyncPending",
		"SortOrder",
		"BrandKey",
		"BrandSummary",
		"ActiveProductCount",
		"WholesaleProductCount",
		"NotificationQuietHours",
		"QuietHours",
		"FCMDestination",
		"FCMDestinations",
		"PushDestination",
		"PushDestinations",
		"CouponPreview",
		"CouponPreviews",

		// Geography, depot, product, packing, and legacy availability types.
		"PostcodeRule",
		"DepotProduct",
		"ProductPlacement",
		"StorefrontPricing",
		"PackingBoxPlan",
		"PackingBoxContent",
		"StockAdjustedEvent",
		"PreorderStockArrivalEvent",
		"PreorderAvailabilityEvent",
		"BackInStockRestockEvent",
		"ProductSellabilityChangedEvent",
		"DeliveryRegion",
		"PackingSessionStatus",

		// Removed Melbourne-specific delivery values.
		"DeliveryRegionLocalMelbourne",
		"DeliveryRegionRegionalVIC",
		"DeliveryRegionInterstate",

		// Removed expiry-merchandising values.

		// Removed packing-session values.
		"PackingSessionStatusPending",
		"PackingSessionStatusPacking",
		"PackingSessionStatusPacked",
		"PackingSessionStatusResolved",

		// Logical reserve/release and split transfer values are not physical
		// movement types in v22.
		"StockMovementTypePurchaseReceipt",
		"StockMovementTypeSaleReserve",
		"StockMovementTypeSaleRelease",
		"StockMovementTypeDamage",
		"StockMovementTypeTransferIn",
		"StockMovementTypeTransferOut",
		"StockMovementTypeStocktake",

		// Competing stock-arrival and availability event values.
		"EventTypeOrderPreorderAvailable",
		"EventTypeStockArrived",
		"EventTypeStockAdjusted",
		"EventTypeProductSellabilityChanged",
	)

	removedTypes := v22StringSet(
		"contracts/notification.BackInStockRestockEvent",
		"contracts/product.Physical",
		"contracts/product.Pricing",
		"contracts/product.ProductPlacement",
		"contracts/product.StorefrontPricing",
		"contracts/sales.PreorderAvailabilityEvent",
		"contracts/sales.PreorderStockArrivalEvent",
		"contracts/warehouse.DepotProduct",
		"contracts/warehouse.OrderPackingProgress",
		"contracts/warehouse.PackingBoxContent",
		"contracts/warehouse.PackingBoxPlan",
		"contracts/warehouse.PostcodeRule",
		"contracts/warehouse.StockAdjustedEvent",
		"enums/shipping.DeliveryRegion",
		"enums/warehouse.PackingSessionStatus",
	)

	removedFields := map[string]map[string]struct{}{
		"common.Address": v22StringSet(
			"City", "State", "Postcode",
		),
		"contracts/analytics.OrderItemFact": v22StringSet(
			"Quantity",
		),
		"contracts/analytics.RefundItemFact": v22StringSet(
			"Quantity",
		),
		"contracts/analytics.SKUDemandForecast": v22StringSet(
			"SKUCode", "CurrentStockAtRun",
		),
		"contracts/campaign.Audience": v22StringSet(
			"Region",
		),
		"contracts/campaign.CampaignProductPrediction": v22StringSet(
			"PredictedDemandUnits", "SellableAvailableUnits", "ConfirmedInboundUnits",
			"NetRequiredUnits", "SuggestedOrderUnits", "MinimumOrderQuantity",
			"SuggestedCartons", "CartonSize",
		),
		"contracts/campaign.CampaignPredictionEvidence": v22StringSet(
			"RawNetUnits", "NormalizedUnits",
		),
		"contracts/campaign.CampaignSupplierPrediction": v22StringSet(
			"TotalUnits",
		),
		"contracts/category.SKU": v22StringSet(
			"Storage",
		),
		"contracts/pos.CatalogProduct": v22StringSet(
			"ID", "SKU", "Barcode", "Storage", "Price", "CurrentStock",
		),
		"contracts/product.Product": v22StringSet(
			"ID", "SKU", "Barcode", "Storage", "PlacingArea", "CurrentStock",
			"RestockedAt", "ExpiredAt", "Pricing", "Physical",
		),
		"contracts/product.Snapshot": v22StringSet(
			"ID", "SKU", "Storage", "DisplayStatus", "Barcode",
		),
		"contracts/product.StorefrontDisplay":       v22StringSet(),
		"contracts/product.StorefrontMerchandising": v22StringSet(),
		"contracts/product.StorefrontProduct": v22StringSet(
			"SKU", "Barcode", "Storage", "CurrentStock", "Pricing", "ExpiryDate", "DisplayStatus",
		),
		"contracts/purchase.OrderItem": v22StringSet(
			"OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode", "ExpireAt",
		),
		"contracts/purchase.ReceiptItem": v22StringSet(
			"SKU", "OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode",
		),
		"contracts/sales.CartItem": v22StringSet(
			"Price", "Quantity",
		),
		"contracts/sales.DemandBucket": v22StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"contracts/sales.OpenDemandLine": v22StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"contracts/sales.Order": v22StringSet(
			"DeliveryRegion",
		),
		"contracts/sales.OrderItem": v22StringSet(
			"UnitPrice", "Quantity", "CartonQty", "CartonSize",
		),
		"contracts/sales.OrderLineSummary": v22StringSet(
			"Quantity", "UnitPrice",
		),
		"contracts/sales.OrderPackingProgress": v22StringSet(
			"BoxPlan",
		),
		"contracts/sales.PreorderItemState": v22StringSet(
			"OrderedQuantity", "AllocatedQuantity",
		),
		"contracts/sales.VolumeDiscountTier": v22StringSet(
			"MinCartons",
		),
		"contracts/shipping.DeliveryAreaRate": v22StringSet(
			"Postcode", "Suburb", "DeliveryRegion",
		),
		"contracts/shipping.Zone": v22StringSet(
			"States", "Postcodes", "IsLocal",
		),
		"contracts/warehouse.DamageReport": v22StringSet(
			"DamagedQty",
		),
		"contracts/warehouse.Depot": v22StringSet(
			"PostcodeRules",
		),
		"contracts/warehouse.InboundItem": v22StringSet(
			"Barcode", "Storage", "ExpectedQty", "ReceivedQty", "LocationCode",
		),
		"contracts/warehouse.PackingDamage": v22StringSet(
			"SKU", "DamagedQty", "DamageReportID", "StockMovementID",
		),
		"contracts/warehouse.PackingDiscrepancy": v22StringSet(
			"OrderedQty", "ScannedQty", "DiffQty", "UnitPrice", "ReturnToStock",
			"DamageReportID", "StockMovementID", "DamagedQty",
		),
		"contracts/warehouse.PackingLine": v22StringSet(
			"SKU", "OrderedQty", "ScannedQty", "DamagedQty",
		),
		"contracts/warehouse.PickingListItem": v22StringSet(
			"Barcode", "Location", "QuantityRequired", "QuantityPicked",
		),
		"contracts/warehouse.StockLocation": v22StringSet(
			"Code", "Zone",
		),
		"contracts/warehouse.StockLocationProductBalance": v22StringSet(
			"AvailabilityRevision",
		),
		"contracts/warehouse.StockMovement": v22StringSet(
			"SKU", "ProductName", "DepotCode", "LocationCode", "QtyDelta", "BalanceAfter",
			"CreatedBy", "SalesOrderNumber", "ReferenceType", "ReferenceID", "Metadata",
		),
		"contracts/warehouse.WMSDraft": v22StringSet(
			"TotalQty",
		),
		"contracts/warehouse.WMSDraftItem": v22StringSet(
			"Barcode", "LocationCode", "Qty", "ExpiryYM",
		),
	}

	// These JSON keys are unambiguously retired everywhere. Reused names such
	// as state, sku, barcode, quantity, unit_price, pricing, storage, and
	// location_code are checked against their exact owning type below.
	removedJSONKeys := v22StringSet(
		"sort_order",
		"quiet_hours",
		"fcm_destination",
		"fcm_destinations",
		"fcm_token",
		"coupon_preview",
		"coupon_previews",
		"placing_area",
		"current_stock",
		"current_stock_at_run",
		"restocked_at",
		"carton_qty",
		"carton_size",
		"min_cartons",
		"suggested_cartons",
		"expiry_ym",
		"delivery_region",
		"box_plan",
		"ambient_boxes",
		"frozen_boxes",
		"manual_ambient_count",
		"manual_frozen_count",
		"postcode_rules",
	)

	removedJSONKeysByType := map[string]map[string]struct{}{
		"common.Address": v22StringSet(
			"city", "state", "postcode",
		),
		"contracts/analytics.OrderItemFact": v22StringSet(
			"quantity",
		),
		"contracts/analytics.RefundItemFact": v22StringSet(
			"quantity",
		),
		"contracts/analytics.SKUDemandForecast": v22StringSet(
			"sku_code", "current_stock_at_run",
		),
		"contracts/campaign.Audience": v22StringSet(
			"region",
		),
		"contracts/campaign.CampaignProductPrediction": v22StringSet(
			"predicted_demand_units", "sellable_available_units", "confirmed_inbound_units",
			"net_required_units", "suggested_order_units", "minimum_order_quantity",
			"suggested_cartons", "carton_size",
		),
		"contracts/campaign.CampaignPredictionEvidence": v22StringSet(
			"raw_net_units", "normalized_units",
		),
		"contracts/campaign.CampaignSupplierPrediction": v22StringSet(
			"total_units",
		),
		"contracts/category.SKU": v22StringSet(
			"storage",
		),
		"contracts/pos.CatalogProduct": v22StringSet(
			"id", "sku", "barcode", "storage", "price", "current_stock",
		),
		"contracts/product.Product": v22StringSet(
			"id", "sku", "barcode", "storage", "placing_area", "current_stock",
			"restocked_at", "expired_at", "pricing", "physical",
		),
		"contracts/product.Snapshot": v22StringSet(
			"id", "sku", "storage", "display_status", "barcode",
		),
		"contracts/product.StorefrontDisplay":       v22StringSet(),
		"contracts/product.StorefrontMerchandising": v22StringSet(),
		"contracts/product.StorefrontProduct": v22StringSet(
			"sku", "barcode", "storage", "current_stock", "pricing", "expiry_date", "display_status",
		),
		"contracts/purchase.OrderItem": v22StringSet(
			"ordered_qty", "received_qty", "rejected_qty", "location_code", "expire_at",
		),
		"contracts/purchase.ReceiptItem": v22StringSet(
			"sku", "ordered_qty", "received_qty", "rejected_qty", "location_code",
		),
		"contracts/sales.CartItem": v22StringSet(
			"price", "quantity",
		),
		"contracts/sales.DemandBucket": v22StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"contracts/sales.OpenDemandLine": v22StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"contracts/sales.Order": v22StringSet(
			"delivery_region",
		),
		"contracts/sales.OrderItem": v22StringSet(
			"unit_price", "quantity", "carton_qty", "carton_size",
		),
		"contracts/sales.OrderLineSummary": v22StringSet(
			"quantity", "unit_price",
		),
		"contracts/sales.OrderPackingProgress": v22StringSet(
			"box_plan",
		),
		"contracts/sales.PreorderItemState": v22StringSet(
			"ordered_quantity", "allocated_quantity",
		),
		"contracts/sales.VolumeDiscountTier": v22StringSet(
			"min_cartons",
		),
		"contracts/shipping.DeliveryAreaRate": v22StringSet(
			"postcode", "suburb", "delivery_region",
		),
		"contracts/shipping.Zone": v22StringSet(
			"states", "postcodes", "is_local",
		),
		"contracts/warehouse.DamageReport": v22StringSet(
			"damaged_qty",
		),
		"contracts/warehouse.Depot": v22StringSet(
			"postcode_rules",
		),
		"contracts/warehouse.InboundItem": v22StringSet(
			"barcode", "storage", "expected_qty", "received_qty", "location_code",
		),
		"contracts/warehouse.PackingDamage": v22StringSet(
			"sku", "damaged_qty", "damage_report_id", "stock_movement_id",
		),
		"contracts/warehouse.PackingDiscrepancy": v22StringSet(
			"ordered_qty", "scanned_qty", "diff_qty", "unit_price", "return_to_stock",
			"damage_report_id", "stock_movement_id", "damaged_qty",
		),
		"contracts/warehouse.PackingLine": v22StringSet(
			"sku", "ordered_qty", "scanned_qty", "damaged_qty",
		),
		"contracts/warehouse.PickingListItem": v22StringSet(
			"barcode", "location", "quantity_required", "quantity_picked",
		),
		"contracts/warehouse.StockLocation": v22StringSet(
			"code", "zone",
		),
		"contracts/warehouse.StockLocationProductBalance": v22StringSet(
			"availability_revision",
		),
		"contracts/warehouse.StockMovement": v22StringSet(
			"sku", "product_name", "depot_code", "location_code", "qty_delta", "balance_after",
			"created_by", "sales_order_number", "reference_type", "reference_id", "metadata",
		),
		"contracts/warehouse.WMSDraft": v22StringSet(
			"total_qty",
		),
		"contracts/warehouse.WMSDraftItem": v22StringSet(
			"barcode", "location_code", "qty", "expiry_ym",
		),
	}

	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if parseErr != nil {
			return parseErr
		}
		for _, group := range file.Comments {
			if strings.Contains(group.Text(), "Deprecated:") || strings.Contains(group.Text(), "@deprecated") {
				t.Errorf("%s contains a deprecated production declaration", path)
			}
		}

		ast.Inspect(file, func(node ast.Node) bool {
			identifier, ok := node.(*ast.Ident)
			if !ok {
				return true
			}
			if _, removed := removedIdentifiers[identifier.Name]; removed {
				t.Errorf("%s contains removed v22 identifier %s", path, identifier.Name)
			}
			return true
		})

		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if !ok {
					continue
				}

				typeKey := v22ProductionTypeKey(path, typeSpecification.Name.Name)
				if _, removed := removedTypes[typeKey]; removed {
					t.Errorf("%s declares removed v22 type %s", path, typeKey)
				}

				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, removed := removedFields[typeKey][name.Name]; removed {
							t.Errorf("%s declares removed v22 field %s.%s", path, typeKey, name.Name)
						}
					}

					jsonKey, present := v22JSONFieldName(t, path, field)
					if !present || jsonKey == "-" {
						continue
					}
					if _, removed := removedJSONKeys[jsonKey]; removed {
						t.Errorf("%s contains removed v22 JSON key %q", path, jsonKey)
					}
					if _, removed := removedJSONKeysByType[typeKey][jsonKey]; removed {
						t.Errorf("%s contains removed v22 JSON key %q on %s", path, jsonKey, typeKey)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestV22GoSourcesContainNoOlderContractImports(t *testing.T) {
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if parseErr != nil {
			return parseErr
		}
		for _, spec := range file.Imports {
			importPath, unquoteErr := strconv.Unquote(spec.Path.Value)
			if unquoteErr != nil {
				return unquoteErr
			}
			for _, oldMajor := range []string{"/v19/", "/v20/", "/v21/"} {
				if strings.Contains(importPath, "Backend-Shared-Contract"+oldMajor) {
					t.Errorf("%s imports older shared-contract major %s", path, importPath)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func v22StringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func v22ProductionTypeKey(path string, typeName string) string {
	directory := filepath.ToSlash(filepath.Dir(path))
	directory = strings.TrimPrefix(directory, "./")
	if directory == "." || directory == "" {
		return typeName
	}
	return directory + "." + typeName
}

func v22JSONFieldName(t *testing.T, path string, field *ast.Field) (string, bool) {
	t.Helper()
	if field.Tag == nil {
		return "", false
	}
	rawTag, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		t.Errorf("%s contains an invalid struct tag literal %s: %v", path, field.Tag.Value, err)
		return "", false
	}
	jsonTag, present := reflect.StructTag(rawTag).Lookup("json")
	if !present {
		return "", false
	}
	return strings.Split(jsonTag, ",")[0], true
}
