package ads

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

// AdConnection is how one platform is wired up for one market: which accounts
// are in use, whether the conversion plumbing is switched on, and whether the
// credentials still work.
//
// It carries no credential and no part of one. Tokens, developer tokens and
// keys live only in the secret store, and CredentialStatus reports their health
// without revealing them. MaskedPixelID and ConversionActionName are
// recognition labels, not identifiers to send to a platform.
// MetaConversionsAPIEnabled, GoogleAutoTagging and
// GoogleEnhancedConversionsEnabled are the platform-side switches our
// server-side conversions depend on; an unset switch means the platform
// feature is off, not unknown. LastSyncedAt and LastErrorCode mirror the
// platform's sync state so a degraded connection is visible without a
// separate read.
type AdConnection struct {
	Platform                         ads_enums.AdPlatform         `json:"platform,omitempty"`
	Status                           ads_enums.AdConnectionStatus `json:"status,omitempty"`
	Accounts                         []AdAccount                  `json:"accounts,omitempty"`
	MaskedPixelID                    string                       `json:"masked_pixel_id,omitempty"`
	MetaConversionsAPIEnabled        bool                         `json:"meta_conversions_api_enabled,omitempty"`
	ConversionActionName             string                       `json:"conversion_action_name,omitempty"`
	GoogleAutoTagging                bool                         `json:"google_auto_tagging,omitempty"`
	GoogleEnhancedConversionsEnabled bool                         `json:"google_enhanced_conversions_enabled,omitempty"`
	CredentialStatus                 ads_enums.AdCredentialStatus `json:"credential_status,omitempty"`
	CredentialExpiresAt              *time.Time                   `json:"credential_expires_at,omitempty"`
	PagePostSourcesEnabled           []string                     `json:"page_post_sources_enabled,omitempty"`
	LastSyncedAt                     *time.Time                   `json:"last_synced_at,omitempty"`
	LastErrorCode                    string                       `json:"last_error_code,omitempty"`
	SourcePlatforms                  []ads_enums.AdSourcePlatform `json:"source_platforms,omitempty"`
}
