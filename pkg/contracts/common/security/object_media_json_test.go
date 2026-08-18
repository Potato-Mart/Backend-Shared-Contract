package security

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestObjectMediaJSONIsSafeRenderProjection(t *testing.T) {
	media := ObjectMedia{Code: "media_1", URL: "https://cdn.example.test/products/potato.png"}
	payload, err := json.Marshal(media)
	if err != nil {
		t.Fatalf("marshal object media: %v", err)
	}

	var shape map[string]any
	if err := json.Unmarshal(payload, &shape); err != nil {
		t.Fatalf("unmarshal object media: %v", err)
	}
	if len(shape) != 2 || shape["code"] != "media_1" || shape["url"] != "https://cdn.example.test/products/potato.png" {
		t.Fatalf("ObjectMedia JSON = %s, want only code and url", payload)
	}
	for _, forbidden := range []string{"filename", "bucket", "storage_path", "visibility", "mime_type", "size_bytes", "references"} {
		if _, exists := shape[forbidden]; exists {
			t.Fatalf("ObjectMedia JSON leaked %q: %s", forbidden, payload)
		}
	}

	typ := reflect.TypeOf(ObjectMedia{})
	if typ.NumField() != 2 || typ.Field(0).Name != "Code" || typ.Field(1).Name != "URL" {
		t.Fatalf("ObjectMedia fields = %v, want exactly Code and URL", typ)
	}
}

func TestObjectMediaOmitsEmptyURL(t *testing.T) {
	payload, err := json.Marshal(ObjectMedia{Code: "media_1"})
	if err != nil {
		t.Fatalf("marshal object media without URL: %v", err)
	}
	if string(payload) != `{"code":"media_1"}` {
		t.Fatalf("ObjectMedia JSON = %s, want code only", payload)
	}
}

func TestObjectMediaAssetUsesObjectMediaReferences(t *testing.T) {
	asset := ObjectMediaAsset{
		ID:          "media_1",
		Code:        "MED-SHA256",
		Filename:    "potato.png",
		Bucket:      "catalogue-public",
		StoragePath: "products/A0001/potato.png",
		MIMEType:    "image/png",
		SizeBytes:   42,
		References: []ObjectMediaReference{{
			EntityType: "product",
			EntityID:   "A0001",
			Field:      "images.cover",
		}},
	}
	payload, err := json.Marshal(asset)
	if err != nil {
		t.Fatalf("marshal object media asset: %v", err)
	}

	var shape map[string]any
	if err := json.Unmarshal(payload, &shape); err != nil {
		t.Fatalf("unmarshal object media asset: %v", err)
	}
	references, ok := shape["references"].([]any)
	if !ok || len(references) != 1 {
		t.Fatalf("ObjectMediaAsset references = %#v, want one object-media reference", shape["references"])
	}
	reference := references[0].(map[string]any)
	if reference["entity_type"] != "product" || reference["entity_id"] != "A0001" || reference["field"] != "images.cover" {
		t.Fatalf("ObjectMediaReference JSON = %#v", reference)
	}
}
