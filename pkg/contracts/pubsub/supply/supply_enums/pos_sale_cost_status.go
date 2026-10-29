package supply_enums

// POSSaleCostStatus reports whether the movement has verified valuation evidence.
// Unresolved includes absent, provisional or otherwise unverified cost evidence.
type POSSaleCostStatus string

const (
	POSSaleCostStatusValued     POSSaleCostStatus = "valued"
	POSSaleCostStatusUnresolved POSSaleCostStatus = "unresolved"
)

func (status POSSaleCostStatus) IsValid() bool {
	return status == POSSaleCostStatusValued || status == POSSaleCostStatusUnresolved
}

func (status POSSaleCostStatus) String() string { return string(status) }
