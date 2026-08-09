package product

import (
	"encoding/json"

	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/localization"
)

func TestProductSupplyAndDetailImagesJSONShape(t *testing.T) {
	zero := int64(0)
	want := Product{
		SKUCode:         "A0001",
		CategorySKUCode: "A1",
		Name:            "Golden potato",
		Supply: &ProductSupply{
			Supplier: &ProductSupplierRef{
				Code: "SUP-1",
				Name: "Taiwan Foods",
			},
			Manufacturing: &ProductManufacturing{
				CompanyName: "Taiwan Foods Factory",
				Location:    "Taichung, Taiwan",
			},
		},
		DisplaySellingCount: &zero,
		Media: Media{DetailImages: []DetailImage{
			{
				URL: "https://cdn.example.test/products/A0001/front.jpg",
				AltText: []localization.LocalizedText{
					{Language: "en", Text: "Front of package"},
					{Language: "zh-TW", Text: "包裝正面"},
				},
			},
			{
				URL: "https://cdn.example.test/products/A0001/back.jpg",
				Caption: []localization.LocalizedText{
					{Language: "en", Text: "Cooking instructions"},
				},
			},
		}},
	}

	body, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal product story fields: %v", err)
	}

	var shape map[string]any
	if err := json.Unmarshal(body, &shape); err != nil {
		t.Fatalf("unmarshal product JSON shape: %v", err)
	}
	if got, ok := shape["display_selling_count"]; !ok || got != float64(0) {
		t.Fatalf("Product JSON = %s, want explicit display_selling_count zero", body)
	}

	supply, ok := shape["supply"].(map[string]any)
	if !ok {
		t.Fatalf("Product JSON = %s, want supply object", body)
	}
	supplier, ok := supply["supplier"].(map[string]any)
	if !ok || supplier["code"] != "SUP-1" || supplier["name"] != "Taiwan Foods" {
		t.Fatalf("Product supply supplier = %#v, want safe code and name", supply["supplier"])
	}
	if len(supplier) != 2 {
		t.Fatalf("Product supplier JSON = %#v, want only customer-safe code and name", supplier)
	}
	manufacturing, ok := supply["manufacturing"].(map[string]any)
	if !ok || manufacturing["company_name"] != "Taiwan Foods Factory" || manufacturing["location"] != "Taichung, Taiwan" {
		t.Fatalf("Product supply manufacturing = %#v, want company_name and location", supply["manufacturing"])
	}

	media, ok := shape["media"].(map[string]any)
	if !ok {
		t.Fatalf("Product JSON = %s, want media object", body)
	}
	images, ok := media["detail_images"].([]any)
	if !ok || len(images) != 2 {
		t.Fatalf("Product detail_images = %#v, want two ordered images", media["detail_images"])
	}
	first := images[0].(map[string]any)
	second := images[1].(map[string]any)
	if first["url"] != "https://cdn.example.test/products/A0001/front.jpg" || second["url"] != "https://cdn.example.test/products/A0001/back.jpg" {
		t.Fatalf("Product detail_images order changed: %#v", images)
	}
	altText := first["alt_text"].([]any)
	if altText[1].(map[string]any)["text"] != "包裝正面" {
		t.Fatalf("Product detail image alt_text = %#v, want localized text", first["alt_text"])
	}
	caption := second["caption"].([]any)
	if caption[0].(map[string]any)["text"] != "Cooking instructions" {
		t.Fatalf("Product detail image caption = %#v, want localized text", second["caption"])
	}

	var got Product
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product story fields: %v", err)
	}
	if got.DisplaySellingCount == nil || *got.DisplaySellingCount != 0 {
		t.Fatalf("display_selling_count did not round-trip: %#v", got.DisplaySellingCount)
	}
	if got.Supply == nil || got.Supply.Supplier == nil || got.Supply.Manufacturing == nil {
		t.Fatalf("supply did not round-trip: %+v", got.Supply)
	}
	if len(got.Media.DetailImages) != 2 || got.Media.DetailImages[0].URL != want.Media.DetailImages[0].URL || got.Media.DetailImages[1].URL != want.Media.DetailImages[1].URL {
		t.Fatalf("ordered detail_images did not round-trip: %+v", got.Media.DetailImages)
	}
}

