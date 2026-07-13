package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/contracts/membership"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v16/pkg/enums/wallet"
)

// CustomerWallet is a read-model aggregate of every value instrument a customer
// (retail or wholesale) holds: loyalty points, gift cards, vouchers, held
// coupons, and membership rewards. It is owner-keyed via
// membership.MembershipOwnerRef (the same owner key the membership programme
// uses) and links each instrument by its business key. The per-instrument
// ledgers remain the source of truth; this is a projection.
type CustomerWallet struct {
	Owner               membership.MembershipOwnerRef `json:"owner"`
	MembershipAccountID string                        `json:"membership_account_id,omitempty"`
	Instruments         []WalletInstrument            `json:"instruments,omitempty"`
	Summary             CustomerWalletSummary         `json:"summary"`
	CalculatedAt        time.Time                     `json:"calculated_at"`
}

// WalletInstrument is a uniform link row pointing at one value instrument by its
// business key (gift_card_code, voucher_code, coupon_code, reward_code, or the
// membership_account_id for points). Balance is set only for stored-value
// instruments (gift cards); it is nil for count/state instruments.
type WalletInstrument struct {
	Type      walletenum.WalletInstrumentType `json:"type"`
	Code      string                          `json:"code"`
	Status    string                          `json:"status,omitempty"`
	Balance   *common.Money                   `json:"balance,omitempty"`
	ExpiresAt *time.Time                      `json:"expires_at,omitempty"`
}

// CustomerWalletSummary is a thin headline projection. The points figure mirrors
// membership.MembershipWalletSummary; the gift-card/voucher/coupon/reward
// figures aggregate the linked instruments.
type CustomerWalletSummary struct {
	AvailablePoints      int          `json:"available_points"`
	GiftCardBalanceTotal common.Money `json:"gift_card_balance_total"`
	ActiveGiftCards      int          `json:"active_gift_cards"`
	ActiveVouchers       int          `json:"active_vouchers"`
	ActiveCoupons        int          `json:"active_coupons"`
	AvailableRewards     int          `json:"available_rewards"`
}
