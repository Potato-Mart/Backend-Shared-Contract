// Package campaign defines the Campaign content model: admin-authored
// banners, announcements, modals and notices surfaced across every storefront
// surface (3 web + 2 mobile clients). A campaign is pure presentational
// content with scheduling and audience targeting; it carries no pricing logic
// (that is promotion's job — a campaign may LINK to a promotion via CTAHref).
package campaign

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/common"
	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/enums/campaign"
)

// Audience narrows who a campaign is shown to. An empty field means "any";
// clients pass their own context (customer_type/platform/region) when querying
// and the server filters server-side.
type Audience struct {
	CustomerType campaignenum.CampaignCustomerType `json:"customer_type,omitempty"`
	Platform     campaignenum.CampaignPlatform     `json:"platform,omitempty"`
	Region       string                            `json:"region,omitempty"`
}

// Campaign is one piece of scheduled, targeted storefront content.
type Campaign struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message,omitempty"`
	CTAText string `json:"cta_text,omitempty"`
	CTAHref string `json:"cta_href,omitempty"`

	// MediaURL is the slide/hero image (used by home_hero, modal). BackgroundToken
	// is an optional theme token name for the banner background.
	MediaURL        string `json:"media_url,omitempty"`
	BackgroundToken string `json:"background_token,omitempty"`

	Placement campaignenum.CampaignPlacement `json:"placement"`
	Severity  campaignenum.CampaignSeverity  `json:"severity"`

	// Priority orders competing campaigns for the same placement (higher first).
	Priority int `json:"priority"`

	IsActive    bool   `json:"is_active"`
	Dismissible bool   `json:"dismissible"`
	DismissKey  string `json:"dismiss_key,omitempty"`

	// StartsAt/EndsAt bound the active window; nil means open-ended.
	StartsAt *time.Time `json:"starts_at,omitempty"`
	EndsAt   *time.Time `json:"ends_at,omitempty"`

	Audience *Audience `json:"audience,omitempty"`

	// TargetScope scopes a product_notice to a product or category tag.
	TargetScope string `json:"target_scope,omitempty"`

	common.AuditFields
}
