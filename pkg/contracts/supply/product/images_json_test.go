package product

import (
	"reflect"
	"testing"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v26/pkg/contracts/common/security"
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
		{name: "Cover", typ: reflect.TypeOf((*security.ObjectMedia)(nil))},
		{name: "Gallery", typ: reflect.TypeOf([]security.ObjectMedia{})},
		{name: "Details", typ: reflect.TypeOf([]security.ObjectMedia{})},
	}
	for index, want := range expected {
		field := typ.Field(index)
		if field.Name != want.name || field.Type != want.typ {
			t.Fatalf("Images field %d = %s %s, want %s %s", index, field.Name, field.Type, want.name, want.typ)
		}
	}

	productType := reflect.TypeOf(Product{})
	if _, exists := productType.FieldByName("Media"); exists {
		t.Fatal("Product must not retain the legacy Media field")
	}
	images, exists := productType.FieldByName("Images")
	if !exists || images.Type != reflect.TypeOf((*Images)(nil)) || images.Tag.Get("json") != "images,omitempty" {
		t.Fatalf("Product.Images = %#v, want *Images json:images,omitempty", images)
	}
}
