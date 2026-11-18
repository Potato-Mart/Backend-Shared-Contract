package fulfilment

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"time"
)

// PackageSubstitutionSnapshot records the exact loose-item replacement evidence
// for one requested sealed-case substitution.
//
// ReasonCode is a warehouse_enums.PackageSubstitutionReason value held as a
// string. For CASE_SHORTAGE_EXACT_EACH the replacement comes from one lot and
// one bucket, named by LotID and SourceBucketID, and Parts is empty.
// PACKING_CASE_BREAK also records loose units at a single destination bucket
// after a sealed case was opened at the station.
//
// For CASE_ASSEMBLED_FROM_LOOSE the snapshot records ONE case made up from
// loose units, which may come from several lots and buckets:
// RequestedCaseCount is 1, ReplacementBaseUnits is the units per case,
// LotID and SourceBucketID are empty, and Parts lists where every unit came
// from. The Parts' BaseUnits add up exactly to ReplacementBaseUnits, each part
// names an allocation, a bucket and a lot with positive units, and no two
// parts share an allocation id. No buyer consent is needed for an assembled
// case. Parts is set only for CASE_ASSEMBLED_FROM_LOOSE. Supply enforces these
// rules; this record only carries the evidence.
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
	Parts                        []PackageSubstitutionPart  `json:"parts,omitempty"`
	ReasonCode                   string                     `json:"reason_code"`
	Operator                     string                     `json:"operator"`
	CapturedAt                   time.Time                  `json:"captured_at"`
}
