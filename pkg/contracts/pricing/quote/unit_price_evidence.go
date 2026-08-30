package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// UnitPriceEvidence is the comparison-unit price shown beside an ordinary
// grocery listing. Exempt records a listing that is not required to display a
// comparison price, with the reason a markdown created that exemption.
type UnitPriceEvidence struct {
	NetContent       measurement.NetContent `json:"net_content"`
	ComparisonAmount money.Money            `json:"comparison_amount"`
	Exempt           bool                   `json:"exempt"`
	ExemptionReason  string                 `json:"exemption_reason,omitempty"`
}
