package account_enums

// DeviceType is the coarse class of client device observed for a login
// session or device record.
type DeviceType string

const (
	DeviceTypeDesktop DeviceType = "desktop"
	DeviceTypeMobile  DeviceType = "mobile"
	DeviceTypeTablet  DeviceType = "tablet"
	DeviceTypeAPI     DeviceType = "api"
)

// IsValid reports whether d is a known DeviceType.
func (d DeviceType) IsValid() bool {
	switch d {
	case DeviceTypeDesktop, DeviceTypeMobile, DeviceTypeTablet, DeviceTypeAPI:
		return true
	}
	return false
}

func (d DeviceType) String() string { return string(d) }
