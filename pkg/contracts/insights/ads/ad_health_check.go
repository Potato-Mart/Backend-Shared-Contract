package ads

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdHealthCheck is one tracking health finding: a code, how bad it is, and the
// numbers behind it. The staff-facing wording belongs to admin, so no text is
// carried here.
//
// Which counters a finding fills depends on Code, and a counter left at zero
// means "not part of this check" as readily as it means zero. ObservedCount
// and MatchedCount are the pair behind a gap check; SentCount, FailedCount,
// DeduplicatedCount and WithdrawnCount are the server-side conversion
// counters; RatioBasisPoints carries a share as basis points, so 12.5 % is
// 1250 and no ratio is ever a float. EntityIDs names the offending objects at
// Level, for example the delivering ads that carry no tracking parameters.
// ExpiresAt is set only by a credential finding, and never alongside the
// credential itself.
type AdHealthCheck struct {
	Code              ads_enums.AdHealthCheckCode `json:"code,omitempty"`
	Severity          ads_enums.AdHealthSeverity  `json:"severity,omitempty"`
	Platform          ads_enums.AdPlatform        `json:"platform,omitempty"`
	Level             ads_enums.AdLevel           `json:"level,omitempty"`
	EntityIDs         []string                    `json:"entity_ids,omitempty"`
	ObservedCount     int64                       `json:"observed_count,omitempty"`
	MatchedCount      int64                       `json:"matched_count,omitempty"`
	SentCount         int64                       `json:"sent_count,omitempty"`
	FailedCount       int64                       `json:"failed_count,omitempty"`
	DeduplicatedCount int64                       `json:"deduplicated_count,omitempty"`
	WithdrawnCount    int64                       `json:"withdrawn_count,omitempty"`
	RatioBasisPoints  int32                       `json:"ratio_basis_points,omitempty"`
	ExpiresAt         *time.Time                  `json:"expires_at,omitempty"`
	CheckedAt         *time.Time                  `json:"checked_at,omitempty"`
}
