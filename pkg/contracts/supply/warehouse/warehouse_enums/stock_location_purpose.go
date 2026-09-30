package warehouse_enums

// StockLocationPurpose describes the operational purpose of a stock location.
type StockLocationPurpose string

const (
	StockLocationPurposeExpiryHold         StockLocationPurpose = "EXPIRY_HOLD"
	StockLocationPurposeStandard           StockLocationPurpose = "STANDARD"
	StockLocationPurposeQualityHold        StockLocationPurpose = "QUALITY_HOLD"
	StockLocationPurposeOnlineOrderStaging StockLocationPurpose = "ONLINE_ORDER_STAGING"
)

func (p StockLocationPurpose) IsValid() bool {
	switch p {
	case StockLocationPurposeExpiryHold, StockLocationPurposeStandard, StockLocationPurposeQualityHold,
		StockLocationPurposeOnlineOrderStaging:
		return true
	}
	return false
}

func (p StockLocationPurpose) String() string { return string(p) }
