package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// PackagePriceEntry is one independently priced package option for a SKU in a
// price book. PackageAmount buys exactly one package identified by the
// immutable PackageOptionCode; it is not a base-unit PriceEntry amount.
type PackagePriceEntry struct {
	ID                     string                           `json:"id"`
	PriceBookCode          string                           `json:"price_book_code"`
	SKUCode                string                           `json:"sku_code"`
	PackageOptionCode      string                           `json:"package_option_code"`
	PackageAmount          money.Money                      `json:"package_amount"`
	Status                 pricebook_enums.PriceEntryStatus `json:"status"`
	Derivation             pricebook_enums.PriceDerivation  `json:"derivation"`
	ValidFrom              time.Time                        `json:"valid_from"`
	ValidUntil             *time.Time                       `json:"valid_until,omitempty"`
	Approval               *audit.LifecycleAction           `json:"approval,omitempty"`
	Rejection              *audit.LifecycleAction           `json:"rejection,omitempty"`
	Withdrawal             *audit.LifecycleAction           `json:"withdrawal,omitempty"`
	SourceBaseCostRevision *int64                           `json:"source_base_cost_revision,omitempty"`
	Revision               int64                            `json:"revision"`

	audit.AuditFields
}
