package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// MixedAdminDiscountLine binds the global allocation to a priced line. Added
// lines omit native source fields; retained lines include their complete
// authenticated predecessor component and whole-package range.
type MixedAdminDiscountLine struct {
	LineID              string            `json:"line_id"`
	Kind                string            `json:"kind"`
	SourceVersionID     string            `json:"source_version_id,omitempty"`
	SourceItemID        string            `json:"source_item_id,omitempty"`
	SourceComponentID   string            `json:"source_component_id,omitempty"`
	SourceLineID        string            `json:"source_line_id,omitempty"`
	FirstPackageOrdinal int64             `json:"first_package_ordinal,omitempty"`
	LastPackageOrdinal  int64             `json:"last_package_ordinal,omitempty"`
	SourcePackageCount  int64             `json:"source_package_count,omitempty"`
	UnitsPerPackage     int64             `json:"units_per_package,omitempty"`
	PreAdminGross       money.Money       `json:"pre_admin_gross"`
	PreAdminTax         money.Money       `json:"pre_admin_tax"`
	CarriedGross        money.Money       `json:"carried_gross"`
	CarriedTax          money.Money       `json:"carried_tax"`
	AppliedGross        money.Money       `json:"applied_gross"`
	AppliedTax          money.Money       `json:"applied_tax"`
	AppliedRounding     *RoundingEvidence `json:"applied_rounding,omitempty"`
}
