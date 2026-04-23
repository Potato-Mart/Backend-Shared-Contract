package enums

// DiscountType controls how the discount's Value is interpreted.
//
//	PERCENTAGE    -> Value is a percentage 0..100 applied to the
//	                 eligible subtotal, capped by MaxDiscountAmount.
//	FIXED_AMOUNT  -> Value is a flat amount in the store's currency.
type DiscountType string

const (
	DiscountTypePercentage   DiscountType = "PERCENTAGE"
	DiscountTypeFixedAmount  DiscountType = "FIXED_AMOUNT"
	DiscountTypeFreeShipping DiscountType = "FREE_SHIPPING"
)

func (d DiscountType) IsValid() bool {
	switch d {
	case DiscountTypePercentage, DiscountTypeFixedAmount, DiscountTypeFreeShipping:
		return true
	}
	return false
}

func (d DiscountType) String() string { return string(d) }

// DiscountScope controls what the discount applies to.
type DiscountScope string

const (
	DiscountScopeAll      DiscountScope = "ALL"
	DiscountScopeCategory DiscountScope = "CATEGORY"
	DiscountScopeProduct  DiscountScope = "PRODUCT"
)

func (d DiscountScope) IsValid() bool {
	switch d {
	case DiscountScopeAll, DiscountScopeCategory, DiscountScopeProduct:
		return true
	}
	return false
}

func (d DiscountScope) String() string { return string(d) }
