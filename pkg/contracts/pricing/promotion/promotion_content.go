package promotion

import "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/localization"

// PromotionContent contains approved localized customer-facing copy.
type PromotionContent struct {
	Names           []localization.LocalizedName        `json:"names,omitempty"`
	Descriptions    []localization.LocalizedDescription `json:"descriptions,omitempty"`
	DisplayMessages []localization.LocalizedText        `json:"display_messages,omitempty"`
	ReceiptMessages []localization.LocalizedText        `json:"receipt_messages,omitempty"`
}
