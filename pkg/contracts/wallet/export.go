package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/contracts/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/contracts/promotion"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v13/pkg/enums/wallet"
)

const WalletExportSchemaVersion = "wallet_export_v1"

// WalletExportSpec describes how Management should build a durable export.
type WalletExportSpec struct {
	Format          walletenum.WalletExportFormat `json:"format,omitempty"`
	IncludeHistory  *bool                         `json:"include_history,omitempty"`
	From            *time.Time                    `json:"from,omitempty"`
	To              *time.Time                    `json:"to,omitempty"`
	IncludeExpired  *bool                         `json:"include_expired,omitempty"`
	IncludeRedeemed *bool                         `json:"include_redeemed,omitempty"`
	IncludeVoided   *bool                         `json:"include_voided,omitempty"`
	Locale          string                        `json:"locale,omitempty"`
}

// WalletExportRecord is returned by export status endpoints.
type WalletExportRecord struct {
	ID            string                        `json:"id"`
	SchemaVersion string                        `json:"schema_version"`
	Owner         membership.MembershipOwnerRef `json:"owner"`
	RequestedBy   string                        `json:"requested_by,omitempty"`
	Format        walletenum.WalletExportFormat `json:"format"`
	Status        walletenum.WalletExportStatus `json:"status"`
	Filters       WalletExportFilters           `json:"filters"`
	RecordCounts  WalletExportRecordCounts      `json:"record_counts,omitempty"`
	Checksum      string                        `json:"checksum,omitempty"`
	FailureReason string                        `json:"failure_reason,omitempty"`
	DownloadURL   string                        `json:"download_url,omitempty"`
	RequestedAt   time.Time                     `json:"requested_at"`
	CompletedAt   *time.Time                    `json:"completed_at,omitempty"`
	ExpiresAt     time.Time                     `json:"expires_at"`
}

type WalletExportFilters struct {
	IncludeHistory  bool       `json:"include_history"`
	From            *time.Time `json:"from,omitempty"`
	To              *time.Time `json:"to,omitempty"`
	IncludeExpired  bool       `json:"include_expired"`
	IncludeRedeemed bool       `json:"include_redeemed"`
	IncludeVoided   bool       `json:"include_voided"`
	Locale          string     `json:"locale,omitempty"`
}

type WalletExportRecordCounts struct {
	MembershipPointLedgerEntries int `json:"membership_point_ledger_entries"`
	CouponAssignments            int `json:"coupon_assignments"`
	CouponUsageRows              int `json:"coupon_usage_rows"`
	Vouchers                     int `json:"vouchers"`
	GiftCards                    int `json:"gift_cards"`
	GiftCardTransactions         int `json:"gift_card_transactions"`
	RewardRedemptions            int `json:"reward_redemptions"`
	HistoryEvents                int `json:"history_events"`
}

// WalletExport is the canonical JSON payload for wallet downloads.
type WalletExport struct {
	SchemaVersion string                        `json:"schema_version"`
	Owner         membership.MembershipOwnerRef `json:"owner"`
	Membership    WalletExportMembership        `json:"membership"`
	Coupons       []WalletExportCoupon          `json:"coupons,omitempty"`
	Vouchers      []Voucher                     `json:"vouchers,omitempty"`
	GiftCards     []WalletExportGiftCard        `json:"gift_cards,omitempty"`
	Rewards       []membership.RewardRedemption `json:"rewards,omitempty"`
	History       []WalletExportHistoryEvent    `json:"history,omitempty"`
	Summary       WalletExportSummary           `json:"summary"`
	GeneratedAt   time.Time                     `json:"generated_at"`
	GeneratedBy   string                        `json:"generated_by,omitempty"`
	Filters       WalletExportFilters           `json:"filters"`
	RecordCounts  WalletExportRecordCounts      `json:"record_counts"`
	Checksum      string                        `json:"checksum,omitempty"`
}

type WalletExportMembership struct {
	Account      *membership.MembershipAccount `json:"account,omitempty"`
	PointLedger  []membership.PointLedgerEntry `json:"point_ledger,omitempty"`
	PointBuckets []membership.PointBucket      `json:"point_buckets,omitempty"`
}

type WalletExportCoupon struct {
	Assignment promotion.CouponAssignment `json:"assignment"`
	Coupon     *promotion.Coupon          `json:"coupon,omitempty"`
	Usage      []WalletExportCouponUsage  `json:"usage,omitempty"`
}

type WalletExportCouponUsage struct {
	ID                  string        `json:"id"`
	CouponCode          string        `json:"coupon_code"`
	CustomerNumber      string        `json:"customer_number,omitempty"`
	RedeemedOrderNumber string        `json:"redeemed_order_number,omitempty"`
	DiscountAmount      *common.Money `json:"discount_amount,omitempty"`
	RedeemedAt          *time.Time    `json:"redeemed_at,omitempty"`
	CreatedAt           time.Time     `json:"created_at,omitempty"`
}

type WalletExportGiftCard struct {
	GiftCard     GiftCard              `json:"gift_card"`
	Transactions []GiftCardTransaction `json:"transactions,omitempty"`
}

type WalletExportHistoryEvent struct {
	At                 time.Time                       `json:"at"`
	Type               walletenum.WalletInstrumentType `json:"type"`
	Code               string                          `json:"code,omitempty"`
	Status             string                          `json:"status,omitempty"`
	DeltaPoints        int                             `json:"delta_points,omitempty"`
	DeltaMoney         *common.Money                   `json:"delta_money,omitempty"`
	BalanceAfter       *common.Money                   `json:"balance_after,omitempty"`
	RelatedOrderNumber string                          `json:"related_order_number,omitempty"`
	Source             string                          `json:"source,omitempty"`
	Note               string                          `json:"note,omitempty"`
}

type WalletExportSummary struct {
	AvailablePoints      int          `json:"available_points"`
	GiftCardBalanceTotal common.Money `json:"gift_card_balance_total"`
	ActiveGiftCards      int          `json:"active_gift_cards"`
	ActiveVouchers       int          `json:"active_vouchers"`
	ActiveCoupons        int          `json:"active_coupons"`
	AvailableRewards     int          `json:"available_rewards"`
	RedeemedItems        int          `json:"redeemed_items"`
	ExpiredItems         int          `json:"expired_items"`
	VoidedItems          int          `json:"voided_items"`
}