func TestStorefrontProductStoryJSONShape(t *testing.T) {
	zero := int64(0)
	want := StorefrontProduct{
		SKUCode:         "A0001",
		CategorySKUCode: "A1",
		Name:            "Golden potato",
		Supply: &ProductSupply{
			Manufacturing: &ProductManufacturing{Location: "Taiwan"},
		},
		DetailImages: []DetailImage{
			{URL: "https://cdn.example.test/products/A0001/first.jpg"},
			{URL: "https://cdn.example.test/products/A0001/second.jpg"},
		},
		DisplaySellingCount: &zero,
	}

	body, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal storefront product story fields: %v", err)
	}
	var shape map[string]any
	if err := json.Unmarshal(body, &shape); err != nil {
		t.Fatalf("unmarshal storefront product JSON shape: %v", err)
	}
	if got, ok := shape["display_selling_count"]; !ok || got != float64(0) {
		t.Fatalf("StorefrontProduct JSON = %s, want explicit display_selling_count zero", body)
	}
	supply := shape["supply"].(map[string]any)
	if _, exists := supply["supplier"]; exists {
		t.Fatalf("manufacturing-only supply must omit supplier: %s", body)
	}
	manufacturing := supply["manufacturing"].(map[string]any)
	if manufacturing["location"] != "Taiwan" {
		t.Fatalf("StorefrontProduct manufacturing = %#v, want location", manufacturing)
	}
	if _, exists := manufacturing["company_name"]; exists {
		t.Fatalf("partial manufacturing must omit empty company_name: %s", body)
	}
	images := shape["detail_images"].([]any)
	if len(images) != 2 || images[0].(map[string]any)["url"] != want.DetailImages[0].URL || images[1].(map[string]any)["url"] != want.DetailImages[1].URL {
		t.Fatalf("StorefrontProduct detail_images order changed: %#v", images)
	}

	var got StorefrontProduct
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal storefront product story fields: %v", err)
	}
	if got.DisplaySellingCount == nil || *got.DisplaySellingCount != 0 || got.Supply == nil || got.Supply.Supplier != nil {
		t.Fatalf("StorefrontProduct story fields did not round-trip: %+v", got)
	}
}

func TestProductStoryFieldsAreOptional(t *testing.T) {
	tests := []struct {
		name  string
		value any
	}{
		{
			name:  "product",
			value: Product{SKUCode: "A0001", CategorySKUCode: "A1", Name: "Product"},
		},
		{
			name:  "storefront product",
			value: StorefrontProduct{SKUCode: "A0001", CategorySKUCode: "A1", Name: "Product"},
		},
		{
			name:  "snapshot",
			value: Snapshot{SKUCode: "A0001", Name: "Product"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.value)
			if err != nil {
				t.Fatalf("marshal product shape: %v", err)
			}
			for _, optionalKey := range []string{`"supply"`, `"display_selling_count"`, `"detail_images"`} {
				if strings.Contains(string(body), optionalKey) {
					t.Fatalf("optional JSON unexpectedly contains %s: %s", optionalKey, body)
				}
			}
		})
	}
}

func TestDisplaySellingCountAcceptsAbsentNullAndZero(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		wantValue *int64
	}{
		{name: "absent", payload: `{"sku_code":"A0001","category_sku_code":"A1","name":"Product"}`},
		{name: "null", payload: `{"sku_code":"A0001","category_sku_code":"A1","name":"Product","display_selling_count":null}`},
		{name: "zero", payload: `{"sku_code":"A0001","category_sku_code":"A1","name":"Product","display_selling_count":0}`, wantValue: int64Pointer(0)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got Product
			if err := json.Unmarshal([]byte(tc.payload), &got); err != nil {
				t.Fatalf("unmarshal display_selling_count %s: %v", tc.name, err)
			}
			if tc.wantValue == nil {
				if got.DisplaySellingCount != nil {
					t.Fatalf("display_selling_count %s = %v, want nil", tc.name, *got.DisplaySellingCount)
				}
				return
			}
			if got.DisplaySellingCount == nil || *got.DisplaySellingCount != *tc.wantValue {
				t.Fatalf("display_selling_count %s = %v, want %d", tc.name, got.DisplaySellingCount, *tc.wantValue)
			}
		})
	}
}

func int64Pointer(value int64) *int64 {
	return &value
}

func TestProductSupplySectionsAreIndependent(t *testing.T) {
	tests := []struct {
		name string
		want string
		got  ProductSupply
	}{
		{
			name: "supplier only",
			want: `{"supplier":{"code":"SUP-1","name":"Taiwan Foods"}}`,
			got: ProductSupply{
				Supplier: &ProductSupplierRef{Code: "SUP-1", Name: "Taiwan Foods"},
			},
		},
		{
			name: "manufacturing only",
			want: `{"manufacturing":{"location":"Taiwan"}}`,
			got: ProductSupply{
				Manufacturing: &ProductManufacturing{Location: "Taiwan"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.got)
			if err != nil {
				t.Fatalf("marshal product supply: %v", err)
			}
			if string(body) != tc.want {
				t.Fatalf("ProductSupply JSON = %s, want %s", body, tc.want)
			}
		})
	}
}

func TestSnapshotUsesCanonicalNestedSupply(t *testing.T) {
	want := Snapshot{
		SKUCode: "A0001",
		Name:    "Golden potato",
		Supply: &ProductSupply{
			Supplier: &ProductSupplierRef{Code: "SUP-1", Name: "Taiwan Foods"},
		},
	}

	body, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal snapshot supply: %v", err)
	}
	if strings.Contains(string(body), `"supplier_code"`) || !strings.Contains(string(body), `"supply":{"supplier":{"code":"SUP-1","name":"Taiwan Foods"}}`) {
		t.Fatalf("Snapshot JSON = %s, want canonical nested supply only", body)
	}
}
