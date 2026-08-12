package access_test

import (
	"encoding/json"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/identity/identity_enums"
	identity "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/access"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/access/access_enums"
	accountenum "github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/account"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/account/account_enums"
)

func TestPortalAccessJSONGroupsLifecycleAndKeepsCoreFieldsTopLevel(t *testing.T) {
	grantedAt := time.Date(2026, 6, 18, 4, 30, 0, 0, time.UTC)
	access := identity.PortalAccess{
		ID:          "portal_access_1",
		UserID:      "user_1",
		AccountID:   "acct_1",
		AccountType: account_enums.AccountTypeRetailCustomer,
		Portal:      identity_enums.PortalRetail,
		Status:      access_enums.PortalAccessStatusActive,
		Grant:       &audit.LifecycleAction{By: "admin_1", At: &grantedAt, Reason: "approved"},
	}

	payload, err := json.Marshal(access)
	if err != nil {
		t.Fatalf("marshal portal access: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal portal access JSON: %v", err)
	}

	for _, key := range []string{"id", "user_id", "account_id", "account_type", "portal", "status"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("PortalAccess JSON missing top-level %q: %s", key, payload)
		}
	}
	if _, ok := got["grant"]; !ok {
		t.Fatalf("PortalAccess JSON missing nested grant: %s", payload)
	}
	for _, key := range []string{"granted_by", "granted_at", "granted_reason"} {
		if _, ok := got[key]; ok {
			t.Fatalf("PortalAccess JSON should not include flat lifecycle key %q: %s", key, payload)
		}
	}

	var decoded identity.PortalAccess
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal portal access: %v", err)
	}
	if decoded.Grant == nil || decoded.Grant.By != "admin_1" || decoded.Grant.At == nil || !decoded.Grant.At.Equal(grantedAt) {
		t.Fatalf("grant lifecycle did not round-trip: %+v", decoded.Grant)
	}
}

func TestIdentityWholesaleAccessJSONUsesOrganisationAccessID(t *testing.T) {
	claims := identity.AccessTokenClaims{
		Subject:                   "user_1",
		UserID:                    "user_1",
		SessionID:                 "session_1",
		AccountID:                 "acct_1",
		Portal:                    identity_enums.PortalWholesale,
		WholesaleOrganisationCode: "org_1",
		OrganisationAccessID:      "access_1",
		RoleKey:                   "buyer",
	}
	session := identity.LoginSession{
		ID:                        "session_1",
		UserID:                    "user_1",
		Portal:                    identity_enums.PortalWholesale,
		WholesaleOrganisationCode: "org_1",
		OrganisationAccessID:      "access_1",
		RoleKey:                   "buyer",
		IssuedAt:                  time.Date(2026, 6, 23, 0, 0, 0, 0, time.UTC),
		LastSeenAt:                time.Date(2026, 6, 23, 0, 1, 0, 0, time.UTC),
		ExpiresAt:                 time.Date(2026, 6, 23, 1, 0, 0, 0, time.UTC),
	}

	payload, err := json.Marshal(struct {
		Claims  identity.AccessTokenClaims `json:"claims"`
		Session identity.LoginSession      `json:"session"`
	}{Claims: claims, Session: session})
	if err != nil {
		t.Fatalf("marshal identity wholesale access payload: %v", err)
	}

	var got map[string]map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal identity wholesale access JSON: %v", err)
	}
	for group, fields := range got {
		if _, ok := fields["organisation_access_id"]; !ok {
			t.Fatalf("%s missing organisation_access_id: %s", group, payload)
		}
		if _, ok := fields["membership_id"]; ok {
			t.Fatalf("%s should not include legacy membership_id: %s", group, payload)
		}
	}
	if got["claims"]["session_id"] != "session_1" {
		t.Fatalf("claims missing session_id: %s", payload)
	}
}

func TestUserDeviceExactLastLoginIPRoundTrip(t *testing.T) {
	now := time.Date(2026, 7, 17, 1, 2, 3, 0, time.UTC)
	device := accountenum.UserDevice{
		ID:          "device_1",
		UserID:      "user_1",
		FirstSeenAt: now,
		LastSeenAt:  now,
		LastLoginAt: &now,
		LastLoginIP: "203.0.113.7",
	}
	payload, err := json.Marshal(device)
	if err != nil {
		t.Fatalf("marshal user device: %v", err)
	}
	var decoded accountenum.UserDevice
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal user device: %v", err)
	}
	if decoded.LastLoginIP != "203.0.113.7" || decoded.LastLoginAt == nil || !decoded.LastLoginAt.Equal(now) {
		t.Fatalf("last login fields did not round-trip: %+v", decoded)
	}
}

func TestOptionalSessionAndLastLoginIPFieldsOmitZeroValues(t *testing.T) {
	payload, err := json.Marshal(struct {
		Claims identity.AccessTokenClaims `json:"claims"`
		Device accountenum.UserDevice     `json:"device"`
	}{})
	if err != nil {
		t.Fatalf("marshal zero-value identity projections: %v", err)
	}
	for _, key := range []string{`"session_id"`, `"last_login_ip"`} {
		if strings.Contains(string(payload), key) {
			t.Fatalf("zero-value optional field %s must be omitted: %s", key, payload)
		}
	}
}
