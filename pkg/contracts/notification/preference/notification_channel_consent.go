package preference

import (
	"time"

	security "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notification/core/notification_enums"
)

// NotificationChannelConsent is the latest consent state held by the
// preference aggregate, not the legal system of record. The owning service
// stamps evidence fields and ignores client-supplied values for them. Social
// media consent is destination-scoped and requires DestinationCode; non-social
// channel consent leaves DestinationCode empty.
type NotificationChannelConsent struct {
	Channel         notification_enums.NotificationChannel `json:"channel"`
	DestinationCode string                                 `json:"destination_code,omitempty"`
	Granted         bool                                   `json:"granted"`
	Actor           security.ActorRef                      `json:"actor"`
	Source          string                                 `json:"source"`
	PolicyVersion   string                                 `json:"policy_version"`
	RequestID       string                                 `json:"request_id"`
	ChangedAt       time.Time                              `json:"changed_at"`
	Reason          string                                 `json:"reason,omitempty"`
}
