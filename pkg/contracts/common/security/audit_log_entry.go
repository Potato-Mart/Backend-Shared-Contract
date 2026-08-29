package security

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/metadata"

// AuditLogEntry is one immutable record of an administrative action.
//
// Entries are written by middleware on every successful write request
// (HTTP method POST/PUT/PATCH/DELETE) plus explicit business events
// (login, role change, refund issued, etc.). Reads are never audited
// here – use access logs for that.
type AuditLogEntry struct {
	SecurityLogEntryFields
	// Diff carries arbitrary before/after fragments for an administrative
	// write; reads continue to use AccessLogEntry.
	Diff metadata.Metadata `json:"diff,omitempty"`
}
