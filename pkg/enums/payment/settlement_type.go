package paymentenum

// SettlementType distinguishes a real settlement (closes the batch and
// transfers funds to the merchant's bank) from an enquiry (read-only
// summary of transactions for a given date).
type SettlementType string

const (
	SettlementTypeSettlement SettlementType = "settlement"
	SettlementTypeEnquiry    SettlementType = "enquiry"
)

// IsValid reports whether t is a known SettlementType.
func (t SettlementType) IsValid() bool {
	switch t {
	case SettlementTypeSettlement, SettlementTypeEnquiry:
		return true
	}
	return false
}

func (t SettlementType) String() string { return string(t) }
