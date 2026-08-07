package retail_test

import (
	"encoding/json"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	customers "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/customers/retail"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/orders/shipping"
	event "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/pubsub/event"
	"strings"
	"testing"
	"time"
)

func TestRetailCustomerJSONShape(t *testing.T) {
	customer := customers.RetailCustomer{
		ID:                    "retail_123",
		CustomerNumber:        "RC-123",
		UserID:                "user_123",
		AccountID:             "acct_123",
		PrimaryAuthIdentityID: "auth_123",
		AuthIdentityIDs:       []string{"auth_123"},
		BasicInfo: customers.RetailCustomerBasicInfo{
			Name:     common.PersonName{DisplayName: "Retail Customer"},
			Contacts: common.ContactChannels{Email: "retail@example.com"},
			Gender:   customers.CustomerGenderNonBinary,
		},
		Lifecycle: customers.RetailCustomerLifecycle{
			Status: customers.CustomerStatusActive,
		},
		Marketing: customers.RetailCustomerMarketingProfile{EmailOptIn: true},
		Commerce:  customers.RetailCustomerCommerceProfile{TotalOrders: 2},
		Referral: &customers.RetailCustomerReferralProfile{
			Code:                      "REF-ADA",
			ReferrerCustomerNumber:    "RC-REF",
			UsedReferralCodeConfirmed: true,
			UsedByCount:               2,
			RewardVouchersIssued:      1,
			RewardVoucherCodes:        []string{"REF-ABC"},
		},
		ProfileCompletion: &customers.RetailCustomerProfileCompletion{
			Percent:         70,
			CompletedFields: []string{"email", "full_name", "date_of_birth", "phone", "gender"},
			MissingFields:   []string{"default_shipping", "default_billing"},
		},
		DefaultBilling: &common.ContactAddress{
			Address: &common.Address{Label: "Billing", Line1: "1 Account St", Locality: "Springvale", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-VIC"}, PostalCode: "3171", Country: geography.CountryRef{Code: "AU"}},
		},
		ShippingAddresses: []common.ContactAddress{
			{
				ID:      "addr_123",
				Contact: &common.Recipient{Name: "Retail Customer"},
				Address: &common.Address{Label: "Home", Line1: "2 Account St", Locality: "Springvale", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-VIC"}, PostalCode: "3171", Country: geography.CountryRef{Code: "AU"}},
			},
		},
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
		"customer_number",
		"user_id",
		"account_id",
		"primary_auth_identity_id",
		"auth_identity_ids",
		"basic_info",
		"lifecycle",
		"marketing",
		"commerce",
		"referral",
		"profile_completion",
		"default_billing",
		"shipping_addresses",
	} {
		if _, ok := got[key]; !ok {
			t.Fatalf("RetailCustomer JSON missing %q: %s", key, payload)
		}
	}
	if _, ok := got["identity"]; ok {
		t.Fatalf("RetailCustomer JSON should not include nested identity: %s", payload)
	}
	if _, ok := got["recent_orders"]; ok {
		t.Fatalf("RetailCustomer JSON should not include embedded recent orders: %s", payload)
	}
	basicInfo, ok := got["basic_info"].(map[string]any)
	if !ok {
		t.Fatalf("RetailCustomer JSON basic_info should be an object: %s", payload)
	}
	if _, ok := basicInfo["customer_number"]; ok {
		t.Fatalf("RetailCustomer JSON should not include nested customer_number: %s", payload)
	}
	if basicInfo["gender"] != "non_binary" {
		t.Fatalf("RetailCustomer JSON gender = %v, want non_binary: %s", basicInfo["gender"], payload)
	}
	referral := got["referral"].(map[string]any)
	if _, ok := referral["credited"]; ok {
		t.Fatalf("RetailCustomer referral should not include credited: %s", payload)
	}
	if referral["used_referral_code_confirmed"] != true || referral["used_by_count"].(float64) != 2 {
		t.Fatalf("RetailCustomer referral did not expose replacement fields: %+v", referral)
	}
	addresses := got["shipping_addresses"].([]any)
	address := addresses[0].(map[string]any)
	if address["id"] != "addr_123" {
		t.Fatalf("RetailCustomer saved address id = %v, want addr_123: %s", address["id"], payload)
	}
}

func TestRetailCustomerDeliveryPreferenceJSONShape(t *testing.T) {
	customer := customers.RetailCustomer{
		ID:                    "customer_1",
		BillingSameAsDelivery: true,
		PreferredDeliverySlot: &shipping.PreferredDeliverySlot{
			Date: "2026-08-01", SlotID: "slot_7", ScheduleRevision: 7,
		},
	}
	payload, err := json.Marshal(customer)
	if err != nil {
		t.Fatalf("marshal retail customer delivery preference: %v", err)
	}
	for _, field := range []string{`"billing_same_as_delivery":true`, `"preferred_delivery_slot"`, `"schedule_revision":7`} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("retail customer delivery preference missing %s: %s", field, payload)
		}
	}
}

