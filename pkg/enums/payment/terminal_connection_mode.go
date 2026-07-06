package paymentenum

// TerminalConnectionMode describes how the POS backend talks to the
// payment terminal.
type TerminalConnectionMode string

const (
	TerminalConnectionModeCloudSync  TerminalConnectionMode = "cloud_sync"
	TerminalConnectionModeCloudAsync TerminalConnectionMode = "cloud_async"
	TerminalConnectionModeLocal      TerminalConnectionMode = "local"
)

// IsValid reports whether m is a known TerminalConnectionMode.
func (m TerminalConnectionMode) IsValid() bool {
	switch m {
	case TerminalConnectionModeCloudSync, TerminalConnectionModeCloudAsync,
		TerminalConnectionModeLocal:
		return true
	}
	return false
}

func (m TerminalConnectionMode) String() string { return string(m) }
