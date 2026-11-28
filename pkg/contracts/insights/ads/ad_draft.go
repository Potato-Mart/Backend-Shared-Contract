package ads

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdDraft is an ad built in our admin: where it will run, what it shows, what
// it says in each language, and where it sends people.
//
// A submitted draft is always created paused on the platform. Resuming it is a
// separate delivery change with its own stronger permission, so nothing built
// here starts spending on its own. Status follows the platform review:
// approved means approved and still paused, and a rejected draft carries
// RejectionReasonCodes, which are our own codes rather than platform wording.
// ValidationCodes are the problems found before submission, for example an
// image ratio or a landing page that cannot be reached. PlatformAdID is filled
// once the platform has created the ad.
type AdDraft struct {
	DraftID              string                   `json:"draft_id,omitempty"`
	Platform             ads_enums.AdPlatform     `json:"platform,omitempty"`
	Status               ads_enums.AdDraftStatus  `json:"status,omitempty"`
	ReviewStatus         ads_enums.AdReviewStatus `json:"review_status,omitempty"`
	Placement            *AdDraftPlacement        `json:"placement,omitempty"`
	Creative             *AdCreative              `json:"creative,omitempty"`
	Copy                 []AdCopy                 `json:"copy,omitempty"`
	Landing              *AdLanding               `json:"landing,omitempty"`
	ValidationCodes      []string                 `json:"validation_codes,omitempty"`
	RejectionReasonCodes []string                 `json:"rejection_reason_codes,omitempty"`
	PlatformAdID         string                   `json:"platform_ad_id,omitempty"`
	MarketCode           string                   `json:"market_code,omitempty"`
	CreatedBy            string                   `json:"created_by,omitempty"`
	CreatedAt            *time.Time               `json:"created_at,omitempty"`
	UpdatedAt            *time.Time               `json:"updated_at,omitempty"`
	SubmittedAt          *time.Time               `json:"submitted_at,omitempty"`
}
