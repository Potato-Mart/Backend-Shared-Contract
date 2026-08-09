package promotion_enums

// PromotionStatus is the lifecycle status of a promotion definition.
type PromotionStatus string

const (
	PromotionStatusDraft    PromotionStatus = "draft"
	PromotionStatusActive   PromotionStatus = "active"
	PromotionStatusInactive PromotionStatus = "inactive"
	PromotionStatusArchived PromotionStatus = "archived"
)

// IsValid reports whether s is a supported promotion lifecycle status.
func (s PromotionStatus) IsValid() bool {
	switch s {
	case PromotionStatusDraft, PromotionStatusActive, PromotionStatusInactive, PromotionStatusArchived:
		return true
	}
	return false
}

// String returns the wire value for s.
func (s PromotionStatus) String() string { return string(s) }
