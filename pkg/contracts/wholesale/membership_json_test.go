package wholesale_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/contracts/wholesale"
	"github.com/Potato-Mart/Backend-Shared-Contract/v6/pkg/enums"
)

func TestWholesaleMembershipJSONGroupsLifecycleAndKeepsCoreFieldsTopLevel(t *testing.T) {
	invitedAt := time.Date(2026, 6, 18, 5, 15, 0, 0, time.UTC)
	joinedAt := time.Date(2026, 6, 18, 5, 30, 0, 0, time.UTC)
	membership := wholesale.WholesaleMembership{
		ID:                      "member_1",
		WholesaleOrganisationID: "org_1",
		UserID:                  "user_1",
		AccountID:               "acct_1",
		RoleKey:                 "buyer",
		Status:                  enums.WholesaleMembershipStatusActive,
		Invitation:              &common.LifecycleAction{By: "admin_1", At: &invitedAt},
		JoinedAt:                &joinedAt,
	}

	payload, err := json.Marshal(membership)
	if err != nil {
		t.Fatalf("marshal wholesale membership: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal wholesale membership JSON: %v", err)
	}

	for _, key := range []string{"id", "wholesale_organisation_id", "user_id", "account_id", "role_key", "status", "joined_at"} {
		if _, ok := got[key]; !ok {
			t.Fatalf("WholesaleMembership JSON missing top-level %q: %s", key, payload)
		}
	}
	if _, ok := got["invitation"]; !ok {
		t.Fatalf("WholesaleMembership JSON missing nested invitation: %s", payload)
	}
	for _, key := range []string{"invited_by", "invited_at", "revoked_by", "revoked_at"} {
		if _, ok := got[key]; ok {
			t.Fatalf("WholesaleMembership JSON should not include flat lifecycle key %q: %s", key, payload)
		}
	}

	var decoded wholesale.WholesaleMembership
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal wholesale membership: %v", err)
	}
	if decoded.Invitation == nil || decoded.Invitation.By != "admin_1" || decoded.JoinedAt == nil || !decoded.JoinedAt.Equal(joinedAt) {
		t.Fatalf("membership lifecycle did not round-trip: %+v", decoded)
	}
}
