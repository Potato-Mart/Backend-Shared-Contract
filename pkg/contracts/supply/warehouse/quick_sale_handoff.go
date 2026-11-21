package warehouse

import "time"

// QuickSaleHandoff is carried by quick-sale admissions, offers, and line
// snapshots. Delivery and pickup must finish before HandoffBefore. When it is
// absent on a legacy record, consumers use the record's selling-phase cutoff.
type QuickSaleHandoff struct {
	HandoffBefore *time.Time `json:"handoff_before,omitempty"`
}
