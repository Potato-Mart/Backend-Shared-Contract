package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/common"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v19/pkg/enums/wallet"
)

// CustomerWallet is the retail read-model aggregate of every value instrument
// a customer holds: loyalty points, gift cards, vouchers, assigned coupons,
// and membership rewards. CustomerNumber is the canonical membership and
// wallet key. The per-instrument ledgers remain the source of truth; this is a
// projection. Wholesale APIs use a separate benefits projection and must not
// expose this wallet model.
type CustomerWallet struct {
	CustomerNumber string                `json:"customer_number"`
	Instruments    []WalletInstrument    `json:"instruments,omitempty"`
	Summary        CustomerWalletSummary `json:"summary"`
	CalculatedAt   time.Time             `json:"calculated_at"`
}

// WalletInstrument is a uniform link row pointing at one value instrument by its
// business key (gift_card_code, voucher_code, coupon_code, reward_code, or the
// customer_number for points). Gift cards expose their committed ledger
// balance, amount held by live reservations, and amount currently available to
// checkout. Value is the single-use face value for vouchers.
type WalletInstrument struct {
	Type             walletenum.WalletInstrumentType `json:"type"`
	Code             string                          `json:"code"`
	Status           string                          `json:"status,omitempty"`
	Value            *common.Money                   `json:"value,omitempty"`
	CommittedBalance *common.Money                   `json:"committed_balance,omitempty"`
	ReservedBalance  *common.Money                   `json:"reserved_balance,omitempty"`
	AvailableBalance *common.Money                   `json:"available_balance,omitempty"`
	ExpiresAt        *time.Time                      `json:"expires_at,omitempty"`
}

// CustomerWalletSummary is a thin headline projection. The points figure mirrors
// membership.MembershipWalletSummary. Gift-card totals deliberately separate
// committed ledger value from live reservations and checkout-available value;
// storefront headline totals use GiftCardAvailableBalanceTotal.
type CustomerWalletSummary struct {
	AvailablePoints               int          `json:"available_points"`
	GiftCardCommittedBalanceTotal common.Money `json:"gift_card_committed_balance_total"`
	GiftCardReservedBalanceTotal  common.Money `json:"gift_card_reserved_balance_total"`
	GiftCardAvailableBalanceTotal common.Money `json:"gift_card_available_balance_total"`
	ActiveGiftCards               int          `json:"active_gift_cards"`
	ActiveVouchers                int          `json:"active_vouchers"`
	ActiveCoupons                 int          `json:"active_coupons"`
	AvailableRewards              int          `json:"available_rewards"`
}
