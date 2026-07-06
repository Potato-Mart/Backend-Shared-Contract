package sales_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/product"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/contracts/sales"
	"github.com/Potato-Mart/Backend-Shared-Contract/v11/pkg/enums"
)

func TestPreorderJSONRoundTrip(t *testing.T) {
	requestedAt := time.Date(2026, 7, 6, 10, 0, 0, 0, time.UTC)
	expectedAt := time.Date(2026, 7, 20, 0, 0, 0, 0, time.UTC)
	preorder := sales.Preorder{
		ID:             "pre_1",
		PreorderNumber: "PO-1001",
		Channel:        enums.OrderTypeOnline,
		Status:         sales.PreorderStatusRequested,
		Customer: common.PartyRef{
			ID:    "retail_1",
			Code:  "RC-1001",
			Name:  "Retail Customer",
			Email: "customer@example.com",
		},
		ProductSKUCode:      "A00084",
		Product:             product.Snapshot{ID: "prd_1", SKUCode: "A00084", Name: "Frozen fries", Taxed: true},
		Quantity:            2,
		RequestedAt:         requestedAt,
		ExpectedAvailableAt: &expectedAt,
		CustomerNote:        "Please notify me",
	}

	body, err := json.Marshal(preorder)
	if err != nil {
		t.Fatalf("marshal preorder: %v", err)
	}

	var decoded sales.Preorder
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal preorder: %v", err)
	}
	if decoded.Status != sales.PreorderStatusRequested || decoded.ProductSKUCode != "A00084" || decoded.Quantity != 2 {
		t.Fatalf("preorder did not round-trip: %+v", decoded)
	}
	if decoded.ExpectedAvailableAt == nil || !decoded.ExpectedAvailableAt.Equal(expectedAt) {
		t.Fatalf("expected_available_at did not round-trip: %+v", decoded.ExpectedAvailableAt)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal preorder JSON: %v", err)
	}
	for _, key := range []string{"preorder_number", "status", "customer", "product_sku_code", "product", "quantity", "requested_at", "expected_available_at"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("preorder JSON missing %q: %s", key, body)
		}
	}
}

func TestPreorderSummaryJSONShape(t *testing.T) {
	requestedAt := time.Date(2026, 7, 6, 10, 0, 0, 0, time.UTC)
	updatedAt := requestedAt.Add(2 * time.Hour)
	summary := sales.PreorderSummary{
		PreorderNumber:       "PO-1001",
		Status:               sales.PreorderStatusConverted,
		Channel:              enums.OrderTypeOnline,
		ProductSKUCode:       "A00084",
		Product:              product.Snapshot{SKUCode: "A00084", Name: "Frozen fries", Taxed: true},
		Quantity:             1,
		RequestedAt:          requestedAt,
		UpdatedAt:            updatedAt,
		ConvertedOrderNumber: "SO-2001",
	}

	body, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("marshal preorder summary: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal preorder summary JSON: %v", err)
	}
	if got["status"] != "converted" || got["converted_order_number"] != "SO-2001" {
		t.Fatalf("preorder summary JSON mismatch: %s", body)
	}
	if _, ok := got["internal_note"]; ok {
		t.Fatalf("summary should not expose internal_note: %s", body)
	}
}
