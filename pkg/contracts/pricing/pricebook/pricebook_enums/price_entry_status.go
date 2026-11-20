package pricebook_enums

type PriceEntryStatus string

const (
	PriceEntryStatusDraft           PriceEntryStatus = "draft"
	PriceEntryStatusPendingApproval PriceEntryStatus = "pending_approval"
	PriceEntryStatusApproved        PriceEntryStatus = "approved"
	// PriceEntryStatusScheduled is approved with a future valid_from; it
	// becomes current at valid_from.
	PriceEntryStatusScheduled  PriceEntryStatus = "scheduled"
	PriceEntryStatusRejected   PriceEntryStatus = "rejected"
	PriceEntryStatusSuperseded PriceEntryStatus = "superseded"
	PriceEntryStatusWithdrawn  PriceEntryStatus = "withdrawn"
	PriceEntryStatusExpired    PriceEntryStatus = "expired"
)

func (s PriceEntryStatus) IsValid() bool {
	switch s {
	case PriceEntryStatusDraft, PriceEntryStatusPendingApproval, PriceEntryStatusApproved, PriceEntryStatusScheduled, PriceEntryStatusRejected, PriceEntryStatusSuperseded, PriceEntryStatusWithdrawn, PriceEntryStatusExpired:
		return true
	}
	return false
}
func (s PriceEntryStatus) String() string { return string(s) }
