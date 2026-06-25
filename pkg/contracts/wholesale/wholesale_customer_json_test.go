package wholesale_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/contracts/wholesale"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

func TestWholesaleCustomerJSONShape(t *testing.T) {
	customer := wholesale.WholesaleCustomer{
		ID:                   "wholesale_123",
		Identity:             common.IdentityLink{UserID: "user_123", AccountID: "acct_123"},
		OrganisationID:       "org_123",
		OrganisationAccessID: "access_123",
		BasicInfo: wholesale.WholesaleCustomerBasicInfo{
			CustomerNumber: "WC-123",
			Name:           common.PersonName{DisplayName: "Wholesale Buyer"},
			Contacts:       common.ContactChannels{Email: "buyer@example.com"},
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
		"identity",
		"organisation_id",
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
	if _, ok := got["membership_id"]; ok {
		t.Fatalf("WholesaleCustomer JSON should not include legacy membership_id: %s", payload)
	}
}
