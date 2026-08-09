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

func TestV26NonSupplyProductionModelsRejectLegacyMediaShapes(t *testing.T) {
	legacyTypes := v25StringSet("Media", "MediaReference")
	legacyFields := v25StringSet(
		"MediaID", "MediaURL",
		"AvatarMediaID", "AvatarURL",
		"LogoURL",
		"ImageID", "ImageURL",
		"CoverMediaID", "CoverURL",
		"ImageMediaIDs", "ImageURLs",
	)
	legacyJSONKeys := v25StringSet(
		"media_id", "media_url",
		"avatar_media_id", "avatar_url",
		"logo_url",
		"image_id", "image_url",
		"cover_media_id", "cover_url",
		"image_media_ids", "image_urls",
	)

	pkgRoot := sharedContractPkgRoot(t)
	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		relativePath, relativeErr := filepath.Rel(pkgRoot, path)
		if relativeErr != nil {
			return relativeErr
		}
		relativePath = filepath.ToSlash(relativePath)
		// Supply-only image records are deliberately exempt from the v26
		// ObjectMedia cutover, except for the canonical Product.Images contract
		// covered by its package test.
		if strings.HasPrefix(relativePath, "contracts/supply/") {
			return nil
		}
		if relativePath == "contracts/common/security/media.go" {
			t.Errorf("%s retains the retired common/security media.go path", path)
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
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
				if _, legacy := legacyTypes[typeSpecification.Name.Name]; legacy {
					t.Errorf("%s declares retired v26 media type %s", path, typeSpecification.Name.Name)
				}
				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, legacy := legacyFields[name.Name]; legacy {
							t.Errorf("%s retains legacy media field %s.%s", path, typeSpecification.Name.Name, name.Name)
						}
					}
					jsonKey, present := v25JSONFieldName(t, path, field)
					if present {
						if _, legacy := legacyJSONKeys[jsonKey]; legacy {
							t.Errorf("%s retains legacy media JSON key %q on %s", path, jsonKey, typeSpecification.Name.Name)
						}
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

func TestV25ProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := v25StringSet(
		// Earlier hard cut-overs that remain forbidden in v25.
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
		// movement types in v25.
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

	removedTypes := v25StringSet(
		// The v26 canonical-product cut-over removes all endpoint-specific
		// product projections and product-owned snapshot duplicates.
		"contracts/orders/pos.CatalogProduct",
		"contracts/supply/product.DetailImage",
		"contracts/supply/product.ProductBarcodeAssignmentSnapshot",
		"contracts/supply/product.ProductPackageOptionSnapshot",
		"contracts/supply/product.PreorderPolicy",
		"contracts/supply/product.Snapshot",
		"contracts/supply/product.SoonExpiryMerchandisingPolicy",
		"contracts/supply/product.StorefrontCommercial",
		"contracts/supply/product.StorefrontDisplay",
		"contracts/supply/product.StorefrontExpiryDisplay",
		"contracts/supply/product.StorefrontMerchandising",
		"contracts/supply/product.StorefrontOrigin",
		"contracts/supply/product.StorefrontPreorderDisplay",
		"contracts/supply/product.StorefrontProduct",
		"contracts/supply/product.StorefrontPromotionBadge",
		"contracts/supply/product/product_enums.StorefrontExpiryStatus",
		"contracts/supply/product/product_enums.StorefrontPreorderStatus",
		"contracts/supply/product/product_enums.WholesalePriceMode",

		"notifications/backinstock.BackInStockRestockEvent",
		"common/packaging.Physical",
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
		"common/geography.Address": v25StringSet(
			"City", "State", "Postcode",
		),
		"insights/analytics.OrderItemFact": v25StringSet(
			"Quantity",
		),
		"insights/analytics.RefundItemFact": v25StringSet(
			"Quantity",
		),
		"insights/analytics.SKUDemandForecast": v25StringSet(
			"SKUCode", "CurrentStockAtRun",
		),
		"customers/campaign.Audience": v25StringSet(
			"Region",
		),
		"customers/campaign.CampaignProductPrediction": v25StringSet(
			"PredictedDemandUnits", "SellableAvailableUnits", "ConfirmedInboundUnits",
			"NetRequiredUnits", "SuggestedOrderUnits", "MinimumOrderQuantity",
			"SuggestedCartons", "CartonSize",
		),
		"customers/campaign.CampaignPredictionEvidence": v25StringSet(
			"RawNetUnits", "NormalizedUnits",
		),
		"customers/campaign.CampaignSupplierPrediction": v25StringSet(
			"TotalUnits",
		),
		"supply/classification.SKU": v25StringSet(
			"Storage",
		),
		"orders/pos.CatalogProduct": v25StringSet(
			"ID", "SKU", "Barcode", "Storage", "Price", "CurrentStock",
		),
		"supply/product.Product": v25StringSet(
			"ID", "SKU", "Barcode", "Storage", "PlacingArea", "CurrentStock",
			"RestockedAt", "ExpiredAt", "Pricing", "Physical",
		),
		"supply/product.Snapshot": v25StringSet(
			"ID", "SKU", "Storage", "DisplayStatus", "Barcode",
		),
		"supply/product.StorefrontDisplay":       v25StringSet(),
		"supply/product.StorefrontMerchandising": v25StringSet(),
		"supply/product.StorefrontProduct": v25StringSet(
			"SKU", "Barcode", "Storage", "CurrentStock", "Pricing", "ExpiryDate", "DisplayStatus",
		),
		"supply/purchase.OrderItem": v25StringSet(
			"OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode", "ExpireAt",
		),
		"supply/purchase.ReceiptItem": v25StringSet(
			"SKU", "OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode",
		),
		"orders/order.CartItem": v25StringSet(
			"Price", "Quantity",
		),
		"orders/order.DemandBucket": v25StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.OpenDemandLine": v25StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.Order": v25StringSet(
			"DeliveryRegion",
		),
		"orders/order.OrderItem": v25StringSet(
			"UnitPrice", "Quantity", "CartonQty", "CartonSize",
		),
		"orders/order.OrderLineSummary": v25StringSet(
			"Quantity", "UnitPrice",
		),
		"orders/order.OrderPackingProgress": v25StringSet(
			"BoxPlan",
		),
		"orders/order.PreorderItemState": v25StringSet(
			"OrderedQuantity", "AllocatedQuantity",
		),
		"orders/order.VolumeDiscountTier": v25StringSet(
			"MinCartons",
		),
		"orders/shipping.DeliveryAreaRate": v25StringSet(
			"Postcode", "Suburb", "DeliveryRegion",
		),
		"orders/shipping.Zone": v25StringSet(
			"States", "Postcodes", "IsLocal",
		),
		"supply/warehouse.DamageReport": v25StringSet(
			"DamagedQty",
		),
		"supply/warehouse.Depot": v25StringSet(
			"PostcodeRules",
		),
		"supply/operations.InboundItem": v25StringSet(
			"Barcode", "Storage", "ExpectedQty", "ReceivedQty", "LocationCode",
		),
		"supply/operations.PackingDamage": v25StringSet(
			"SKU", "DamagedQty", "DamageReportID", "StockMovementID",
		),
		"orders/shipping.PackingDiscrepancy": v25StringSet(
			"OrderedQty", "ScannedQty", "DiffQty", "UnitPrice", "ReturnToStock",
			"DamageReportID", "StockMovementID", "DamagedQty",
		),
		"supply/operations.PackingLine": v25StringSet(
			"SKU", "OrderedQty", "ScannedQty", "DamagedQty",
		),
		"supply/operations.PickingListItem": v25StringSet(
			"Barcode", "Location", "QuantityRequired", "QuantityPicked",
		),
		"supply/warehouse.StockLocation": v25StringSet(
			"Code", "Zone",
		),
		"supply/warehouse.StockLocationProductBalance": v25StringSet(
			"AvailabilityRevision",
		),
		"supply/operations.StockMovement": v25StringSet(
			"SKU", "ProductName", "DepotCode", "LocationCode", "QtyDelta", "BalanceAfter",
			"CreatedBy", "SalesOrderNumber", "ReferenceType", "ReferenceID", "Metadata",
		),
		"supply/warehouse.WMSDraft": v25StringSet(
			"TotalQty",
		),
		"supply/warehouse.WMSDraftItem": v25StringSet(
			"Barcode", "LocationCode", "Qty", "ExpiryYM",
		),
	}

	// These JSON keys are unambiguously retired everywhere. Reused names such
	// as state, sku, barcode, quantity, unit_price, pricing, storage, and
	// location_code are checked against their exact owning type below.
	removedJSONKeys := v25StringSet(
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
		"common/geography.Address": v25StringSet(
			"city", "state", "postcode",
		),
		"insights/analytics.OrderItemFact": v25StringSet(
			"quantity",
		),
		"insights/analytics.RefundItemFact": v25StringSet(
			"quantity",
		),
		"insights/analytics.SKUDemandForecast": v25StringSet(
			"sku_code", "current_stock_at_run",
		),
		"customers/campaign.Audience": v25StringSet(
			"region",
		),
		"customers/campaign.CampaignProductPrediction": v25StringSet(
			"predicted_demand_units", "sellable_available_units", "confirmed_inbound_units",
			"net_required_units", "suggested_order_units", "minimum_order_quantity",
			"suggested_cartons", "carton_size",
		),
		"customers/campaign.CampaignPredictionEvidence": v25StringSet(
			"raw_net_units", "normalized_units",
		),
		"customers/campaign.CampaignSupplierPrediction": v25StringSet(
			"total_units",
		),
		"supply/classification.SKU": v25StringSet(
			"storage",
		),
		"orders/pos.CatalogProduct": v25StringSet(
			"id", "sku", "barcode", "storage", "price", "current_stock",
		),
		"supply/product.Product": v25StringSet(
			"id", "sku", "barcode", "storage", "placing_area", "current_stock",
			"restocked_at", "expired_at", "pricing", "physical",
		),
		"supply/product.Snapshot": v25StringSet(
			"id", "sku", "storage", "display_status", "barcode",
		),
		"supply/product.StorefrontDisplay":       v25StringSet(),
		"supply/product.StorefrontMerchandising": v25StringSet(),
		"supply/product.StorefrontProduct": v25StringSet(
			"sku", "barcode", "storage", "current_stock", "pricing", "expiry_date", "display_status",
		),
		"supply/purchase.OrderItem": v25StringSet(
			"ordered_qty", "received_qty", "rejected_qty", "location_code", "expire_at",
		),
		"supply/purchase.ReceiptItem": v25StringSet(
			"sku", "ordered_qty", "received_qty", "rejected_qty", "location_code",
		),
		"orders/order.CartItem": v25StringSet(
			"price", "quantity",
		),
		"orders/order.DemandBucket": v25StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.OpenDemandLine": v25StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.Order": v25StringSet(
			"delivery_region",
		),
		"orders/order.OrderItem": v25StringSet(
			"unit_price", "quantity", "carton_qty", "carton_size",
		),
		"orders/order.OrderLineSummary": v25StringSet(
			"quantity", "unit_price",
		),
		"orders/order.OrderPackingProgress": v25StringSet(
			"box_plan",
		),
		"orders/order.PreorderItemState": v25StringSet(
			"ordered_quantity", "allocated_quantity",
		),
		"orders/order.VolumeDiscountTier": v25StringSet(
			"min_cartons",
		),
		"orders/shipping.DeliveryAreaRate": v25StringSet(
			"postcode", "suburb", "delivery_region",
		),
		"orders/shipping.Zone": v25StringSet(
			"states", "postcodes", "is_local",
		),
		"supply/warehouse.DamageReport": v25StringSet(
			"damaged_qty",
		),
		"supply/warehouse.Depot": v25StringSet(
			"postcode_rules",
		),
		"supply/operations.InboundItem": v25StringSet(
			"barcode", "storage", "expected_qty", "received_qty", "location_code",
		),
		"supply/operations.PackingDamage": v25StringSet(
			"sku", "damaged_qty", "damage_report_id", "stock_movement_id",
		),
		"orders/shipping.PackingDiscrepancy": v25StringSet(
			"ordered_qty", "scanned_qty", "diff_qty", "unit_price", "return_to_stock",
			"damage_report_id", "stock_movement_id", "damaged_qty",
		),
		"supply/operations.PackingLine": v25StringSet(
			"sku", "ordered_qty", "scanned_qty", "damaged_qty",
		),
		"supply/operations.PickingListItem": v25StringSet(
			"barcode", "location", "quantity_required", "quantity_picked",
		),
		"supply/warehouse.StockLocation": v25StringSet(
			"code", "zone",
		),
		"supply/warehouse.StockLocationProductBalance": v25StringSet(
			"availability_revision",
		),
		"supply/operations.StockMovement": v25StringSet(
			"sku", "product_name", "depot_code", "location_code", "qty_delta", "balance_after",
			"created_by", "sales_order_number", "reference_type", "reference_id", "metadata",
		),
		"supply/warehouse.WMSDraft": v25StringSet(
			"total_qty",
		),
		"supply/warehouse.WMSDraftItem": v25StringSet(
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
				t.Errorf("%s contains removed v25 identifier %s", path, identifier.Name)
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

				typeKey := v25ProductionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				if _, removed := removedTypes[typeKey]; removed {
					t.Errorf("%s declares removed v25 type %s", path, typeKey)
				}

				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, removed := removedFields[typeKey][name.Name]; removed {
							t.Errorf("%s declares removed v25 field %s.%s", path, typeKey, name.Name)
						}
					}

					jsonKey, present := v25JSONFieldName(t, path, field)
					if !present || jsonKey == "-" {
						continue
					}
					if _, removed := removedJSONKeys[jsonKey]; removed {
						t.Errorf("%s contains removed v25 JSON key %q", path, jsonKey)
					}
					if _, removed := removedJSONKeysByType[typeKey][jsonKey]; removed {
						t.Errorf("%s contains removed v25 JSON key %q on %s", path, jsonKey, typeKey)
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

func TestV25GoSourcesContainNoOlderContractImports(t *testing.T) {
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

func v25StringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func v25ProductionTypeKey(pkgRoot string, path string, typeName string) string {
	relativePath, _ := filepath.Rel(pkgRoot, path)
	directory := filepath.ToSlash(filepath.Dir(relativePath))
	if directory == "." || directory == "" {
		return typeName
	}
	return directory + "." + typeName
}

func v25JSONFieldName(t *testing.T, path string, field *ast.Field) (string, bool) {
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
