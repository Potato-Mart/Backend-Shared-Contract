package ads

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdChangeRecord is one durable row of the delivery change history: what was
// changed, on which object, by whom, from where, and how it ended.
//
// Source separates our own writes from edits made directly in Ads Manager or
// Google Ads and discovered by a later sync; a platform-sourced row has no
// staff actor. Before and After are the normalised settings on either side of
// the change, so the row is readable without re-reading the platform.
// ExpectedRevision is the object revision the change was validated against: a
// row in the conflict state means the platform had moved on and nothing was
// applied. ErrorCode is our own platform-independent code; platform error
// text, request bodies and credentials never appear here. Warnings are the
// codes raised while validating, for example a large increase that needed
// confirming.
type AdChangeRecord struct {
	ChangeID         string                     `json:"change_id,omitempty"`
	Platform         ads_enums.AdPlatform       `json:"platform,omitempty"`
	Level            ads_enums.AdLevel          `json:"level,omitempty"`
	EntityID         string                     `json:"entity_id,omitempty"`
	EntityName       string                     `json:"entity_name,omitempty"`
	Status           ads_enums.AdChangeStatus   `json:"status,omitempty"`
	Source           ads_enums.AdChangeSource   `json:"source,omitempty"`
	ExpectedRevision string                     `json:"expected_revision,omitempty"`
	RequestedBy      string                     `json:"requested_by,omitempty"`
	RequestedAt      *time.Time                 `json:"requested_at,omitempty"`
	ResolvedAt       *time.Time                 `json:"resolved_at,omitempty"`
	Before           *AdChangeSet               `json:"before,omitempty"`
	After            *AdChangeSet               `json:"after,omitempty"`
	ErrorCode        string                     `json:"error_code,omitempty"`
	Warnings         []string                   `json:"warnings,omitempty"`
	MarketCode       string                     `json:"market_code,omitempty"`
	SourcePlatform   ads_enums.AdSourcePlatform `json:"source_platform,omitempty"`
}
