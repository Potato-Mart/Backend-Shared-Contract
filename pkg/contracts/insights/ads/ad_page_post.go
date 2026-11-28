package ads

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdPagePost is one of our own existing page or Instagram posts, offered as
// ready-made ad creative.
//
// Source is an open code naming where the post lives, for example a Facebook
// page or an Instagram account, because that list grows with the platform.
// ThumbnailURL is a short-lived platform preview link, so it is for display
// only and must never be stored as if it were durable. Excerpt is the opening
// of the post's own text, which a promoted post cannot change; only a call to
// action and a tracked link may be added. Eligible says whether the post can
// be promoted at all, and IneligibilityReason gives the code when it cannot;
// the reason is a code, never the platform's own wording. Likes, Comments and
// Shares are engagement counts at read time and carry no information about any
// individual.
type AdPagePost struct {
	PostID              string                              `json:"post_id,omitempty"`
	Platform            ads_enums.AdPlatform                `json:"platform,omitempty"`
	Source              string                              `json:"source,omitempty"`
	Type                ads_enums.AdPostType                `json:"type,omitempty"`
	PublishedAt         *time.Time                          `json:"published_at,omitempty"`
	Excerpt             string                              `json:"excerpt,omitempty"`
	ThumbnailURL        string                              `json:"thumbnail_url,omitempty"`
	Likes               int64                               `json:"likes,omitempty"`
	Comments            int64                               `json:"comments,omitempty"`
	Shares              int64                               `json:"shares,omitempty"`
	Eligible            bool                                `json:"eligible,omitempty"`
	IneligibilityReason ads_enums.AdPostIneligibilityReason `json:"ineligibility_reason,omitempty"`
}
