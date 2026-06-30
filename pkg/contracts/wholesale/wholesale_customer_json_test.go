package wholesale_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/contracts/wholesale"
	"github.com/Potato-Mart/Backend-Shared-Contract/v10/pkg/enums"
)

func TestWholesaleCustomerJSONShape(t *testing.T) {
	customer := wholesale.WholesaleCustomer{
		ID:                    "wholesale_123",
		CustomerNumber:        "WC-123",
		UserID:                "user_123",
		AccountID:             "acct_123",
		PrimaryAuthIdentityID: "auth_123",
		AuthIdentityIDs:       []string{"auth_123"},
		OrganisationCode:      "org_123",
		OrganisationAccessID:  "access_123",
		BasicInfo: wholesale.WholesaleCustomerBasicInfo{
			Name:     common.PersonName{DisplayName: "Wholesale Buyer"},
			Contacts: common.ContactChannels{Email: "buyer@example.com"},
		},
		Commercial:     wholesale.WholesaleCustomerCommercialProfile{SalesRep: "sales_123"},
		AccountProfile: wholesale.WholesaleCustomerAccountProfile{Status: enums.CustomerStatusActive, RoleKey: "buyer"},
		Terms:          wholesale.WholesaleTerms{TierKey: "standard", PriceTier: 1},
	}

	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Marshal WholesaleCustomer: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("Unmarshal WholesaleCustomer JSON: %v", err)
	}

	for _, key := range []string{
		"id",
		"customer_number",
		"user_id",
		"account_id",
		"primary_auth_identity_id",
		"auth_identity_ids",
		"organisation_code",
		"organisation_access_id",
		"basic_info",
		"commercial",
		"account_profile",
		"terms",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("WholesaleCustomer JSON missing %q: %s", key, payload)
		}
	}
	if _, ok := got["identity"]; ok {
		t.Fatalf("WholesaleCustomer JSON should not include nested identity: %s", payload)
	}
	basicInfo, ok := got["basic_info"].(map[string]any)
	if !ok {
		t.Fatalf("WholesaleCustomer JSON basic_info should be an object: %s", payload)
	}
	if _, ok := basicInfo["customer_number"]; ok {
		t.Fatalf("WholesaleCustomer JSON should not include nested customer_number: %s", payload)
	}
	if _, ok := got["membership_id"]; ok {
		t.Fatalf("WholesaleCustomer JSON should not include legacy membership_id: %s", payload)
	}
}
