package membership

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pricing/membership/membership_enums"

// MembershipPassBarcode is the canonical scannable membership identifier.
type MembershipPassBarcode struct {
	Format        membership_enums.WalletPassBarcodeFormat `json:"format"`
	Value         string                                   `json:"value"`
	AlternateText string                                   `json:"alternate_text,omitempty"`
}
