package geography

// GeographicPath is one explicit prefix of Country -> Market -> State -> Depot.
// StateCode is an official subdivision, never an operational depot region.
// Services validate contiguous parents and authoritative ancestry. Multiple
// paths are independent selections, not Cartesian combinations.
type GeographicPath struct {
	CountryCode CountryCode     `json:"country_code"`
	MarketCode  string          `json:"market_code,omitempty"`
	StateCode   SubdivisionCode `json:"state_code,omitempty"`
	DepotCode   string          `json:"depot_code,omitempty"`
}
