package import_compliance_enums

type RFISubmissionState string

const (
	RFISubmissionStateNotSubmitted RFISubmissionState = "not_submitted"
	RFISubmissionStateSubmitted    RFISubmissionState = "submitted"
	RFISubmissionStateConfirmed    RFISubmissionState = "confirmed"
	RFISubmissionStateRescheduled  RFISubmissionState = "rescheduled"
	RFISubmissionStateCancelled    RFISubmissionState = "cancelled"
	RFISubmissionStateClosed       RFISubmissionState = "closed"
)

func (s RFISubmissionState) IsValid() bool {
	switch s {
	case RFISubmissionStateNotSubmitted, RFISubmissionStateSubmitted,
		RFISubmissionStateConfirmed, RFISubmissionStateRescheduled,
		RFISubmissionStateCancelled, RFISubmissionStateClosed:
		return true
	}
	return false
}

func (s RFISubmissionState) String() string { return string(s) }
