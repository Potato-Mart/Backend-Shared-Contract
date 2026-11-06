package promotion

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/promotion/promotion_enums"
)

// PromotionProgress reports how far a cart is from one promotion relation.
// Names repeat approved promotion copy and are omitted when the promotion has
// none. LineIDs lists the cart lines counted toward the relation and is empty
// when no line has progress. Each optional quantity or amount is omitted when
// it does not apply to the relation or could not be established.
//
// PotentialSaving is indicative: stock, limits and other offers may change it.
//
// GiftOptions and GiftDecision are present only for gift-type relations
// (spend_gift, addon_purchase, bogo). GiftOptions names every SKU the relation
// can award as its gift or add-on; a till offers exactly these and never
// guesses from catalogue displays. GiftDecision says whether staff must accept
// or decline the gift and how many gift base units the relation allows.
type PromotionProgress struct {
	PromotionID       string                                  `json:"promotion_id"`
	PromotionRevision int64                                   `json:"promotion_revision"`
	PromotionKind     string                                  `json:"promotion_kind"`
	RelationID        string                                  `json:"relation_id"`
	Names             []localization.LocalizedName            `json:"names,omitempty"`
	Status            promotion_enums.PromotionProgressStatus `json:"status"`
	LineIDs           []string                                `json:"line_ids"`
	RequiredQuantity  *int64                                  `json:"required_quantity,omitempty"`
	CurrentQuantity   *int64                                  `json:"current_quantity,omitempty"`
	MissingQuantity   *int64                                  `json:"missing_quantity,omitempty"`
	MissingQualifiers []PromotionMissingQualifier             `json:"missing_qualifiers,omitempty"`
	SpendMissing      *money.Money                            `json:"spend_missing,omitempty"`
	PotentialSaving   *money.Money                            `json:"potential_saving,omitempty"`
	GiftOptions       []PromotionGiftOption                   `json:"gift_options,omitempty"`
	GiftDecision      *PromotionGiftDecision                  `json:"gift_decision,omitempty"`
}
