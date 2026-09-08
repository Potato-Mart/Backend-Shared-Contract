package procurement

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/temporal"
)

func netGoodsJSONFields(t *testing.T, value any) map[string]json.RawMessage {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v", value, err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatalf("decode %T JSON fields: %v", value, err)
	}
	return fields
}

func netGoodsRequireJSON(t *testing.T, fields map[string]json.RawMessage, key, want string) {
	t.Helper()
	if got := string(fields[key]); got != want {
		t.Fatalf("JSON %s = %s, want %s", key, got, want)
	}
}

func netGoodsRoundTrip[T any](t *testing.T, value T) T {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v", value, err)
	}
	var decoded T
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("decode %T: %v", value, err)
	}
	if !reflect.DeepEqual(decoded, value) {
		t.Fatalf("%T changed through JSON: got %+v, want %+v", value, decoded, value)
	}
	return decoded
}

func TestNetGoodsAdditionsOmitAbsentLegacyEvidence(t *testing.T) {
	cases := []struct {
		name  string
		value any
		keys  []string
	}{
		{"acquisition", BaseAcquisitionCost{}, []string{"sources", "manual_lock"}},
		{"invoice", SupplierInvoice{}, []string{"cancellation"}},
		{"invoice_line", SupplierInvoiceLine{}, []string{"net_goods_amount", "receipt_allocations"}},
		{"receipt_item", PurchaseReceiptItem{}, []string{"purchase_order_item_id"}},
		{"order_item", PurchaseOrderItem{}, []string{"net_goods_amount"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fields := netGoodsJSONFields(t, tc.value)
			for _, key := range tc.keys {
				if _, found := fields[key]; found {
					t.Fatalf("legacy %s gained absent field %s", tc.name, key)
				}
			}
		})
	}

	var legacy SupplierInvoiceLine
	if err := json.Unmarshal([]byte(`{"id":"line_legacy","sku_code":"A00001","base_units":3,"line_amount":{"amount_minor":101,"currency":"AUD"},"receipt_id":"receipt_legacy"}`), &legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.NetGoodsAmount != nil || legacy.ReceiptAllocations != nil || legacy.ReceiptID != "receipt_legacy" || legacy.LineAmount.AmountMinor != 101 {
		t.Fatalf("legacy invoice evidence changed: %+v", legacy)
	}
}

func TestAcquisitionCostPreservesMultipleExactSourcesAndManualLock(t *testing.T) {
	confirmedAt := time.Date(2026, 9, 8, 4, 5, 6, 0, time.UTC)
	issueDate := temporal.Date("2026-09-07")
	value := BaseAcquisitionCost{
		ID: "cost_1", SKUCode: "A00001", Currency: "AUD",
		Amount: money.Money{AmountMinor: 40, Currency: "AUD"}, Revision: 4, EffectiveFrom: confirmedAt,
		Sources: []NetGoodsCostSourceSnapshot{
			{SourceType: "supplier_invoice", SourceID: "invoice_1", SourceLineID: "line_1", SourceRevision: 3,
				IssueDate: &issueDate, ConfirmedAt: &confirmedAt, ReceiptID: "receipt_1", ReceiptItemID: "item_1",
				BaseUnits: 3, NetGoodsAmount: &money.Money{AmountMinor: 101, Currency: "AUD"}},
			{SourceType: "supplier_invoice", SourceID: "invoice_1", SourceLineID: "line_2", SourceRevision: 3,
				IssueDate: &issueDate, ConfirmedAt: &confirmedAt, ReceiptID: "receipt_1", ReceiptItemID: "item_2",
				BaseUnits: 2, NetGoodsAmount: &money.Money{AmountMinor: 100, Currency: "AUD"}},
		},
		ManualLock: &audit.LifecycleAction{By: "admin_1", At: confirmedAt, Reason: "Agreed supplier cost pending replacement invoice"},
	}
	netGoodsRoundTrip(t, value)
	fields := netGoodsJSONFields(t, value.Sources[0])
	netGoodsRequireJSON(t, fields, "source_line_id", `"line_1"`)
	netGoodsRequireJSON(t, fields, "source_revision", `3`)
	netGoodsRequireJSON(t, fields, "base_units", `3`)
	netGoodsRequireJSON(t, fields, "net_goods_amount", `{"amount_minor":101,"currency":"AUD"}`)
	netGoodsRequireJSON(t, fields, "issue_date", `"2026-09-07"`)
	netGoodsRequireJSON(t, fields, "confirmed_at", `"2026-09-08T04:05:06Z"`)
	netGoodsRequireJSON(t, fields, "provisional", `false`)
}

func TestNetGoodsEvidenceRetainsMissingZeroAndExactMinorUnits(t *testing.T) {
	unknown := NetGoodsCostSourceSnapshot{SourceType: "purchase_receipt", SourceID: "receipt_1", SourceRevision: 1, BaseUnits: 1}
	if _, exists := netGoodsJSONFields(t, unknown)["net_goods_amount"]; exists {
		t.Fatal("unknown source acquired a zero cost")
	}
	knownZero := unknown
	knownZero.NetGoodsAmount = &money.Money{Currency: "AUD"}
	netGoodsRequireJSON(t, netGoodsJSONFields(t, knownZero), "net_goods_amount", `{"amount_minor":0,"currency":"AUD"}`)
	netGoodsRoundTrip(t, knownZero)

	large := unknown
	large.BaseUnits = 9007199254740993
	large.NetGoodsAmount = &money.Money{AmountMinor: 9007199254740993, Currency: "AUD"}
	netGoodsRequireJSON(t, netGoodsJSONFields(t, large), "net_goods_amount", `{"amount_minor":9007199254740993,"currency":"AUD"}`)
	netGoodsRoundTrip(t, large)
}

func TestInvoiceReceiptAllocationsAndPOEvidenceRoundTrip(t *testing.T) {
	line := SupplierInvoiceLine{
		ID: "line_1", SKUCode: "A00001", BaseUnits: 4,
		LineAmount:     money.Money{AmountMinor: 111, Currency: "AUD"},
		NetGoodsAmount: &money.Money{AmountMinor: 101, Currency: "AUD"},
		ReceiptID:      "receipt_1",
		ReceiptAllocations: []SupplierInvoiceReceiptAllocation{
			{ReceiptID: "receipt_1", ReceiptItemID: "item_1", BaseUnits: 3, NetGoodsAmount: &money.Money{AmountMinor: 101, Currency: "AUD"}},
			{ReceiptID: "receipt_1", ReceiptItemID: "item_2", BaseUnits: 1, NetGoodsAmount: &money.Money{Currency: "AUD"}},
		},
	}
	netGoodsRoundTrip(t, line)
	netGoodsRequireJSON(t, netGoodsJSONFields(t, line), "receipt_allocations", `[{"receipt_id":"receipt_1","receipt_item_id":"item_1","base_units":3,"net_goods_amount":{"amount_minor":101,"currency":"AUD"}},{"receipt_id":"receipt_1","receipt_item_id":"item_2","base_units":1,"net_goods_amount":{"amount_minor":0,"currency":"AUD"}}]`)
	allocation := netGoodsJSONFields(t, SupplierInvoiceReceiptAllocation{ReceiptID: "receipt_1", ReceiptItemID: "item_1", BaseUnits: 3})
	if _, exists := allocation["net_goods_amount"]; exists {
		t.Fatal("unknown allocation cost must stay absent")
	}

	item := PurchaseOrderItem{ID: "po_item_1", OrderedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 4}, NetGoodsAmount: &money.Money{AmountMinor: 101, Currency: "AUD"}}
	netGoodsRoundTrip(t, item)
	netGoodsRequireJSON(t, netGoodsJSONFields(t, item), "net_goods_amount", `{"amount_minor":101,"currency":"AUD"}`)
	receiptItem := PurchaseReceiptItem{ID: "item_1", PurchaseOrderItemID: "po_item_1"}
	netGoodsRoundTrip(t, receiptItem)
	netGoodsRequireJSON(t, netGoodsJSONFields(t, receiptItem), "purchase_order_item_id", `"po_item_1"`)
}

