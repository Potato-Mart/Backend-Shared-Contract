package notifications

import security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"

// PushNotification is the provider-neutral authored content for one push
// notification delivery.
type PushNotification struct {
	Title     string                `json:"title"`
	Body      string                `json:"body"`
	ActionURL string                `json:"action_url,omitempty"`
	Image     *security.ObjectMedia `json:"image,omitempty"`
}
