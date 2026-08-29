package compliance_enums

type RFIRequestedTime string

const (
	RFIRequestedTimeAnytime RFIRequestedTime = "anytime"
	RFIRequestedTimeAM      RFIRequestedTime = "am"
	RFIRequestedTimePM      RFIRequestedTime = "pm"
)

func (t RFIRequestedTime) IsValid() bool {
	switch t {
	case RFIRequestedTimeAnytime, RFIRequestedTimeAM, RFIRequestedTimePM:
		return true
	}
	return false
}

func (t RFIRequestedTime) String() string { return string(t) }
