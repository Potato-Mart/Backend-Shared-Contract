package warehouseenum

// InventoryDateMarkKind identifies the meaning of a lot's date mark.
type InventoryDateMarkKind string

const (
	InventoryDateMarkBestBefore InventoryDateMarkKind = "BEST_BEFORE"
	InventoryDateMarkUseBy      InventoryDateMarkKind = "USE_BY"
	InventoryDateMarkExpiry     InventoryDateMarkKind = "EXPIRY"
)

func (k InventoryDateMarkKind) IsValid() bool {
	switch k {
	case InventoryDateMarkBestBefore, InventoryDateMarkUseBy, InventoryDateMarkExpiry:
		return true
	}
	return false
}

func (k InventoryDateMarkKind) String() string { return string(k) }

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

// InventoryDisposition describes how inventory may be used or sold.
type InventoryDisposition string

const (
	InventoryDispositionStandardSellable InventoryDisposition = "STANDARD_SELLABLE"
	InventoryDispositionReducedSellable  InventoryDisposition = "REDUCED_SELLABLE"
	InventoryDispositionQualityHold      InventoryDisposition = "QUALITY_HOLD"
	InventoryDispositionQuarantined      InventoryDisposition = "QUARANTINED"
	InventoryDispositionBlocked          InventoryDisposition = "BLOCKED"
	InventoryDispositionRecalled         InventoryDisposition = "RECALLED"
	InventoryDispositionReturnToVendor   InventoryDisposition = "RETURN_TO_VENDOR"
	InventoryDispositionScrap            InventoryDisposition = "SCRAP"
)

func (d InventoryDisposition) IsValid() bool {
	switch d {
	case InventoryDispositionStandardSellable, InventoryDispositionReducedSellable,
		InventoryDispositionQualityHold, InventoryDispositionQuarantined,
		InventoryDispositionBlocked, InventoryDispositionRecalled,
		InventoryDispositionReturnToVendor, InventoryDispositionScrap:
		return true
	}
	return false
}

func (d InventoryDisposition) String() string { return string(d) }

// StockReservationStatus describes the logical lifecycle of a reservation.
type StockReservationStatus string

const (
	StockReservationStatusReserved        StockReservationStatus = "RESERVED"
	StockReservationStatusPartiallyStaged StockReservationStatus = "PARTIALLY_STAGED"
	StockReservationStatusStaged          StockReservationStatus = "STAGED"
	StockReservationStatusCommitted       StockReservationStatus = "COMMITTED"
	StockReservationStatusReleased        StockReservationStatus = "RELEASED"
	StockReservationStatusExpired         StockReservationStatus = "EXPIRED"
	StockReservationStatusCancelled       StockReservationStatus = "CANCELLED"
)

func (s StockReservationStatus) IsValid() bool {
	switch s {
	case StockReservationStatusReserved, StockReservationStatusPartiallyStaged,
		StockReservationStatusStaged, StockReservationStatusCommitted,
		StockReservationStatusReleased, StockReservationStatusExpired,
		StockReservationStatusCancelled:
		return true
	}
	return false
}

func (s StockReservationStatus) String() string { return string(s) }

// StockAvailabilityDirection identifies a customer-visible zero crossing.
type StockAvailabilityDirection string

const (
	StockAvailabilityOutOfStock StockAvailabilityDirection = "OUT_OF_STOCK"
	StockAvailabilityRestocked  StockAvailabilityDirection = "RESTOCKED"
)

func (d StockAvailabilityDirection) IsValid() bool {
	switch d {
	case StockAvailabilityOutOfStock, StockAvailabilityRestocked:
		return true
	}
	return false
}

func (d StockAvailabilityDirection) String() string { return string(d) }

// InventoryDateMarkThreshold identifies a date-mark threshold crossing.
type InventoryDateMarkThreshold string

const (
	InventoryDateMarkThresholdApproaching InventoryDateMarkThreshold = "APPROACHING"
	InventoryDateMarkThresholdReached     InventoryDateMarkThreshold = "REACHED"
	InventoryDateMarkThresholdPassed      InventoryDateMarkThreshold = "PASSED"
)

func (t InventoryDateMarkThreshold) IsValid() bool {
	switch t {
	case InventoryDateMarkThresholdApproaching, InventoryDateMarkThresholdReached,
		InventoryDateMarkThresholdPassed:
		return true
	}
	return false
}

func (t InventoryDateMarkThreshold) String() string { return string(t) }
