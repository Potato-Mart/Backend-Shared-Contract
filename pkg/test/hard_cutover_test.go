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

func TestV23ProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := v23StringSet(
		// Earlier hard cut-overs that remain forbidden in v23.
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
		// movement types in v23.
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

	removedTypes := v23StringSet(
		"notifications/backinstock.BackInStockRestockEvent",
		"common/shared.Physical",
		"supply/product.Pricing",
		"supply/product.ProductPlacement",
		"supply/product.StorefrontPricing",
		"orders/order.PreorderAvailabilityEvent",
		"pubsub/event.PreorderStockArrivalEvent",
		"supply/product.DepotProduct",
		"supply/warehouse.PackingBoxContent",
		"supply/warehouse.PackingBoxPlan",
		"supply/warehouse.PostcodeRule",
		"pubsub/event.StockAdjustedEvent",
		"orders/shipping.DeliveryRegion",
		"supply/warehouse.PackingSessionStatus",
	)

	removedFields := map[string]map[string]struct{}{
		"common/shared.Address": v23StringSet(
			"City", "State", "Postcode",
		),
		"insights/analytics.OrderItemFact": v23StringSet(
			"Quantity",
		),
		"insights/analytics.RefundItemFact": v23StringSet(
			"Quantity",
		),
		"insights/analytics.SKUDemandForecast": v23StringSet(
			"SKUCode", "CurrentStockAtRun",
		),
		"customers/campaign.Audience": v23StringSet(
			"Region",
		),
		"customers/campaign.CampaignProductPrediction": v23StringSet(
			"PredictedDemandUnits", "SellableAvailableUnits", "ConfirmedInboundUnits",
			"NetRequiredUnits", "SuggestedOrderUnits", "MinimumOrderQuantity",
			"SuggestedCartons", "CartonSize",
		),
		"customers/campaign.CampaignPredictionEvidence": v23StringSet(
			"RawNetUnits", "NormalizedUnits",
		),
		"customers/campaign.CampaignSupplierPrediction": v23StringSet(
			"TotalUnits",
		),
		"supply/category.SKU": v23StringSet(
			"Storage",
		),
		"orders/pos.CatalogProduct": v23StringSet(
			"ID", "SKU", "Barcode", "Storage", "Price", "CurrentStock",
		),
		"supply/product.Product": v23StringSet(
			"ID", "SKU", "Barcode", "Storage", "PlacingArea", "CurrentStock",
			"RestockedAt", "ExpiredAt", "Pricing", "Physical",
		),
		"supply/product.Snapshot": v23StringSet(
			"ID", "SKU", "Storage", "DisplayStatus", "Barcode",
		),
		"supply/product.StorefrontDisplay":       v23StringSet(),
		"supply/product.StorefrontMerchandising": v23StringSet(),
		"supply/product.StorefrontProduct": v23StringSet(
			"SKU", "Barcode", "Storage", "CurrentStock", "Pricing", "ExpiryDate", "DisplayStatus",
		),
		"supply/purchase.OrderItem": v23StringSet(
			"OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode", "ExpireAt",
		),
		"supply/purchase.ReceiptItem": v23StringSet(
			"SKU", "OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode",
		),
		"orders/order.CartItem": v23StringSet(
			"Price", "Quantity",
		),
		"orders/order.DemandBucket": v23StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.OpenDemandLine": v23StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.Order": v23StringSet(
			"DeliveryRegion",
		),
		"orders/order.OrderItem": v23StringSet(
			"UnitPrice", "Quantity", "CartonQty", "CartonSize",
		),
		"orders/order.OrderLineSummary": v23StringSet(
			"Quantity", "UnitPrice",
		),
		"orders/order.OrderPackingProgress": v23StringSet(
			"BoxPlan",
		),
		"orders/order.PreorderItemState": v23StringSet(
			"OrderedQuantity", "AllocatedQuantity",
		),
		"orders/order.VolumeDiscountTier": v23StringSet(
			"MinCartons",
		),
		"orders/shipping.DeliveryAreaRate": v23StringSet(
			"Postcode", "Suburb", "DeliveryRegion",
		),
		"orders/shipping.Zone": v23StringSet(
			"States", "Postcodes", "IsLocal",
		),
		"supply/warehouse.DamageReport": v23StringSet(
			"DamagedQty",
		),
		"supply/warehouse.Depot": v23StringSet(
			"PostcodeRules",
		),
		"supply/warehouse.InboundItem": v23StringSet(
			"Barcode", "Storage", "ExpectedQty", "ReceivedQty", "LocationCode",
		),
		"supply/warehouse.PackingDamage": v23StringSet(
			"SKU", "DamagedQty", "DamageReportID", "StockMovementID",
		),
		"supply/warehouse.PackingDiscrepancy": v23StringSet(
			"OrderedQty", "ScannedQty", "DiffQty", "UnitPrice", "ReturnToStock",
			"DamageReportID", "StockMovementID", "DamagedQty",
		),
		"supply/warehouse.PackingLine": v23StringSet(
			"SKU", "OrderedQty", "ScannedQty", "DamagedQty",
		),
		"supply/warehouse.PickingListItem": v23StringSet(
			"Barcode", "Location", "QuantityRequired", "QuantityPicked",
		),
		"supply/warehouse.StockLocation": v23StringSet(
			"Code", "Zone",
		),
		"supply/warehouse.StockLocationProductBalance": v23StringSet(
			"AvailabilityRevision",
		),
		"supply/warehouse.StockMovement": v23StringSet(
			"SKU", "ProductName", "DepotCode", "LocationCode", "QtyDelta", "BalanceAfter",
			"CreatedBy", "SalesOrderNumber", "ReferenceType", "ReferenceID", "Metadata",
		),
		"supply/warehouse.WMSDraft": v23StringSet(
			"TotalQty",
		),
		"supply/warehouse.WMSDraftItem": v23StringSet(
			"Barcode", "LocationCode", "Qty", "ExpiryYM",
		),
	}

	// These JSON keys are unambiguously retired everywhere. Reused names such
	// as state, sku, barcode, quantity, unit_price, pricing, storage, and
	// location_code are checked against their exact owning type below.
	removedJSONKeys := v23StringSet(
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
		"common/shared.Address": v23StringSet(
			"city", "state", "postcode",
		),
		"insights/analytics.OrderItemFact": v23StringSet(
			"quantity",
		),
		"insights/analytics.RefundItemFact": v23StringSet(
			"quantity",
		),
		"insights/analytics.SKUDemandForecast": v23StringSet(
			"sku_code", "current_stock_at_run",
		),
		"customers/campaign.Audience": v23StringSet(
			"region",
		),
		"customers/campaign.CampaignProductPrediction": v23StringSet(
			"predicted_demand_units", "sellable_available_units", "confirmed_inbound_units",
			"net_required_units", "suggested_order_units", "minimum_order_quantity",
			"suggested_cartons", "carton_size",
		),
		"customers/campaign.CampaignPredictionEvidence": v23StringSet(
			"raw_net_units", "normalized_units",
		),
		"customers/campaign.CampaignSupplierPrediction": v23StringSet(
			"total_units",
		),
		"supply/category.SKU": v23StringSet(
			"storage",
		),
		"orders/pos.CatalogProduct": v23StringSet(
			"id", "sku", "barcode", "storage", "price", "current_stock",
		),
		"supply/product.Product": v23StringSet(
			"id", "sku", "barcode", "storage", "placing_area", "current_stock",
			"restocked_at", "expired_at", "pricing", "physical",
		),
		"supply/product.Snapshot": v23StringSet(
			"id", "sku", "storage", "display_status", "barcode",
		),
		"supply/product.StorefrontDisplay":       v23StringSet(),
		"supply/product.StorefrontMerchandising": v23StringSet(),
		"supply/product.StorefrontProduct": v23StringSet(
			"sku", "barcode", "storage", "current_stock", "pricing", "expiry_date", "display_status",
		),
		"supply/purchase.OrderItem": v23StringSet(
			"ordered_qty", "received_qty", "rejected_qty", "location_code", "expire_at",
		),
		"supply/purchase.ReceiptItem": v23StringSet(
			"sku", "ordered_qty", "received_qty", "rejected_qty", "location_code",
		),
		"orders/order.CartItem": v23StringSet(
			"price", "quantity",
		),
		"orders/order.DemandBucket": v23StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.OpenDemandLine": v23StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.Order": v23StringSet(
			"delivery_region",
		),
		"orders/order.OrderItem": v23StringSet(
			"unit_price", "quantity", "carton_qty", "carton_size",
		),
		"orders/order.OrderLineSummary": v23StringSet(
			"quantity", "unit_price",
		),
		"orders/order.OrderPackingProgress": v23StringSet(
			"box_plan",
		),
		"orders/order.PreorderItemState": v23StringSet(
			"ordered_quantity", "allocated_quantity",
		),
		"orders/order.VolumeDiscountTier": v23StringSet(
			"min_cartons",
		),
		"orders/shipping.DeliveryAreaRate": v23StringSet(
			"postcode", "suburb", "delivery_region",
		),
		"orders/shipping.Zone": v23StringSet(
			"states", "postcodes", "is_local",
		),
		"supply/warehouse.DamageReport": v23StringSet(
			"damaged_qty",
		),
		"supply/warehouse.Depot": v23StringSet(
			"postcode_rules",
		),
		"supply/warehouse.InboundItem": v23StringSet(
			"barcode", "storage", "expected_qty", "received_qty", "location_code",
		),
		"supply/warehouse.PackingDamage": v23StringSet(
			"sku", "damaged_qty", "damage_report_id", "stock_movement_id",
		),
		"supply/warehouse.PackingDiscrepancy": v23StringSet(
			"ordered_qty", "scanned_qty", "diff_qty", "unit_price", "return_to_stock",
			"damage_report_id", "stock_movement_id", "damaged_qty",
		),
		"supply/warehouse.PackingLine": v23StringSet(
			"sku", "ordered_qty", "scanned_qty", "damaged_qty",
		),
		"supply/warehouse.PickingListItem": v23StringSet(
			"barcode", "location", "quantity_required", "quantity_picked",
		),
		"supply/warehouse.StockLocation": v23StringSet(
			"code", "zone",
		),
		"supply/warehouse.StockLocationProductBalance": v23StringSet(
			"availability_revision",
		),
		"supply/warehouse.StockMovement": v23StringSet(
			"sku", "product_name", "depot_code", "location_code", "qty_delta", "balance_after",
			"created_by", "sales_order_number", "reference_type", "reference_id", "metadata",
		),
		"supply/warehouse.WMSDraft": v23StringSet(
			"total_qty",
		),
		"supply/warehouse.WMSDraftItem": v23StringSet(
			"barcode", "location_code", "qty", "expiry_ym",
		),
	}

	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
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
				t.Errorf("%s contains removed v23 identifier %s", path, identifier.Name)
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

				typeKey := v23ProductionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				if _, removed := removedTypes[typeKey]; removed {
					t.Errorf("%s declares removed v23 type %s", path, typeKey)
				}

				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, removed := removedFields[typeKey][name.Name]; removed {
							t.Errorf("%s declares removed v23 field %s.%s", path, typeKey, name.Name)
						}
					}

					jsonKey, present := v23JSONFieldName(t, path, field)
					if !present || jsonKey == "-" {
						continue
					}
					if _, removed := removedJSONKeys[jsonKey]; removed {
						t.Errorf("%s contains removed v23 JSON key %q", path, jsonKey)
					}
					if _, removed := removedJSONKeysByType[typeKey][jsonKey]; removed {
						t.Errorf("%s contains removed v23 JSON key %q on %s", path, jsonKey, typeKey)
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

func TestV23GoSourcesContainNoOlderContractImports(t *testing.T) {
	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
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
			for _, oldMajor := range []string{"/v19/", "/v20/", "/v21/", "/v22/"} {
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

func v23StringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func v23ProductionTypeKey(pkgRoot string, path string, typeName string) string {
	relativePath, _ := filepath.Rel(pkgRoot, path)
	directory := filepath.ToSlash(filepath.Dir(relativePath))
	if directory == "." || directory == "" {
		return typeName
	}
	return directory + "." + typeName
}

func v23JSONFieldName(t *testing.T, path string, field *ast.Field) (string, bool) {
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
