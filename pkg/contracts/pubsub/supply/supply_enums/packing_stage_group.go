package supply_enums

// PackingStageGroup is the storage-temperature group a packing station packs
// and labels as one stage. Ambient and chilled lines share a stage; frozen
// lines are packed separately.
type PackingStageGroup string

const (
	PackingStageGroupAmbientChilled PackingStageGroup = "AMBIENT_CHILLED"
	PackingStageGroupFrozen         PackingStageGroup = "FROZEN"
)

func (group PackingStageGroup) IsValid() bool {
	return group == PackingStageGroupAmbientChilled || group == PackingStageGroupFrozen
}

func (group PackingStageGroup) String() string { return string(group) }
