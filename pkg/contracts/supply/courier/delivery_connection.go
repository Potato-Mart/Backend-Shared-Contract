package courier

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/supply/courier/courier_enums"
)

// DeliveryConnection carries admin-safe connection evidence only. Credential
// values, secret resource/version references and provider diagnostics remain
// in backend-owned storage and DTOs, never this projection.
// Health and LastCheckedAt are backend-observed evidence, not editable claims.
type DeliveryConnection struct {
	CredentialConfigured bool                                   `json:"credential_configured"`
	Health               courier_enums.DeliveryConnectionHealth `json:"health"`
	LastCheckedAt        *time.Time                             `json:"last_checked_at,omitempty"`
}
