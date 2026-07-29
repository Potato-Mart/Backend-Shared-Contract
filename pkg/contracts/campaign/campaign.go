// Package campaign defines the Campaign content model: admin-authored
// banners, announcements, modals and notices surfaced across every storefront
// surface (3 web + 2 mobile clients). A campaign is pure presentational
// content with scheduling and audience targeting; it carries no pricing logic.
// PromotionID links optional Pricing-owned promotion content without embedding
// any discount rules.
package campaign

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/campaign"
)

// Audience narrows who a campaign is shown to. An empty field means "any";
// clients pass their own context (customer_type/platform/region) when querying
// and the server filters server-side.
type Audience struct {
	CustomerType campaignenum.CampaignCustomerType `json:"customer_type,omitempty"`
	Platform     campaignenum.CampaignPlatform     `json:"platform,omitempty"`
	Region       string                            `json:"region,omitempty"`
}

// CTADestination is the parsed, allowlisted route for a campaign call to
// action. Services validate that only fields relevant to Type are populated.
// CTAHref remains the canonical authored href for existing storefront clients.
type CTADestination struct {
	Type           campaignenum.CampaignCTADestinationType `json:"type"`
	ProductSKUCode string                                  `json:"product_sku_code,omitempty"`
	CollectionSlug string                                  `json:"collection_slug,omitempty"`
	CategorySlug   string                                  `json:"category_slug,omitempty"`
	PromotionID    string                                  `json:"promotion_id,omitempty"`
}

// Campaign is one piece of scheduled, targeted storefront content.
type Campaign struct {
	ID          string          `json:"id,omitempty"`
	CampaignKey string          `json:"campaign_key"`
	SeriesKey   string          `json:"series_key,omitempty"`
	PromotionID string          `json:"promotion_id,omitempty"`
	Title       string          `json:"title"`
	Message     string          `json:"message,omitempty"`
	CTAText     string          `json:"cta_text,omitempty"`
	CTAHref     string          `json:"cta_href,omitempty"`
	CTA         *CTADestination `json:"cta,omitempty"`

	// MediaID is the managed media identity. MediaURL is the stable backend
	// content URL projected for storefronts (used by home_hero and modal).
	// BackgroundToken is an optional theme token name for banner styling.
	MediaID         string `json:"media_id,omitempty"`
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

	Audience           *Audience                   `json:"audience,omitempty"`
	Targets            CampaignTarget              `json:"targets,omitempty"`
	Planning           *CampaignPlanning           `json:"planning,omitempty"`
	Status             campaignenum.CampaignStatus `json:"status"`
	Revision           int64                       `json:"revision"`
	ActivationRevision int64                       `json:"activation_revision"`
	ActivatedAt        *time.Time                  `json:"activated_at,omitempty"`
	DeactivatedAt      *time.Time                  `json:"deactivated_at,omitempty"`
	ArchivedAt         *time.Time                  `json:"archived_at,omitempty"`

	common.AuditFields
}

type CampaignCategoryTarget struct {
	CollectionSlug string `json:"collection_slug"`
	CategorySlug   string `json:"category_slug"`
}

type CampaignTarget struct {
	ProductSKUCodes []string                 `json:"product_sku_codes,omitempty"`
	Categories      []CampaignCategoryTarget `json:"categories,omitempty"`
}

type CampaignPlanning struct {
	ResolvedProductSKUCodes []string   `json:"resolved_product_sku_codes,omitempty"`
	PredictionKey           string     `json:"prediction_key,omitempty"`
	PredictionRevision      int        `json:"prediction_revision,omitempty"`
	AlgorithmVersion        string     `json:"algorithm_version,omitempty"`
	PredictedAt             *time.Time `json:"predicted_at,omitempty"`
}
