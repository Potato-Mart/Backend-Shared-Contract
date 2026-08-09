package product

import (
	"encoding/json"
	"strings"
	"testing"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/supply/classification"
)

func TestProductSupplyImagesAndMetricsJSONShape(t *testing.T) {
	zero := int64(0)
	want := Product{
		SKUCode: "A0001",
		Content: ProductContent{
			Name: "Golden potato",
			Images: &Images{
				Cover: &security.ObjectMedia{ID: "media_cover", URL: "https://cdn.example.test/products/A0001/cover.jpg"},
				Gallery: []security.ObjectMedia{
					{ID: "media_gallery_1", URL: "https://cdn.example.test/products/A0001/gallery-1.jpg"},
					{ID: "media_gallery_2", URL: "https://cdn.example.test/products/A0001/gallery-2.jpg"},
				},
				Details: []security.ObjectMedia{
					{ID: "media_detail_1", URL: "https://cdn.example.test/products/A0001/detail-1.jpg"},
					{ID: "media_detail_2", URL: "https://cdn.example.test/products/A0001/detail-2.jpg"},
				},
			},
		},
		Metrics: ProductMetrics{DisplaySellingCount: &zero},
		Supply: &classification.ProductSupply{
			Supplier:      &classification.ProductSupplierRef{Code: "SUP-1", Name: "Taiwan Foods"},
			Manufacturing: &classification.ProductManufacturing{CompanyName: "Taiwan Foods Factory", Location: "Taichung, Taiwan"},
		},
	}

	body, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	var shape map[string]any
	if err := json.Unmarshal(body, &shape); err != nil {
		t.Fatalf("unmarshal product JSON shape: %v", err)
	}
	metrics := shape["metrics"].(map[string]any)
	if got, ok := metrics["display_selling_count"]; !ok || got != float64(0) {
		t.Fatalf("Product JSON = %s, want explicit nested display_selling_count zero", body)
	}
	supply, ok := shape["supply"].(map[string]any)
	if !ok || supply["supplier"].(map[string]any)["code"] != "SUP-1" {
		t.Fatalf("Product JSON = %s, want supply object", body)
	}
	content, ok := shape["content"].(map[string]any)
	if !ok {
		t.Fatalf("Product JSON = %s, want content object", body)
	}
	images, ok := content["images"].(map[string]any)
	if !ok || len(images) != 3 {
		t.Fatalf("Product images = %#v, want exactly cover, gallery, and details", content["images"])
	}
	cover, ok := images["cover"].(map[string]any)
	if !ok || cover["id"] != "media_cover" || cover["url"] != "https://cdn.example.test/products/A0001/cover.jpg" {
		t.Fatalf("Product images.cover = %#v", images["cover"])
	}
	if gallery, ok := images["gallery"].([]any); !ok || len(gallery) != 2 || gallery[0].(map[string]any)["id"] != "media_gallery_1" {
		t.Fatalf("Product images.gallery = %#v, want ordered object media", images["gallery"])
	}
	if details, ok := images["details"].([]any); !ok || len(details) != 2 || details[1].(map[string]any)["id"] != "media_detail_2" {
		t.Fatalf("Product images.details = %#v, want ordered object media", images["details"])
	}
	for _, legacy := range []string{"media", "cover_media_id", "cover_url", "image_media_ids", "image_urls", "detail_images"} {
		if _, exists := shape[legacy]; exists {
			t.Fatalf("Product JSON retained legacy %s: %s", legacy, body)
		}
	}

	var got Product
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal product: %v", err)
	}
	if got.Metrics.DisplaySellingCount == nil || *got.Metrics.DisplaySellingCount != 0 {
		t.Fatalf("display_selling_count did not round-trip: %#v", got.Metrics.DisplaySellingCount)
	}
	if got.Supply == nil || got.Supply.Supplier == nil || got.Supply.Manufacturing == nil {
		t.Fatalf("supply did not round-trip: %+v", got.Supply)
	}
	if got.Content.Images == nil || got.Content.Images.Cover == nil || len(got.Content.Images.Gallery) != 2 || len(got.Content.Images.Details) != 2 {
		t.Fatalf("Product images did not round-trip: %+v", got.Content.Images)
	}
}

func TestProductNestedOptionalFields(t *testing.T) {
	body, err := json.Marshal(Product{SKUCode: "A0001", Content: ProductContent{Name: "Product"}})
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}
	for _, optionalKey := range []string{`"supply"`, `"display_selling_count"`, `"images"`, `"administration"`} {
		if strings.Contains(string(body), optionalKey) {
			t.Fatalf("optional JSON unexpectedly contains %s: %s", optionalKey, body)
		}
	}
}

func TestProductMetricsDisplaySellingCountAcceptsAbsentNullAndZero(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		wantValue *int64
	}{
		{name: "absent", payload: `{"sku_code":"A0001","metrics":{}}`},
		{name: "null", payload: `{"sku_code":"A0001","metrics":{"display_selling_count":null}}`},
		{name: "zero", payload: `{"sku_code":"A0001","metrics":{"display_selling_count":0}}`, wantValue: int64Pointer(0)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got Product
			if err := json.Unmarshal([]byte(tc.payload), &got); err != nil {
				t.Fatalf("unmarshal display_selling_count %s: %v", tc.name, err)
			}
			if tc.wantValue == nil {
				if got.Metrics.DisplaySellingCount != nil {
					t.Fatalf("display_selling_count %s = %v, want nil", tc.name, *got.Metrics.DisplaySellingCount)
				}
				return
			}
			if got.Metrics.DisplaySellingCount == nil || *got.Metrics.DisplaySellingCount != *tc.wantValue {
				t.Fatalf("display_selling_count %s = %v, want %d", tc.name, got.Metrics.DisplaySellingCount, *tc.wantValue)
			}
		})
	}
}

func int64Pointer(value int64) *int64 {
	return &value
}
