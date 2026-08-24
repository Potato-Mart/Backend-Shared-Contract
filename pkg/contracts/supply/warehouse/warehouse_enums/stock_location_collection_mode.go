package warehouse_enums

// StockLocationCollectionMode describes whether collection access is limited
// to an explicit collection list.
type StockLocationCollectionMode string

const (
	StockLocationCollectionAllowList    StockLocationCollectionMode = "ALLOW_LIST"
	StockLocationCollectionUnrestricted StockLocationCollectionMode = "UNRESTRICTED"
)

func (m StockLocationCollectionMode) IsValid() bool {
	switch m {
	case StockLocationCollectionAllowList, StockLocationCollectionUnrestricted:
		return true
	}
	return false
}

func (m StockLocationCollectionMode) String() string { return string(m) }
