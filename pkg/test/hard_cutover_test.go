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

func TestV27NonSupplyProductionModelsRejectLegacyMediaShapes(t *testing.T) {
	legacyTypes := v27StringSet("Media", "MediaReference")
	legacyFields := v27StringSet(
		"MediaCode", "MediaURL",
		"AvatarMediaCode", "AvatarURL",
		"LogoURL",
		"ImageID", "ImageURL",
		"CoverMediaCode", "CoverURL",
		"ImageMediaCodes", "ImageURLs",
	)
	legacyJSONKeys := v27StringSet(
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
		// Supply-only image records are deliberately exempt from the v27
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
					t.Errorf("%s declares retired v27 media type %s", path, typeSpecification.Name.Name)
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
					jsonKey, present := v27JSONFieldName(t, path, field)
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

func TestV27ProductionModelsContainNoRemovedFieldsOrDeprecations(t *testing.T) {
	removedIdentifiers := v27StringSet(
		// v30 workforce-role cut-over: admin, warehouse, and cashier are
		// replaced by countryAdmin, depotManager, and warehouseManager, and
		// POS sign-in moves to every staff role instead of a cashier role.
		// sales goes the same way: till duty follows the geographic scope a
		// principal holds, so a dedicated selling rank granted no authority
		// the remaining six ranks did not already carry.
		"UserRoleAdmin",
		"UserRoleWarehouse",
		"UserRoleCashier",
		"UserRoleSales",

		// v30 POS cut-over: per-operator shifts are replaced by one daily
		// session per register, shared by every operator on it.
		"RegisterShift",
		"ShiftTotalsSnapshot",
		"ShiftStatus",
		"ShiftStatusOpen",
		"ShiftStatusClosed",
		"ShiftID",

		// Earlier hard cut-overs that remain forbidden in v27.
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

		// v27 global product/SKU split and market pricing cutover.
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
		// movement types in v27.
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

	removedTypes := v27StringSet(
		// v30 replaces the per-operator POS shift with a per-register daily
		// session.
		"orders/pos.RegisterShift",
		"orders/pos.ShiftTotalsSnapshot",

		// The v27 canonical-product cut-over removes all endpoint-specific
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

		// v27 removes the catalogue commerce projections, the combined
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
		// v29.0.1 catalog relationships and listing evidence persist immutable
		// business codes only. Mutable display projections and listing database
		// identifiers are resolved by their owning service.
		"supply/classification.CollectionRef": v27StringSet(
			"Name",
		),
		"supply/classification.CategoryTagRef": v27StringSet(
			"Name",
		),
		"supply/classification.BrandRef": v27StringSet(
			"Name", "Logo",
		),
		"supply/classification.ProductSupplierRef": v27StringSet(
			"Name",
		),
		"pubsub/event.CatalogListingChangedEvent": v27StringSet(
			"ListingID",
		),
		"supply/listing.SaleEligibilitySnapshot": v27StringSet(
			"ListingID",
		),
		"common/geography.Address": v27StringSet(
			"City", "State", "Postcode",
		),
		"insights/analytics.OrderItemFact": v27StringSet(
			"Quantity",
		),
		"insights/analytics.RefundItemFact": v27StringSet(
			"Quantity",
		),
		"insights/analytics.SKUDemandForecast": v27StringSet(
			"CurrentStockAtRun",
		),
		"supply/classification.SKUSeries": v27StringSet(
			"Storage",
		),
		"orders/pos.CatalogProduct": v27StringSet(
			"ID", "SKU", "Barcode", "Storage", "Price", "CurrentStock",
		),
		// v27 reinstates a stable Product.ID as the global product identity;
		// the v26 removal of "ID"/"id" is deliberately retired here. Every
		// other v26 removal on Product remains forbidden, and v27 adds the
		// commerce/metrics/packaging projections it retired.
		"supply/product.Product": v27StringSet(
			"SKU", "Barcode", "Storage", "PlacingArea", "CurrentStock",
			"RestockedAt", "ExpiredAt", "Pricing", "Physical",
			"Packaging", "Commerce", "Metrics", "Selling", "Taxed",
		),
		"supply/product.Snapshot": v27StringSet(
			"ID", "SKU", "Storage", "DisplayStatus", "Barcode",
		),
		"supply/product.StorefrontDisplay":       v27StringSet(),
		"supply/product.StorefrontMerchandising": v27StringSet(),
		"supply/product.StorefrontProduct": v27StringSet(
			"SKU", "Barcode", "Storage", "CurrentStock", "Pricing", "ExpiryDate", "DisplayStatus",
		),
		"supply/import_compliance.LabelMaster": v27StringSet(
			"SourceProduct", "SKU",
		),
		"supply/import_compliance.DeclarationLine": v27StringSet(
			"ProductReference",
		),
		"supply/import_compliance.TariffLineSnapshot": v27StringSet(
			"SKU",
		),
		"supply/import_compliance.TariffProfile":     v27StringSet(),
		"supply/import_compliance.TrademarkEvidence": v27StringSet(),
		"supply/purchase.OrderItem": v27StringSet(
			"OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode", "ExpireAt",
		),
		"supply/purchase.ReceiptItem": v27StringSet(
			"SKU", "OrderedQty", "ReceivedQty", "RejectedQty", "LocationCode",
		),
		"orders/order.CartItem": v27StringSet(
			"Price", "Quantity",
		),
		"orders/order.DemandBucket": v27StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.OpenDemandLine": v27StringSet(
			"SKU", "QtyUnits", "CartonQty", "CartonSize",
		),
		"orders/order.Order": v27StringSet(
			"DeliveryRegion",
		),
		"orders/order.OrderItem": v27StringSet(
			"UnitPrice", "Quantity", "CartonQty", "CartonSize",
		),
		"orders/order.OrderLineSummary": v27StringSet(
			"Quantity", "UnitPrice",
		),
		"orders/order.OrderPackingProgress": v27StringSet(
			"BoxPlan",
		),
		"orders/order.PreorderItemState": v27StringSet(
			"OrderedQuantity", "AllocatedQuantity",
		),
		"orders/order.VolumeDiscountTier": v27StringSet(
			"MinCartons",
		),
		"orders/shipping.DeliveryAreaRate": v27StringSet(
			"Postcode", "Suburb", "DeliveryRegion",
		),
		"orders/shipping.Zone": v27StringSet(
			"States", "Postcodes", "IsLocal",
		),
		"supply/warehouse.DamageReport": v27StringSet(
			"DamagedQty",
		),
		"supply/warehouse.Depot": v27StringSet(
			"PostcodeRules",
		),
		"supply/operations.InboundItem": v27StringSet(
			"Barcode", "Storage", "ExpectedQty", "ReceivedQty", "LocationCode",
		),
		"supply/operations.PackingDamage": v27StringSet(
			"SKU", "DamagedQty", "DamageReportID", "StockMovementID",
		),
		"orders/shipping.PackingDiscrepancy": v27StringSet(
			"OrderedQty", "ScannedQty", "DiffQty", "UnitPrice", "ReturnToStock",
			"DamageReportID", "StockMovementID", "DamagedQty",
		),
		"supply/operations.PackingLine": v27StringSet(
			"SKU", "OrderedQty", "ScannedQty", "DamagedQty",
		),
		"supply/operations.PickingListItem": v27StringSet(
			"Barcode", "Location", "QuantityRequired", "QuantityPicked",
		),
		"supply/warehouse.StockLocation": v27StringSet(
			"Code", "Zone",
		),
		"supply/warehouse.StockLocationProductBalance": v27StringSet(
			"AvailabilityRevision",
		),
		"supply/operations.StockMovement": v27StringSet(
			"SKU", "ProductName", "DepotCode", "LocationCode", "QtyDelta", "BalanceAfter",
			"CreatedBy", "SalesOrderNumber", "ReferenceType", "ReferenceID", "Metadata",
		),
		"supply/warehouse.WMSDraft": v27StringSet(
			"TotalQty",
		),
		"supply/warehouse.WMSDraftItem": v27StringSet(
			"Barcode", "LocationCode", "Qty", "ExpiryYM",
		),
	}

	// These JSON keys are unambiguously retired everywhere. Reused names such
	// as state, sku, barcode, quantity, unit_price, pricing, storage, and
	// location_code are checked against their exact owning type below.
	removedJSONKeys := v27StringSet(
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

		// v27 foreign-key and package-pricing cutover. The frozen
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
		"supply/classification.CollectionRef": v27StringSet(
			"name",
		),
		"supply/classification.CategoryTagRef": v27StringSet(
			"name",
		),
		"supply/classification.BrandRef": v27StringSet(
			"name", "logo",
		),
		"supply/classification.ProductSupplierRef": v27StringSet(
			"name",
		),
		"pubsub/event.CatalogListingChangedEvent": v27StringSet(
			"listing_id",
		),
		"supply/listing.SaleEligibilitySnapshot": v27StringSet(
			"listing_id",
		),
		"common/geography.Address": v27StringSet(
			"city", "state", "postcode",
		),
		"insights/analytics.OrderItemFact": v27StringSet(
			"quantity",
		),
		"insights/analytics.RefundItemFact": v27StringSet(
			"quantity",
		),
		"insights/analytics.SKUDemandForecast": v27StringSet(
			"current_stock_at_run",
		),
		"supply/classification.SKUSeries": v27StringSet(
			"storage",
		),
		"orders/pos.CatalogProduct": v27StringSet(
			"id", "sku", "barcode", "storage", "price", "current_stock",
		),
		"supply/product.Product": v27StringSet(
			"sku", "barcode", "storage", "placing_area", "current_stock",
			"restocked_at", "expired_at", "pricing", "physical",
			"packaging", "commerce", "metrics", "selling", "taxed",
		),
		"supply/product.Snapshot": v27StringSet(
			"id", "sku", "storage", "display_status", "barcode",
		),
		"supply/product.StorefrontDisplay":       v27StringSet(),
		"supply/product.StorefrontMerchandising": v27StringSet(),
		"supply/product.StorefrontProduct": v27StringSet(
			"sku", "barcode", "storage", "current_stock", "pricing", "expiry_date", "display_status",
		),
		"supply/import_compliance.LabelMaster": v27StringSet(
			"source_product", "sku",
		),
		"supply/import_compliance.DeclarationLine": v27StringSet(
			"product_reference",
		),
		"supply/import_compliance.TariffLineSnapshot": v27StringSet(
			"sku",
		),
		"supply/import_compliance.TariffProfile":     v27StringSet(),
		"supply/import_compliance.TrademarkEvidence": v27StringSet(),
		"supply/purchase.OrderItem": v27StringSet(
			"ordered_qty", "received_qty", "rejected_qty", "location_code", "expire_at",
		),
		"supply/purchase.ReceiptItem": v27StringSet(
			"sku", "ordered_qty", "received_qty", "rejected_qty", "location_code",
		),
		"orders/order.CartItem": v27StringSet(
			"price", "quantity",
		),
		"orders/order.DemandBucket": v27StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.OpenDemandLine": v27StringSet(
			"sku", "qty_units", "carton_qty", "carton_size",
		),
		"orders/order.Order": v27StringSet(
			"delivery_region",
		),
		"orders/order.OrderItem": v27StringSet(
			"unit_price", "quantity", "carton_qty", "carton_size",
		),
		"orders/order.OrderLineSummary": v27StringSet(
			"quantity", "unit_price",
		),
		"orders/order.OrderPackingProgress": v27StringSet(
			"box_plan",
		),
		"orders/order.PreorderItemState": v27StringSet(
			"ordered_quantity", "allocated_quantity",
		),
		"orders/order.VolumeDiscountTier": v27StringSet(
			"min_cartons",
		),
		"orders/shipping.DeliveryAreaRate": v27StringSet(
			"postcode", "suburb", "delivery_region",
		),
		"orders/shipping.Zone": v27StringSet(
			"states", "postcodes", "is_local",
		),
		"supply/warehouse.DamageReport": v27StringSet(
			"damaged_qty",
		),
		"supply/warehouse.Depot": v27StringSet(
			"postcode_rules",
		),
		// v30: depots are the only site identity, so no record carries a
		// store code, and POS records key on the register's daily session.
		"orders/pos.Register": v27StringSet(
			"store_id",
		),
		"orders/pos.CashMovement": v27StringSet(
			"shift_id",
		),
		"orders/order.POSAttribution": v27StringSet(
			"store_id", "shift_id",
		),
		"payments/terminal.Terminal": v27StringSet(
			"store_id",
		),
		"supply/operations.InboundItem": v27StringSet(
			"barcode", "storage", "expected_qty", "received_qty", "location_code",
		),
		"supply/operations.PackingDamage": v27StringSet(
			"sku", "damaged_qty", "damage_report_id", "stock_movement_id",
		),
		"orders/shipping.PackingDiscrepancy": v27StringSet(
			"ordered_qty", "scanned_qty", "diff_qty", "unit_price", "return_to_stock",
			"damage_report_id", "stock_movement_id", "damaged_qty",
		),
		"supply/operations.PackingLine": v27StringSet(
			"sku", "ordered_qty", "scanned_qty", "damaged_qty",
		),
		"supply/operations.PickingListItem": v27StringSet(
			"barcode", "location", "quantity_required", "quantity_picked",
		),
		"supply/warehouse.StockLocation": v27StringSet(
			"code", "zone",
		),
		"supply/warehouse.StockLocationProductBalance": v27StringSet(
			"availability_revision",
		),
		"supply/operations.StockMovement": v27StringSet(
			"sku", "product_name", "depot_code", "location_code", "qty_delta", "balance_after",
			"created_by", "sales_order_number", "reference_type", "reference_id", "metadata",
		),
		"supply/warehouse.WMSDraft": v27StringSet(
			"total_qty",
		),
		"supply/warehouse.WMSDraftItem": v27StringSet(
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
				t.Errorf("%s contains removed v27 identifier %s", path, identifier.Name)
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

				typeKey := v27ProductionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				seenTypes[typeKey] = struct{}{}
				if _, removed := removedTypes[typeKey]; removed {
					t.Errorf("%s declares removed v27 type %s", path, typeKey)
				}

				structure, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						if _, removed := removedFields[typeKey][name.Name]; removed {
							t.Errorf("%s declares removed v27 field %s.%s", path, typeKey, name.Name)
						}
					}

					jsonKey, present := v27JSONFieldName(t, path, field)
					if !present || jsonKey == "-" {
						continue
					}
					if _, removed := removedJSONKeys[jsonKey]; removed {
						t.Errorf("%s contains removed v27 JSON key %q", path, jsonKey)
					}
					if _, removed := removedJSONKeysByType[typeKey][jsonKey]; removed {
						t.Errorf("%s contains removed v27 JSON key %q on %s", path, jsonKey, typeKey)
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

func TestV31GoSourcesContainNoOlderContractImports(t *testing.T) {
	const contractImportRoot = "github.com/Potato-Mart/Backend-Shared-Contract/"
	const currentContractImportPrefix = contractImportRoot + "v31/"
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
				t.Errorf("%s imports non-v31 shared-contract path %s", path, importPath)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestV27ProductionModelsRejectRetiredPromotionOfferAndImportComplianceTerms(t *testing.T) {
	legacyImportCompliance := "import" + "compliance"
	legacyImportComplianceEnums := legacyImportCompliance + "_enums"
	retiredIdentifiers := v27StringSet(
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
		legacyImportCompliance,
		legacyImportComplianceEnums,
	)
	retiredTypes := v27StringSet(
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
		"contracts/supply/" + legacyImportCompliance,
		"contracts/supply/" + legacyImportCompliance + "/" + legacyImportComplianceEnums,
		"contracts/supply/product/offer.go",
		"contracts/supply/classification/sku.go",
		"contracts/pricing/promotion/package_pricing.go",
	}

	pkgRoot := sharedContractPkgRoot(t)
	for _, retiredPath := range retiredPaths {
		absolutePath := filepath.Join(pkgRoot, filepath.FromSlash(retiredPath))
		if _, err := os.Stat(absolutePath); err == nil {
			t.Errorf("retired v27 production path remains: %s", retiredPath)
		} else if !os.IsNotExist(err) {
			t.Errorf("inspect retired v27 production path %s: %v", retiredPath, err)
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
		// The additive v30 marketing foundation deliberately reintroduces its
		// own closed promotion and discount vocabulary under the separate
		// canonical marketing namespace. The v27 removal still applies to the
		// legacy Pricing-owned promotion grammar and every other namespace.
		marketingFoundation := strings.HasPrefix(relativePath, "contracts/marketing/")
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
				typeKey := v27ProductionTypeKey(pkgRoot, path, typeSpecification.Name.Name)
				if _, retired := retiredTypes[typeKey]; retired {
					t.Errorf("%s declares retired v27 type %s", path, typeKey)
				}
			}
		}

		ast.Inspect(file, func(node ast.Node) bool {
			switch value := node.(type) {
			case *ast.Ident:
				if _, retired := retiredIdentifiers[value.Name]; retired && !marketingFoundation {
					t.Errorf("%s retains retired v27 identifier %s", path, value.Name)
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
						t.Errorf("%s retains retired v27 wire fragment %q", path, fragment)
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

func v27StringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func v27ProductionTypeKey(pkgRoot string, path string, typeName string) string {
	relativePath, _ := filepath.Rel(pkgRoot, path)
	directory := filepath.ToSlash(filepath.Dir(relativePath))
	directory = strings.TrimPrefix(directory, "contracts/")
	if directory == "." || directory == "" {
		return typeName
	}
	return directory + "." + typeName
}

func v27JSONFieldName(t *testing.T, path string, field *ast.Field) (string, bool) {
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
