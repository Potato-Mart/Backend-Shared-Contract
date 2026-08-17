// Package pos_enums holds typed enums for in-store point-of-sale operations.
package pos_enums

// SessionStatus is the lifecycle state of a register's daily trading session.
type SessionStatus string

const (
	SessionStatusOpen   SessionStatus = "open"
	SessionStatusClosed SessionStatus = "closed"
)

func (s SessionStatus) IsValid() bool {
	switch s {
	case SessionStatusOpen, SessionStatusClosed:
		return true
	default:
		return false
	}
}

func (s SessionStatus) String() string { return string(s) }

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
