package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/geography"
	security "github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/pricing/wallet/wallet_enums"
)

// RewardRedemption records a customer's redemption of a Membership reward.
type RewardRedemption struct {
	ID             string `json:"id"`
	CustomerNumber string `json:"customer_number"`
	// MarketCode and CountryCode are the denormalized market and country used
	// when this redemption was created, so geographically scoped reads are
	// plain indexed matches.
	MarketCode         string                              `json:"market_code,omitempty"`
	CountryCode        geography.CountryCode               `json:"country_code,omitempty"`
	RewardCode         string                              `json:"reward_code"`
	ReservationID      string                              `json:"reservation_id,omitempty"`
	PointsSpent        int                                 `json:"points_spent"`
	Status             wallet_enums.RewardRedemptionStatus `json:"status"`
	Outcome            *RewardRedemptionOutcome            `json:"outcome,omitempty"`
	RelatedOrderNumber string                              `json:"related_order_number,omitempty"`
	FulfilledAt        *time.Time                          `json:"fulfilled_at,omitempty"`
	ExpiresAt          *time.Time                          `json:"expires_at,omitempty"`
	CreatedBy          string                              `json:"created_by,omitempty"`
	CreatedAt          time.Time                           `json:"created_at"`
	History            []security.HistoryEntry             `json:"history,omitempty"`
}
