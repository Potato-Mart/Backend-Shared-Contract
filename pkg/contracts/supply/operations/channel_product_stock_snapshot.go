package operations

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/commerce/commerce_enums"
)

// ChannelProductStockSnapshot qualifies product stock by sales channel.
type ChannelProductStockSnapshot struct {
	Channel    commerce_enums.OrderType     `json:"channel"`
	Quantities ProductStockQuantitySnapshot `json:"quantities"`
}
