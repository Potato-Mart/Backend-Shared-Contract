package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// PriceEntry is one SKU's commercial amount inside one price book.
type PriceEntry struct {
	ID                     string                           `json:"id"`
	PriceBookCode          string                           `json:"price_book_code"`
	SKUCode                string                           `json:"sku_code"`
	Amount                 money.Money                      `json:"amount"`
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
