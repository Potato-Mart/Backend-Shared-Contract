package wallet_enums

// PointAwardStatus is the durable processing state of a wallet points award.
type PointAwardStatus string

const (
	PointAwardStatusIneligible PointAwardStatus = "ineligible"
	PointAwardStatusDisabled   PointAwardStatus = "disabled"
	PointAwardStatusPending    PointAwardStatus = "pending"
	PointAwardStatusAwarded    PointAwardStatus = "awarded"
	PointAwardStatusFailed     PointAwardStatus = "failed"
)

func (s PointAwardStatus) IsValid() bool {
	switch s {
	case PointAwardStatusIneligible, PointAwardStatusDisabled, PointAwardStatusPending, PointAwardStatusAwarded, PointAwardStatusFailed:
		return true
	}
	return false
}
func (s PointAwardStatus) String() string { return string(s) }
