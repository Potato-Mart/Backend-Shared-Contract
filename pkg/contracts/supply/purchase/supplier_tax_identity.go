package purchase

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"
)

// SupplierTaxIdentity is the supplier's own registration evidence recorded on
// an invoice. Registration is an entity fact about the supplier and never a
// statement about a product's taxability.
type SupplierTaxIdentity struct {
	SupplierCode string `json:"supplier_code"`
	SupplierName string `json:"supplier_name"`
	// BusinessNumberScheme names the register the number belongs to, for
	// example "abn".
	BusinessNumberScheme string         `json:"business_number_scheme,omitempty"`
	BusinessNumber       string         `json:"business_number,omitempty"`
	TaxRegistered        bool           `json:"tax_registered"`
	TaxRegisteredFrom    *temporal.Date `json:"tax_registered_from,omitempty"`
}
