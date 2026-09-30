package courier

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDeliveryCompanyConfigurationJSON(t *testing.T) {
	// This is a read projection: Supply has already masked the sensitive value.
	const fixture = `{"code":"custom-courier","name":"Custom Courier","enabled":true,"revision":2,"country_codes":["AU"],"capabilities":{"booking":true,"tracking":false,"proof_of_delivery":false,"refrigerated":false,"provider_coverage":true,"provider_slots":true},"configuration":[{"key":"availability_endpoint","value":"https://courier.example/availability","sensitive":false},{"key":"api_key","value":"********","sensitive":true},{"key":"options","value":{"retries":2,"enabled":true,"regions":["AU"]},"sensitive":false}],"service_areas":[]}`
	var company DeliveryCompany
	if err := json.Unmarshal([]byte(fixture), &company); err != nil {
		t.Fatal(err)
	}
	if len(company.Configuration) != 3 || !company.Configuration[1].Sensitive || company.Configuration[1].Value != "********" {
		t.Fatal("configuration entries or sensitivity lost")
	}
	if !company.Capabilities.Booking || !company.Capabilities.ProviderSlots {
		t.Fatal("admin-authored capabilities lost")
	}
	encoded, err := json.Marshal(company)
	if err != nil {
		t.Fatal(err)
	}
	var got, want map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(fixture), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got["configuration"], want["configuration"]) {
		t.Fatalf("typed configuration changed: %s", encoded)
	}
	for _, removed := range []string{"credential_requirements", "dispatch_capable", "connection", "provider_settings", "slot_source", "schedules"} {
		if _, ok := got[removed]; ok {
			t.Errorf("removed field %s serialized", removed)
		}
	}
}
