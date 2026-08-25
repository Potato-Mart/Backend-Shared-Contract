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

func TestNonSupplyProductionModelsRejectRetiredMediaShapes(t *testing.T) {
	retiredMediaTypes := stringSet("Media", "MediaReference")
	retiredMediaFields := stringSet(
		"MediaCode", "MediaURL",
		"AvatarMediaCode", "AvatarURL",
		"LogoURL",
		"ImageID", "ImageURL",
		"CoverMediaCode", "CoverURL",
		"ImageMediaCodes", "ImageURLs",
	)
	retiredMediaJSONKeys := stringSet(
		"media_code", "media_url",
		"avatar_media_code", "avatar_url",
		"logo_url",
		"image_id", "image_url",
		"cover_media_code", "cover_url",
		"image_media_codes", "image_urls",
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
		// Supply-only image records are deliberately exempt from the ObjectMedia
		// boundary, except for the canonical Product.Images contract
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
				if _, retired := retiredMediaTypes[typeSpecification.Name.Name]; retired {
					t.Errorf("%s declares retired media type %s", path, typeSpecification.Name.Name)
				}
				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, retired := retiredMediaFields[name.Name]; retired {
							t.Errorf("%s retains retired media field %s.%s", path, typeSpecification.Name.Name, name.Name)
						}
					}
					jsonKey, present := jsonFieldName(t, path, field)
					if present {
						if _, retired := retiredMediaJSONKeys[jsonKey]; retired {
							t.Errorf("%s retains retired media JSON key %q on %s", path, jsonKey, typeSpecification.Name.Name)
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

func TestProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := stringSet(
		// Workforce roles are represented by countryAdmin, depotManager, and
		// warehouseManager. POS sign-in and till duty follow a principal's
		// geographic scope, so a dedicated selling role grants no authority.
		"UserRoleAdmin",
		"UserRoleWarehouse",
		"UserRoleCashier",
		"UserRoleSales",

		// Per-operator shifts are replaced by one daily
		// session per register, shared by every operator on it.
		"RegisterShift",
		"ShiftTotalsSnapshot",
		"ShiftStatus",
		"ShiftStatusOpen",
		"ShiftStatusClosed",
		"ShiftID",

		// Retired identifiers that must remain forbidden.
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

		// Geography, depot, product, packing, and retired availability types.
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

		// Product/SKU and market-pricing boundary.
		"ProductSKUCode",
		"ProductSKUCodes",
		"CategorySKUCode",
		"ProductPackaging",
		"ProductCommerce",
		"ProductPackageCommerce",
		"ProductMetrics",
		"Selling",
		"PackagePricing",
		"AcceptedPackagePricing",
		"PackagePricingID",
		"PackagePricingRevision",
		"PackagePricingAvailabilityChangedEvent",
		"EventTypeInventoryPackagePricingAvailable",
		"EventTypeInventoryPackagePricingWithdrawn",

		// Removed expiry-merchandising values.

		// Removed packing-session values.
		"PackingSessionStatusPending",
		"PackingSessionStatusPacking",
		"PackingSessionStatusPacked",
		"PackingSessionStatusResolved",

		// Logical reserve/release and split transfer values are not physical
		// movement types.
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

	removedTypes := stringSet(
		// A per-register daily session replaces the per-operator POS shift.
		"orders/pos.RegisterShift",
		"orders/pos.ShiftTotalsSnapshot",

		// Canonical product models exclude all endpoint-specific
		// product projections and product-owned snapshot duplicates.
		"orders/pos.CatalogProduct",
		"orders/order.VolumeDiscountTier",
		"supply/product.DetailImage",
		"supply/product.ProductBarcodeAssignmentSnapshot",
		"supply/product.ProductPackageOptionSnapshot",
		"supply/product.PreorderPolicy",
		"supply/product.Snapshot",
		"supply/product.SoonExpiryMerchandisingPolicy",
		"supply/product.StorefrontCommercial",
		"supply/product.StorefrontDisplay",
		"supply/product.StorefrontExpiryDisplay",
		"supply/product.StorefrontMerchandising",
		"supply/product.StorefrontOrigin",
		"supply/product.StorefrontPreorderDisplay",
		"supply/product.StorefrontProduct",
		"supply/product.StorefrontPromotionBadge",
		"supply/product/product_enums.StorefrontExpiryStatus",
		"supply/product/product_enums.StorefrontPreorderStatus",
		"supply/product/product_enums.WholesalePriceMode",
		"supply/import_compliance.ProductSnapshot",

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

		// The contract excludes catalogue commerce projections, the combined
		// package-pricing record, and the category-shaped classification SKU.
		"supply/product.ProductPackaging",
		"supply/product.ProductCommerce",
		"supply/product.ProductPackageCommerce",
		"supply/product.ProductMetrics",
		"supply/product.Selling",
		"supply/classification.SKU",
		"pricing/promotion.PackagePricing",
		"pubsub/event.PackagePricingAvailabilityChangedEvent",
	)

	removedFields := map[string]map[string]struct{}{
		// Catalog relationships and listing evidence persist immutable
		// business codes only. Mutable display projections and listing database
		// identifiers are resolved by their owning service.
		"supply/classification.CollectionRef": stringSet(
			"Name",
		),
		"supply/classification.CategoryTagRef": stringSet(
			"Name",
		),
		"supply/classification.BrandRef": stringSet(
			"Name", "Logo",
		),
		"supply/classification.ProductSupplierRef": stringSet(
			"Name",
		),
		"pubsub/event.CatalogListingChangedEvent": stringSet(
			"ListingID",
		),
		"supply/listing.SaleEligibilitySnapshot": stringSet(
			"ListingID",
		),
		"common/geography.Address": stringSet(
			"City", "State", "Postcode",
		),
		"insights/analytics.OrderItemFact": stringSet(
			"Quantity",
		),
		"insights/analytics.RefundItemFact": stringSet(
			"Quantity",
		),
		"insights/analytics.SKUDemandForecast": stringSet(
			"CurrentStockAtRun",
		),
		"supply/classification.SKUSeries": stringSet(
			"Storage",
		),
		"orders/pos.CatalogProduct": stringSet(
			"ID", "SKU", "Barcode", "Storage", "Price", "CurrentStock",
		),
		// Product.ID is the global product identity.
		// Retired Product fields remain forbidden, while
		// the current surface excludes
		// commerce/metrics/packaging projections it retired.
		"supply/product.Product": stringSet(
			"SKU", "Barcode", "Storage", "PlacingArea", "CurrentStock",
			"RestockedAt", "ExpiredAt", "Pricing", "Physical",
			"Packaging", "Commerce", "Metrics", "Selling", "Taxed",
		),
		"supply/product.Snapshot": stringSet(
			"ID", "SKU", "Storage", "DisplayStatus", "Barcode",
		),
		"supply/product.StorefrontDisplay":       stringSet(),
		"supply/product.StorefrontMerchandising": stringSet(),
		"supply/product.StorefrontProduct": stringSet(
			"SKU", "Barcode", "Storage", "CurrentStock", "Pricing", "ExpiryDate", "DisplayStatus",
		),
		"supply/import_compliance.LabelMaster": stringSet(
			"SourceProduct", "SKU",
		),
		"supply/import_compliance.DeclarationLine": stringSet(
			"ProductReference",
		),
		"supply/import_compliance.TariffLineSnapshot": stringSet(
			"SKU",
		),
		"supply/import_compliance.TariffProfile":     stringSet(),
		"supply/import_compliance.TrademarkEvidence": stringSet(),
		"supply/purchase.OrderItem": stringSet(
			"OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode", "ExpireAt",
		),
		"supply/purchase.ReceiptItem": stringSet(
			"SKU", "OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode",
		),
		"orders/order.CartItem": stringSet(
			"Price", "Quantity",
		),
		"orders/order.DemandBucket": stringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.OpenDemandLine": stringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.Order": stringSet(
			"DeliveryRegion",
		),
		"orders/order.OrderItem": stringSet(
			"UnitPrice", "Quantity", "CartonQty", "CartonSize",
		),
		"orders/order.OrderLineSummary": stringSet(
			"Quantity", "UnitPrice",
		),
		"orders/order.OrderPackingProgress": stringSet(
			"BoxPlan",
		),
		"orders/order.PreorderItemState": stringSet(
			"OrderedQuantity", "AllocatedQuantity",
		),
		"orders/order.VolumeDiscountTier": stringSet(
			"MinCartons",
		),
		"orders/shipping.DeliveryAreaRate": stringSet(
			"Postcode", "Suburb", "DeliveryRegion",
		),
		"orders/shipping.Zone": stringSet(
			"States", "Postcodes", "IsLocal",
		),
		"supply/warehouse.DamageReport": stringSet(
			"DamagedQty",
		),
		"supply/warehouse.Depot": stringSet(
			"PostcodeRules",
		),
		"supply/operations.InboundItem": stringSet(
			"Barcode", "Storage", "ExpectedQty", "ReceivedQty", "LocationCode",
		),
		"supply/operations.PackingDamage": stringSet(
			"SKU", "DamagedQty", "DamageReportID", "StockMovementID",
		),
		"orders/shipping.PackingDiscrepancy": stringSet(
			"OrderedQty", "ScannedQty", "DiffQty", "UnitPrice", "ReturnToStock",
			"DamageReportID", "StockMovementID", "DamagedQty",
		),
		"supply/operations.PackingLine": stringSet(
			"SKU", "OrderedQty", "ScannedQty", "DamagedQty",
		),
		"supply/operations.PickingListItem": stringSet(
			"Barcode", "Location", "QuantityRequired", "QuantityPicked",
		),
		"supply/warehouse.StockLocation": stringSet(
			"Code", "Zone",
		),
		"supply/warehouse.StockLocationProductBalance": stringSet(
			"AvailabilityRevision",
		),
		"supply/operations.StockMovement": stringSet(
			"SKU", "ProductName", "DepotCode", "LocationCode", "QtyDelta", "BalanceAfter",
			"CreatedBy", "SalesOrderNumber", "ReferenceType", "ReferenceID", "Metadata",
		),
		"supply/warehouse.WMSDraft": stringSet(
			"TotalQty",
		),
		"supply/warehouse.WMSDraftItem": stringSet(
			"Barcode", "LocationCode", "Qty", "ExpiryYM",
		),
	}

	// These JSON keys are unambiguously retired everywhere. Reused names such
	// as state, sku, barcode, quantity, unit_price, pricing, storage, and
	// location_code are checked against their exact owning type below.
	removedJSONKeys := stringSet(
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

		// Foreign-key and package-pricing boundary. The frozen
		// "sku_code" key is deliberately NOT listed here: it survives on
		// the seven transaction-evidence types and is policed by the
		// canonical-product allowlist test instead.
		"product_sku_code",
		"product_sku_codes",
		"resolved_product_sku_codes",
		"resolved_qualifier_product_sku_codes",
		"resolved_target_product_sku_codes",
		"created_product_sku_code",
		"category_sku_code",
		"accepted_package_pricing",
		"package_pricing_id",
		"package_pricing_revision",
		"available_package_count",
		"taxed",
		"selling",
		"first_listed_at",
		"display_selling_count",
	)

	removedJSONKeysByType := map[string]map[string]struct{}{
		"supply/classification.CollectionRef": stringSet(
			"name",
		),
		"supply/classification.CategoryTagRef": stringSet(
			"name",
		),
		"supply/classification.BrandRef": stringSet(
			"name", "logo",
		),
		"supply/classification.ProductSupplierRef": stringSet(
			"name",
		),
		"pubsub/event.CatalogListingChangedEvent": stringSet(
			"listing_id",
		),
		"supply/listing.SaleEligibilitySnapshot": stringSet(
			"listing_id",
		),
		"common/geography.Address": stringSet(
			"city", "state", "postcode",
		),
		"insights/analytics.OrderItemFact": stringSet(
			"quantity",
		),
		"insights/analytics.RefundItemFact": stringSet(
			"quantity",
		),
		"insights/analytics.SKUDemandForecast": stringSet(
			"current_stock_at_run",
		),
		"supply/classification.SKUSeries": stringSet(
			"storage",
		),
		"orders/pos.CatalogProduct": stringSet(
			"id", "sku", "barcode", "storage", "price", "current_stock",
		),
		"supply/product.Product": stringSet(
			"sku", "barcode", "storage", "placing_area", "current_stock",
			"restocked_at", "expired_at", "pricing", "physical",
			"packaging", "commerce", "metrics", "selling", "taxed",
		),
		"supply/product.Snapshot": stringSet(
			"id", "sku", "storage", "display_status", "barcode",
		),
		"supply/product.StorefrontDisplay":       stringSet(),
		"supply/product.StorefrontMerchandising": stringSet(),
		"supply/product.StorefrontProduct": stringSet(
			"sku", "barcode", "storage", "current_stock", "pricing", "expiry_date", "display_status",
		),
		"supply/import_compliance.LabelMaster": stringSet(
			"source_product", "sku",
		),
		"supply/import_compliance.DeclarationLine": stringSet(
			"product_reference",
		),
		"supply/import_compliance.TariffLineSnapshot": stringSet(
			"sku",
		),
		"supply/import_compliance.TariffProfile":     stringSet(),
		"supply/import_compliance.TrademarkEvidence": stringSet(),
		"supply/purchase.OrderItem": stringSet(
			"ordered_qty", "received_qty", "rejected_qty", "location_code", "expire_at",
		),
		"supply/purchase.ReceiptItem": stringSet(
			"sku", "ordered_qty", "received_qty", "rejected_qty", "location_code",
		),
		"orders/order.CartItem": stringSet(
			"price", "quantity",
		),
		"orders/order.DemandBucket": stringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.OpenDemandLine": stringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.Order": stringSet(
			"delivery_region",
		),
		"orders/order.OrderItem": stringSet(
			"unit_price", "quantity", "carton_qty", "carton_size",
		),
		"orders/order.OrderLineSummary": stringSet(
			"quantity", "unit_price",
		),
		"orders/order.OrderPackingProgress": stringSet(
			"box_plan",
		),
		"orders/order.PreorderItemState": stringSet(
			"ordered_quantity", "allocated_quantity",
		),
		"orders/order.VolumeDiscountTier": stringSet(
			"min_cartons",
		),
		"orders/shipping.DeliveryAreaRate": stringSet(
			"postcode", "suburb", "delivery_region",
		),
		"orders/shipping.Zone": stringSet(
			"states", "postcodes", "is_local",
		),
		"supply/warehouse.DamageReport": stringSet(
			"damaged_qty",
		),
		"supply/warehouse.Depot": stringSet(
			"postcode_rules",
		),
		// Depots are the only site identity, so no record carries a
		// store code, and POS records key on the register's daily session.
		"orders/pos.Register": stringSet(
			"store_id",
		),
		"orders/pos.CashMovement": stringSet(
			"shift_id",
		),
		"orders/order.POSAttribution": stringSet(
			"store_id", "shift_id",
		),
		"payments/terminal.Terminal": stringSet(
			"store_id",
		),
		"supply/operations.InboundItem": stringSet(
			"barcode", "storage", "expected_qty", "received_qty", "location_code",
		),
		"supply/operations.PackingDamage": stringSet(
			"sku", "damaged_qty", "damage_report_id", "stock_movement_id",
		),
		"orders/shipping.PackingDiscrepancy": stringSet(
			"ordered_qty", "scanned_qty", "diff_qty", "unit_price", "return_to_stock",
			"damage_report_id", "stock_movement_id", "damaged_qty",
		),
		"supply/operations.PackingLine": stringSet(
			"sku", "ordered_qty", "scanned_qty", "damaged_qty",
		),
		"supply/operations.PickingListItem": stringSet(
			"barcode", "location", "quantity_required", "quantity_picked",
		),
		"supply/warehouse.StockLocation": stringSet(
			"code", "zone",
		),
		"supply/warehouse.StockLocationProductBalance": stringSet(
			"availability_revision",
		),
		"supply/operations.StockMovement": stringSet(
			"sku", "product_name", "depot_code", "location_code", "qty_delta", "balance_after",
			"created_by", "sales_order_number", "reference_type", "reference_id", "metadata",
		),
		"supply/warehouse.WMSDraft": stringSet(
			"total_qty",
		),
		"supply/warehouse.WMSDraftItem": stringSet(
			"barcode", "location_code", "qty", "expiry_ym",
		),
	}

	pkgRoot := sharedContractPkgRoot(t)
	seenTypes := make(map[string]struct{})
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
				t.Errorf("%s contains removed identifier %s", path, identifier.Name)
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

				typeKey := productionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				seenTypes[typeKey] = struct{}{}
				if _, removed := removedTypes[typeKey]; removed {
					t.Errorf("%s declares removed type %s", path, typeKey)
				}

				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, removed := removedFields[typeKey][name.Name]; removed {
							t.Errorf("%s declares removed field %s.%s", path, typeKey, name.Name)
						}
					}

					jsonKey, present := jsonFieldName(t, path, field)
					if !present || jsonKey == "-" {
						continue
					}
					if _, removed := removedJSONKeys[jsonKey]; removed {
						t.Errorf("%s contains removed JSON key %q", path, jsonKey)
					}
					if _, removed := removedJSONKeysByType[typeKey][jsonKey]; removed {
						t.Errorf("%s contains removed JSON key %q on %s", path, jsonKey, typeKey)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for mapName, typeMap := range map[string]map[string]map[string]struct{}{
		"removedFields":         removedFields,
		"removedJSONKeysByType": removedJSONKeysByType,
	} {
		for typeKey := range typeMap {
			if _, exists := seenTypes[typeKey]; exists {
				continue
			}
			if _, explicitlyRemoved := removedTypes[typeKey]; explicitlyRemoved {
				continue
			}
			t.Errorf("%s contains unresolved type key %q", mapName, typeKey)
		}
	}
}

func TestGoSourcesContainNoOlderContractImports(t *testing.T) {
	const contractImportRoot = "github.com/Potato-Mart/Backend-Shared-Contract/"
	const currentContractImportPrefix = contractImportRoot + "v32/"
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
			if strings.HasPrefix(importPath, contractImportRoot) && !strings.HasPrefix(importPath, currentContractImportPrefix) {
				t.Errorf("%s imports a non-current shared-contract path %s", path, importPath)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestProductionModelsRejectRetiredPromotionOfferAndImportComplianceTerms(t *testing.T) {
	retiredImportCompliance := "import" + "compliance"
	retiredImportComplianceEnums := retiredImportCompliance + "_enums"
	retiredIdentifiers := stringSet(
		"AcceptedOffer",
		"AppliedPromotion",
		"AppliedPromotions",
		"CatalogProduct",
		"CouponAppliesTo",
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
		"ProductSnapshot",
		"PointPromotion",
		"MembershipPromotionTarget",
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
		retiredImportCompliance,
		retiredImportComplianceEnums,
	)
	retiredTypes := stringSet(
		"common/security.Media",
		"common/security.MediaReference",
		"orders/pos.CatalogProduct",
		"orders/order.VolumeDiscountTier",
		"pricing/promotion.ActiveWindow",
		"pricing/promotion.DiscountSpec",
		"pricing/promotion.EffectivePromotion",
		"pricing/promotion.GroupOrderDiscountApplication",
		"pricing/promotion.GroupOrderDiscountDecision",
		"pricing/promotion.GroupOrderDiscountDecisionLine",
		"pricing/promotion.GroupOrderDiscountProposal",
		"pricing/promotion.ReceiptOffer",
		"pricing/promotion.StorefrontPromotion",
		"pricing/promotion.UsageLimits",
		"pricing/promotion/promotion_enums.DiscountScope",
		"pricing/promotion/promotion_enums.CouponAppliesTo",
		"pricing/promotion/promotion_enums.DiscountType",
		"pricing/promotion/promotion_enums.GroupOrderDiscountState",
		"pricing/promotion/promotion_enums.PromotionAddonTrigger",
		"pricing/promotion/promotion_enums.PromotionClass",
		"pricing/promotion/promotion_enums.PromotionDiscountTarget",
		"pricing/promotion/promotion_enums.PromotionQtyMode",
		"pricing/promotion/promotion_enums.PromotionType",
		"pricing/promotion/promotion_enums.VolumeDiscountAppliesTo",
		"pricing/membership.PointPromotion",
		"pricing/membership/membership_enums.MembershipPromotionTarget",
		"pubsub/event.SellableOfferAvailabilityChangedEvent",
		"supply/product.DetailImage",
		"supply/product.PreorderPolicy",
		"supply/product.SellableOffer",
		"supply/product.SellableOfferDateMarkSnapshot",
		"supply/product.SellableOfferDiscountSnapshot",
		"supply/product.SellableOfferSnapshot",
		"supply/product.Snapshot",
		"supply/product.SoonExpiryMerchandisingPolicy",
		"supply/product.StorefrontCommercial",
		"supply/product.StorefrontDisplay",
		"supply/product.StorefrontMerchandising",
		"supply/product.StorefrontProduct",
		"supply/import_compliance.ProductSnapshot",
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
		"contracts/supply/" + retiredImportCompliance,
		"contracts/supply/" + retiredImportCompliance + "/" + retiredImportComplianceEnums,
		"contracts/supply/product/offer.go",
		"contracts/supply/classification/sku.go",
		"contracts/pricing/promotion/package_pricing.go",
	}

	pkgRoot := sharedContractPkgRoot(t)
	for _, retiredPath := range retiredPaths {
		absolutePath := filepath.Join(pkgRoot, filepath.FromSlash(retiredPath))
		if _, err := os.Stat(absolutePath); err == nil {
			t.Errorf("retired production path remains: %s", retiredPath)
		} else if !os.IsNotExist(err) {
			t.Errorf("inspect retired production path %s: %v", retiredPath, err)
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
		// Marketing owns its closed promotion and discount vocabulary under its
		// canonical namespace. The retired Pricing-owned grammar remains
		// forbidden in every other namespace.

		marketingFoundation := strings.HasPrefix(relativePath, "contracts/marketing/")
		if strings.Contains(relativePath, "/"+retiredImportCompliance+"/") || strings.HasSuffix(relativePath, "/"+retiredImportCompliance) {
			t.Errorf("%s retains retired import-compliance package path", path)
		}

		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		if file.Name != nil && file.Name.Name == retiredImportCompliance {
			t.Errorf("%s retains retired import-compliance package identifier", path)
		}
		for _, specification := range file.Imports {
			importPath, unquoteErr := strconv.Unquote(specification.Path.Value)
			if unquoteErr != nil {
				return unquoteErr
			}
			if strings.Contains(importPath, "/"+retiredImportCompliance) {
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
				typeKey := productionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				if _, retired := retiredTypes[typeKey]; retired {
					t.Errorf("%s declares retired type %s", path, typeKey)
				}
			}
		}

		ast.Inspect(file, func(node ast.Node) bool {
			switch value := node.(type) {
			case *ast.Ident:
				if _, retired := retiredIdentifiers[value.Name]; retired && !marketingFoundation {
					t.Errorf("%s retains retired identifier %s", path, value.Name)
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
						t.Errorf("%s retains retired wire fragment %q", path, fragment)
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

func stringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func productionTypeKey(pkgRoot string, path string, typeName string) string {
	relativePath, _ := filepath.Rel(pkgRoot, path)
	directory := filepath.ToSlash(filepath.Dir(relativePath))
	directory = strings.TrimPrefix(directory, "contracts/")
	if directory == "." || directory == "" {
		return typeName
	}
	return directory + "." + typeName
}

func jsonFieldName(t *testing.T, path string, field *ast.Field) (string, bool) {
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
