package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/import_compliance/import_compliance_enums"
)

// LabelLayout stores the local preview controls. FontScaleBasisPoints uses
// 10,000 as a scale of 1.0.
type LabelLayout struct {
	Size                 import_compliance_enums.LabelSize        `json:"size"`
	Orientation          import_compliance_enums.LabelOrientation `json:"orientation"`
	FontScaleBasisPoints int64                                    `json:"font_scale_basis_points"`
	IncludeBarcode       bool                                     `json:"include_barcode"`
}
