package wallet

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/pricing/wallet/wallet_enums"
)

// MembershipPassContent is the provider-neutral snapshot used to issue a
// customer-owned membership pass. Google save-JWT responses, Apple pass
// packages, signing policy, routes, and media types remain backend-owned.
type MembershipPassContent struct {
	CustomerNumber  string                `json:"customer_number"`
	TierKey         string                `json:"tier_key,omitempty"`
	AvailablePoints int                   `json:"available_points"`
	Barcode         MembershipPassBarcode `json:"barcode"`
	GeneratedAt     time.Time             `json:"generated_at"`
}

// MembershipPassBarcode is the canonical scannable membership identifier.
// Value is the raw retail customer number; AlternateText is safe display copy.
type MembershipPassBarcode struct {
	Format        wallet_enums.WalletPassBarcodeFormat `json:"format"`
	Value         string                               `json:"value"`
	AlternateText string                               `json:"alternate_text,omitempty"`
}
