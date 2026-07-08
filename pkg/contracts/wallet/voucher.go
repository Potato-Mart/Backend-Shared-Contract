package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v14/pkg/contracts/membership"
)

// Voucher is a customer-held, single-redemption instrument (often issued by a
// membership reward). Unlike a gift card it is not re-spendable stored value:
// Value, when set, is a fixed amount applied once on redemption. Status is one
// of "issued", "redeemed", "expired", or "void". Referenced by Code.
type Voucher struct {
	ID                  string                        `json:"id"`
	Code                string                        `json:"code"`
	Owner               membership.MembershipOwnerRef `json:"owner"`
	Value               *common.Money                 `json:"value,omitempty"`
	Status              string                        `json:"status,omitempty"`
	SourceRewardCode    string                        `json:"source_reward_code,omitempty"`
	IssuedAt            time.Time                     `json:"issued_at"`
	ExpiresAt           *time.Time                    `json:"expires_at,omitempty"`
	RedeemedAt          *time.Time                    `json:"redeemed_at,omitempty"`
	RedeemedOrderNumber string                        `json:"redeemed_order_number,omitempty"`
	Note                string                        `json:"note,omitempty"`

	common.AuditFields
}
