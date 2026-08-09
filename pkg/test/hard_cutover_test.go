package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestV26NonSupplyProductionModelsRejectLegacyMediaShapes(t *testing.T) {
	legacyTypes := v26StringSet("Media", "MediaReference")
	legacyFields := v26StringSet(
		"MediaID", "MediaURL",
		"AvatarMediaID", "AvatarURL",
		"LogoURL",
		"ImageID", "ImageURL",
		"CoverMediaID", "CoverURL",
		"ImageMediaIDs", "ImageURLs",
	)
	legacyJSONKeys := v26StringSet(
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
					jsonKey, present := v26JSONFieldName(t, path, field)
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

func TestV26ProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := v26StringSet(
		// Earlier hard cut-overs that remain forbidden in v26.
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
		// movement types in v26.
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

	removedTypes := v26StringSet(
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
		"common/geography.Address": v26StringSet(
			"City", "State", "Postcode",
		),
		"insights/analytics.OrderItemFact": v26StringSet(
			"Quantity",
		),
		"insights/analytics.RefundItemFact": v26StringSet(
			"Quantity",
		),
		"insights/analytics.SKUDemandForecast": v26StringSet(
			"SKUCode", "CurrentStockAtRun",
		),
		"customers/campaign.Audience": v26StringSet(
			"Region",
		),
		"customers/campaign.CampaignProductPrediction": v26StringSet(
			"PredictedDemandUnits", "SellableAvailableUnits", "ConfirmedInboundUnits",
			"NetRequiredUnits", "SuggestedOrderUnits", "MinimumOrderQuantity",
			"SuggestedCartons", "CartonSize",
		),
		"customers/campaign.CampaignPredictionEvidence": v26StringSet(
			"RawNetUnits", "NormalizedUnits",
		),
		"customers/campaign.CampaignSupplierPrediction": v26StringSet(
			"TotalUnits",
		),
		"supply/classification.SKU": v26StringSet(
			"Storage",
		),
		"orders/pos.CatalogProduct": v26StringSet(
			"ID", "SKU", "Barcode", "Storage", "Price", "CurrentStock",
		),
		"supply/product.Product": v26StringSet(
			"ID", "SKU", "Barcode", "Storage", "PlacingArea", "CurrentStock",
			"RestockedAt", "ExpiredAt", "Pricing", "Physical",
		),
		"supply/product.Snapshot": v26StringSet(
			"ID", "SKU", "Storage", "DisplayStatus", "Barcode",
		),
		"supply/product.StorefrontDisplay":       v26StringSet(),
		"supply/product.StorefrontMerchandising": v26StringSet(),
		"supply/product.StorefrontProduct": v26StringSet(
			"SKU", "Barcode", "Storage", "CurrentStock", "Pricing", "ExpiryDate", "DisplayStatus",
		),
		"supply/purchase.OrderItem": v26StringSet(
			"OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode", "ExpireAt",
		),
		"supply/purchase.ReceiptItem": v26StringSet(
			"SKU", "OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode",
		),
		"orders/order.CartItem": v26StringSet(
			"Price", "Quantity",
		),
		"orders/order.DemandBucket": v26StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.OpenDemandLine": v26StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.Order": v26StringSet(
			"DeliveryRegion",
		),
		"orders/order.OrderItem": v26StringSet(
			"UnitPrice", "Quantity", "CartonQty", "CartonSize",
		),
		"orders/order.OrderLineSummary": v26StringSet(
			"Quantity", "UnitPrice",
		),
		"orders/order.OrderPackingProgress": v26StringSet(
			"BoxPlan",
		),
		"orders/order.PreorderItemState": v26StringSet(
			"OrderedQuantity", "AllocatedQuantity",
		),
		"orders/order.VolumeDiscountTier": v26StringSet(
			"MinCartons",
		),
		"orders/shipping.DeliveryAreaRate": v26StringSet(
			"Postcode", "Suburb", "DeliveryRegion",
		),
		"orders/shipping.Zone": v26StringSet(
			"States", "Postcodes", "IsLocal",
		),
		"supply/warehouse.DamageReport": v26StringSet(
			"DamagedQty",
		),
		"supply/warehouse.Depot": v26StringSet(
			"PostcodeRules",
		),
		"supply/operations.InboundItem": v26StringSet(
			"Barcode", "Storage", "ExpectedQty", "ReceivedQty", "LocationCode",
		),
		"supply/operations.PackingDamage": v26StringSet(
			"SKU", "DamagedQty", "DamageReportID", "StockMovementID",
		),
		"orders/shipping.PackingDiscrepancy": v26StringSet(
			"OrderedQty", "ScannedQty", "DiffQty", "UnitPrice", "ReturnToStock",
			"DamageReportID", "StockMovementID", "DamagedQty",
		),
		"supply/operations.PackingLine": v26StringSet(
			"SKU", "OrderedQty", "ScannedQty", "DamagedQty",
		),
		"supply/operations.PickingListItem": v26StringSet(
			"Barcode", "Location", "QuantityRequired", "QuantityPicked",
		),
		"supply/warehouse.StockLocation": v26StringSet(
			"Code", "Zone",
		),
		"supply/warehouse.StockLocationProductBalance": v26StringSet(
			"AvailabilityRevision",
		),
		"supply/operations.StockMovement": v26StringSet(
			"SKU", "ProductName", "DepotCode", "LocationCode", "QtyDelta", "BalanceAfter",
			"CreatedBy", "SalesOrderNumber", "ReferenceType", "ReferenceID", "Metadata",
		),
		"supply/warehouse.WMSDraft": v26StringSet(
			"TotalQty",
		),
		"supply/warehouse.WMSDraftItem": v26StringSet(
			"Barcode", "LocationCode", "Qty", "ExpiryYM",
		),
	}

	// These JSON keys are unambiguously retired everywhere. Reused names such
	// as state, sku, barcode, quantity, unit_price, pricing, storage, and
	// location_code are checked against their exact owning type below.
	removedJSONKeys := v26StringSet(
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
		"common/geography.Address": v26StringSet(
			"city", "state", "postcode",
		),
		"insights/analytics.OrderItemFact": v26StringSet(
			"quantity",
		),
		"insights/analytics.RefundItemFact": v26StringSet(
			"quantity",
		),
		"insights/analytics.SKUDemandForecast": v26StringSet(
			"sku_code", "current_stock_at_run",
		),
		"customers/campaign.Audience": v26StringSet(
			"region",
		),
		"customers/campaign.CampaignProductPrediction": v26StringSet(
			"predicted_demand_units", "sellable_available_units", "confirmed_inbound_units",
			"net_required_units", "suggested_order_units", "minimum_order_quantity",
			"suggested_cartons", "carton_size",
		),
		"customers/campaign.CampaignPredictionEvidence": v26StringSet(
			"raw_net_units", "normalized_units",
		),
		"customers/campaign.CampaignSupplierPrediction": v26StringSet(
			"total_units",
		),
		"supply/classification.SKU": v26StringSet(
			"storage",
		),
		"orders/pos.CatalogProduct": v26StringSet(
			"id", "sku", "barcode", "storage", "price", "current_stock",
		),
		"supply/product.Product": v26StringSet(
			"id", "sku", "barcode", "storage", "placing_area", "current_stock",
			"restocked_at", "expired_at", "pricing", "physical",
		),
		"supply/product.Snapshot": v26StringSet(
			"id", "sku", "storage", "display_status", "barcode",
		),
		"supply/product.StorefrontDisplay":       v26StringSet(),
		"supply/product.StorefrontMerchandising": v26StringSet(),
		"supply/product.StorefrontProduct": v26StringSet(
			"sku", "barcode", "storage", "current_stock", "pricing", "expiry_date", "display_status",
		),
		"supply/purchase.OrderItem": v26StringSet(
			"ordered_qty", "received_qty", "rejected_qty", "location_code", "expire_at",
		),
		"supply/purchase.ReceiptItem": v26StringSet(
			"sku", "ordered_qty", "received_qty", "rejected_qty", "location_code",
		),
		"orders/order.CartItem": v26StringSet(
			"price", "quantity",
		),
		"orders/order.DemandBucket": v26StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.OpenDemandLine": v26StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.Order": v26StringSet(
			"delivery_region",
		),
		"orders/order.OrderItem": v26StringSet(
			"unit_price", "quantity", "carton_qty", "carton_size",
		),
		"orders/order.OrderLineSummary": v26StringSet(
			"quantity", "unit_price",
		),
		"orders/order.OrderPackingProgress": v26StringSet(
			"box_plan",
		),
		"orders/order.PreorderItemState": v26StringSet(
			"ordered_quantity", "allocated_quantity",
		),
		"orders/order.VolumeDiscountTier": v26StringSet(
			"min_cartons",
		),
		"orders/shipping.DeliveryAreaRate": v26StringSet(
			"postcode", "suburb", "delivery_region",
		),
		"orders/shipping.Zone": v26StringSet(
			"states", "postcodes", "is_local",
		),
		"supply/warehouse.DamageReport": v26StringSet(
			"damaged_qty",
		),
		"supply/warehouse.Depot": v26StringSet(
			"postcode_rules",
		),
		"supply/operations.InboundItem": v26StringSet(
			"barcode", "storage", "expected_qty", "received_qty", "location_code",
		),
		"supply/operations.PackingDamage": v26StringSet(
			"sku", "damaged_qty", "damage_report_id", "stock_movement_id",
		),
		"orders/shipping.PackingDiscrepancy": v26StringSet(
			"ordered_qty", "scanned_qty", "diff_qty", "unit_price", "return_to_stock",
			"damage_report_id", "stock_movement_id", "damaged_qty",
		),
		"supply/operations.PackingLine": v26StringSet(
			"sku", "ordered_qty", "scanned_qty", "damaged_qty",
		),
		"supply/operations.PickingListItem": v26StringSet(
			"barcode", "location", "quantity_required", "quantity_picked",
		),
		"supply/warehouse.StockLocation": v26StringSet(
			"code", "zone",
		),
		"supply/warehouse.StockLocationProductBalance": v26StringSet(
			"availability_revision",
		),
		"supply/operations.StockMovement": v26StringSet(
			"sku", "product_name", "depot_code", "location_code", "qty_delta", "balance_after",
			"created_by", "sales_order_number", "reference_type", "reference_id", "metadata",
		),
		"supply/warehouse.WMSDraft": v26StringSet(
			"total_qty",
		),
		"supply/warehouse.WMSDraftItem": v26StringSet(
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
				t.Errorf("%s contains removed v26 identifier %s", path, identifier.Name)
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

				typeKey := v26ProductionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				if _, removed := removedTypes[typeKey]; removed {
					t.Errorf("%s declares removed v26 type %s", path, typeKey)
				}

				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, removed := removedFields[typeKey][name.Name]; removed {
							t.Errorf("%s declares removed v26 field %s.%s", path, typeKey, name.Name)
						}
					}

					jsonKey, present := v26JSONFieldName(t, path, field)
					if !present || jsonKey == "-" {
						continue
					}
					if _, removed := removedJSONKeys[jsonKey]; removed {
						t.Errorf("%s contains removed v26 JSON key %q", path, jsonKey)
					}
					if _, removed := removedJSONKeysByType[typeKey][jsonKey]; removed {
						t.Errorf("%s contains removed v26 JSON key %q on %s", path, jsonKey, typeKey)
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

func TestV26GoSourcesContainNoOlderContractImports(t *testing.T) {
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
			for _, oldMajor := range []string{"/v19/", "/v20/", "/v21/", "/v22/", "/v23/", "/v24/", "/v25/"} {
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

func TestV26ProductionModelsRejectRetiredPromotionOfferAndImportComplianceTerms(t *testing.T) {
	legacyImportCompliance := "import" + "compliance"
	legacyImportComplianceEnums := legacyImportCompliance + "_enums"
	retiredIdentifiers := v26StringSet(
		"AcceptedOffer",
		"AppliedPromotion",
		"AppliedPromotions",
		"CatalogProduct",
		"DetailImage",
		"DiscountScope",
		"DiscountSpec",
		"DiscountType",
		"EffectivePromotion",
		"EventTypeInventorySellableOfferAvailable",
		"EventTypeInventorySellableOfferWithdrawn",
		"GroupOrderDiscountApplication",
		"GroupOrderDiscountDecision",
		"GroupOrderDiscountDecisionLine",
		"GroupOrderDiscountProposal",
		"GroupOrderDiscountState",
		"MediaReference",
		"OfferID",
		"OfferRevision",
		"Offers",
		"PreorderPolicy",
		"ProductBarcodeAssignmentSnapshot",
		"ProductPackageOptionSnapshot",
		"PromotionAddonTrigger",
		"PromotionClass",
		"PromotionDiscountTarget",
		"PromotionProduct",
		"PromotionQtyMode",
		"PromotionType",
		"ReceiptOffer",
		"SellableOffer",
		"SellableOfferAvailabilityChangedEvent",
		"SellableOfferDateMarkSnapshot",
		"SellableOfferDiscountSnapshot",
		"SellableOfferSnapshot",
		"SameSaleOrder",
		"StorefrontPromotion",
		"StorefrontPromotionBadge",
		"StorefrontCommercial",
		"StorefrontDisplay",
		"StorefrontExpiryDisplay",
		"StorefrontMerchandising",
		"StorefrontOrigin",
		"StorefrontPreorderDisplay",
		"StorefrontProduct",
		"SoonExpiryMerchandisingPolicy",
		"UsageLimits",
		"VolumeDiscountAppliesTo",
		"VolumeDiscountTier",
		legacyImportCompliance,
		legacyImportComplianceEnums,
	)
	retiredTypes := v26StringSet(
		"contracts/common/security.Media",
		"contracts/common/security.MediaReference",
		"contracts/orders/pos.CatalogProduct",
		"contracts/orders/order.VolumeDiscountTier",
		"contracts/pricing/promotion.ActiveWindow",
		"contracts/pricing/promotion.DiscountSpec",
		"contracts/pricing/promotion.EffectivePromotion",
		"contracts/pricing/promotion.GroupOrderDiscountApplication",
		"contracts/pricing/promotion.GroupOrderDiscountDecision",
		"contracts/pricing/promotion.GroupOrderDiscountDecisionLine",
		"contracts/pricing/promotion.GroupOrderDiscountProposal",
		"contracts/pricing/promotion.ReceiptOffer",
		"contracts/pricing/promotion.StorefrontPromotion",
		"contracts/pricing/promotion.UsageLimits",
		"contracts/pricing/promotion/promotion_enums.DiscountScope",
		"contracts/pricing/promotion/promotion_enums.DiscountType",
		"contracts/pricing/promotion/promotion_enums.GroupOrderDiscountState",
		"contracts/pricing/promotion/promotion_enums.PromotionAddonTrigger",
		"contracts/pricing/promotion/promotion_enums.PromotionClass",
		"contracts/pricing/promotion/promotion_enums.PromotionDiscountTarget",
		"contracts/pricing/promotion/promotion_enums.PromotionQtyMode",
		"contracts/pricing/promotion/promotion_enums.PromotionType",
		"contracts/pricing/promotion/promotion_enums.VolumeDiscountAppliesTo",
		"contracts/pubsub/event.SellableOfferAvailabilityChangedEvent",
		"contracts/supply/product.DetailImage",
		"contracts/supply/product.PreorderPolicy",
		"contracts/supply/product.SellableOffer",
		"contracts/supply/product.SellableOfferDateMarkSnapshot",
		"contracts/supply/product.SellableOfferDiscountSnapshot",
		"contracts/supply/product.SellableOfferSnapshot",
		"contracts/supply/product.Snapshot",
		"contracts/supply/product.SoonExpiryMerchandisingPolicy",
		"contracts/supply/product.StorefrontCommercial",
		"contracts/supply/product.StorefrontDisplay",
		"contracts/supply/product.StorefrontMerchandising",
		"contracts/supply/product.StorefrontProduct",
	)
	retiredWireFragments := []string{
		"accepted_offer",
		"applied_promotions",
		"offer_id",
		"offer_revision",
		"same_sale_order",
		"sellable_offer",
	}
	retiredPaths := []string{
		"contracts/common/security/media.go",
		"contracts/pricing/promotion/effective_promotion.go",
		"contracts/pricing/promotion/group_order_discount.go",
		"contracts/pricing/promotion/receipt_offer.go",
		"contracts/pricing/promotion/shared.go",
		"contracts/pricing/promotion/storefront_promotion.go",
		"contracts/supply/" + legacyImportCompliance,
		"contracts/supply/" + legacyImportCompliance + "/" + legacyImportComplianceEnums,
		"contracts/supply/product/offer.go",
	}

	pkgRoot := sharedContractPkgRoot(t)
	for _, retiredPath := range retiredPaths {
		absolutePath := filepath.Join(pkgRoot, filepath.FromSlash(retiredPath))
		if _, err := os.Stat(absolutePath); err == nil {
			t.Errorf("retired v26 production path remains: %s", retiredPath)
		} else if !os.IsNotExist(err) {
			t.Errorf("inspect retired v26 production path %s: %v", retiredPath, err)
		}
	}

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
		if strings.Contains(relativePath, "/"+legacyImportCompliance+"/") || strings.HasSuffix(relativePath, "/"+legacyImportCompliance) {
			t.Errorf("%s retains retired import-compliance package path", path)
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		if file.Name != nil && file.Name.Name == legacyImportCompliance {
			t.Errorf("%s retains retired import-compliance package identifier", path)
		}
		for _, specification := range file.Imports {
			importPath, unquoteErr := strconv.Unquote(specification.Path.Value)
			if unquoteErr != nil {
				return unquoteErr
			}
			if strings.Contains(importPath, "/"+legacyImportCompliance) {
				t.Errorf("%s imports retired import-compliance path %s", path, importPath)
			}
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
				typeKey := v26ProductionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				if _, retired := retiredTypes[typeKey]; retired {
					t.Errorf("%s declares retired v26 type %s", path, typeKey)
				}
			}
		}

		ast.Inspect(file, func(node ast.Node) bool {
			switch value := node.(type) {
			case *ast.Ident:
				if _, retired := retiredIdentifiers[value.Name]; retired {
					t.Errorf("%s retains retired v26 identifier %s", path, value.Name)
				}
			case *ast.BasicLit:
				if value.Kind != token.STRING {
					return true
				}
				text, unquoteErr := strconv.Unquote(value.Value)
				if unquoteErr != nil {
					return true
				}
				lowerText := strings.ToLower(text)
				for _, fragment := range retiredWireFragments {
					if strings.Contains(lowerText, fragment) {
						t.Errorf("%s retains retired v26 wire fragment %q", path, fragment)
					}
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func v26StringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func v26ProductionTypeKey(pkgRoot string, path string, typeName string) string {
	relativePath, _ := filepath.Rel(pkgRoot, path)
	directory := filepath.ToSlash(filepath.Dir(relativePath))
	if directory == "." || directory == "" {
		return typeName
	}
	return directory + "." + typeName
}

func v26JSONFieldName(t *testing.T, path string, field *ast.Field) (string, bool) {
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
