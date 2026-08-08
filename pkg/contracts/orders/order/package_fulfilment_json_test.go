package order_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/geography"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/supply/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/supply/warehouse"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/supply/warehouse/warehouse_enums"
)

func TestRetailOrderItemJSONPreservesMixedCaseAndEachPricing(t *testing.T) {
	now := time.Date(2026, 8, 4, 4, 5, 6, 0, time.UTC)
	caseOption := packageOptionSnapshot("pkg_case_12", "CASE-12", packaging_enums.PackageHandlingUnitCase, 12, now)
	eachOption := packageOptionSnapshot("pkg_each", "EACH", packaging_enums.PackageHandlingUnitEach, 1, now)
	caseOffer := offerSnapshot("offer_case", caseOption, money.Money{AmountMinor: 1800, Currency: "AUD"}, now)
	eachOffer := offerSnapshot("offer_each", eachOption, money.Money{AmountMinor: 175, Currency: "AUD"}, now)

	item := sales.OrderItem{
		ID: "item_1",
		Components: []sales.PricedPackageComponent{
			{AcceptedOffer: caseOffer, RequestedPackageCount: 2, RequestedBaseUnits: 24, PackagePrice: caseOffer.PackagePrice, ComponentTotal: money.Money{AmountMinor: 3600, Currency: "AUD"}},
			{AcceptedOffer: eachOffer, RequestedPackageCount: 3, RequestedBaseUnits: 3, PackagePrice: eachOffer.PackagePrice, ComponentTotal: money.Money{AmountMinor: 525, Currency: "AUD"}},
		},
		TotalBaseUnits:     27,
		SubstitutionPolicy: sales.LooseSubstitutionPolicySnapshot{Allowed: true, Source: order_enums.LooseSubstitutionPolicySourceBuyerSelected, CapturedAt: now},
		RequestedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 27, Components: []packaging.PackageComponentSnapshot{
			{PackageOptionID: "pkg_case_12", HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24},
			{PackageOptionID: "pkg_each", HandlingUnit: packaging_enums.PackageHandlingUnitEach, PackageCount: 3, UnitsPerPackage: 1, BaseUnits: 3},
		}},
		Substitutions:  []warehouse.PackageSubstitutionSnapshot{{ID: "sub_1", RequestedCasePackageOptionID: "pkg_case_12", RequestedCaseCount: 1, RequestedUnitsPerCase: 12, FulfilledSealedCaseCount: 0, ReplacementEachPackageOptionID: "pkg_each", ReplacementBaseUnits: 12, LotID: "lot_1", SourceBucketID: "bucket_each_1", ReasonCode: "NO_SEALED_CASE", Operator: "packer_1", CapturedAt: now}},
		DiscountAmount: money.Money{Currency: "AUD"},
		Total:          money.Money{AmountMinor: 4125, Currency: "AUD"},
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("marshal mixed package order item: %v", err)
	}
	for _, want := range []string{`"handling_unit":"CASE"`, `"handling_unit":"EACH"`, `"requested_package_count":2`, `"requested_package_count":3`, `"total_base_units":27`, `"allowed":true`, `"source":"BUYER_SELECTED"`, `"requested_case_count":1`, `"replacement_base_units":12`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("mixed package JSON missing %s: %s", want, body)
		}
	}
	for _, removed := range []string{`"quantity":`, `"unit_price":`, `"carton_qty":`, `"carton_size":`} {
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
	caseOption := packageOptionSnapshot("pkg_case_12", "CASE-12", packaging_enums.PackageHandlingUnitCase, 12, now)
	caseOffer := offerSnapshot("offer_case", caseOption, money.Money{AmountMinor: 1800, Currency: "AUD"}, now)
	aggregateComponent := sales.PricedPackageComponent{
		AcceptedOffer: caseOffer, RequestedPackageCount: 2, RequestedBaseUnits: 24,
		PackagePrice: money.Money{AmountMinor: 1800, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 300, Currency: "AUD"},
		DiscountAmount: money.Money{AmountMinor: 200, Currency: "AUD"}, ComponentTotal: money.Money{AmountMinor: 3700, Currency: "AUD"},
	}
	participantComponent := sales.PricedPackageComponent{
		AcceptedOffer: caseOffer, RequestedPackageCount: 1, RequestedBaseUnits: 12,
		PackagePrice: money.Money{AmountMinor: 1800, Currency: "AUD"}, TaxAmount: money.Money{AmountMinor: 150, Currency: "AUD"},
		DiscountAmount: money.Money{AmountMinor: 100, Currency: "AUD"}, ComponentTotal: money.Money{AmountMinor: 1850, Currency: "AUD"},
	}
	plan := sales.GroupOrderFulfilmentPlan{
		ID: "group_fulfilment_1", GroupOrderCode: "GROUP-1", ParentOrderNumber: "PARENT-1", ParentFulfilmentID: "fulfilment_1",
		AggregateLines: []sales.GroupOrderAggregateLine{{
			ID: "aggregate_1", ProductSKUCode: "A00001", OfferID: "offer_case", OfferRevision: 3, PackageOptionID: "pkg_case_12",
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
	for _, want := range []string{`"role":"CONSOLIDATED_PARENT"`, `"aggregate_lines"`, `"participant_shares"`, `"captured_at":"2026-08-04T06:07:08Z"`} {
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
	for _, key := range []string{"components", "returned_composition", "refunded_composition", "discount_amount", "tax_amount", "total", "refund_amount"} {
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

func packageOptionSnapshot(id string, code string, handling packaging_enums.PackageHandlingUnit, units int64, capturedAt time.Time) product.ProductPackageOptionSnapshot {
	return product.ProductPackageOptionSnapshot{ID: id, Code: code, ProductSKUCode: "A00001", HandlingUnit: handling, UnitsPerPackage: units, EffectiveFrom: capturedAt, CapturedAt: capturedAt}
}

func offerSnapshot(id string, option product.ProductPackageOptionSnapshot, price money.Money, capturedAt time.Time) product.SellableOfferSnapshot {
	return product.SellableOfferSnapshot{
		ID: id, ProductSKUCode: "A00001", DepotCode: "AU-VIC-MEL-DC-01", PackageOption: option,
		AvailablePackageCount: 10, AvailableBaseUnits: 120,
		Condition: warehouse_enums.InventoryConditionGood, Disposition: warehouse_enums.InventoryDispositionStandardSellable,
		Revision: 3, InventoryRevision: 9, PackagePrice: price, TaxAmount: money.Money{Currency: "AUD"},
		ValidFrom: capturedAt, Timezone: "Etc/UTC", CapturedAt: capturedAt,
		GeographicContext: geography.GeographicContext{Source: geography_enums.GeographicContextSourceRetailCustomerProfile, CountryCode: "AU", ScopeRevision: 1, RuleRevision: 3, EvaluationTimezone: "Australia/Melbourne"},
	}
}
