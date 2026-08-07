package campaign

// CampaignPlacement is where a storefront campaign renders. One campaign
// targets exactly one placement; clients query the placement they render.
type CampaignPlacement string

const (
	// CampaignPlacementTopBanner is the thin full-width announcement row in the header.
	CampaignPlacementTopBanner CampaignPlacement = "top_banner"
	// CampaignPlacementHomeHero is the full-length home-page slideshow banner.
	CampaignPlacementHomeHero CampaignPlacement = "home_hero"
	// CampaignPlacementModal is a pop-out content modal on entry.
	CampaignPlacementModal CampaignPlacement = "modal"
	// CampaignPlacementCheckoutNotice is a notice shown on the checkout page.
	CampaignPlacementCheckoutNotice CampaignPlacement = "checkout_notice"
	// CampaignPlacementProductNotice is a notice shown on product detail.
	CampaignPlacementProductNotice CampaignPlacement = "product_notice"
)

// IsValid reports whether p is a known campaign placement.
func (p CampaignPlacement) IsValid() bool {
	switch p {
	case CampaignPlacementTopBanner, CampaignPlacementHomeHero, CampaignPlacementModal,
		CampaignPlacementCheckoutNotice, CampaignPlacementProductNotice:
		return true
	}
	return false
}

func (p CampaignPlacement) String() string { return string(p) }

// CampaignSeverity drives the visual tone of a campaign.
type CampaignSeverity string

const (
	CampaignSeverityInfo     CampaignSeverity = "info"
	CampaignSeveritySuccess  CampaignSeverity = "success"
	CampaignSeverityWarning  CampaignSeverity = "warning"
	CampaignSeverityCritical CampaignSeverity = "critical"
)

// IsValid reports whether s is a known campaign severity.
func (s CampaignSeverity) IsValid() bool {
	switch s {
	case CampaignSeverityInfo, CampaignSeveritySuccess, CampaignSeverityWarning, CampaignSeverityCritical:
		return true
	}
	return false
}

func (s CampaignSeverity) String() string { return string(s) }

// CampaignCustomerType targets a campaign at a buyer segment.
type CampaignCustomerType string

const (
	CampaignCustomerTypeGuest     CampaignCustomerType = "guest"
	CampaignCustomerTypeRetail    CampaignCustomerType = "retail"
	CampaignCustomerTypeWholesale CampaignCustomerType = "wholesale"
)

// IsValid reports whether c is a known campaign customer type.
func (c CampaignCustomerType) IsValid() bool {
	switch c {
	case CampaignCustomerTypeGuest, CampaignCustomerTypeRetail, CampaignCustomerTypeWholesale:
		return true
	}
	return false
}

func (c CampaignCustomerType) String() string { return string(c) }

// CampaignPlatform targets a campaign at a client family.
type CampaignPlatform string

const (
	CampaignPlatformWeb    CampaignPlatform = "web"
	CampaignPlatformMobile CampaignPlatform = "mobile"
)

// IsValid reports whether p is a known campaign platform.
func (p CampaignPlatform) IsValid() bool {
	switch p {
	case CampaignPlatformWeb, CampaignPlatformMobile:
		return true
	}
	return false
}

func (p CampaignPlatform) String() string { return string(p) }

type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusActive    CampaignStatus = "active"
	CampaignStatusCompleted CampaignStatus = "completed"
	CampaignStatusArchived  CampaignStatus = "archived"
)

func (s CampaignStatus) IsValid() bool {
	switch s {
	case CampaignStatusDraft, CampaignStatusScheduled, CampaignStatusActive,
		CampaignStatusCompleted, CampaignStatusArchived:
		return true
	default:
		return false
	}
}

func (s CampaignStatus) String() string { return string(s) }

type CampaignPredictionStatus string

const (
	CampaignPredictionStatusNotApplicable CampaignPredictionStatus = "not_applicable"
	CampaignPredictionStatusReady         CampaignPredictionStatus = "ready"
	CampaignPredictionStatusWarning       CampaignPredictionStatus = "warning"
)

func (s CampaignPredictionStatus) IsValid() bool {
	switch s {
	case CampaignPredictionStatusNotApplicable, CampaignPredictionStatusReady,
		CampaignPredictionStatusWarning:
		return true
	default:
		return false
	}
}

func (s CampaignPredictionStatus) String() string { return string(s) }

type CampaignPredictionSource string

const (
	CampaignPredictionSourceSameSeries        CampaignPredictionSource = "same_series"
	CampaignPredictionSourceSimilarEvent      CampaignPredictionSource = "similar_event"
	CampaignPredictionSourceLast14DaysDoubled CampaignPredictionSource = "last_14_days_doubled"
)

func (s CampaignPredictionSource) IsValid() bool {
	switch s {
	case CampaignPredictionSourceSameSeries, CampaignPredictionSourceSimilarEvent,
		CampaignPredictionSourceLast14DaysDoubled:
		return true
	default:
		return false
	}
}

func (s CampaignPredictionSource) String() string { return string(s) }
