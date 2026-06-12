package enums

// DiscountType controls how the discount's Value is interpreted.
//
//	percentage    -> Value is a percentage 0..100 applied to the
//	                 eligible subtotal, capped by MaxDiscountAmount.
//	fixed_amount  -> Value is a flat amount in the store's currency.
//	fixed_price   -> Value is the final sale price (major units) of the
//	                 targeted product; only meaningful on promotions with
//	                 target_scope=product or target_scope=category.
type DiscountType string

const (
	DiscountTypePercentage   DiscountType = "percentage"
	DiscountTypeFixedAmount  DiscountType = "fixed_amount"
	DiscountTypeFreeShipping DiscountType = "free_shipping"
	DiscountTypeFixedPrice   DiscountType = "fixed_price"
)

func (d DiscountType) IsValid() bool {
	switch d {
	case DiscountTypePercentage, DiscountTypeFixedAmount, DiscountTypeFreeShipping, DiscountTypeFixedPrice:
		return true
	}
	return false
}

func (d DiscountType) String() string { return string(d) }

// DiscountScope controls what the discount applies to.
type DiscountScope string

const (
	DiscountScopeAll      DiscountScope = "all"
	DiscountScopeCategory DiscountScope = "category"
	DiscountScopeProduct  DiscountScope = "product"
)

func (d DiscountScope) IsValid() bool {
	switch d {
	case DiscountScopeAll, DiscountScopeCategory, DiscountScopeProduct:
		return true
	}
	return false
}

func (d DiscountScope) String() string { return string(d) }
