package ads_enums

// AdPlatform names the paid-advertising platform an ads record belongs to.
// Spend, performance and object state always come from exactly one platform;
// amounts held in different account currencies are never added together.
type AdPlatform string

const (
	// AdPlatformMeta is Meta: Facebook, Instagram, Messenger and the
	// Audience Network, reported through the Meta Marketing API.
	AdPlatformMeta AdPlatform = "meta"
	// AdPlatformGoogle is Google Ads, reported through the Google Ads API.
	AdPlatformGoogle AdPlatform = "google"
)

// IsValid reports whether p is a known AdPlatform.
func (p AdPlatform) IsValid() bool {
	switch p {
	case AdPlatformMeta, AdPlatformGoogle:
		return true
	}
	return false
}

func (p AdPlatform) String() string { return string(p) }
