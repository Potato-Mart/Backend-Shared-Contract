package enums

// SalesPerformance classifies how well a product is selling relative to
// the rest of the product set. It replaces the former free-form
// "freshness_status" string (which the storefront had repurposed for an
// expiry indicator). It is a merchandising signal — typically derived
// from sales velocity (e.g. avg_weekly_sales) by an analytics job — used
// to surface best-sellers and flag slow movers.
type SalesPerformance string

const (
	// SalesPerformanceHot is a best-seller / fast mover (熱賣).
	SalesPerformanceHot SalesPerformance = "hot"
	// SalesPerformanceNormal is a steady, average-selling product.
	SalesPerformanceNormal SalesPerformance = "normal"
	// SalesPerformanceSlow is a slow mover that may need attention (滯銷).
	SalesPerformanceSlow SalesPerformance = "slow"
)

// IsValid reports whether s is a known SalesPerformance.
func (s SalesPerformance) IsValid() bool {
	switch s {
	case SalesPerformanceHot, SalesPerformanceNormal, SalesPerformanceSlow:
		return true
	}
	return false
}

func (s SalesPerformance) String() string { return string(s) }
