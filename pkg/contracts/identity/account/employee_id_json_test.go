package account

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWorkforceProfileEmployeeIDJSONRoundTrip(t *testing.T) {
	for _, employeeID := range []string{"00123456", "00000000", ""} {
		for _, tc := range []struct {
			name     string
			original any
			decoded  any
		}{
			{"user_profile", &UserProfile{ID: "user_1", Email: "staff@example.test", EmployeeID: employeeID, Active: true}, &UserProfile{}},
			{"admin_account_profile", &AdminAccountProfile{ID: "profile_1", UserID: "user_1", AccountID: "account_1", EmployeeID: employeeID}, &AdminAccountProfile{}},
		} {
			t.Run(tc.name+"/"+employeeID, func(t *testing.T) {
				payload, err := json.Marshal(tc.original)
				if err != nil {
					t.Fatalf("marshal profile: %v", err)
				}
				var fields map[string]json.RawMessage
				if err := json.Unmarshal(payload, &fields); err != nil {
					t.Fatalf("decode profile fields: %v", err)
				}
				rawID, present := fields["employee_id"]
				if employeeID == "" {
					if present {
						t.Fatal("empty employee_id must be omitted")
					}
				} else {
					var wireID string
					if err := json.Unmarshal(rawID, &wireID); err != nil {
						t.Fatalf("employee_id must be a JSON string: %v", err)
					}
					if wireID != employeeID {
						t.Fatalf("employee_id = %q, want %q", wireID, employeeID)
					}
				}
				if err := json.Unmarshal(payload, tc.decoded); err != nil {
					t.Fatalf("unmarshal profile: %v", err)
				}
				if !reflect.DeepEqual(tc.decoded, tc.original) {
					t.Fatalf("profile did not round-trip: got %+v, want %+v", tc.decoded, tc.original)
				}
			})
		}
	}
}

func TestLegacyAndCustomerProfilesOmitEmployeeID(t *testing.T) {
	for _, tc := range []struct {
		name    string
		payload string
		profile any
	}{
		{"legacy_user", `{"id":"user_legacy","email":"legacy@example.test","active":true}`, &UserProfile{}},
		{"customer", `{"id":"user_customer","email":"customer@example.test","active":true,"primary_account_type":"retailCustomer"}`, &UserProfile{}},
		{"legacy_admin_account", `{"user_id":"user_legacy","account_id":"account_legacy"}`, &AdminAccountProfile{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.payload), tc.profile); err != nil {
				t.Fatalf("decode existing profile: %v", err)
			}
			payload, err := json.Marshal(tc.profile)
			if err != nil {
				t.Fatalf("marshal existing profile: %v", err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(payload, &fields); err != nil {
				t.Fatalf("decode existing profile fields: %v", err)
			}
			if _, present := fields["employee_id"]; present {
				t.Fatal("existing profile must not gain employee_id")
			}
		})
	}
}

func TestWorkforceProfilesRejectNumericEmployeeID(t *testing.T) {
	for _, profile := range []any{&UserProfile{}, &AdminAccountProfile{}} {
		if err := json.Unmarshal([]byte(`{"employee_id":12345678}`), profile); err == nil {
			t.Errorf("%T must reject a numeric employee_id", profile)
		}
	}
}
