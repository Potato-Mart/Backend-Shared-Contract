package measurement

// MeasurementUnit names the unit of a measured quantity, for example "g",
// "kg", "mL", "L", "m", or "ea". It is an open typed string: the unit
// vocabulary a market requires for comparison pricing is configuration, not a
// closed contract enum.
type MeasurementUnit string
