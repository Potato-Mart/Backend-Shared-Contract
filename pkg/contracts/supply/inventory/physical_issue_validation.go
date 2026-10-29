package inventory

import (
	"math"
	"strings"

	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/packaging"
)

// ValidPhysicalIssueComposition checks a positive package-grain quantity. It
// does not establish physical availability, source identity or valuation.
func ValidPhysicalIssueComposition(sku string, option packaging.PackageOptionRef, value packaging.PackageCompositionSnapshot) bool {
	if sku == "" || strings.TrimSpace(sku) != sku || option.SKUCode != sku || option.Code == "" || strings.TrimSpace(option.Code) != option.Code || option.Version <= 0 || value.TotalBaseUnits <= 0 || len(value.Components) == 0 {
		return false
	}
	var total int64
	seen := map[string]bool{}
	for _, part := range value.Components {
		if part.PackageOption != option || !part.HandlingUnit.IsValid() || part.PackageCount <= 0 || part.UnitsPerPackage <= 0 || part.PackageCount > math.MaxInt64/part.UnitsPerPackage || part.BaseUnits != part.PackageCount*part.UnitsPerPackage || total > math.MaxInt64-part.BaseUnits {
			return false
		}
		key := string(part.HandlingUnit)
		if seen[key] {
			return false
		}
		seen[key] = true
		total += part.BaseUnits
	}
	return total == value.TotalBaseUnits
}
