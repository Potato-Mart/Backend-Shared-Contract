package order_enums

// OfflineCashierAuthKind distinguishes the evidence captured at the till.
type OfflineCashierAuthKind string

const (
	OfflineCashierAuthKindPIN             OfflineCashierAuthKind = "offline_pin"
	OfflineCashierAuthKindPasswordSession OfflineCashierAuthKind = "password_session"
)

func (k OfflineCashierAuthKind) IsValid() bool {
	return k == OfflineCashierAuthKindPIN || k == OfflineCashierAuthKindPasswordSession
}

func (k OfflineCashierAuthKind) String() string { return string(k) }
