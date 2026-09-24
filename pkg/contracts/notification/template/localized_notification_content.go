package template

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/notification/push"
	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/notification/sms"
	"github.com/Potato-Mart/Backend-Shared-Contract/v34/pkg/contracts/notification/template/template_enums"
)

// LocalizedNotificationContent carries exactly one channel arm matching the
// enclosing NotificationContent. Language uses the same BCP 47 representation
// as common localization records. Source copy is authored; translated copy may
// include human refinements. Review is absent for the source and required for
// each confirmed target. Notification owns all one-of and freshness checks.
type LocalizedNotificationContent struct {
	Language string                               `json:"language"`
	Email    *EmailTemplateContent                `json:"email,omitempty"`
	SMS      *sms.SMSNotification                 `json:"sms,omitempty"`
	Push     *push.PushNotification               `json:"push,omitempty"`
	Origin   template_enums.TemplateContentOrigin `json:"origin"`
	Review   *TranslationReview                   `json:"review,omitempty"`
}
