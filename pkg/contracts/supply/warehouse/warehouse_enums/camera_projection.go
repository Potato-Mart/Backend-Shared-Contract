package warehouse_enums

// CameraProjection selects perspective or orthographic rendering for a
// saved camera preset.
type CameraProjection string

const (
	CameraPerspective  CameraProjection = "perspective"
	CameraOrthographic CameraProjection = "orthographic"
)

// IsValid reports whether p is a known CameraProjection.
func (p CameraProjection) IsValid() bool {
	switch p {
	case CameraPerspective, CameraOrthographic:
		return true
	}
	return false
}

func (p CameraProjection) String() string { return string(p) }
