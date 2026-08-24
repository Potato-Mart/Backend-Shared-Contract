package warehouse_enums

// InventoryCondition describes the physical condition of inventory.
type InventoryCondition string

const (
	InventoryConditionGood                  InventoryCondition = "GOOD"
	InventoryConditionPackagingDamagedMinor InventoryCondition = "PACKAGING_DAMAGED_MINOR"
	InventoryConditionPackagingDamagedMajor InventoryCondition = "PACKAGING_DAMAGED_MAJOR"
	InventoryConditionProductDamaged        InventoryCondition = "PRODUCT_DAMAGED"
	InventoryConditionOpened                InventoryCondition = "OPENED"
	InventoryConditionContaminated          InventoryCondition = "CONTAMINATED"
	InventoryConditionSpoiled               InventoryCondition = "SPOILED"
)

func (c InventoryCondition) IsValid() bool {
	switch c {
	case InventoryConditionGood, InventoryConditionPackagingDamagedMinor,
		InventoryConditionPackagingDamagedMajor, InventoryConditionProductDamaged,
		InventoryConditionOpened, InventoryConditionContaminated,
		InventoryConditionSpoiled:
		return true
	}
	return false
}

func (c InventoryCondition) String() string { return string(c) }
