package deletion_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/deletion"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/identity/deletion/deletion_enums"
)

func TestServiceOperationAndReceiptUseAStableIdempotentEnvelope(t *testing.T) {
	retryAfter := 30
	recordedAt := time.Date(2026, 8, 14, 2, 3, 4, 0, time.UTC)
	envelope := deletion.CommandEnvelope{
		RequestID: "adr_01J4S6RK1BJ9Y3KSTN8YQ3PGVR", UserID: "usr_01J4S6RTV0T0WGWBQW2SXN3THS",
		PolicyVersion: "account-deletion-v28.0", SnapshotVersion: "snapshot_01J4S6S2", IdempotencyKey: "erase:pricing:01J4S6S2",
	}
	value := struct {
		Command deletion.ServiceOperation `json:"command"`
		Receipt deletion.CommandReceipt   `json:"receipt"`
	}{
		Command: deletion.ServiceOperation{Envelope: envelope, Operation: deletion_enums.AccountDeletionOperationEraseOrDeidentify},
		Receipt: deletion.CommandReceipt{
			Envelope: envelope, Participant: deletion_enums.AccountDeletionParticipantPricing,
			Operation:   deletion_enums.AccountDeletionOperationEraseOrDeidentify,
			Status:      deletion_enums.CommandReceiptStatusCompleted,
			Disposition: deletion_enums.ErasureDispositionMixed,
			Retryable:   false, RetryAfterSeconds: &retryAfter, Replayed: true, RecordedAt: recordedAt,
		},
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal deletion command and receipt: %v", err)
	}
	text := string(payload)
	for _, expected := range []string{
		`"operation":"erase_or_deidentify"`, `"request_id":"adr_01J4S6RK1BJ9Y3KSTN8YQ3PGVR"`,
		`"user_id":"usr_01J4S6RTV0T0WGWBQW2SXN3THS"`, `"policy_version":"account-deletion-v28.0"`,
		`"snapshot_version":"snapshot_01J4S6S2"`, `"idempotency_key":"erase:pricing:01J4S6S2"`,
		`"participant":"pricing"`, `"status":"completed"`, `"disposition":"mixed"`,
		`"retry_after_seconds":30`, `"replayed":true`,
	} {
		if !strings.Contains(text, expected) {
			t.Fatalf("deletion JSON missing %s: %s", expected, text)
		}
	}
	for _, forbidden := range []string{
		`"_id"`, `"mongo_id"`, `"account_id"`, `"email"`, `"phone"`, `"address"`,
		`"provider_id"`, `"provider_reference"`, `"order_number"`, `"payment_id"`, `"error_detail"`,
	} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("deletion JSON leaked %s: %s", forbidden, text)
		}
	}

	var decoded struct {
		Command deletion.ServiceOperation `json:"command"`
		Receipt deletion.CommandReceipt   `json:"receipt"`
	}
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal deletion command and receipt: %v", err)
	}
	if decoded.Command.Envelope != envelope || decoded.Receipt.Envelope != envelope || !decoded.Receipt.Replayed || decoded.Receipt.RecordedAt != recordedAt {
		t.Fatalf("deletion command/receipt did not round-trip: %+v", decoded)
	}
}

func TestWorkflowStatusIsMinimalAndDoesNotExposeCommandSecrets(t *testing.T) {
	status := deletion.WorkflowStatus{
		RequestID: "adr_01J4S6RK1BJ9Y3KSTN8YQ3PGVR", UserID: "usr_01J4S6RTV0T0WGWBQW2SXN3THS",
		PolicyVersion: "account-deletion-v28.0", SnapshotVersion: "snapshot_01J4S6S2",
		State:     deletion_enums.AccountDeletionStateActionRequired,
		UpdatedAt: time.Date(2026, 8, 14, 2, 3, 4, 0, time.UTC),
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatalf("marshal workflow status: %v", err)
	}
	text := string(payload)
	for _, expected := range []string{`"state":"action_required"`, `"request_id"`, `"user_id"`, `"snapshot_version"`} {
		if !strings.Contains(text, expected) {
			t.Fatalf("workflow status JSON missing %s: %s", expected, text)
		}
	}
	for _, forbidden := range []string{`"idempotency_key"`, `"email"`, `"phone"`, `"account_id"`, `"blockers"`} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("workflow status JSON leaked %s: %s", forbidden, text)
		}
	}
}

func TestBlockedReceiptUsesCodeOnlyBlockers(t *testing.T) {
	receipt := deletion.CommandReceipt{
		Envelope:    deletion.CommandEnvelope{RequestID: "adr_1", UserID: "usr_1", PolicyVersion: "v28", SnapshotVersion: "s1", IdempotencyKey: "eligibility:payments:1"},
		Participant: deletion_enums.AccountDeletionParticipantPayments,
		Operation:   deletion_enums.AccountDeletionOperationEligibility,
		Status:      deletion_enums.CommandReceiptStatusBlocked,
		Blockers: []deletion_enums.AccountDeletionBlocker{
			deletion_enums.AccountDeletionBlockerPayToProcessing,
			deletion_enums.AccountDeletionBlockerProtectedStoredValue,
		},
		RecordedAt: time.Date(2026, 8, 14, 2, 3, 4, 0, time.UTC),
	}
	payload, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("marshal blocked receipt: %v", err)
	}
	for _, expected := range []string{`"blockers":["payto_processing","protected_stored_value"]`, `"retryable":false`, `"replayed":false`} {
		if !strings.Contains(string(payload), expected) {
			t.Fatalf("blocked receipt JSON missing %s: %s", expected, payload)
		}
	}
}
