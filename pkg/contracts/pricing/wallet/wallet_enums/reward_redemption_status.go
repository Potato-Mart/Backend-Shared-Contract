package wallet_enums

// RewardRedemptionStatus is the lifecycle state of a wallet reward redemption.
type RewardRedemptionStatus string

const (
	RewardRedemptionStatusReserved  RewardRedemptionStatus = "RESERVED"
	RewardRedemptionStatusRedeemed  RewardRedemptionStatus = "REDEEMED"
	RewardRedemptionStatusCancelled RewardRedemptionStatus = "CANCELLED"
	RewardRedemptionStatusExpired   RewardRedemptionStatus = "EXPIRED"
)

func (s RewardRedemptionStatus) IsValid() bool {
	switch s {
	case RewardRedemptionStatusReserved, RewardRedemptionStatusRedeemed, RewardRedemptionStatusCancelled, RewardRedemptionStatusExpired:
		return true
	}
	return false
}
func (s RewardRedemptionStatus) String() string { return string(s) }
