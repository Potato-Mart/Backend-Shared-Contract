package warehouseenum

type DamageStage string

const (
	DamageStageInbound DamageStage = "inbound"
	DamageStagePicking DamageStage = "picking"
	DamageStagePacking DamageStage = "packing"
	DamageStageStorage DamageStage = "storage"
)

// IsValid reports whether s is a known DamageStage.
func (s DamageStage) IsValid() bool {
	switch s {
	case DamageStageInbound, DamageStagePicking, DamageStagePacking, DamageStageStorage:
		return true
	}
	return false
}

func (s DamageStage) String() string { return string(s) }
