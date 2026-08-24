package pricebook_enums

type PriceEntryStatus string

const (
	PriceEntryStatusDraft           PriceEntryStatus = "draft"
	PriceEntryStatusPendingApproval PriceEntryStatus = "pending_approval"
	PriceEntryStatusApproved        PriceEntryStatus = "approved"
	PriceEntryStatusRejected        PriceEntryStatus = "rejected"
	PriceEntryStatusSuperseded      PriceEntryStatus = "superseded"
	PriceEntryStatusWithdrawn       PriceEntryStatus = "withdrawn"
	PriceEntryStatusExpired         PriceEntryStatus = "expired"
)

func (s PriceEntryStatus) IsValid() bool {
	switch s {
	case PriceEntryStatusDraft, PriceEntryStatusPendingApproval, PriceEntryStatusApproved, PriceEntryStatusRejected, PriceEntryStatusSuperseded, PriceEntryStatusWithdrawn, PriceEntryStatusExpired:
		return true
	}
	return false
}
func (s PriceEntryStatus) String() string { return string(s) }
