package courier

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestDeliveryCompanyProviderSettingsAndMetadataRoundTrip(t *testing.T) {
	company := DeliveryCompany{
		Code: "provider_example",
		ProviderSettings: &DeliveryProviderSettings{SchemaVersion: 3, Values: map[string]any{
			"client_id":     "example-client",
			"max_options":   12,
			"feature_flags": map[string]any{"cold_chain": true},
		}},
		CustomMetadata: []DeliveryCompanyCustomMetadataEntry{{
			Key:       "dispatch_group",
			ValueType: "string",
			Value:     "north",
		}, {
			Key:       "future_value",
			ValueType: "custom_scalar_v2",
			Value:     map[string]any{"nested": []any{1, true}},
		}},
	}

	encoded, err := json.Marshal(company)
	if err != nil {
		t.Fatalf("marshal company: %v", err)
	}

	var decoded DeliveryCompany
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unmarshal company: %v", err)
	}
	if decoded.ProviderSettings == nil || decoded.ProviderSettings.SchemaVersion != company.ProviderSettings.SchemaVersion {
		t.Fatalf("provider settings schema version changed: %#v", decoded.ProviderSettings)
	}
	for key, want := range company.ProviderSettings.Values {
		assertJSONValueEqual(t, decoded.ProviderSettings.Values[key], want, "provider setting "+key)
	}
	if len(decoded.CustomMetadata) != len(company.CustomMetadata) {
		t.Fatalf("custom metadata entries changed: got %d, want %d", len(decoded.CustomMetadata), len(company.CustomMetadata))
	}
	for i, want := range company.CustomMetadata {
		got := decoded.CustomMetadata[i]
		if got.Key != want.Key || got.ValueType != want.ValueType {
			t.Errorf("custom metadata entry %d changed: got %#v, want %#v", i, got, want)
		}
		assertJSONValueEqual(t, got.Value, want.Value, "custom metadata entry value")
	}
}

func TestDeliveryCompanyOmitsAbsentExtensionsForLegacyJSON(t *testing.T) {
	encoded, err := json.Marshal(DeliveryCompany{Code: "provider_example"})
	if err != nil {
		t.Fatalf("marshal legacy company: %v", err)
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &fields); err != nil {
		t.Fatalf("decode legacy company fields: %v", err)
	}
	for _, key := range []string{"provider_settings", "custom_metadata"} {
		if _, exists := fields[key]; exists {
			t.Errorf("absent extension %q was serialized: %s", key, encoded)
		}
	}
}

func TestDeliveryProviderCredentialExtensionRoundTrip(t *testing.T) {
	credentials := DeliveryProviderCredentials{
		ProviderExtension: &DeliveryProviderCredentialExtension{
			SchemaVersion: 2,
			Values: map[string]any{
				"api_key":    "synthetic-test-value",
				"account_id": "example-account",
			},
		},
	}

	encoded, err := json.Marshal(credentials)
	if err != nil {
		t.Fatalf("marshal credentials: %v", err)
	}
	var decoded DeliveryProviderCredentials
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unmarshal credentials: %v", err)
	}
	if decoded.ProviderExtension == nil || decoded.ProviderExtension.SchemaVersion != credentials.ProviderExtension.SchemaVersion {
		t.Fatalf("provider credential extension changed: %#v", decoded.ProviderExtension)
	}
	for key, want := range credentials.ProviderExtension.Values {
		assertJSONValueEqual(t, decoded.ProviderExtension.Values[key], want, "provider credential value "+key)
	}
}

func assertJSONValueEqual(t *testing.T, got, want any, field string) {
	t.Helper()
	gotJSON, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("marshal decoded %s: %v", field, err)
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal expected %s: %v", field, err)
	}
	if !bytes.Equal(gotJSON, wantJSON) {
		t.Errorf("%s changed: got %s, want %s", field, gotJSON, wantJSON)
	}
}
