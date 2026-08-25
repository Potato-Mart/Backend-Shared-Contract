package measurement

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
