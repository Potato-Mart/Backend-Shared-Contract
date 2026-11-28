package ads_enums

// AdSourcePlatform names the client an ad click was captured on, and the
// client a paid order was placed from. Ads reporting can be read and grouped
// by it, and it decides which device rules apply to a server-side conversion.
type AdSourcePlatform string

const (
	// AdSourcePlatformWeb is the storefront website.
	AdSourcePlatformWeb AdSourcePlatform = "web"
	// AdSourcePlatformIOS is the iOS app, where app tracking authorisation
	// also applies.
	AdSourcePlatformIOS AdSourcePlatform = "ios"
	// AdSourcePlatformAndroid is the Android app.
	AdSourcePlatformAndroid AdSourcePlatform = "android"
)

// IsValid reports whether p is a known AdSourcePlatform.
func (p AdSourcePlatform) IsValid() bool {
	switch p {
	case AdSourcePlatformWeb, AdSourcePlatformIOS, AdSourcePlatformAndroid:
		return true
	}
	return false
}

func (p AdSourcePlatform) String() string { return string(p) }
