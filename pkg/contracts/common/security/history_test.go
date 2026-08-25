package security_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/identity/role/role_enums"
	order "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order"
	terminal "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/payments/terminal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/payments/terminal/terminal_enums"
)

func TestHistoryOmittedWhenEmpty(t *testing.T) {
	payload, err := json.Marshal(order.Order{ID: "ord_1"})
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}

	if strings.Contains(string(payload), `"history"`) {
		t.Fatalf("empty history should be omitted, got %s", payload)
	}
}

func TestHistoryEntryRoundTrip(t *testing.T) {
	occurredAt := time.Date(2026, 6, 17, 10, 20, 0, 0, time.UTC)
	entry := security.HistoryEntry{
		ID:         "hist_1",
		Sequence:   7,
		OccurredAt: occurredAt,
		Type:       "status_change",
		Summary:    "Terminal transaction moved from unknown to override_pending",
		Changes: []security.HistoryChange{
			{Field: "status", FromValue: "unknown", ToValue: "override_pending"},
		},
		Source: "provider",
		ActorRef: security.ActorRef{
			ActorID:    "usr_1",
			ActorEmail: "ops@example.com",
			ActorRole:  role_enums.UserRoleDepotManager,
		},
		ReasonCode:        "terminal_timeout",
		Note:              "Provider status check could not confirm the outcome.",
		RiskLevel:         security_enums.SecurityRiskLevelHigh,
		RiskFlags:         []string{"payment_unknown", "manual_recovery_required"},
		RelatedResource:   "terminal_transaction",
		RelatedResourceID: "ttx_123",
		Metadata: metadata.Metadata{
			"provider_result": "timeout",
		},
	}

	tx := terminal.TerminalTransaction{
		ID:      "ttx_123",
		Status:  terminal_enums.TerminalTxStatusOverridePending,
		History: []security.HistoryEntry{entry},
	}

	payload, err := json.Marshal(tx)
	if err != nil {
		t.Fatalf("marshal terminal transaction: %v", err)
	}
	if !strings.Contains(string(payload), `"history"`) {
		t.Fatalf("history should be present, got %s", payload)
	}

	var decoded terminal.TerminalTransaction
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal terminal transaction: %v", err)
	}
	if len(decoded.History) != 1 {
		t.Fatalf("decoded history length = %d, want 1", len(decoded.History))
	}

	got := decoded.History[0]
	if !got.OccurredAt.Equal(occurredAt) {
		t.Fatalf("occurred_at = %s, want %s", got.OccurredAt, occurredAt)
	}
	if got.ActorEmail != "ops@example.com" {
		t.Fatalf("actor reference did not round-trip: %+v", got)
	}
	if got.RiskLevel != security_enums.SecurityRiskLevelHigh {
		t.Fatalf("risk_level = %q, want %q", got.RiskLevel, security_enums.SecurityRiskLevelHigh)
	}
	if got.Metadata["provider_result"] != "timeout" {
		t.Fatalf("metadata provider_result = %v, want timeout", got.Metadata["provider_result"])
	}
	if len(got.Changes) != 1 || got.Changes[0].Field != "status" || got.Changes[0].ToValue != "override_pending" {
		t.Fatalf("changes did not round-trip: %+v", got.Changes)
	}
}

func TestAuditDomainRecordsExcludeTransportCorrelationFields(t *testing.T) {
	models := []any{
		security.HistoryEntry{},
		security.SecurityEvent{},
		security.AuditLogEntry{},
		security.AccessLogEntry{},
	}
	fields := []string{
		"DeviceKey", "SessionID", "IPAddress", "UserAgent", "RequestID", "CorrelationID", "TraceID",
	}
	jsonKeys := []string{
		`"device_key"`, `"session_id"`, `"ip_address"`, `"user_agent"`, `"request_id"`, `"correlation_id"`, `"trace_id"`,
	}

	for _, model := range models {
		modelType := reflect.TypeOf(model)
		for _, field := range fields {
			if _, found := modelType.FieldByName(field); found {
				t.Errorf("%s retains transport/correlation field %s", modelType, field)
			}
		}

		payload, err := json.Marshal(model)
		if err != nil {
			t.Errorf("marshal %s: %v", modelType, err)
			continue
		}
		for _, key := range jsonKeys {
			if strings.Contains(string(payload), key) {
				t.Errorf("%s retains transport/correlation JSON key %s: %s", modelType, key, payload)
			}
		}
	}
}

func TestExistingSalesStatusHistoryRemainsAvailable(t *testing.T) {
	history := order.StatusHistory{
		FromValue: "pending",
		ToValue:   "paid",
		CreatedAt: time.Date(2026, 6, 17, 11, 0, 0, 0, time.UTC),
	}

	if history.ToValue != "paid" {
		t.Fatalf("sales.StatusHistory.ToValue = %q, want paid", history.ToValue)
	}
}
