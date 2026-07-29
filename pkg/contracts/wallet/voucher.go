package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/benefit"
	walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/wallet"
)

// Voucher is a customer-held, single-redemption instrument (often issued by a
// membership reward). Unlike a gift card it is not re-spendable stored value:
// Value, when set, is a fixed amount applied once on redemption. Reserved is a
// temporary checkout hold. The voucher is referenced by Code.
type Voucher struct {
	ID                       string                   `json:"id"`
	Code                     string                   `json:"code"`
	Owner                    benefit.OwnerRef         `json:"owner"`
	Value                    *common.Money            `json:"value,omitempty"`
	Status                   walletenum.VoucherStatus `json:"status"`
	SourceRewardCode         string                   `json:"source_reward_code,omitempty"`
	SourceRewardRedemptionID string                   `json:"source_reward_redemption_id,omitempty"`
	IssuedAt                 time.Time                `json:"issued_at"`
	ExpiresAt                *time.Time               `json:"expires_at,omitempty"`
	ReservationID            string                   `json:"reservation_id,omitempty"`
	ReservedAt               *time.Time               `json:"reserved_at,omitempty"`
	ReservationExpiresAt     *time.Time               `json:"reservation_expires_at,omitempty"`
	RedeemedAt               *time.Time               `json:"redeemed_at,omitempty"`
	RedeemedOrderNumber      string                   `json:"redeemed_order_number,omitempty"`
	Note                     string                   `json:"note,omitempty"`

	common.AuditFields
}
