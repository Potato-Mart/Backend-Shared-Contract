package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/wallet/wallet_enums"
)

// WalletInstrument is a uniform link to one customer value instrument.
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
	RedeemedAt       *time.Time                        `json:"redeemed_at,omitempty"`
	ExpiresAt        *time.Time                        `json:"expires_at,omitempty"`
}
