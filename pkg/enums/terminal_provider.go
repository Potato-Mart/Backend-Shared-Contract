package enums

// TerminalProvider identifies which payment-terminal integration backs a
// Terminal record. Different providers have different pairing flows,
// transaction shapes, and authentication. Keeping this on the Terminal
// lets a single TerminalTransaction model serve every backend.
type TerminalProvider string

const (
	TerminalProviderMX51SCI   TerminalProvider = "mx51_sci"
	TerminalProviderMX51Spice TerminalProvider = "mx51_spice"
	TerminalProviderMX51SPI   TerminalProvider = "mx51_spi"
)

// IsValid reports whether p is a known TerminalProvider.
func (p TerminalProvider) IsValid() bool {
	switch p {
	case TerminalProviderMX51SCI, TerminalProviderMX51Spice, TerminalProviderMX51SPI:
		return true
	}
	return false
}

func (p TerminalProvider) String() string { return string(p) }
