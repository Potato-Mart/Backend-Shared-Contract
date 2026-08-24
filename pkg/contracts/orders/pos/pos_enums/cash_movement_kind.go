package pos_enums

// CashMovementKind classifies a cash-drawer movement within a session.
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
