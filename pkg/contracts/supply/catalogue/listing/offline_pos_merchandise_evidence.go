package listing

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v48/pkg/contracts/common/packaging/packaging_enums"
)

// OfflinePOSMerchandiseEvidence identifies the merchandise captured in a POS
// journal against a cached signed book. The depot is the till's captured sale
// location; it is not a claim that the book signs the depot. It carries no
// online validity token and asserts neither stock nor independent signature
// validity. Service owners verify the book and journal separately.
type OfflinePOSMerchandiseEvidence struct {
	ClientSaleID     string                              `json:"client_sale_id"`
	SKUCode          string                              `json:"sku_code"`
	MarketCode       string                              `json:"market_code"`
	DepotCode        string                              `json:"depot_code"`
	PackageOption    packaging.PackageOptionRef          `json:"package_option"`
	HandlingUnit     packaging_enums.PackageHandlingUnit `json:"handling_unit"`
	PackageCount     int64                               `json:"package_count"`
	UnitsPerPackage  int64                               `json:"units_per_package"`
	BaseUnits        int64                               `json:"base_units"`
	PriceBookVersion string                              `json:"price_book_version"`
	PriceBookSHA256  string                              `json:"price_book_sha256"`
	JournalEntryHash string                              `json:"journal_entry_hash"`
	SoldAt           time.Time                           `json:"sold_at"`
}
