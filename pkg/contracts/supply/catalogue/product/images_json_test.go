package product

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification"
)

func TestImagesCarryOptionalAltTextAndCaptionOnEveryImage(t *testing.T) {
	images := Images{
		Cover:   &classification.ObjectMediaRef{Code: "MED-COVER", AltText: "Potato bag", Caption: "2 kg bag"},
		Gallery: []classification.ObjectMediaRef{{Code: "MED-GALLERY", AltText: "Potatoes in a bowl"}, {Code: "MED-GALLERY-2"}},
		Details: []classification.ObjectMediaRef{{Code: "MED-DETAIL", Caption: "Nutrition panel"}},
	}
	payload, err := json.Marshal(images)
	if err != nil {
		t.Fatalf("marshal images: %v", err)
	}
	want := `{"cover":{"code":"MED-COVER","alt_text":"Potato bag","caption":"2 kg bag"},` +
		`"gallery":[{"code":"MED-GALLERY","alt_text":"Potatoes in a bowl"},{"code":"MED-GALLERY-2"}],` +
		`"details":[{"code":"MED-DETAIL","caption":"Nutrition panel"}]}`
	if string(payload) != want {
		t.Fatalf("Images JSON = %s, want %s", payload, want)
	}
	var decoded Images
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal images: %v", err)
	}
	if !reflect.DeepEqual(decoded, images) {
		t.Fatalf("Images round trip = %+v, want %+v", decoded, images)
	}

	selling := SellingProductImages{
		Cover:   &security.ObjectMedia{Code: "MED-COVER", URL: "/media/cover", AltText: "Potato bag", Caption: "2 kg bag"},
		Gallery: []security.ObjectMedia{{Code: "MED-GALLERY", URL: "/media/gallery"}},
		Details: []security.ObjectMedia{{Code: "MED-DETAIL", URL: "/media/detail", AltText: "Nutrition panel"}},
	}
	payload, err = json.Marshal(selling)
	if err != nil {
		t.Fatalf("marshal selling images: %v", err)
	}
	want = `{"cover":{"code":"MED-COVER","url":"/media/cover","alt_text":"Potato bag","caption":"2 kg bag"},` +
		`"gallery":[{"code":"MED-GALLERY","url":"/media/gallery"}],` +
		`"details":[{"code":"MED-DETAIL","url":"/media/detail","alt_text":"Nutrition panel"}]}`
	if string(payload) != want {
		t.Fatalf("SellingProductImages JSON = %s, want %s", payload, want)
	}
}

func TestImagesHasExactlyCanonicalObjectMediaFields(t *testing.T) {
	typ := reflect.TypeOf(Images{})
	if typ.NumField() != 3 {
		t.Fatalf("Images has %d fields, want exactly Cover, Gallery, and Details", typ.NumField())
	}

	expected := []struct {
		name string
		typ  reflect.Type
	}{
		{name: "Cover", typ: reflect.TypeOf((*classification.ObjectMediaRef)(nil))},
		{name: "Gallery", typ: reflect.TypeOf([]classification.ObjectMediaRef{})},
		{name: "Details", typ: reflect.TypeOf([]classification.ObjectMediaRef{})},
	}
	for index, want := range expected {
		field := typ.Field(index)
		if field.Name != want.name || field.Type != want.typ {
			t.Fatalf("Images field %d = %s %s, want %s %s", index, field.Name, field.Type, want.name, want.typ)
		}
	}

	contentType := reflect.TypeOf(ProductContent{})
	images, exists := contentType.FieldByName("Images")
	if !exists || images.Type != reflect.TypeOf((*Images)(nil)) || images.Tag.Get("json") != "images,omitempty" {
		t.Fatalf("ProductContent.Images = %#v, want *Images json:images,omitempty", images)
	}
	if _, exists := reflect.TypeOf(Product{}).FieldByName("Media"); exists {
		t.Fatal("Product must not retain the retired Media field")
	}
}
