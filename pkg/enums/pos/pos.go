// Package posenum holds typed enums for in-store point-of-sale operations.
package posenum

// ShiftStatus is the lifecycle state of a register shift.
type ShiftStatus string

const (
	ShiftStatusOpen   ShiftStatus = "open"
	ShiftStatusClosed ShiftStatus = "closed"
)

func (s ShiftStatus) IsValid() bool {
	switch s {
	case ShiftStatusOpen, ShiftStatusClosed:
		return true
	default:
		return false
	}
}

func (s ShiftStatus) String() string { return string(s) }

// CashMovementKind classifies a cash-drawer movement within a shift.
type CashMovementKind string

const (
	CashMovementKindCashIn          CashMovementKind = "cash_in"
	CashMovementKindCashOut         CashMovementKind = "cash_out"
	CashMovementKindFloatAdjustment CashMovementKind = "float_adjustment"
	CashMovementKindDrop            CashMovementKind = "drop"
)

func (k CashMovementKind) IsValid() bool {
	switch k {
	case CashMovementKindCashIn, CashMovementKindCashOut,
		CashMovementKindFloatAdjustment, CashMovementKindDrop:
		return true
	default:
		return false
	}
}

func (k CashMovementKind) String() string { return string(k) }
