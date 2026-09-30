package fulfilment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v39/pkg/contracts/common/packaging"
	"time"
)

// PackageSubstitutionSnapshot records the exact loose-item replacement evidence
// for one requested sealed-case substitution.
type PackageSubstitutionSnapshot struct {
	ID                           string                     `json:"id"`
	RequestedCasePackageOption   packaging.PackageOptionRef `json:"requested_case_package_option"`
	RequestedCaseCount           int64                      `json:"requested_case_count"`
	RequestedUnitsPerCase        int64                      `json:"requested_units_per_case"`
	FulfilledSealedCaseCount     int64                      `json:"fulfilled_sealed_case_count"`
	ReplacementEachPackageOption packaging.PackageOptionRef `json:"replacement_each_package_option"`
	ReplacementBaseUnits         int64                      `json:"replacement_base_units"`
	LotID                        string                     `json:"lot_id"`
	SourceBucketID               string                     `json:"source_bucket_id"`
	StockUnitIDs                 []string                   `json:"stock_unit_ids,omitempty"`
	ReasonCode                   string                     `json:"reason_code"`
	Operator                     string                     `json:"operator"`
	CapturedAt                   time.Time                  `json:"captured_at"`
}
