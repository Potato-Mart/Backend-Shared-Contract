package measurement

// MeasurementUnit names the unit of a measured quantity, for example "g",
// "kg", "mL", "L", "m", or "ea". It is an open typed string: the unit
// vocabulary a market requires for comparison pricing is configuration, not a
// closed contract enum.
type MeasurementUnit string

// Measure is an exact quantity in a named unit.
//
// Amount is scaled by Exponent so the value is Amount * 10^Exponent Unit; for
// example {Amount: 375, Exponent: 0, Unit: "mL"} is 375 mL and {Amount: 125,
// Exponent: -1, Unit: "g"} is 12.5 g. Keeping the mantissa an integer avoids
// float drift in unit-price evidence.
type Measure struct {
	Amount   int64           `json:"amount"`
	Exponent int32           `json:"exponent"`
	Unit     MeasurementUnit `json:"unit"`
}

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
