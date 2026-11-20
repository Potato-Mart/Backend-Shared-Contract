package classification_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/catalogue/classification"
)

func TestCatalogReferencesAreCodeOnly(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{name: "country", value: classification.CountryCodeRef{Code: geography.CountryCode("TW")}, want: `{"code":"TW"}`},
		{name: "collection", value: classification.CollectionRef{Code: "COL0001"}, want: `{"code":"COL0001"}`},
		{name: "category tag", value: classification.CategoryTagRef{Code: "TAG0001"}, want: `{"code":"TAG0001"}`},
		{name: "brand", value: classification.BrandRef{Code: "BRD000001"}, want: `{"code":"BRD000001"}`},
		{name: "supplier", value: classification.ProductSupplierRef{Code: "SUP0001"}, want: `{"code":"SUP0001"}`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(test.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(body) != test.want {
				t.Fatalf("JSON = %s, want %s", body, test.want)
			}
			if typ := reflect.TypeOf(test.value); typ.NumField() != 1 || typ.Field(0).Name != "Code" {
				t.Fatalf("%s fields = %v, want exactly Code", typ, typ)
			}
		})
	}
}

func TestObjectMediaRefCarriesOptionalPlacementCopy(t *testing.T) {
	body, err := json.Marshal(classification.ObjectMediaRef{Code: "MED-SHA256"})
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"code":"MED-SHA256"}` {
		t.Fatalf("ObjectMediaRef without copy JSON = %s, want code only", body)
	}

	withCopy := classification.ObjectMediaRef{Code: "MED-SHA256", AltText: "Bag of washed potatoes", Caption: "Washed, 2 kg"}
	body, err = json.Marshal(withCopy)
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"code":"MED-SHA256","alt_text":"Bag of washed potatoes","caption":"Washed, 2 kg"}`; string(body) != want {
		t.Fatalf("ObjectMediaRef JSON = %s, want %s", body, want)
	}
	var decoded classification.ObjectMediaRef
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded != withCopy {
		t.Fatalf("ObjectMediaRef round trip = %+v, want %+v", decoded, withCopy)
	}

	typ := reflect.TypeOf(classification.ObjectMediaRef{})
	if typ.NumField() != 3 || typ.Field(0).Name != "Code" || typ.Field(1).Name != "AltText" || typ.Field(2).Name != "Caption" {
		t.Fatalf("ObjectMediaRef fields = %v, want exactly Code, AltText and Caption", typ)
	}
}

func TestCommonGeographyAndMediaRenderTypesRemainRich(t *testing.T) {
	countryType := reflect.TypeOf(geography.CountryRef{})
	if _, ok := countryType.FieldByName("Name"); !ok {
		t.Fatal("geography.CountryRef must retain its optional display name")
	}
	mediaType := reflect.TypeOf(security.ObjectMedia{})
	if _, ok := mediaType.FieldByName("URL"); !ok {
		t.Fatal("security.ObjectMedia must retain its render URL")
	}
}
