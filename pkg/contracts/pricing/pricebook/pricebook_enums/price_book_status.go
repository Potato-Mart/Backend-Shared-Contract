package pricebook_enums

type PriceBookStatus string

const (
	PriceBookStatusDraft    PriceBookStatus = "draft"
	PriceBookStatusActive   PriceBookStatus = "active"
	PriceBookStatusInactive PriceBookStatus = "inactive"
	PriceBookStatusArchived PriceBookStatus = "archived"
)

func (s PriceBookStatus) IsValid() bool {
	switch s {
	case PriceBookStatusDraft, PriceBookStatusActive, PriceBookStatusInactive, PriceBookStatusArchived:
		return true
	}
	return false
}
func (s PriceBookStatus) String() string { return string(s) }
