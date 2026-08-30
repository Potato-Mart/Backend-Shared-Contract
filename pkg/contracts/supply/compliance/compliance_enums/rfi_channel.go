package compliance_enums

type RFIChannel string

const (
	RFIChannelBiosecurityPortal RFIChannel = "biosecurity_portal"
	RFIChannelEmailException    RFIChannel = "email_exception"
)

func (c RFIChannel) IsValid() bool {
	switch c {
	case RFIChannelBiosecurityPortal, RFIChannelEmailException:
		return true
	}
	return false
}

func (c RFIChannel) String() string { return string(c) }
