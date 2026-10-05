package procurement_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/common/packaging/packaging_enums"
	purchase "github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/supply/procurement"
)

func TestSelectedPackageCostPreservesExactMinorUnitsWithoutBaseUnitGuess(t *testing.T) {
	ref := packaging.PackageOptionRef{SKUCode: "A00001", Code: "CASE12", Version: 3}
	line := purchase.PurchaseOrderItem{
		SKUCode: ref.SKUCode, PackageOption: ref,
		SelectedPackageCost: &money.Money{AmountMinor: 1000, Currency: "AUD"},
		OrderedComposition: packaging.PackageCompositionSnapshot{TotalBaseUnits: 24, Components: []packaging.PackageComponentSnapshot{
			{PackageOption: ref, HandlingUnit: packaging_enums.PackageHandlingUnitCase, PackageCount: 2, UnitsPerPackage: 12, BaseUnits: 24},
		}},
		LineTotal: money.Money{AmountMinor: 2000, Currency: "AUD"},
	}
	payload, err := json.Marshal(line)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(payload), `"unit_cost"`) {
		t.Fatalf("invented per-base-unit price: %s", payload)
	}
	var got purchase.PurchaseOrderItem
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatal(err)
	}
	if got.UnitCost != nil || got.SelectedPackageCost == nil || got.SelectedPackageCost.AmountMinor != 1000 || got.SelectedPackageCost.Currency != "AUD" || got.LineTotal.AmountMinor != 2000 || got.OrderedComposition.TotalBaseUnits != 24 {
		t.Fatalf("exact package evidence changed: %+v", got)
	}
	component := got.OrderedComposition.Components[0]
	if component.PackageOption != ref || component.PackageCount != 2 || component.UnitsPerPackage != 12 || component.BaseUnits != 24 {
		t.Fatalf("frozen package basis changed: %+v", component)
	}
}

func TestBaseUnitCostJSONRetainsMeaningAndOptionalEvidence(t *testing.T) {
	var line purchase.PurchaseOrderItem
	if err := json.Unmarshal([]byte(`{"unit_cost":{"amount_minor":200,"currency":"AUD"},"line_total":{"amount_minor":4800,"currency":"AUD"}}`), &line); err != nil {
		t.Fatal(err)
	}
	if line.UnitCost == nil || line.UnitCost.AmountMinor != 200 || line.UnitCost.Currency != "AUD" || line.SelectedPackageCost != nil {
		t.Fatalf("base-unit evidence changed: %+v", line)
	}
	for _, field := range []string{"unit_cost", "selected_package_cost"} {
		var zero purchase.PurchaseOrderItem
		if err := json.Unmarshal([]byte(`{"`+field+`":{"amount_minor":0,"currency":"AUD"}}`), &zero); err != nil {
			t.Fatal(err)
		}
		payload, err := json.Marshal(zero)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(payload), `"`+field+`":{"amount_minor":0,"currency":"AUD"}`) {
			t.Fatalf("explicit free-cost evidence lost: %s", payload)
		}
	}
	empty, err := json.Marshal(purchase.PurchaseOrderItem{})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(empty), `"unit_cost"`) || strings.Contains(string(empty), `"selected_package_cost"`) {
		t.Fatalf("absence fabricated price evidence: %s", empty)
	}
}
