package message_enums

// MarketingTemplateStatus is the lifecycle state of an immutable template.
type MarketingTemplateStatus string

const (
	MarketingTemplateStatusActive   MarketingTemplateStatus = "active"
	MarketingTemplateStatusArchived MarketingTemplateStatus = "archived"
)

func (s MarketingTemplateStatus) IsValid() bool {
	return s == MarketingTemplateStatusActive || s == MarketingTemplateStatusArchived
}
func (s MarketingTemplateStatus) String() string { return string(s) }
