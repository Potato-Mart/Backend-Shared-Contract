package product

import (
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/supply/classification"
)

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
		t.Fatal("Product must not retain the legacy Media field")
	}
}
