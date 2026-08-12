package market_enums

// MarketStatus is the admin-controlled lifecycle state of a commercial market.
type MarketStatus string

const (
	// MarketStatusDraft is a market being configured that may not be quoted
	// against or referenced by an active listing or price book.
	MarketStatusDraft MarketStatus = "draft"
	// MarketStatusActive is a market open for commercial trading.
	MarketStatusActive MarketStatus = "active"
	// MarketStatusSuspended is a market temporarily closed to new trading
	// while its configuration and history are retained.
	MarketStatusSuspended MarketStatus = "suspended"
	// MarketStatusRetired is a market permanently withdrawn from trading.
	MarketStatusRetired MarketStatus = "retired"
)

// IsValid reports whether s is a known MarketStatus.
func (s MarketStatus) IsValid() bool {
	switch s {
	case MarketStatusDraft, MarketStatusActive, MarketStatusSuspended, MarketStatusRetired:
		return true
	}
	return false
}

func (s MarketStatus) String() string { return string(s) }
