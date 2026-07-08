package purchase_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/contracts/purchase"
)

func TestSupplierOrganisationDetailJSONShape(t *testing.T) {
	supplier := purchase.Supplier{
		OrganisationDetail: common.OrganisationDetail{
			PartyRef: common.PartyRef{
				ID:   "supplier_123",
				Name: "Supplier Co",
			},
			LegalName: "Supplier Legal Pty Ltd",
			ABN:       "10987654321",
			Website:   "https://supplier.example.com",
			RegisteredAddress: &common.ContactAddress{
				Address: &common.Address{
					Label:    "HQ",
					Line1:    "2 Supply Road",
					City:     "Melbourne",
					State:    "VIC",
					Postcode: "3000",
					Country:  "AU",
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
