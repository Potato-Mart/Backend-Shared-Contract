package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// PriceUnitAllocation freezes identical per-unit monetary values for one
// contiguous one-based ordinal range on this snapshot's requested line. The
// ranges partition the parent source allocation and their unit amounts times
// BaseUnits reconcile exactly to its list, discount, consideration, taxable
// base, tax and payable totals. Differing discounts or allocated minor-unit
// remainders require separate ranges. These values describe the captured quote;
// they do not prescribe a refund or amendment pricing policy. For a retained
// amendment these ordinals are rebased to the new line; its predecessor binding
// records the native selection from which the retained units originated.
type PriceUnitAllocation struct {
	FirstUnit          int64       `json:"first_unit"`
	BaseUnits          int64       `json:"base_units"`
	UnitListAmount     money.Money `json:"unit_list_amount"`
	UnitDiscountAmount money.Money `json:"unit_discount_amount"`
	UnitConsideration  money.Money `json:"unit_consideration"`
	UnitTaxableBase    money.Money `json:"unit_taxable_base"`
	UnitTaxAmount      money.Money `json:"unit_tax_amount"`
	UnitLineTotal      money.Money `json:"unit_line_total"`
}
