package terminal_enums

type TerminalStatus string

const (
	TerminalStatusRegistered   TerminalStatus = "registered"
	TerminalStatusActive       TerminalStatus = "active"
	TerminalStatusDeregistered TerminalStatus = "deregistered"
	TerminalStatusExpired      TerminalStatus = "expired"
	TerminalStatusError        TerminalStatus = "error"
)

func (s TerminalStatus) IsValid() bool {
	switch s {
	case TerminalStatusRegistered, TerminalStatusActive, TerminalStatusDeregistered,
		TerminalStatusExpired, TerminalStatusError:
		return true
	default:
		return false
	}
}

func (s TerminalStatus) String() string { return string(s) }
