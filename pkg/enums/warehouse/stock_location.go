package warehouseenum

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

// StockLocationHandlingMode describes the package forms accepted by a stock
// location.
type StockLocationHandlingMode string

const (
	StockLocationHandlingEach  StockLocationHandlingMode = "EACH"
	StockLocationHandlingCase  StockLocationHandlingMode = "CASE"
	StockLocationHandlingMixed StockLocationHandlingMode = "MIXED"
)

func (m StockLocationHandlingMode) IsValid() bool {
	switch m {
	case StockLocationHandlingEach, StockLocationHandlingCase, StockLocationHandlingMixed:
		return true
	}
	return false
}

func (m StockLocationHandlingMode) String() string { return string(m) }

// StockLocationAccess describes who may access a stock location.
type StockLocationAccess string

const (
	StockLocationAccessCustomerAccessible StockLocationAccess = "CUSTOMER_ACCESSIBLE"
	StockLocationAccessStaffOnly          StockLocationAccess = "STAFF_ONLY"
)

func (a StockLocationAccess) IsValid() bool {
	switch a {
	case StockLocationAccessCustomerAccessible, StockLocationAccessStaffOnly:
		return true
	}
	return false
}

func (a StockLocationAccess) String() string { return string(a) }

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

// StockLocationCollectionRole identifies the primary or overflow role of a
// collection's eligible location.
type StockLocationCollectionRole string

const (
	StockLocationCollectionPrimary  StockLocationCollectionRole = "PRIMARY"
	StockLocationCollectionOverflow StockLocationCollectionRole = "OVERFLOW"
)

func (r StockLocationCollectionRole) IsValid() bool {
	switch r {
	case StockLocationCollectionPrimary, StockLocationCollectionOverflow:
		return true
	}
	return false
}

func (r StockLocationCollectionRole) String() string { return string(r) }
