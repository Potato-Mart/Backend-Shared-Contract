package procurement

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v36/pkg/contracts/common/temporal"
)

// NetGoodsCostSourceSnapshot is immutable private acquisition/valuation
// evidence for one source line or its exact allocated receipt portion. It
// preserves the numerator (NetGoodsAmount) and denominator (BaseUnits) used
// by Supply, without storing a rounded per-base-unit cost as source truth.
// Source identity, revision, issue date and confirmation time preserve the
// evidence for selection; selection and fractional allocation stay in Supply.
type NetGoodsCostSourceSnapshot struct {
	SourceType     string         `json:"source_type"`
	SourceID       string         `json:"source_id"`
	SourceLineID   string         `json:"source_line_id,omitempty"`
	SourceRevision int64          `json:"source_revision"`
	IssueDate      *temporal.Date `json:"issue_date,omitempty"`
	ConfirmedAt    *time.Time     `json:"confirmed_at,omitempty"`

	ReceiptID     string `json:"receipt_id,omitempty"`
	ReceiptItemID string `json:"receipt_item_id,omitempty"`
	BaseUnits     int64  `json:"base_units"`
	// NetGoodsAmount excludes all tax (recoverable or not), freight, and
	// duty, after goods discounts. Absence means unvalued/unknown evidence;
	// an explicit zero amount retains a known zero-cost source.
	NetGoodsAmount *money.Money `json:"net_goods_amount,omitempty"`
	// Provisional identifies cost evidence awaiting final invoice valuation.
	Provisional bool `json:"provisional"`
}
