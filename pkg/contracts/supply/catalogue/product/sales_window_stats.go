package product

// SalesWindowStats identifies the bounded period represented by its totals.
type SalesWindowStats struct {
	WindowDays int `json:"window_days"`
	SalesTotals
}
