package warehouse

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/supply/warehouse/warehouse_enums"
)

// MarketSalePolicy is the Supply-owned stock sale policy of one commercial
// market. SellPastDateMark is the market default for selling stock at or past
// its date mark; a product's own override replaces it for that product.
// Stored is false when the market has no saved policy and the value shown is
// Supply's default. History lists every change, oldest first.
type MarketSalePolicy struct {
	MarketCode       string                                 `json:"market_code"`
	SellPastDateMark warehouse_enums.SellPastDateMarkPolicy `json:"sell_past_date_mark"`
	Stored           bool                                   `json:"stored"`
	Revision         int64                                  `json:"revision"`
	History          []security.HistoryEntry                `json:"history,omitempty"`

	audit.AuditFields
}
