package ads_enums

// AdChangeStatus is the lifecycle of one recorded delivery change. A change
// starts pending and reaches exactly one terminal state. The contract carries
// the state; applying the change against the platform is service-owned.
type AdChangeStatus string

const (
	// AdChangeStatusPending was accepted and is still being applied on the
	// platform.
	AdChangeStatusPending AdChangeStatus = "pending"
	// AdChangeStatusApplied was accepted by the platform.
	AdChangeStatusApplied AdChangeStatus = "applied"
	// AdChangeStatusFailed was refused or errored. The record carries our
	// own platform-independent error code.
	AdChangeStatusFailed AdChangeStatus = "failed"
	// AdChangeStatusConflict was abandoned because the object on the
	// platform no longer matched the expected revision, so nothing was
	// applied.
	AdChangeStatusConflict AdChangeStatus = "conflict"
)

// IsValid reports whether s is a known AdChangeStatus.
func (s AdChangeStatus) IsValid() bool {
	switch s {
	case AdChangeStatusPending, AdChangeStatusApplied, AdChangeStatusFailed, AdChangeStatusConflict:
		return true
	}
	return false
}

func (s AdChangeStatus) String() string { return string(s) }
