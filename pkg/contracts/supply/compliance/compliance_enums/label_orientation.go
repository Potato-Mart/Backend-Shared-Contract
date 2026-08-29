package compliance_enums

type LabelOrientation string

const (
	LabelOrientationPortrait  LabelOrientation = "portrait"
	LabelOrientationLandscape LabelOrientation = "landscape"
)

func (o LabelOrientation) IsValid() bool {
	switch o {
	case LabelOrientationPortrait, LabelOrientationLandscape:
		return true
	}
	return false
}

func (o LabelOrientation) String() string { return string(o) }
