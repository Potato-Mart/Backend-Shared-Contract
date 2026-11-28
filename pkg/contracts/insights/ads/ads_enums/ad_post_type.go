package ads_enums

// AdPostType is the media shape of an existing page or Instagram post offered
// for promotion. Video and reel posts are promoted as they stand, so no media
// is re-uploaded for them.
type AdPostType string

const (
	// AdPostTypeImage is a single still image.
	AdPostTypeImage AdPostType = "image"
	// AdPostTypeCarousel is several images or videos in one swipeable post.
	AdPostTypeCarousel AdPostType = "carousel"
	// AdPostTypeVideo is a feed video.
	AdPostTypeVideo AdPostType = "video"
	// AdPostTypeReel is a short vertical video reel.
	AdPostTypeReel AdPostType = "reel"
)

// IsValid reports whether t is a known AdPostType.
func (t AdPostType) IsValid() bool {
	switch t {
	case AdPostTypeImage, AdPostTypeCarousel, AdPostTypeVideo, AdPostTypeReel:
		return true
	}
	return false
}

func (t AdPostType) String() string { return string(t) }
