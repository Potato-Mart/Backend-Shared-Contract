package audience_enums

// CustomerType is an offer classification, independent of campaign audience and checkout authentication.
type CustomerType string

const (
	CustomerTypeGuest     CustomerType = "guest"
	CustomerTypeRetail    CustomerType = "retail"
	CustomerTypeWholesale CustomerType = "wholesale"
)

func (v CustomerType) IsValid() bool {
	return v == CustomerTypeGuest || v == CustomerTypeRetail || v == CustomerTypeWholesale
}
func (v CustomerType) String() string { return string(v) }