func TestRetailCustomerMarketingConsentProvenanceRoundTrip(t *testing.T) {
	consentedAt := time.Date(2026, 7, 30, 1, 2, 3, 0, time.UTC)
	profile := customers.RetailCustomerMarketingProfile{
		EmailOptIn:            true,
		SMSOptIn:              true,
		LineOptIn:             true,
		PushOptIn:             true,
		EmailConsentUpdatedAt: &consentedAt,
		EmailConsentSource:    "checkout",
		SMSConsentUpdatedAt:   &consentedAt,
		SMSConsentSource:      "account_preferences",
		LineConsentUpdatedAt:  &consentedAt,
		LineConsentSource:     "account_preferences",
		PushConsentUpdatedAt:  &consentedAt,
		PushConsentSource:     "account_preferences",
	}

	payload, err := json.Marshal(profile)
	if err != nil {
		t.Fatalf("marshal retail customer marketing profile: %v", err)
	}
	for _, field := range []string{
		`"push_opt_in":true`,
		`"email_consent_updated_at":"2026-07-30T01:02:03Z"`,
		`"email_consent_source":"checkout"`,
		`"sms_consent_updated_at":"2026-07-30T01:02:03Z"`,
		`"sms_consent_source":"account_preferences"`,
		`"line_consent_updated_at":"2026-07-30T01:02:03Z"`,
		`"line_consent_source":"account_preferences"`,
		`"push_consent_updated_at":"2026-07-30T01:02:03Z"`,
		`"push_consent_source":"account_preferences"`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("marketing profile JSON missing %s: %s", field, payload)
		}
	}

	var decoded customers.RetailCustomerMarketingProfile
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal retail customer marketing profile: %v", err)
	}
	if !decoded.PushOptIn || decoded.EmailConsentUpdatedAt == nil || decoded.SMSConsentUpdatedAt == nil ||
		decoded.LineConsentUpdatedAt == nil || decoded.PushConsentUpdatedAt == nil ||
		!decoded.PushConsentUpdatedAt.Equal(consentedAt) || decoded.PushConsentSource != "account_preferences" {
		t.Fatalf("marketing consent provenance did not round-trip: %+v", decoded)
	}
}

func TestCustomerConsentChangedEventIncludesPushOptIn(t *testing.T) {
	changedAt := time.Date(2026, 7, 30, 2, 3, 4, 0, time.UTC)
	consentEvent := event.CustomerConsentChangedEvent{
		CustomerNumber: "RC-1",
		EmailOptIn:     true,
		PushOptIn:      true,
		Source:         "account_preferences",
		ChangedAt:      changedAt,
	}

	payload, err := json.Marshal(consentEvent)
	if err != nil {
		t.Fatalf("marshal customer consent changed event: %v", err)
	}
	for _, field := range []string{
		`"push_opt_in":true`,
		`"source":"account_preferences"`,
		`"changed_at":"2026-07-30T02:03:04Z"`,
	} {
		if !strings.Contains(string(payload), field) {
			t.Fatalf("customer consent event missing %s: %s", field, payload)
		}
	}

	var decoded event.CustomerConsentChangedEvent
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal customer consent changed event: %v", err)
	}
	if !decoded.PushOptIn || !decoded.ChangedAt.Equal(changedAt) {
		t.Fatalf("customer consent event did not round-trip: %+v", decoded)
	}
}
