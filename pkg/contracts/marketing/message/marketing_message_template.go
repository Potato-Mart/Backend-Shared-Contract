package message

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/marketing/message/message_enums"
)

// MarketingMessageTemplate is an immutable, versioned message payload.
type MarketingMessageTemplate struct {
	TemplateCode string                                `json:"template_code"`
	Version      int                                   `json:"version"`
	Status       message_enums.MarketingTemplateStatus `json:"status"`
	Payload      MarketingMessage                      `json:"payload"`

	audit.AuditFields
}
