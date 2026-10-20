package selection_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v45/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v45/pkg/contracts/supply/catalogue/classification"
	"github.com/Potato-Mart/Backend-Shared-Contract/v45/pkg/contracts/supply/catalogue/selection"
	"github.com/Potato-Mart/Backend-Shared-Contract/v45/pkg/contracts/supply/catalogue/selection/selection_enums"
)

func TestProductSelectionJSONShapeAndRoundTrip(t *testing.T) {
	original := selection.ProductSelection{
		ID: "64c13ab08edf48a008793cb1", Code: "SEL0001", Slug: "staff-picks",
		Kind:         selection_enums.ProductSelectionKindBadge,
		Status:       selection_enums.ProductSelectionStatusDraft,
		Name:         []localization.LocalizedName{{Language: "en", Name: "Staff Picks"}},
		Description:  []localization.LocalizedDescription{{Language: "en", Description: "Chosen by our team"}},
		Icon:         &classification.ObjectMediaRef{Code: "MED-SEL-ICON"},
		Cover:        &classification.ObjectMediaRef{Code: "MED-SEL-COVER"},
		DisplayOrder: 3,
		Revision:     7,
	}
	body, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`"code":"SEL0001"`, `"slug":"staff-picks"`, `"kind":"badge"`, `"status":"draft"`,
		`"display_order":3`, `"revision":7`,
		`"icon":{"code":"MED-SEL-ICON"}`, `"cover":{"code":"MED-SEL-COVER"}`,
		`"description":[{"language":"en","description":"Chosen by our team"}]`,
	} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("ProductSelection JSON = %s, want %s", body, want)
		}
	}
	if strings.Contains(string(body), `"url"`) {
		t.Fatalf("ProductSelection media must persist only the media code: %s", body)
	}

	var decoded selection.ProductSelection
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, original) {
		t.Fatalf("ProductSelection round trip = %+v, want %+v", decoded, original)
	}
}

func TestProductSelectionOmitsOptionalFields(t *testing.T) {
	body, err := json.Marshal(selection.ProductSelection{
		Code: "SEL0002", Slug: "weekend",
		Kind:   selection_enums.ProductSelectionKindSection,
		Status: selection_enums.ProductSelectionStatusPublished,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, absent := range []string{`"description"`, `"icon"`, `"cover"`, `"deleted_at"`, `"deleted_by"`} {
		if strings.Contains(string(body), absent) {
			t.Fatalf("ProductSelection JSON = %s, must omit %s when empty", body, absent)
		}
	}
	for _, want := range []string{`"kind":"section"`, `"status":"published"`, `"display_order":0`, `"revision":0`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("ProductSelection JSON = %s, want %s", body, want)
		}
	}
}

func TestProductSelectionItemJSONShape(t *testing.T) {
	item := selection.ProductSelectionItem{SelectionCode: "SEL0001", SKUCode: "SKU-0001", Position: 0}
	body, err := json.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"selection_code":"SEL0001","sku_code":"SKU-0001","position":0}` {
		t.Fatalf("ProductSelectionItem JSON = %s", body)
	}
	var decoded selection.ProductSelectionItem
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded != item {
		t.Fatalf("ProductSelectionItem round trip = %+v, want %+v", decoded, item)
	}
}
