package warehouse_enums

// StockLocationPurpose describes the operational purpose of a stock location.
type StockLocationPurpose string

const (
	StockLocationPurposeStandard           StockLocationPurpose = "STANDARD"
	StockLocationPurposeQualityHold        StockLocationPurpose = "QUALITY_HOLD"
	StockLocationPurposeOnlineOrderStaging StockLocationPurpose = "ONLINE_ORDER_STAGING"
)

func (p StockLocationPurpose) IsValid() bool {
	switch p {
	case StockLocationPurposeStandard, StockLocationPurposeQualityHold,
		StockLocationPurposeOnlineOrderStaging:
		return true
	}
	return false
}

func (p StockLocationPurpose) String() string { return string(p) }
