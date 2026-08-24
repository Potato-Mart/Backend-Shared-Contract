package wallet

import "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"

// CustomerWalletSummary is the customer-safe headline projection of wallet
// holdings. Membership policy is composed by Membership, not imported here.
type CustomerWalletSummary struct {
	Points                        PointsSummary `json:"points"`
	GiftCardCommittedBalanceTotal money.Money   `json:"gift_card_committed_balance_total"`
	GiftCardReservedBalanceTotal  money.Money   `json:"gift_card_reserved_balance_total"`
	GiftCardAvailableBalanceTotal money.Money   `json:"gift_card_available_balance_total"`
	ActiveGiftCards               int           `json:"active_gift_cards"`
	ActiveVouchers                int           `json:"active_vouchers"`
	ActiveCoupons                 int           `json:"active_coupons"`
	AvailableRewards              int           `json:"available_rewards"`
}
