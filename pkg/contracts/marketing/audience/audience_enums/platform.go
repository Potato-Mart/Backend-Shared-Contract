package audience_enums

type Platform string

const (
	PlatformWeb    Platform = "web"
	PlatformMobile Platform = "mobile"
)

func (v Platform) IsValid() bool  { return v == PlatformWeb || v == PlatformMobile }
func (v Platform) String() string { return string(v) }
