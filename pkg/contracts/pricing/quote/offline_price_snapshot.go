package quote

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pricing/pricebook"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/pricing/quote/quote_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/supply/catalogue/listing"
)

// OfflinePriceSnapshot captures what the till charged for one offline POS
// line. LineID is stable across journal upload, order creation, Pricing review
// and deterministic document-level tax allocation. LineNumber retains the
// journal's display order; it is not an online Pricing quote line ID.
// Book and merchandise references are facts to verify against retained signed
// bytes and journal evidence, not independent proof of price or eligibility.
// The tax snapshot states whether SourcePrice and ChargedPrice include tax.
type OfflinePriceSnapshot struct {
	LineID              string                                `json:"line_id"`
	LineNumber          int64                                 `json:"line_number"`
	Merchandise         listing.OfflinePOSMerchandiseEvidence `json:"merchandise"`
	PriceBook           pricebook.OfflinePriceBookEvidence    `json:"price_book"`
	CurrencyExponent    money.CurrencyExponent                `json:"currency_exponent"`
	SourcePrice         money.Money                           `json:"source_price"`
	SourcePriceBasis    quote_enums.PriceAmountBasis          `json:"source_price_basis"`
	ChargedPrice        money.Money                           `json:"charged_price"`
	ChargedPriceBasis   quote_enums.PriceAmountBasis          `json:"charged_price_basis"`
	Tax                 OfflineTaxSnapshot                    `json:"tax"`
	AppliedPromotionIDs []string                              `json:"applied_promotion_ids,omitempty"`
	CustomOverride      *CustomPriceOverrideEvidence          `json:"custom_override,omitempty"`
	CapturedAt          time.Time                             `json:"captured_at"`
}
