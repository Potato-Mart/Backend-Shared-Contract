package wallet_enums

// ExternalRewardFulfilmentStatus is the partner-side provisioning state of an
// EXTERNAL reward redemption.
type ExternalRewardFulfilmentStatus string

const (
	ExternalRewardFulfilmentStatusPending     ExternalRewardFulfilmentStatus = "PENDING"
	ExternalRewardFulfilmentStatusProvisioned ExternalRewardFulfilmentStatus = "PROVISIONED"
	ExternalRewardFulfilmentStatusFailed      ExternalRewardFulfilmentStatus = "FAILED"
	ExternalRewardFulfilmentStatusRevoked     ExternalRewardFulfilmentStatus = "REVOKED"
)

func (s ExternalRewardFulfilmentStatus) IsValid() bool {
	switch s {
	case ExternalRewardFulfilmentStatusPending, ExternalRewardFulfilmentStatusProvisioned, ExternalRewardFulfilmentStatusFailed, ExternalRewardFulfilmentStatusRevoked:
		return true
	}
	return false
}
func (s ExternalRewardFulfilmentStatus) String() string { return string(s) }
