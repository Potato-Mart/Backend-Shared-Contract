package enums

// TerminalStatus is the lifecycle state of a registered EFTPOS terminal.
//
// Registered -> Active is the success path. Deregistered and Expired are
// terminal states that require the merchant to register or reassign the
// device again. Error captures provider-side failures that are not a
// clean deregistration.
type TerminalStatus string

const (
	TerminalStatusRegistered   TerminalStatus = "registered"
	TerminalStatusActive       TerminalStatus = "active"
	TerminalStatusDeregistered TerminalStatus = "deregistered"
	TerminalStatusExpired      TerminalStatus = "expired"
	TerminalStatusError        TerminalStatus = "error"
)

// IsValid reports whether s is a known TerminalStatus.
func (s TerminalStatus) IsValid() bool {
	switch s {
	case TerminalStatusRegistered, TerminalStatusActive, TerminalStatusDeregistered,
		TerminalStatusExpired, TerminalStatusError:
		return true
	}
	return false
}

// IsTerminal reports whether the terminal is in an end-of-life state
// that requires re-registration.
func (s TerminalStatus) IsTerminal() bool {
	return s == TerminalStatusDeregistered || s == TerminalStatusExpired
}

func (s TerminalStatus) String() string { return string(s) }
