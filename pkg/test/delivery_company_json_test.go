package pkg_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/customers/retail"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/shipping"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/courier"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/courier/courier_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/fulfilment"
)

func TestDeliveryLegacyJSONRemainsUnchanged(t *testing.T) {
	for name, tc := range map[string]struct {
		wire  string
		model any
	}{
		"slot":      {`{"id":"old-slot","start_at":"2026-10-01T21:00:00Z","end_at":"2026-10-02T07:00:00Z","label":"Daytime","availability":"available"}`, &shipping.DeliverySlot{}},
		"schedule":  {`{"availability":"available","revision":4,"timezone":"Australia/Melbourne","carrier":"Be Cool Refrigerated Couriers","date_groups":[]}`, &shipping.DeliverySchedule{}},
		"preferred": {`{"date":"2026-10-02","slot_id":"old-slot","start_at":"2026-10-01T21:00:00Z","end_at":"2026-10-02T07:00:00Z","schedule_revision":4}`, &shipping.PreferredDeliverySlot{}},
	} {
		t.Run(name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.wire), tc.model); err != nil {
				t.Fatal(err)
			}
			assertDeliveryJSON(t, tc.model, tc.wire)
		})
	}
	for name, model := range map[string]any{"order": order.Order{}, "shipment": fulfilment.OutboundShipment{}} {
		t.Run(name, func(t *testing.T) {
			data, err := json.Marshal(model)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(data, &fields); err != nil {
				t.Fatal(err)
			}
			if _, present := fields["delivery_selection"]; present {
				t.Fatal("legacy record unexpectedly emits delivery_selection")
			}
		})
	}
}

func TestDeliverySelectionSurvivesOrdersToSupplyHandoff(t *testing.T) {
	const wire = `{"delivery_company":{"code":"melbourne-fleet","name":"Local delivery","integration":"api","adapter":"detrack","revision":7},"schedule_code":"daytime","schedule_revision":12,"slot_id":"opaque-slot","date":"2026-10-05","start_at":"2026-10-04T20:00:00Z","end_at":"2026-10-05T07:00:00Z","timezone":"Australia/Melbourne","source":"configured_schedule"}`
	var selection shipping.DeliverySelection
	if err := json.Unmarshal([]byte(wire), &selection); err != nil {
		t.Fatal(err)
	}
	assertDeliveryJSON(t, selection, wire)
	orderWire, err := json.Marshal(order.Order{OrderNumber: "test-order", OutsourcedCarrier: "Local delivery", DeliverySelection: &selection})
	if err != nil {
		t.Fatal(err)
	}
	var received order.Order
	if err := json.Unmarshal(orderWire, &received); err != nil {
		t.Fatal(err)
	}
	shipmentWire, err := json.Marshal(fulfilment.OutboundShipment{OrderNumber: received.OrderNumber, Carrier: received.DeliverySelection.DeliveryCompany.Code, DeliverySelection: received.DeliverySelection})
	if err != nil {
		t.Fatal(err)
	}
	var persisted fulfilment.OutboundShipment
	if err := json.Unmarshal(shipmentWire, &persisted); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(persisted.DeliverySelection, &selection) {
		t.Fatalf("selection lost in handoff: %+v", persisted.DeliverySelection)
	}
	if received.OutsourcedCarrier != "Local delivery" || persisted.Carrier != "melbourne-fleet" {
		t.Fatal("legacy carrier meanings changed")
	}
}

func TestCourierExplicitCoverageWireShape(t *testing.T) {
	// These postal strings exercise JSON shape only; they are not approved
	// provider coverage or rollout configuration.
	for _, wire := range []string{
		`{"code":"configured","country_code":"AU","postal_code_mode":"include_only","include_postal_codes":["0800","3000"],"exclude_postal_codes":["3000"],"routing_priority":10,"enabled":true,"provider_coverage_required":false}`,
		`{"code":"provider-checked","country_code":"AU","postal_code_mode":"all_except","exclude_postal_codes":["3000"],"routing_priority":20,"enabled":true,"provider_coverage_required":true}`,
		`{"code":"zone-bound","country_code":"AU","shipping_zone":{"id":"zone_test_1"},"state_codes":["AU-VIC"],"postal_code_mode":"include_only","routing_priority":10,"enabled":true,"provider_coverage_required":false}`,
	} {
		var area courier.DeliveryServiceArea
		if err := json.Unmarshal([]byte(wire), &area); err != nil {
			t.Fatal(err)
		}
		assertDeliveryJSON(t, area, wire)
	}
	// Adapter identifiers are open, independently of the stable company code.
	const reference = `{"code":"new-fleet","name":"Future carrier","integration":"api","adapter":"future-adapter","revision":2}`
	var ref courier.DeliveryCompanyRef
	if err := json.Unmarshal([]byte(reference), &ref); err != nil {
		t.Fatal(err)
	}
	assertDeliveryJSON(t, ref, reference)
}

