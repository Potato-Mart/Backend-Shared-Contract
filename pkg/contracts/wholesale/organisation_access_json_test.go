package wholesale_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/wholesale"
	wholesaleenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/wholesale"
)

func TestOrganisationAccessJSONGroupsLifecycleAndKeepsCoreFieldsTopLevel(t *testing.T) {
	invitedAt := time.Date(2026, 6, 18, 5, 15, 0, 0, time.UTC)
	joinedAt := time.Date(2026, 6, 18, 5, 30, 0, 0, time.UTC)
	access := wholesale.OrganisationAccess{
		ID:                        "access_1",
		WholesaleOrganisationCode: "org_1",
		UserID:                    "user_1",
		AccountID:                 "acct_1",
		RoleKey:                   "buyer",
		Status:                    wholesaleenum.OrganisationAccessStatusActive,
		Invitation:                &common.LifecycleAction{By: "admin_1", At: &invitedAt},
		JoinedAt:                  &joinedAt,
	}

	payload, err := json.Marshal(access)
	if err != nil {
		t.Fatalf("marshal organisation access: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal organisation access JSON: %v", err)
	}

	for _, key := range []string{"id", "wholesale_organisation_code", "user_id", "account_id", "role_key", "status", "joined_at"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("OrganisationAccess JSON missing top-level %q: %s", key, payload)
		}
	}
	if _, ok := got["invitation"]; !ok {
		t.Fatalf("OrganisationAccess JSON missing nested invitation: %s", payload)
	}
	for _, key := range []string{"membership_id", "invited_by", "invited_at", "revoked_by", "revoked_at"} {
		if _, ok := got[key]; ok {
			t.Fatalf("OrganisationAccess JSON should not include legacy or flat key %q: %s", key, payload)
		}
	}

	var decoded wholesale.OrganisationAccess
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal organisation access: %v", err)
	}
	if decoded.Invitation == nil || decoded.Invitation.By != "admin_1" || decoded.JoinedAt == nil || !decoded.JoinedAt.Equal(joinedAt) {
		t.Fatalf("organisation access lifecycle did not round-trip: %+v", decoded)
	}
}
