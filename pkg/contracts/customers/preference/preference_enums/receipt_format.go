package preference_enums

// ReceiptFormat is one receipt delivery format a retail customer elects to
// receive. Electronic is the default election; paper is an explicit opt-in.
type ReceiptFormat string

const (
	ReceiptFormatElectronic ReceiptFormat = "electronic"
	ReceiptFormatPaper      ReceiptFormat = "paper"
)

// IsValid reports whether f is a known ReceiptFormat value.
func (f ReceiptFormat) IsValid() bool {
	switch f {
	case ReceiptFormatElectronic, ReceiptFormatPaper:
		return true
	}
	return false
}

func (f ReceiptFormat) String() string { return string(f) }
