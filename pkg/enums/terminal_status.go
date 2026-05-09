package enums

// TerminalStatus is the lifecycle state of a paired EFTPOS terminal.
//
// Pairing → Active is the success path. Unpaired and Expired are
// terminal states that require the merchant to re-pair from the device.
// Error captures provider-side failures that aren't a clean unpair
// (for example, pairing_route_forbidden from MX51).
type TerminalStatus string

const (
	TerminalStatusPairing  TerminalStatus = "pairing"
	TerminalStatusActive   TerminalStatus = "active"
	TerminalStatusUnpaired TerminalStatus = "unpaired"
	TerminalStatusExpired  TerminalStatus = "expired"
	TerminalStatusError    TerminalStatus = "error"
)

// IsValid reports whether s is a known TerminalStatus.
func (s TerminalStatus) IsValid() bool {
	switch s {
	case TerminalStatusPairing, TerminalStatusActive, TerminalStatusUnpaired,
		TerminalStatusExpired, TerminalStatusError:
		return true
	}
	return false
}

// IsTerminal reports whether the terminal is in an end-of-life state
// that requires re-pairing.
func (s TerminalStatus) IsTerminal() bool {
	return s == TerminalStatusUnpaired || s == TerminalStatusExpired
}

func (s TerminalStatus) String() string { return string(s) }