func TestSupplierInvoiceCancellationDoesNotReplaceConfirmationEvidence(t *testing.T) {
	at := time.Date(2026, 9, 8, 4, 5, 6, 0, time.UTC)
	value := SupplierInvoice{
		ID:             "invoice_1",
		Reconciliation: &audit.LifecycleAction{By: "admin_1", At: at, Reason: "Confirmed source allocations"},
		Cancellation:   &audit.LifecycleAction{By: "admin_2", At: at.Add(time.Hour), Reason: "Supplier withdrew document"},
	}
	netGoodsRoundTrip(t, value)
	fields := netGoodsJSONFields(t, value)
	netGoodsRequireJSON(t, fields, "cancellation", `{"by":"admin_2","at":"2026-09-08T05:05:06Z","reason":"Supplier withdrew document"}`)
	if _, exists := fields["reconciliation"]; !exists {
		t.Fatal("cancellation lost existing reconciliation evidence")
	}
}

func TestNetGoodsValuationCoverageAndAverageJSON(t *testing.T) {
	value := NetGoodsValuationSnapshot{
		ID: "value_1", SKUCode: "A00001", DepotCode: "DEPOT-1", Currency: "AUD",
		CurrencyExponent: money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		ValuedBaseUnits:  10, ProvisionalBaseUnits: 4, UnvaluedBaseUnits: 3, NetGoodsCostMinor: 101,
		AverageValuedUnitCost:     &money.Money{AmountMinor: 10, Currency: "AUD"},
		ConsumedCostVarianceMinor: -2, RoundingVarianceMinor: 1, Revision: 4,
		AsOf: time.Date(2026, 9, 8, 4, 5, 6, 0, time.UTC),
	}
	netGoodsRoundTrip(t, value)
	fields := netGoodsJSONFields(t, value)
	for key, want := range map[string]string{"valued_base_units": "10", "provisional_base_units": "4", "unvalued_base_units": "3", "net_goods_cost_minor": "101", "consumed_cost_variance_minor": "-2", "rounding_variance_minor": "1"} {
		netGoodsRequireJSON(t, fields, key, want)
	}
	for _, unknown := range []NetGoodsValuationSnapshot{{}, {UnvaluedBaseUnits: 5}} {
		fields := netGoodsJSONFields(t, unknown)
		if _, exists := fields["average_valued_unit_cost"]; exists {
			t.Fatal("absent average must not become a zero-price claim")
		}
		netGoodsRequireJSON(t, fields, "net_goods_cost_minor", "0")
	}
	free := NetGoodsValuationSnapshot{ValuedBaseUnits: 1, AverageValuedUnitCost: &money.Money{Currency: "AUD"}}
	netGoodsRequireJSON(t, netGoodsJSONFields(t, free), "average_valued_unit_cost", `{"amount_minor":0,"currency":"AUD"}`)
}

