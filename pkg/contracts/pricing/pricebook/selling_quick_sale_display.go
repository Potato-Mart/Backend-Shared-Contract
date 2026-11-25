package pricebook

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
)

// SellingQuickSaleDisplay is one POS Quick sale clearance price that the
// requesting store's POS applies now to one package of the display line's
// package option. Pricing resolves it only for retail POS displays at a store.
// It is shelf-label evidence: it never replaces a checkout PriceSnapshot, and
// the cashier still selects the offer (or scans its lot) to charge it.
//
// Price and RegularPrice are tax-inclusive amounts for one package of the
// line's package option. For a whole-lot EACH offer on a case line they are
// the case amounts POS quotes. RegularPrice is the WAS price.
type SellingQuickSaleDisplay struct {
	// OfferCode and Revision identify the published Quick sale offer and its
	// revision, as a POS Quick sale selection references them.
	OfferCode string `json:"offer_code"`
	Revision  int64  `json:"revision"`
	// LotID is the depot lot the offer clears. Stock owners use it to drop an
	// offer whose lot has nothing left to sell.
	LotID string `json:"lot_id"`
	// WholeLot reports an offer for every package of the lot. A whole-lot
	// offer is published for EACH and also applies to the lot's cases.
	WholeLot bool `json:"whole_lot"`
	// RequiresLotScan reports a scan-only offer: POS charges it only after
	// the cashier scans the lot, and it is absent from the POS offer list.
	RequiresLotScan     bool        `json:"requires_lot_scan"`
	Price               money.Money `json:"price"`
	RegularPrice        money.Money `json:"regular_price"`
	EffectivePercentOff int         `json:"effective_percent_off"`
	// PastBestBefore reports the past-best-before phase, when the offer's
	// after-date terms apply.
	PastBestBefore bool `json:"past_best_before"`
	// PhaseEndsAt is the exclusive instant the current POS phase ends: the
	// price changes or the offer stops selling on POS. A display line that
	// carries this offer is valid no later than it.
	PhaseEndsAt time.Time `json:"phase_ends_at"`
	// SellUntil is the label's sell-until instant for the current POS phase:
	// the lot's date mark before it, or the past-best-before cutoff after it.
	// When an earlier cutoff ends the in-date phase first, it is that cutoff.
	// Timezone is the date mark's IANA timezone, used to print it.
	SellUntil time.Time `json:"sell_until"`
	Timezone  string    `json:"timezone,omitempty"`
}
