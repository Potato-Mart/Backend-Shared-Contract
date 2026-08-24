package pricebook_enums

type PriceBookAssignmentKind string

const (
	PriceBookAssignmentKindChannelDefault       PriceBookAssignmentKind = "channel_default"
	PriceBookAssignmentKindOrganisationCategory PriceBookAssignmentKind = "organisation_category"
	PriceBookAssignmentKindOrganisationOverride PriceBookAssignmentKind = "organisation_override"
)

func (k PriceBookAssignmentKind) IsValid() bool {
	switch k {
	case PriceBookAssignmentKindChannelDefault, PriceBookAssignmentKindOrganisationCategory, PriceBookAssignmentKindOrganisationOverride:
		return true
	}
	return false
}
func (k PriceBookAssignmentKind) String() string { return string(k) }
