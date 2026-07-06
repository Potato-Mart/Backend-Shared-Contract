package wholesale_test

import (
	"encoding/json"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/contracts/wholesale"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/customer"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/wholesale"
)

func TestWholesaleApplicationRequestJSONShape(t *testing.T) {
	req := wholesale.WholesaleApplicationRequest{
		Email:        "buyer@example.com",
		Password:     "long-enough-password",
		ContactName:  "A Buyer",
		ContactPhone: "0400000000",
		ContactRole:  "Owner",
		BusinessName: "Example Foods",
		TradingName:  "Example",
		LegalName:    "Example Foods Pty Ltd",
		ABN:          "12345678901",
		BusinessType: "restaurant",
		BusinessAddress: &common.ContactAddress{
			Address: &common.Address{
				Line1:    "1 Market Street",
				City:     "Sydney",
				State:    "NSW",
				Postcode: "2000",
				Country:  "AU",
			},
		},
		ExpectedOrderVolume:   "weekly",
		StorageRequirements:   []string{"chilled", "frozen"},
		PaymentPreference:     "invoice_terms",
		PreferredDeliveryDays: []string{"Monday", "Thursday"},
		Notes:                 "Deliver to loading dock.",
	}

	raw, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	for _, key := range []string{
		"email",
		"password",
		"contact_name",
		"contact_phone",
		"contact_role",
		"business_name",
		"trading_name",
		"legal_name",
		"abn",
		"business_type",
		"business_address",
		"expected_order_volume",
		"storage_requirements",
		"payment_preference",
		"preferred_delivery_days",
		"notes",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, raw)
		}
	}
}

func TestWholesaleApplicationStatusJSONShape(t *testing.T) {
	status := wholesale.WholesaleApplicationStatusResponse{
		WholesaleApplicationResponse: wholesale.WholesaleApplicationResponse{
			State:                     wholesaleenum.WholesaleApplicationStatePending,
			UserID:                    "usr_1",
			WholesaleOrganisationCode: "org_1",
			OrganisationAccessID:      "acc_1",
			WholesaleCustomerNumber:   "cus_1",
			OrganisationStatus:        wholesaleenum.WholesaleOrganisationStatusPending,
			AccessStatus:              wholesaleenum.OrganisationAccessStatusPending,
			CustomerStatus:            customerenum.CustomerStatusInactive,
			EmailVerificationRequired: true,
		},
		BusinessName: "Example Foods",
		TradingName:  "Example",
		ContactName:  "A Buyer",
		Email:        "buyer@example.com",
	}

	raw, err := json.Marshal(status)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	for _, key := range []string{
		"state",
		"user_id",
		"wholesale_organisation_code",
		"organisation_access_id",
		"wholesale_customer_number",
		"organisation_status",
		"access_status",
		"customer_status",
		"email_verification_required",
		"business_name",
		"trading_name",
		"contact_name",
		"email",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, raw)
		}
	}
}

func TestWholesaleApplicationReviewJSONShape(t *testing.T) {
	detail := wholesale.WholesaleApplicationReviewDetail{
		WholesaleApplicationReviewListItem: wholesale.WholesaleApplicationReviewListItem{
			WholesaleApplicationStatusResponse: wholesale.WholesaleApplicationStatusResponse{
				WholesaleApplicationResponse: wholesale.WholesaleApplicationResponse{
					State:                     wholesaleenum.WholesaleApplicationStatePending,
					UserID:                    "usr_1",
					WholesaleOrganisationCode: "org_1",
					OrganisationAccessID:      "acc_1",
					WholesaleCustomerNumber:   "cus_1",
					OrganisationStatus:        wholesaleenum.WholesaleOrganisationStatusPending,
					AccessStatus:              wholesaleenum.OrganisationAccessStatusPending,
					CustomerStatus:            customerenum.CustomerStatusInactive,
					EmailVerificationRequired: true,
				},
				BusinessName: "Example Foods",
				TradingName:  "Example",
				ContactName:  "A Buyer",
				Email:        "buyer@example.com",
			},
			BusinessType:        "restaurant",
			ContactPhone:        "0400000000",
			ContactRole:         "Owner",
			ABN:                 "12345678901",
			ExpectedOrderVolume: "20 cartons weekly",
			PaymentPreference:   "card",
		},
		BusinessAddress: &common.ContactAddress{
			Address: &common.Address{Line1: "1 Market Street", City: "Sydney", State: "NSW", Postcode: "2000", Country: "AU"},
		},
		StorageRequirements:   []string{"chilled"},
		PreferredDeliveryDays: []string{"Monday"},
		Notes:                 "Loading dock.",
	}

	raw, err := json.Marshal(detail)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	for _, key := range []string{
		"state",
		"wholesale_organisation_code",
		"organisation_access_id",
		"business_name",
		"business_type",
		"contact_phone",
		"contact_role",
		"abn",
		"expected_order_volume",
		"payment_preference",
		"business_address",
		"storage_requirements",
		"preferred_delivery_days",
		"notes",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("missing json key %q in %s", key, raw)
		}
	}
}