func TestCourierIntegrationModeAndAdapterCompatibility(t *testing.T) {
	for _, mode := range []string{"api", "manual"} {
		t.Run(mode, func(t *testing.T) {
			legacy := `{"code":"legacy-company","name":"Legacy","integration":"` + mode + `","revision":1}`
			var ref courier.DeliveryCompanyRef
			if err := json.Unmarshal([]byte(legacy), &ref); err != nil {
				t.Fatal(err)
			}
			assertDeliveryJSON(t, ref, legacy)
			var company courier.DeliveryCompany
			if err := json.Unmarshal([]byte(legacy), &company); err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(company)
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(data, &fields); err != nil {
				t.Fatal(err)
			}
			if _, ok := fields["adapter"]; ok {
				t.Fatal("legacy record gained inferred adapter")
			}
			if string(fields["integration"]) != `"`+mode+`"` {
				t.Fatal("legacy integration mode changed")
			}
		})
	}
	for _, code := range []string{"melbourne-one", "melbourne-two"} {
		company := courier.DeliveryCompany{Code: code, Integration: "api", Adapter: "detrack"}
		data, err := json.Marshal(company)
		if err != nil {
			t.Fatal(err)
		}
		var received courier.DeliveryCompany
		if err := json.Unmarshal(data, &received); err != nil {
			t.Fatal(err)
		}
		if received.Code != code || received.Integration != "api" || received.Adapter != "detrack" {
			t.Fatalf("instance/mode/adapter identity conflated: %+v", received)
		}
	}
}

func TestCourierConnectionAndCustomerReferenceStaySeparate(t *testing.T) {
	company := courier.DeliveryCompany{Code: "fleet", Name: "Fleet", Integration: "api", Adapter: "future-adapter", Revision: 3,
		Connection: &courier.DeliveryConnection{CredentialConfigured: true, Health: courier_enums.DeliveryConnectionHealthHealthy}}
	assertDeliveryJSON(t, company.Connection, `{"credential_configured":true,"health":"healthy"}`)
	ref := courier.DeliveryCompanyRef{Code: company.Code, Name: company.Name, Integration: company.Integration, Adapter: company.Adapter, Revision: company.Revision}
	assertDeliveryJSON(t, ref, `{"code":"fleet","name":"Fleet","integration":"api","adapter":"future-adapter","revision":3}`)
	// An explicit reviewed key set prevents accidental secret/diagnostic fields
	// from entering the shared connection projection, even with omitempty.
	connectionType := reflect.TypeOf(courier.DeliveryConnection{})
	if connectionType.NumField() != 3 {
		t.Fatal("review new connection fields for credential exposure")
	}
}

