package supply_enums

// InventoryLotReceivedCause is why stock arrived in an inventory lot. Receipt
// is a goods receipt; LotSplit is a lot split off from a received lot (for
// example at put-away or transfer); LotReassignment moves stock into a
// different lot; CountGain is a stock count that found more than recorded;
// POSAddBack is stock returned to the shelf from a POS sale.
type InventoryLotReceivedCause string

const (
	InventoryLotReceivedCauseReceipt         InventoryLotReceivedCause = "receipt"
	InventoryLotReceivedCauseLotSplit        InventoryLotReceivedCause = "lot_split"
	InventoryLotReceivedCauseLotReassignment InventoryLotReceivedCause = "lot_reassignment"
	InventoryLotReceivedCauseCountGain       InventoryLotReceivedCause = "count_gain"
	InventoryLotReceivedCausePOSAddBack      InventoryLotReceivedCause = "pos_add_back"
)

func (c InventoryLotReceivedCause) IsValid() bool {
	switch c {
	case InventoryLotReceivedCauseReceipt, InventoryLotReceivedCauseLotSplit, InventoryLotReceivedCauseLotReassignment,
		InventoryLotReceivedCauseCountGain, InventoryLotReceivedCausePOSAddBack:
		return true
	}
	return false
}

func (c InventoryLotReceivedCause) String() string { return string(c) }
