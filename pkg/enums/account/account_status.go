package accountenum

// AccountStatus describes the lifecycle state of an account/persona record.
type AccountStatus string

const (
	// AccountStatusPending means the account exists but is not yet active.
	AccountStatusPending AccountStatus = "pending"
	// AccountStatusActive means the account can be used subject to portal access.
	AccountStatusActive AccountStatus = "active"
	// AccountStatusSuspended means the account is temporarily blocked.
	AccountStatusSuspended AccountStatus = "suspended"
	// AccountStatusClosed means the account was intentionally closed.
	AccountStatusClosed AccountStatus = "closed"
	// AccountStatusDeleted means the account has been deleted or tombstoned.
	AccountStatusDeleted AccountStatus = "deleted"
)

// IsValid reports whether s is a known AccountStatus value.
func (s AccountStatus) IsValid() bool {
	switch s {
	case AccountStatusPending, AccountStatusActive, AccountStatusSuspended,
		AccountStatusClosed, AccountStatusDeleted:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s AccountStatus) String() string { return string(s) }
