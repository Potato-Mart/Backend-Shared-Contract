package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/wallet/wallet_enums"
)

// ExternalRewardFulfilment is the partner-side evidence for an EXTERNAL reward
// redemption. ProviderCode is the backend-configured open code taken from the
// reward's external benefit configuration, and ExternalReference is the
// partner's own identifier for the provisioned subscription or service.
type ExternalRewardFulfilment struct {
	ProviderCode      string                                      `json:"provider_code"`
	ExternalReference string                                      `json:"external_reference,omitempty"`
	Status            wallet_enums.ExternalRewardFulfilmentStatus `json:"status"`
	FailureReason     string                                      `json:"failure_reason,omitempty"`
	ProvisionedAt     *time.Time                                  `json:"provisioned_at,omitempty"`
	RevokedAt         *time.Time                                  `json:"revoked_at,omitempty"`
}
