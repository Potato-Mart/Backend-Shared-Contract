package deletion

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/identity/deletion/deletion_enums"
)

// CommandEnvelope is the immutable correlation and idempotency envelope
// shared by every account-deletion service command. RequestID is an opaque
// workflow identifier, while UserID is the canonical internal subject
// identifier needed by the receiving service to locate its subject-bound
// records. Neither value is a MongoDB ObjectID or a customer-facing profile
// field.
//
// A service must persist and replay its first terminal outcome for the tuple
// (request_id, operation, idempotency_key). PolicyVersion and SnapshotVersion
// let the coordinator reject a response evaluated against a stale policy or
// stored-value/eligibility snapshot.
type CommandEnvelope struct {
	RequestID       string `json:"request_id"`
	UserID          string `json:"user_id"`
	PolicyVersion   string `json:"policy_version"`
	SnapshotVersion string `json:"snapshot_version"`
	IdempotencyKey  string `json:"idempotency_key"`
}

// ServiceOperation is the common, transport-neutral operation sent from the
// Identity coordinator to one participating service. The Operation chooses
// one of the four account-deletion commands defined by the workflow:
// restrict, eligibility, erase_or_deidentify, or unrestrict.
//
// This is intentionally not an endpoint request DTO. The owning service may
// expose the command through a private RPC, a queue consumer, or another
// authenticated internal transport.
type ServiceOperation struct {
	Envelope  CommandEnvelope                         `json:"envelope"`
	Operation deletion_enums.AccountDeletionOperation `json:"operation"`
}

// CommandReceipt is the minimized, durable result returned by one
// participating service. It echoes the entire command envelope so retries can
// be correlated without relying on a provider reference, database ID, or
// mutable service-local revision.
//
// Blockers contains code-only reasons. It must never contain order numbers,
// payment identifiers, recipient data, raw provider errors, or free-form
// details. Disposition is populated by erase_or_deidentify and is a summary
// only; it never names a collection or record.
type CommandReceipt struct {
	Envelope    CommandEnvelope                           `json:"envelope"`
	Participant deletion_enums.AccountDeletionParticipant `json:"participant"`
	Operation   deletion_enums.AccountDeletionOperation   `json:"operation"`
	Status      deletion_enums.CommandReceiptStatus       `json:"status"`
	Blockers    []deletion_enums.AccountDeletionBlocker   `json:"blockers,omitempty"`
	Disposition deletion_enums.ErasureDisposition         `json:"disposition,omitempty"`
	Retryable   bool                                      `json:"retryable"`
	// RetryAfterSeconds is present only when the service asks the coordinator
	// to defer a retry. It is not an obligation countdown or a provider ID.
	RetryAfterSeconds *int      `json:"retry_after_seconds,omitempty"`
	Replayed          bool      `json:"replayed"`
	RecordedAt        time.Time `json:"recorded_at"`
}

// WorkflowStatus is the coordinator's minimal, internal snapshot of the
// deletion workflow. It has no administrative actor, customer contact detail,
// financial record, or service-private data. Consumers can correlate it to a
// CommandReceipt by request_id and the receipt's idempotency key.
type WorkflowStatus struct {
	RequestID       string                              `json:"request_id"`
	UserID          string                              `json:"user_id"`
	PolicyVersion   string                              `json:"policy_version"`
	SnapshotVersion string                              `json:"snapshot_version"`
	State           deletion_enums.AccountDeletionState `json:"state"`
	UpdatedAt       time.Time                           `json:"updated_at"`
}
