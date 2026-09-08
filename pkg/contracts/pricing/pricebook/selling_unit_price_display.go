package pricebook

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
)

// SellingUnitPriceDisplay records the exact measurement basis and optional
// comparison price for customer display, independently of checkout evidence.
// An absent amount must not be rendered as zero. An exemption and its reason
// are evidence supplied by the owning service, not an instruction to derive
// eligibility or a price in a consumer.
type SellingUnitPriceDisplay struct {
	NetContent       measurement.NetContent `json:"net_content"`
	ComparisonAmount *money.Money           `json:"comparison_amount,omitempty"`
	Exempt           bool                   `json:"exempt"`
	ExemptionReason  string                 `json:"exemption_reason,omitempty"`
}
