package product

import security "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/security"

// SellingProductImages is the resolved media projection of the Product image
// groups. It carries render-safe object media rather than storage records or
// code-only master references.
type SellingProductImages struct {
	Cover   *security.ObjectMedia  `json:"cover,omitempty"`
	Gallery []security.ObjectMedia `json:"gallery,omitempty"`
	Details []security.ObjectMedia `json:"details,omitempty"`
}
