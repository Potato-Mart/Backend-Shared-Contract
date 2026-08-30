package security

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/metadata"

// AccessLogEntry records read/list/search/export access to protected data.
// Administrative writes should continue to use AuditLogEntry.
type AccessLogEntry struct {
	SecurityLogEntryFields

	// Action is commonly a dotted key such as "customer.read" or
	// "order.export"; Resource may be "customer:cust_123".
	RecordCount int               `json:"record_count,omitempty"`
	Metadata    metadata.Metadata `json:"metadata,omitempty"`
}
