package import_compliance

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/temporal"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/supply/purchase/purchase_enums"
)

// PurchaseOrderSnapshot freezes the purchase-order header used by a
// declaration or tariff assessment. Each aggregate owns its corresponding
// line snapshots.
type PurchaseOrderSnapshot struct {
	ID                   string                             `json:"id"`
	OrderNumber          string                             `json:"order_number"`
	Status               purchase_enums.PurchaseOrderStatus `json:"status"`
	SupplierCode         string                             `json:"supplier_code,omitempty"`
	SupplierName         string                             `json:"supplier_name,omitempty"`
	ExpectedArrival      temporal.Date                      `json:"expected_arrival,omitempty"`
	CapturedAt           time.Time                          `json:"captured_at"`
	SourceChecksumSHA256 string                             `json:"source_checksum_sha256,omitempty"`
}
