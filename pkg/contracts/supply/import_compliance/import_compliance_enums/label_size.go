package import_compliance_enums

type LabelSize string

const (
	LabelSize100x60 LabelSize = "100x60"
	LabelSize100x80 LabelSize = "100x80"
	LabelSize80x50  LabelSize = "80x50"
	LabelSize65x45  LabelSize = "65x45"
	LabelSize60x40  LabelSize = "60x40"
	LabelSize50x30  LabelSize = "50x30"
)

func (s LabelSize) IsValid() bool {
	switch s {
	case LabelSize100x60, LabelSize100x80, LabelSize80x50,
		LabelSize65x45, LabelSize60x40, LabelSize50x30:
		return true
	}
	return false
}

func (s LabelSize) String() string { return string(s) }
