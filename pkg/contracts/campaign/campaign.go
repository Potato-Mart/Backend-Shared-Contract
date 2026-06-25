// Package campaign defines the Campaign content model: admin-authored
// banners, announcements, modals and notices surfaced across every storefront
// surface (3 web + 2 mobile clients). A campaign is pure presentational
// content with scheduling and audience targeting; it carries no pricing logic
// (that is promotion's job — a campaign may LINK to a promotion via CTAHref).
package campaign

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v9/pkg/enums"
)

// Audience narrows who a campaign is shown to. An empty field means "any";
// clients pass their own context (customer_type/platform/region) when querying
// and the server filters server-side.
type Audience struct {
	CustomerType enums.CampaignCustomerType `json:"customer_type,omitempty" bson:"customer_type,omitempty"`
	Platform     enums.CampaignPlatform     `json:"platform,omitempty" bson:"platform,omitempty"`
	Region       string                     `json:"region,omitempty" bson:"region,omitempty"`
}

// Campaign is one piece of scheduled, targeted storefront content.
type Campaign struct {
	ID      string `json:"id" bson:"_id"`
	Title   string `json:"title" bson:"title"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	CTAText string `json:"cta_text,omitempty" bson:"cta_text,omitempty"`
	CTAHref string `json:"cta_href,omitempty" bson:"cta_href,omitempty"`

	// MediaURL is the slide/hero image (used by home_hero, modal). BackgroundToken
	// is an optional theme token name for the banner background.
	MediaURL        string `json:"media_url,omitempty" bson:"media_url,omitempty"`
	BackgroundToken string `json:"background_token,omitempty" bson:"background_token,omitempty"`

	Placement enums.CampaignPlacement `json:"placement" bson:"placement"`
	Severity  enums.CampaignSeverity  `json:"severity" bson:"severity"`

	// Priority orders competing campaigns for the same placement (higher first).
	Priority int `json:"priority" bson:"priority"`

	IsActive    bool   `json:"is_active" bson:"is_active"`
	Dismissible bool   `json:"dismissible" bson:"dismissible"`
	DismissKey  string `json:"dismiss_key,omitempty" bson:"dismiss_key,omitempty"`

	// StartsAt/EndsAt bound the active window; nil means open-ended.
	StartsAt *time.Time `json:"starts_at,omitempty" bson:"starts_at,omitempty"`
	EndsAt   *time.Time `json:"ends_at,omitempty" bson:"ends_at,omitempty"`

	Audience *Audience `json:"audience,omitempty" bson:"audience,omitempty"`

	// TargetScope scopes a product_notice to a product/category key.
	TargetScope string `json:"target_scope,omitempty" bson:"target_scope,omitempty"`

	common.AuditFields `bson:",inline"`
}
