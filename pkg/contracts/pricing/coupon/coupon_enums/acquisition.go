package coupon_enums

// Acquisition identifies the command family permitted to issue a coupon.
type Acquisition string

const (
	AcquisitionSelfClaim        Acquisition = "self_claim"
	AcquisitionAutomatic        Acquisition = "automatic"
	AcquisitionManualAssignment Acquisition = "manual_assignment"
	AcquisitionReferral         Acquisition = "referral"
)

func (v Acquisition) IsValid() bool {
	return v == AcquisitionSelfClaim || v == AcquisitionAutomatic || v == AcquisitionManualAssignment || v == AcquisitionReferral
}
func (v Acquisition) String() string { return string(v) }
