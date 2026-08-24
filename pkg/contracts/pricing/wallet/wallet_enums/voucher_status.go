package wallet_enums

// VoucherStatus is the lifecycle state of a single-use wallet voucher.
type VoucherStatus string

const (
	VoucherStatusIssued   VoucherStatus = "issued"
	VoucherStatusReserved VoucherStatus = "reserved"
	VoucherStatusRedeemed VoucherStatus = "redeemed"
	VoucherStatusExpired  VoucherStatus = "expired"
	VoucherStatusVoid     VoucherStatus = "void"
)

func (s VoucherStatus) IsValid() bool {
	switch s {
	case VoucherStatusIssued, VoucherStatusReserved, VoucherStatusRedeemed, VoucherStatusExpired, VoucherStatusVoid:
		return true
	}
	return false
}
func (s VoucherStatus) String() string { return string(s) }
