package order_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging/packaging_enums"
	sales "github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/pricing/quote"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/pricing/quote/quote_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/listing"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/operations"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestRetailOrderItemJSONPreservesMixedCaseAndEachPricing(t *testing.T) {
	now := time.Date(2026, 8, 4, 4, 5, 6, 0, time.UTC)
	caseOption := packageOption("pkg_case_12", "CASE-12", packaging_enums.PackageHandlingUnitCase, 12, now)
	casePricing := priceSnapshot("line_case", money.Money{AmountMinor: 1800, Currency: "AUD"}, now)
	eachPricing := priceSnapshot("line_each", money.Money{AmountMinor: 175, Currency: "AUD"}, now)

	item := sales.OrderItem{
		ID:                   "item_1",
		SKUID:                "A00001",
		ProductName:          "Potatoes",
		ProductPackageOption: caseOption,
		CapturedAt:           now,
		Components: []sales.PricedPackageComponent{
			{PriceSnapshot: casePricing, RequestedPackageCount: 2, RequestedBaseUnits: 24, PackagePrice: casePricing.LineTotal, ComponentTotal: money.Money{AmountMinor: 3600, Currency: "AUD"}},
			{PriceSnapshot: eachPricing, RequestedPackageCount: 3, RequestedBaseUnits: 3, PackagePrice: eachPricing.LineTotal, ComponentTotal: money.Money{AmountMinor: 525, Currency: "AUD"}},
		},
		TotalBaseUnits:     27,
		SubstitutionPolicy: sales.LooseSubstitutionPolicySnapshot{Allowed: true, Source: order_enums.LooseSubstitutionPolicySourceBuyerSelected, CapturedAt: now},
		RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 27, Components: []packaging.PackageComponentSnapshot{
			{PackageOptionID: "pkg_case_12", HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24},
			{PackageOptionID: "pkg_each", HandlingUnit: packaging_enums.PackageHandlingUnitEach, PackageCount: 3, UnitsPerPackage: 1, BaseUnits: 3},
		}},
		Substitutions:  []operations.PackageSubstitutionSnapshot{{ID: "sub_1", RequestedCasePackageOptionID: "pkg_case_12", RequestedCaseCount: 1, RequestedUnitsPerCase: 12, FulfilledSealedCaseCount: 0, ReplacementEachPackageOptionID: "pkg_each", ReplacementBaseUnits: 12, LotID: "lot_1", SourceBucketID: "bucket_each_1", ReasonCode: "NO_SEALED_CASE", Operator: "packer_1", CapturedAt: now}},
		DiscountAmount: money.Money{Currency: "AUD"},
		Total:          money.Money{AmountMinor: 4125, Currency: "AUD"},
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("marshal mixed package order item: %v", err)
	}
	for _, want := range []string{`"price_snapshot":{"quote_id":"quote_1","line_id":"line_case"`, `"handling_unit":"CASE"`, `"handling_unit":"EACH"`, `"requested_package_count":2`, `"requested_package_count":3`, `"total_base_units":27`, `"allowed":true`, `"source":"BUYER_SELECTED"`, `"requested_case_count":1`, `"replacement_base_units":12`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("mixed package JSON missing %s: %s", want, body)
		}
	}
	for _, removed := range []string{`"quantity":`, `"unit_price":`, `"carton_qty":`, `"carton_size":`, `"accepted_offer"`, `"offer_id"`, `"accepted_package_pricing"`, `"package_pricing_id"`} {
		if strings.Contains(string(body), removed) {
			t.Fatalf("mixed package JSON contains removed scalar field %s: %s", removed, body)
		}
	}
}

