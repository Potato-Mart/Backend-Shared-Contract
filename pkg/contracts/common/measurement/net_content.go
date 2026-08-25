package measurement

// NetContent is the declared measurement content of one sellable base unit and
// the comparison quantity used for unit pricing.
//
// NetQuantity is the net content stated on the pack. StandardMeasure is the
// quantity the comparison ("unit") price is expressed per, for example per
// 100 g or per 1 L. Both are required before a listing that is configured as
// unit-pricing-required may be activated; the calculation itself is service
// behaviour.
type NetContent struct {
	NetQuantity     Measure `json:"net_quantity"`
	StandardMeasure Measure `json:"standard_measure"`
}
