package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/pricing/membership"

	"github.com/Potato-Mart/Backend-Shared-Contract/v25/pkg/contracts/pricing/wallet/wallet_enums"
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
	Type             wallet_enums.WalletInstrumentType `json:"type"`
	Code             string                            `json:"code"`
	Status           string                            `json:"status,omitempty"`
	Value            *money.Money                      `json:"value,omitempty"`
	CommittedBalance *money.Money                      `json:"committed_balance,omitempty"`
	ReservedBalance  *money.Money                      `json:"reserved_balance,omitempty"`
	AvailableBalance *money.Money                      `json:"available_balance,omitempty"`
	IssuedAt         *time.Time                        `json:"issued_at,omitempty"`
	ActivatedAt      *time.Time                        `json:"activated_at,omitempty"`
	// RedeemedAt applies only to single-use instruments such as vouchers,
	// coupons, and rewards. Re-spendable gift cards expose transaction history
	// instead of a singular redemption timestamp.
	RedeemedAt *time.Time `json:"redeemed_at,omitempty"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
}

// CustomerWalletSummary is a thin headline projection. The points figure mirrors
// membership.MembershipWalletSummary. Gift-card totals deliberately separate
// committed ledger value from live reservations and checkout-available value;
// storefront headline totals use GiftCardAvailableBalanceTotal.
type CustomerWalletSummary struct {
	AvailablePoints               int                      `json:"available_points"`
	PointDebt                     int                      `json:"point_debt"`
	PointsPolicy                  *membership.PointsPolicy `json:"points_policy,omitempty"`
	GiftCardCommittedBalanceTotal money.Money              `json:"gift_card_committed_balance_total"`
	GiftCardReservedBalanceTotal  money.Money              `json:"gift_card_reserved_balance_total"`
	GiftCardAvailableBalanceTotal money.Money              `json:"gift_card_available_balance_total"`
	ActiveGiftCards               int                      `json:"active_gift_cards"`
	ActiveVouchers                int                      `json:"active_vouchers"`
	ActiveCoupons                 int                      `json:"active_coupons"`
	AvailableRewards              int                      `json:"available_rewards"`
}
