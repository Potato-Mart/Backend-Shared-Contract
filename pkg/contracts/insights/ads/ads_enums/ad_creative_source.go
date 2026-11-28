package ads_enums

// AdCreativeSource says where the visual content of a draft ad comes from.
// Newly uploaded creative is images only in the first phase, while an existing
// post can be promoted as it stands, video posts included.
type AdCreativeSource string

const (
	// AdCreativeSourceMedia builds the ad from media-library images, named
	// by their media codes.
	AdCreativeSourceMedia AdCreativeSource = "media"
	// AdCreativeSourcePagePost promotes an existing page or Instagram post,
	// named by its post id. Only the call to action and the link may be
	// added.
	AdCreativeSourcePagePost AdCreativeSource = "page_post"
)

// IsValid reports whether s is a known AdCreativeSource.
func (s AdCreativeSource) IsValid() bool {
	switch s {
	case AdCreativeSourceMedia, AdCreativeSourcePagePost:
		return true
	}
	return false
}

func (s AdCreativeSource) String() string { return string(s) }
