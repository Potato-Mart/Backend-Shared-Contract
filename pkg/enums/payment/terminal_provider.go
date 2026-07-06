package paymentenum

// TerminalProvider identifies which payment-terminal integration backs a
// Terminal record. Different providers have different registration flows,
// transaction shapes, and authentication. Keeping this on the Terminal
// lets a single TerminalTransaction model serve every backend.
type TerminalProvider string

const (
	TerminalProviderMx51 TerminalProvider = "mx51"
)

// IsValid reports whether p is a known TerminalProvider.
func (p TerminalProvider) IsValid() bool {
	switch p {
	case TerminalProviderMx51:
		return true
	}
	return false
}

func (p TerminalProvider) String() string { return string(p) }
