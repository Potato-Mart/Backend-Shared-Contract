package access_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/identity/identity_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v44/pkg/contracts/identity/access"
)

func TestLoginSessionAuthMethodJSONRoundTrip(t *testing.T) {
	for _, method := range []security_enums.AuthMethod{
		security_enums.AuthMethodPassword,
		security_enums.AuthMethodMFA,
		security_enums.AuthMethodPasskey,
		security_enums.AuthMethodSSO,
		security_enums.AuthMethodRefreshToken,
		security_enums.AuthMethodAPIKey,
		security_enums.AuthMethodPIN,
		security_enums.AuthMethod("future_method"),
	} {
		t.Run(method.String(), func(t *testing.T) {
			issuedAt := time.Date(2026, time.October, 6, 0, 0, 0, 0, time.UTC)
			original := access.LoginSession{
				ID:                 "session_1",
				UserID:             "user_1",
				AuthIdentityID:     "identity_1",
				IdentityDomain:     security_enums.IdentityDomainWorkforce,
				Portal:             identity_enums.PortalControl,
				AccountID:          "account_1",
				Audience:           "pos",
				DeviceKey:          "device_1",
				AuthMethod:         method,
				AuthAssuranceLevel: security_enums.AuthAssuranceLevel1,
				IssuedAt:           issuedAt,
				LastSeenAt:         issuedAt.Add(time.Minute),
				ExpiresAt:          issuedAt.Add(time.Hour),
			}
			payload, err := json.Marshal(original)
			if err != nil {
				t.Fatalf("marshal login session: %v", err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(payload, &fields); err != nil {
				t.Fatalf("decode login session fields: %v", err)
			}
			var wireMethod string
			if err := json.Unmarshal(fields["auth_method"], &wireMethod); err != nil {
				t.Fatalf("decode auth_method: %v", err)
			}
			if wireMethod != method.String() {
				t.Fatalf("auth_method = %q, want %q", wireMethod, method.String())
			}
			for _, key := range []string{"pin", "pin_hash", "password_hash", "refresh_token", "refresh_token_hash", "mfa_verified_at"} {
				if _, present := fields[key]; present {
					t.Errorf("login session must not add %q", key)
				}
			}
			var decoded access.LoginSession
			if err := json.Unmarshal(payload, &decoded); err != nil {
				t.Fatalf("unmarshal login session: %v", err)
			}
			if !reflect.DeepEqual(decoded, original) {
				t.Fatalf("login session did not round-trip: got %+v, want %+v", decoded, original)
			}
			if method == "future_method" && decoded.AuthMethod.IsValid() {
				t.Fatal("unknown auth_method must remain invalid after JSON decoding")
			}
		})
	}
}

func TestLegacyLoginSessionJSONOmitsAuthMethodAndAssurance(t *testing.T) {
	var session access.LoginSession
	if err := json.Unmarshal([]byte(`{"id":"legacy_session","user_id":"user_1","device_key":"device_1"}`), &session); err != nil {
		t.Fatalf("decode legacy login session: %v", err)
	}
	if session.AuthMethod != "" || session.AuthMethod.IsValid() {
		t.Fatalf("legacy auth_method = %q, want empty invalid value", session.AuthMethod)
	}
	payload, err := json.Marshal(session)
	if err != nil {
		t.Fatalf("marshal legacy login session: %v", err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatalf("decode legacy login session fields: %v", err)
	}
	for _, key := range []string{"auth_method", "auth_assurance_level", "mfa_verified_at"} {
		if _, present := fields[key]; present {
			t.Errorf("legacy login session must omit %q", key)
		}
	}
}
