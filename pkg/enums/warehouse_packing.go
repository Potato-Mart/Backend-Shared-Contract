package enums

type PackingDiscrepancyKind string

const (
	PackingDiscrepancyKindShortage   PackingDiscrepancyKind = "shortage"
	PackingDiscrepancyKindOverweight PackingDiscrepancyKind = "overweight"
	PackingDiscrepancyKindDamaged    PackingDiscrepancyKind = "damaged"
)

// IsValid reports whether s is a known PackingDiscrepancyKind.
func (s PackingDiscrepancyKind) IsValid() bool {
	switch s {
	case PackingDiscrepancyKindShortage, PackingDiscrepancyKindOverweight, PackingDiscrepancyKindDamaged:
		return true
	}
	return false
}

func (s PackingDiscrepancyKind) String() string { return string(s) }
