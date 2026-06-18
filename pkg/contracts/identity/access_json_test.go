package identity_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/contracts/identity"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

func TestPortalAccessJSONGroupsLifecycleAndKeepsCoreFieldsTopLevel(t *testing.T) {
	grantedAt := time.Date(2026, 6, 18, 4, 30, 0, 0, time.UTC)
	access := identity.PortalAccess{
		ID:          "portal_access_1",
		UserID:      "user_1",
		AccountID:   "acct_1",
		AccountType: enums.AccountTypeGeneralCustomer,
		Portal:      enums.PortalStore,
		Status:      enums.PortalAccessStatusActive,
		Grant:       &common.LifecycleAction{By: "admin_1", At: &grantedAt, Reason: "approved"},
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

func TestIdentityRequestsJSONNestContext(t *testing.T) {
	request := identity.GrantRoleAssignmentRequest{
		UserID:    "user_1",
		AccountID: "acct_1",
		Portal:    enums.PortalControl,
		RoleKey:   "admin",
		Context: &identity.IdentityRequestContext{
			RequestedBy: "super_admin_1",
			RequestID:   "req_1",
			Reason:      "least privilege update",
		},
	}

	payload, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal role assignment request: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal role assignment request JSON: %v", err)
	}

	for _, key := range []string{"user_id", "account_id", "portal", "role_key"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("GrantRoleAssignmentRequest JSON missing top-level %q: %s", key, payload)
		}
	}
	if _, ok := got["context"]; !ok {
		t.Fatalf("GrantRoleAssignmentRequest JSON missing nested context: %s", payload)
	}
	for _, key := range []string{"requested_by", "request_id", "reason"} {
		if _, ok := got[key]; ok {
			t.Fatalf("GrantRoleAssignmentRequest JSON should not include flat context key %q: %s", key, payload)
		}
	}

	var decoded identity.GrantRoleAssignmentRequest
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal role assignment request: %v", err)
	}
	if decoded.Context == nil || decoded.Context.RequestID != "req_1" {
		t.Fatalf("request context did not round-trip: %+v", decoded.Context)
	}
}
