package balance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v40/pkg/contracts/pricing/wallet/wallet_enums"
)

// WalletInstrument is a uniform link to one customer value instrument.
type WalletInstrument struct {
	Type wallet_enums.WalletInstrumentType `json:"type"`
	Code string                            `json:"code"`
	// CustomerStatus is optional authoritative customer usage evidence. Absence
	// means unknown, never infer available/used from Status or RedeemedAt.
	// Available means ready with remaining use; used means consumed with no
	// remaining use; expired means its validity period ended. Pricing owns
	// precedence, underlying outcome correlation and current use eligibility.
	CustomerStatus   wallet_enums.WalletInstrumentCustomerStatus `json:"customer_status,omitempty"`
	Status           string                                      `json:"status,omitempty"`
	Value            *money.Money                                `json:"value,omitempty"`
	CommittedBalance *money.Money                                `json:"committed_balance,omitempty"`
	ReservedBalance  *money.Money                                `json:"reserved_balance,omitempty"`
	AvailableBalance *money.Money                                `json:"available_balance,omitempty"`
	IssuedAt         *time.Time                                  `json:"issued_at,omitempty"`
	ActivatedAt      *time.Time                                  `json:"activated_at,omitempty"`
	RedeemedAt       *time.Time                                  `json:"redeemed_at,omitempty"`
	ExpiresAt        *time.Time                                  `json:"expires_at,omitempty"`
}
