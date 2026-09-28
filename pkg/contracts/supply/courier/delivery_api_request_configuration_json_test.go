package courier

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v37/pkg/contracts/supply/courier/courier_enums"
)

func TestDeliveryAPIRequestConfigurationsPreserveManualRequestFields(t *testing.T) {
	want := DeliveryAPIRequestConfigurations{
		Connection: &DeliveryAPIRequestConfiguration{
			Method:       courier_enums.DeliveryAPIHTTPMethodGet,
			EndpointPath: "/connection/check",
			Headers: []DeliveryAPIRequestField{
				{Key: "X-Client", Value: "client-1"},
				{Key: "Authorization", SecretRef: "api_key"},
			},
			Query: []DeliveryAPIRequestField{{Key: "region", Value: "AU-VIC"}},
		},
		ShippingAreas: &DeliveryAPIRequestConfiguration{
			Method:       courier_enums.DeliveryAPIHTTPMethodPost,
			EndpointPath: "/shipping/areas",
			JSONBody: []DeliveryAPIRequestField{
				{Key: "country_code", Value: "AU"},
				{Key: "include_archived", Value: false},
				{Key: "optional_filter", ValueIsNull: true},
				{Key: "options", Value: map[string]any{"active": true, "labels": []any{"priority"}}},
			},
		},
		TimeSlots: &DeliveryAPIRequestConfiguration{
			Method:       courier_enums.DeliveryAPIHTTPMethodGet,
			EndpointPath: "/delivery/time-slots",
			Query:        []DeliveryAPIRequestField{{Key: "date", Value: "2026-10-05"}},
		},
	}

	const expected = `{"connection":{"method":"GET","endpoint_path":"/connection/check","headers":[{"key":"X-Client","value":"client-1"},{"key":"Authorization","secret_ref":"api_key"}],"query":[{"key":"region","value":"AU-VIC"}]},"shipping_areas":{"method":"POST","endpoint_path":"/shipping/areas","json_body":[{"key":"country_code","value":"AU"},{"key":"include_archived","value":false},{"key":"optional_filter","value_is_null":true},{"key":"options","value":{"active":true,"labels":["priority"]}}]},"time_slots":{"method":"GET","endpoint_path":"/delivery/time-slots","query":[{"key":"date","value":"2026-10-05"}]}}`

	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal request configurations: %v", err)
	}
	if string(encoded) != expected {
		t.Fatalf("request configurations JSON = %s, want %s", encoded, expected)
	}

	var decoded DeliveryAPIRequestConfigurations
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unmarshal request configurations: %v", err)
	}
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("request configurations did not round-trip:\ngot  %#v\nwant %#v", decoded, want)
	}
}

func TestDeliveryAPIRequestFieldDistinguishesNullValueFromOmittedValue(t *testing.T) {
	nullValue := DeliveryAPIRequestField{Key: "optional_filter", ValueIsNull: true}
	secretReference := DeliveryAPIRequestField{Key: "Authorization", SecretRef: "api_key"}

	encodedNull, err := json.Marshal(nullValue)
	if err != nil {
		t.Fatalf("marshal explicit null value: %v", err)
	}
	if string(encodedNull) != `{"key":"optional_filter","value_is_null":true}` {
		t.Fatalf("explicit null JSON = %s", encodedNull)
	}

	encodedSecret, err := json.Marshal(secretReference)
	if err != nil {
		t.Fatalf("marshal omitted value with secret reference: %v", err)
	}
	if string(encodedSecret) != `{"key":"Authorization","secret_ref":"api_key"}` {
		t.Fatalf("secret reference JSON = %s", encodedSecret)
	}

	var decodedNull DeliveryAPIRequestField
	if err := json.Unmarshal(encodedNull, &decodedNull); err != nil {
		t.Fatalf("unmarshal explicit null value: %v", err)
	}
	if !reflect.DeepEqual(decodedNull, nullValue) {
		t.Fatalf("explicit null did not round-trip: got %#v, want %#v", decodedNull, nullValue)
	}
	var decodedSecret DeliveryAPIRequestField
	if err := json.Unmarshal(encodedSecret, &decodedSecret); err != nil {
		t.Fatalf("unmarshal omitted value with secret reference: %v", err)
	}
	if !reflect.DeepEqual(decodedSecret, secretReference) {
		t.Fatalf("omitted value did not round-trip: got %#v, want %#v", decodedSecret, secretReference)
	}
}

func TestDeliveryAPIRequestConfigurationsOmitUnconfiguredPurposes(t *testing.T) {
	encoded, err := json.Marshal(DeliveryAPIRequestConfigurations{})
	if err != nil {
		t.Fatalf("marshal empty request configurations: %v", err)
	}
	if string(encoded) != `{}` {
		t.Fatalf("empty request configurations JSON = %s, want {}", encoded)
	}
}

func TestDeliveryAPIRequestFieldDoesNotSerializeCredentialValues(t *testing.T) {
	field := DeliveryAPIRequestField{Key: "Authorization", SecretRef: "api_key"}
	encoded, err := json.Marshal(field)
	if err != nil {
		t.Fatalf("marshal secret reference: %v", err)
	}
	if string(encoded) != `{"key":"Authorization","secret_ref":"api_key"}` {
		t.Fatalf("secret reference JSON = %s", encoded)
	}
}
