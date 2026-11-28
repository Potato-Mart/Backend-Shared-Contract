package ads

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"

// AdCreative is the visual content of a draft ad, from exactly one source.
//
// When Source is the media library, MediaCodes names the images to upload to
// the platform and PostID is empty. When Source is an existing post, PostID
// names that post and MediaCodes is empty: the post's own media and text
// cannot be changed, so only a call to action and a tracked link are added.
// Cropping happens before an image reaches here, as a new media item, so the
// original is never altered.
type AdCreative struct {
	Source     ads_enums.AdCreativeSource `json:"source,omitempty"`
	MediaCodes []string                   `json:"media_codes,omitempty"`
	PostID     string                     `json:"post_id,omitempty"`
}
