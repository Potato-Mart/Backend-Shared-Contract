package customers_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/contracts/customers"
	"github.com/Potato-Mart/Backend-Shared-Contract/v7/pkg/enums"
)

func TestRetailCustomerJSONShape(t *testing.T) {
	customer := customers.RetailCustomer{
		ID: "retail_123",
		Identity: common.IdentityLink{
			UserID:                "user_123",
			AccountID:             "acct_123",
			PrimaryAuthIdentityID: "auth_123",
		},
		BasicInfo: customers.RetailCustomerBasicInfo{
			CustomerNumber: "RC-123",
			Name:           common.PersonName{DisplayName: "Retail Customer"},
			Contacts:       common.ContactChannels{Email: "retail@example.com"},
		},
		Lifecycle: customers.RetailCustomerLifecycle{
			Status: enums.CustomerStatusActive,
		},
		Loyalty:   customers.RetailCustomerLoyaltyProfile{Points: 120},
		Marketing: customers.RetailCustomerMarketingProfile{EmailOptIn: true},
		Commerce:  customers.RetailCustomerCommerceProfile{TotalOrders: 2},
	}

	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("Marshal RetailCustomer: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("Unmarshal RetailCustomer JSON: %v", err)
	}

	for _, key := range []string{
		"id",
		"identity",
		"basic_info",
		"lifecycle",
		"loyalty",
		"marketing",
		"commerce",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("RetailCustomer JSON missing %q: %s", key, payload)
		}
	}
}
