package enums

type PackingDiscrepancyKind string

const (
	PackingDiscrepancyKindShortage   PackingDiscrepancyKind = "SHORTAGE"
	PackingDiscrepancyKindOverweight PackingDiscrepancyKind = "OVERWEIGHT"
)

// IsValid reports whether s is a known PackingDiscrepancyKind.
func (s PackingDiscrepancyKind) IsValid() bool {
	switch s {
	case PackingDiscrepancyKindShortage, PackingDiscrepancyKindOverweight:
		return true
	}
	return false
}

func (s PackingDiscrepancyKind) String() string { return string(s) }