func TestCourierProviderCredentialModelHasPrivilegedStandaloneShape(t *testing.T) {
	modelType := reflect.TypeOf(courier.DeliveryProviderCredentials{})
	wantFields := []struct {
		name string
		json string
	}{
		{"SignInAccount", "sign_in_account,omitempty"},
		{"Password", "password,omitempty"},
		{"APIBaseURL", "api_base_url,omitempty"},
		{"APIToken", "api_token,omitempty"},
	}
	if modelType.NumField() != len(wantFields) {
		t.Fatalf("DeliveryProviderCredentials has %d fields, want exactly %d", modelType.NumField(), len(wantFields))
	}
	for index, want := range wantFields {
		field := modelType.Field(index)
		if field.Name != want.name || field.Type.Kind() != reflect.String || field.Tag.Get("json") != want.json {
			t.Errorf("DeliveryProviderCredentials field %d = %s %s json:%q, want %s string json:%q", index, field.Name, field.Type, field.Tag.Get("json"), want.name, want.json)
		}
	}
	data, err := json.Marshal(courier.DeliveryProviderCredentials{})
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{}` {
		t.Fatalf("empty credential value should omit unset fields: %s", data)
	}
}

func TestCourierCredentialTypeIsAbsentFromGeneralAndCustomerModels(t *testing.T) {
	credentialType := reflect.TypeOf(courier.DeliveryProviderCredentials{})
	models := map[string]reflect.Type{
		"delivery company":           reflect.TypeOf(courier.DeliveryCompany{}),
		"delivery company reference": reflect.TypeOf(courier.DeliveryCompanyRef{}),
		"delivery connection":        reflect.TypeOf(courier.DeliveryConnection{}),
		"delivery selection":         reflect.TypeOf(shipping.DeliverySelection{}),
		"retail customer":            reflect.TypeOf(retail.RetailCustomer{}),
	}
	for name, model := range models {
		if containsModelType(model, credentialType, map[reflect.Type]bool{}) {
			t.Errorf("%s must not embed or contain privileged courier credentials", name)
		}
	}
}

func TestCourierLegacyConfiguredSchedulesRemainReadable(t *testing.T) {
	const wire = `{"schedules":[{"code":"legacy-day","country_code":"AU","days_of_week":[1],"start_time":"09:00","end_time":"17:00","timezone":"Australia/Melbourne","minimum_lead_time_minutes":0,"enabled":true}]}`
	var company courier.DeliveryCompany
	if err := json.Unmarshal([]byte(wire), &company); err != nil {
		t.Fatal(err)
	}
	if len(company.Schedules) != 1 || company.Schedules[0].Code != "legacy-day" {
		t.Fatalf("legacy schedules were not retained for compatibility: %+v", company.Schedules)
	}
}

func TestDeliveryFeesAndOfferMetadata(t *testing.T) {
	window := courier.DeliveryServiceWindow{Code: "day", CountryCode: "AU", DaysOfWeek: []int{1, 2, 3, 4, 5}, StartTime: "07:00", EndTime: "18:00", Timezone: "Australia/Melbourne", MinimumLeadTimeMinutes: 120, Enabled: true}
	assertDeliveryJSON(t, window, `{"code":"day","country_code":"AU","days_of_week":[1,2,3,4,5],"start_time":"07:00","end_time":"18:00","timezone":"Australia/Melbourne","minimum_lead_time_minutes":120,"enabled":true}`)
	window.Fee = &money.Money{AmountMinor: 0, Currency: "AUD"}
	data, err := json.Marshal(window)
	if err != nil {
		t.Fatal(err)
	}
	var decoded courier.DeliveryServiceWindow
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Fee == nil || decoded.Fee.AmountMinor != 0 || decoded.Fee.Currency != "AUD" {
		t.Fatal("explicit free fee became absent")
	}
	expires := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)
	for _, source := range []courier_enums.DeliverySlotSource{courier_enums.DeliverySlotSourceConfiguredSchedule, courier_enums.DeliverySlotSourceProvider} {
		slot := shipping.DeliverySlot{ID: "offer", Source: source, ScheduleCode: "day", ExpiresAt: &expires, Availability: "unavailable", UnavailableReason: "coverage_unknown", DeliveryCompany: &courier.DeliveryCompanyRef{Code: "fleet", Integration: "api", Adapter: "future", Revision: 4}}
		data, err := json.Marshal(slot)
		if err != nil {
			t.Fatal(err)
		}
		var decoded shipping.DeliverySlot
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(slot, decoded) {
			t.Fatalf("offer metadata changed: %+v", decoded)
		}
	}
}

func assertDeliveryJSON(t *testing.T, model any, want string) {
	t.Helper()
	got, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}
	var actual, expected any
	if err := json.Unmarshal(got, &actual); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(want), &expected); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("JSON mismatch\ngot: %s\nwant: %s", got, want)
	}
}

func containsModelType(model, target reflect.Type, seen map[reflect.Type]bool) bool {
	for model.Kind() == reflect.Pointer || model.Kind() == reflect.Slice || model.Kind() == reflect.Array {
		model = model.Elem()
	}
	if model == target {
		return true
	}
	if seen[model] {
		return false
	}
	seen[model] = true
	switch model.Kind() {
	case reflect.Struct:
		for index := 0; index < model.NumField(); index++ {
			if containsModelType(model.Field(index).Type, target, seen) {
				return true
			}
		}
	case reflect.Map:
		return containsModelType(model.Key(), target, seen) || containsModelType(model.Elem(), target, seen)
	}
	return false
}