func TestNetGoodsMovementKeepsCorrectionsLineageAndRoundingExplicit(t *testing.T) {
	value := NetGoodsValuationMovement{
		ID: "movement_3", SKUCode: "A00001", DepotCode: "DEPOT-1", Currency: "AUD",
		ReferenceType: "supplier_credit", ReferenceID: "credit_1", ReferenceLineID: "line_1",
		Sources:           []NetGoodsCostSourceSnapshot{{SourceType: "supplier_invoice", SourceID: "invoice_1", SourceLineID: "line_1", SourceRevision: 2, BaseUnits: 3, NetGoodsAmount: &money.Money{AmountMinor: 101, Currency: "AUD"}}},
		SourceMovementIDs: []string{"receipt_movement_1", "transfer_movement_2"}, ReversesMovementID: "movement_2",
		SourceCorrectionAmount: &money.Money{AmountMinor: -10, Currency: "AUD"},
		OnHandCostMinorDelta:   -6, ConsumedCostVarianceMinorDelta: -3, RoundingVarianceMinorDelta: -1,
		BalanceValuedBaseUnitsAfter: 2, BalanceNetGoodsCostMinorAfter: 61,
		BalanceConsumedCostVarianceMinorAfter: -3, BalanceRoundingVarianceMinorAfter: -1,
		Revision: 3, OccurredAt: time.Date(2026, 9, 8, 4, 5, 6, 0, time.UTC),
		Actor: security.ActorRef{ActorID: "admin_1"}, RequestID: "request_1", CorrelationID: "correlation_1",
	}
	netGoodsRoundTrip(t, value)
	fields := netGoodsJSONFields(t, value)
	for key, want := range map[string]string{
		"source_correction_amount": `{"amount_minor":-10,"currency":"AUD"}`,
		"on_hand_cost_minor_delta": "-6", "consumed_cost_variance_minor_delta": "-3", "rounding_variance_minor_delta": "-1",
		"valued_base_unit_delta": "0", "provisional_base_unit_delta": "0", "unvalued_base_unit_delta": "0",
		"source_movement_ids": `["receipt_movement_1","transfer_movement_2"]`, "reverses_movement_id": `"movement_2"`,
	} {
		netGoodsRequireJSON(t, fields, key, want)
	}

	issue := NetGoodsValuationMovement{ReferenceType: "sale", ValuedBaseUnitDelta: -1, OnHandCostMinorDelta: -34}
	issueFields := netGoodsJSONFields(t, issue)
	if _, exists := issueFields["source_correction_amount"]; exists {
		t.Fatal("ordinary stock issue must not fabricate a financial source correction")
	}
	netGoodsRequireJSON(t, issueFields, "consumed_cost_variance_minor_delta", "0")
	netGoodsRequireJSON(t, issueFields, "rounding_variance_minor_delta", "0")
	zeroCorrection := NetGoodsValuationMovement{SourceCorrectionAmount: &money.Money{Currency: "AUD"}}
	netGoodsRequireJSON(t, netGoodsJSONFields(t, zeroCorrection), "source_correction_amount", `{"amount_minor":0,"currency":"AUD"}`)
}
