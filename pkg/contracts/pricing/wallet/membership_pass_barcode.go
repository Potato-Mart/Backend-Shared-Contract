package wallet

import "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/wallet/wallet_enums"

// MembershipPassBarcode is the canonical scannable membership identifier.
type MembershipPassBarcode struct {
	Format        wallet_enums.WalletPassBarcodeFormat `json:"format"`
	Value         string                               `json:"value"`
	AlternateText string                               `json:"alternate_text,omitempty"`
}