func TestGroupOrderFulfilmentJSONUsesOneParentAllocation(t *testing.T) {
	now := time.Date(2026, 8, 4, 6, 7, 8, 0, time.UTC)
	composition := packaging.PackageCompositionSnapshot{TotalBaseUnits: 24, Components: []packaging.PackageComponentSnapshot{{PackageOptionID: "pkg_case_12", HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24}}}
	participantComposition := packaging.PackageCompositionSnapshot{TotalBaseUnits: 12, Components: []packaging.PackageComponentSnapshot{{PackageOptionID: "pkg_case_12", HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 1, UnitsPerPackage: 12, BaseUnits: 12}}}
	zeroComposition := packaging.PackageCompositionSnapshot{TotalBaseUnits: 0, Components: []packaging.PackageComponentSnapshot{}}
	casePricing := priceSnapshot("line_case", money.Money{AmountMinor: 1800, Currency: "AUD"}, now)
	aggregateComponent := sales.PricedPackageComponent{
		PriceSnapshot: casePricing, RequestedPackageCount: 2, RequestedBaseUnits: 24,
		PackagePrice: money.Money{AmountMinor: 1800, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 300, Currency: "AUD"},
		DiscountAmount: money.Money{AmountMinor: 200, Currency: "AUD"}, ComponentTotal: money.Money{AmountMinor: 3700, Currency: "AUD"},
	}
	participantComponent := sales.PricedPackageComponent{
		PriceSnapshot: casePricing, RequestedPackageCount: 1, RequestedBaseUnits: 12,
		PackagePrice: money.Money{AmountMinor: 1800, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 150, Currency: "AUD"},
		DiscountAmount: money.Money{AmountMinor: 100, Currency: "AUD"}, ComponentTotal: money.Money{AmountMinor: 1850, Currency: "AUD"},
	}
	plan := sales.GroupOrderFulfilmentPlan{
		ID: "group_fulfilment_1", GroupOrderCode: "GROUP-1", ParentOrderNumber: "PARENT-1", ParentFulfilmentID: "fulfilment_1",
		AggregateLines: []sales.GroupOrderAggregateLine{{
			ID: "aggregate_1", SKUID: "sku_a00001", MarketID: "market_au", PackageOptionID: "pkg_case_12",
			RequestedComposition: composition, AllocatedComposition: composition, ShortageComposition: zeroComposition,
			ReturnedComposition: participantComposition, RefundedComposition: participantComposition,
			Components: []sales.PricedPackageComponent{aggregateComponent}, DiscountAmount: money.Money{AmountMinor: 200, Currency: "AUD"},
			TaxAmount: money.Money{AmountMinor: 300, Currency: "AUD"}, Total: money.Money{AmountMinor: 3700, Currency: "AUD"},
			RefundAmount: money.Money{AmountMinor: 1800, Currency: "AUD"}, ReservationID: "reservation_parent_1", AllocationLineID: "allocation_parent_1",
		}},
		ParticipantShares: []sales.GroupOrderParticipantShare{
			{ParticipantOrderNumber: "PART-1", ParticipantOrderItemID: "item_part_1", ParentAllocationLineID: "allocation_parent_1", Sequence: 1, RequestedComposition: participantComposition, FulfilledComposition: participantComposition, ShortageComposition: zeroComposition, ReturnedComposition: participantComposition, RefundedComposition: participantComposition, Components: []sales.PricedPackageComponent{participantComponent}, DiscountAmount: money.Money{AmountMinor: 100, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 150, Currency: "AUD"}, Total: money.Money{AmountMinor: 1850, Currency: "AUD"}, RefundAmount: money.Money{AmountMinor: 1800, Currency: "AUD"}},
			{ParticipantOrderNumber: "PART-2", ParticipantOrderItemID: "item_part_2", ParentAllocationLineID: "allocation_parent_1", Sequence: 2, RequestedComposition: participantComposition, FulfilledComposition: participantComposition, ShortageComposition: zeroComposition, ReturnedComposition: zeroComposition, RefundedComposition: zeroComposition, Components: []sales.PricedPackageComponent{participantComponent}, DiscountAmount: money.Money{AmountMinor: 100, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 150, Currency: "AUD"}, Total: money.Money{AmountMinor: 1850, Currency: "AUD"}, RefundAmount: money.Money{Currency: "AUD"}},
		},
		Revision: 1, Timezone: "Australia/Melbourne", CapturedAt: now,
	}
	parent := sales.Order{OrderNumber: "PARENT-1", GroupOrder: &sales.GroupOrderContext{GroupOrderCode: "GROUP-1", Role: order_enums.GroupOrderRoleConsolidatedParent, ParentFulfilmentID: "fulfilment_1"}, GroupOrderFulfilment: &plan}

	body, err := json.Marshal(parent)
	if err != nil {
		t.Fatalf("marshal group order fulfilment: %v", err)
	}
	if strings.Count(string(body), `"reservation_id":"reservation_parent_1"`) != 1 {
		t.Fatalf("group fulfilment must expose one parent reservation: %s", body)
	}
	if strings.Count(string(body), `"parent_allocation_line_id":"allocation_parent_1"`) != 2 {
		t.Fatalf("participant shares must reference the parent allocation: %s", body)
	}
	for _, want := range []string{`"market_id":"market_au"`, `"price_snapshot":{"quote_id":"quote_1"`, `"aggregate_lines"`, `"participant_shares"`, `"captured_at":"2026-08-04T06:07:08Z"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("group fulfilment JSON missing %s: %s", want, body)
		}
	}

	planBody, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("marshal group order fulfilment plan: %v", err)
	}
	var planShape map[string]any
	if err := json.Unmarshal(planBody, &planShape); err != nil {
		t.Fatalf("unmarshal group order fulfilment plan: %v", err)
	}
	aggregateShape := planShape["aggregate_lines"].([]any)[0].(map[string]any)
	for _, key := range []string{"market_id", "components", "returned_composition", "refunded_composition", "discount_amount", "tax_amount", "total", "refund_amount"} {
		if _, ok := aggregateShape[key]; !ok {
			t.Fatalf("group aggregate JSON missing %q: %+v", key, aggregateShape)
		}
	}
	if len(aggregateShape["components"].([]any)) != 1 || aggregateShape["refund_amount"].(map[string]any)["amount_minor"] != float64(1800) {
		t.Fatalf("group aggregate JSON lost package component or refund amount: %+v", aggregateShape)
	}
	for i, participantValue := range planShape["participant_shares"].([]any) {
		participantShape := participantValue.(map[string]any)
		for _, key := range []string{"components", "returned_composition", "refunded_composition", "discount_amount", "tax_amount", "total", "refund_amount"} {
			if _, ok := participantShape[key]; !ok {
				t.Fatalf("participant %d JSON missing %q: %+v", i+1, key, participantShape)
			}
		}
		if len(participantShape["components"].([]any)) != 1 {
			t.Fatalf("participant %d JSON lost priced package component: %+v", i+1, participantShape)
		}
	}
}

func packageOption(id string, code string, handling packaging_enums.PackageHandlingUnit, units int64, capturedAt time.Time) product.ProductPackageOption {
	return product.ProductPackageOption{ID: id, Code: code, SKUID: "A00001", HandlingUnit: handling, UnitsPerPackage: units, IsActive: true, EffectiveFrom: capturedAt}
}

func priceSnapshot(lineID string, price money.Money, capturedAt time.Time) quote.PriceSnapshot {
	return quote.PriceSnapshot{
		QuoteID: "quote_1", LineID: lineID, SKUID: "sku_a00001", MarketID: "market_au",
		PriceBookID: "book_au_online", PriceBookRevision: 2,
		PriceEntryID: "entry_1", PriceEntryRevision: 3,
		Currency: "AUD", CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		ListUnitPrice: price, TaxableBase: price, TaxAmount: money.Money{Currency: "AUD"}, LineTotal: price,
		Tax: quote.TaxSnapshot{
			TaxCategoryID: "tax_au_gst", TaxRuleID: "rule_au_gst", TaxRuleRevision: 1,
			InclusionBasis: pricebook_enums.PriceTaxInclusionInclusive,
			RateNumerator:  1, RateDenominator: 11,
			TaxableBase: price, AllocatedTax: money.Money{Currency: "AUD"},
			CalculationSource: quote_enums.TaxCalculationSourceInclusiveExtraction,
			RoundingMethod:    quote_enums.TaxRoundingMethodSumExactThenRound,
		},
		Rounding: quote.RoundingEvidence{
			Mode: quote_enums.RoundingModeHalfUp, PriceEnding: pricebook_enums.PriceEndingPolicyNone,
			Exponent: 2, ExactNumerator: price.AmountMinor, ExactDenominator: 1, RoundedAmount: price,
		},
		Eligibility: listing.SaleEligibilitySnapshot{
			MarketID: "market_au", SKUID: "sku_a00001", ListingID: "listing_1", ListingRevision: 4,
			TaxCategoryID: "tax_au_gst", DepotCode: "AU-VIC-MEL-DC-01",
			StockLocation:      warehouse.StockLocationRef{DepotCode: "AU-VIC-MEL-DC-01", LocationCode: "A-01-03"},
			Condition:          warehouse_enums.InventoryConditionGood,
			Disposition:        warehouse_enums.InventoryDispositionStandardSellable,
			AvailableBaseUnits: 120, InventoryRevision: 9,
			ValidityToken: "eligibility_token_1", ValidUntil: capturedAt, CapturedAt: capturedAt,
		},
		Fingerprint: "fingerprint_" + lineID, CapturedAt: capturedAt,
	}
}
