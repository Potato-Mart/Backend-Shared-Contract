package classification_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/geography"

	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/classification"
)

func TestSupplierOrganisationDetailJSONShape(t *testing.T) {
	supplier := classification.Supplier{
		OrganisationDetail: party.OrganisationDetail{
			PartyRef: party.PartyRef{
				ID:   "supplier_123",
				Name: "Supplier Co",
			},
			LegalName: "Supplier Legal Pty Ltd",
			ABN:       "10987654321",
			Website:   "https://supplier.example.com",
			RegisteredAddress: &party.ContactAddress{
				Address: &geography.Address{
					Label:              "HQ",
					Line1:              "2 Supply Road",
					Locality:           "Melbourne",
					AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-VIC"},
					PostalCode:         "3000",
					Country:            geography.CountryRef{Code: "AU"},
				},
			},
		},
	}

	payload, err := json.Marshal(supplier)
	if err != nil {
		t.Fatalf("Marshal Supplier: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("Unmarshal Supplier JSON: %v", err)
	}

	for _, key := range []string{
		"id",
		"name",
		"legal_name",
		"abn",
		"website",
		"registered_address",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("Supplier JSON missing %q: %s", key, payload)
		}
	}

	for key := range got {
		if strings.HasPrefix(key, "company_") {
			t.Fatalf("Supplier JSON should not include company-prefixed key %q: %s", key, payload)
		}
	}
}

func TestProductSupplySectionsAreIndependent(t *testing.T) {
	tests := []struct {
		name string
		want string
		got  classification.ProductSupply
	}{
		{
			name: "supplier only",
			want: `{"supplier":{"code":"SUP-1","name":"Taiwan Foods"}}`,
			got: classification.ProductSupply{
				Supplier: &classification.ProductSupplierRef{Code: "SUP-1", Name: "Taiwan Foods"},
			},
		},
		{
			name: "manufacturing only",
			want: `{"manufacturing":{"location":"Taiwan"}}`,
			got: classification.ProductSupply{
				Manufacturing: &classification.ProductManufacturing{Location: "Taiwan"},
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
