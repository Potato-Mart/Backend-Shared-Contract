package wallet

import (
	security "github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/benefit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet/wallet_enums"
	"time"
)

// CouponAssignment is an owner-specific issuance of a wallet Coupon.
type CouponAssignment struct {
	ID                  string                    `json:"id"`
	CouponID            string                    `json:"coupon_id"`
	CouponCode          string                    `json:"coupon_code"`
	Owner               benefit.OwnerRef          `json:"owner"`
	Source              wallet_enums.CouponSource `json:"source"`
	Status              string                    `json:"status"`
	ExpiresAt           *time.Time                `json:"expires_at,omitempty"`
	RedeemedAt          *time.Time                `json:"redeemed_at,omitempty"`
	RedeemedOrderNumber string                    `json:"redeemed_order_number,omitempty"`
	VoidedAt            *time.Time                `json:"voided_at,omitempty"`
	Note                string                    `json:"note,omitempty"`
	History             []security.HistoryEntry   `json:"history,omitempty"`
	CreatedAt           time.Time                 `json:"created_at"`
}
