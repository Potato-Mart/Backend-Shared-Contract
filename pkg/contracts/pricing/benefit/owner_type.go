package benefit

// OwnerType identifies the retail customer or wholesale organisation that
// owns a non-membership commercial benefit.
type OwnerType string

const (
	OwnerTypeRetailCustomer        OwnerType = "retail_customer"
	OwnerTypeWholesaleOrganisation OwnerType = "wholesale_organisation"
)

func (o OwnerType) IsValid() bool {
	switch o {
	case OwnerTypeRetailCustomer, OwnerTypeWholesaleOrganisation:
		return true
	default:
		return false
	}
}

func (o OwnerType) String() string { return string(o) }
