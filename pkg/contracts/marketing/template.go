package marketing

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/marketing/marketing_enums"
)

// CampaignTemplate is an immutable, versioned copy of a campaign payload. It
// deliberately contains a copied payload instead of a live aggregate link so
// later campaign edits cannot rewrite an already-versioned template.
type CampaignTemplate struct {
	TemplateCode string                         `json:"template_code"`
	Version      int                            `json:"version"`
	Status       marketing_enums.TemplateStatus `json:"status"`
	Payload      Campaign                       `json:"payload"`

	audit.AuditFields
}

// MarketingMessageTemplate is an immutable, versioned copy of a marketing
// message payload. It contains no recipient or delivery-provider state.
type MarketingMessageTemplate struct {
	TemplateCode string                         `json:"template_code"`
	Version      int                            `json:"version"`
	Status       marketing_enums.TemplateStatus `json:"status"`
	Payload      MarketingMessage               `json:"payload"`

	audit.AuditFields
}
