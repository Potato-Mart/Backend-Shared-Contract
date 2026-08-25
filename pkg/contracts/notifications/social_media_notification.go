package notifications

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/metadata"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"
)

// SocialMediaNotification is a provider-neutral direct or official-account
// message. ProviderCode and MessageMode are backend-configured open strings;
// this contract intentionally declares no provider enum. Metadata may contain
// provider rendering attributes only and must never contain raw destinations,
// access tokens, credentials, webhook secrets, or transport configuration.
type SocialMediaNotification struct {
	ProviderCode          string                 `json:"provider_code"`
	MessageMode           string                 `json:"message_mode"`
	SenderReference       string                 `json:"sender_reference,omitempty"`
	RecipientReference    string                 `json:"recipient_reference,omitempty"`
	ConversationReference string                 `json:"conversation_reference,omitempty"`
	Body                  string                 `json:"body"`
	Media                 []security.ObjectMedia `json:"media,omitempty"`
	ActionURL             string                 `json:"action_url,omitempty"`
	Metadata              metadata.Metadata      `json:"metadata,omitempty"`
}
