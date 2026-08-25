package product

// SalesTotals records paid-order unit activity for one measurement period.
// Completed line refunds are represented separately and subtracted in NetUnits.
type SalesTotals struct {
	PaidOrderCount int64 `json:"paid_order_count"`
	GrossUnits     int64 `json:"gross_units"`
	RefundedUnits  int64 `json:"refunded_units"`
	NetUnits       int64 `json:"net_units"`
}
