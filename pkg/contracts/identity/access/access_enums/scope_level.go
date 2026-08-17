package access_enums

// ScopeLevel is the geographic breadth a workforce principal is granted.
//
// The levels are ordered widest to narrowest. Depots are the only site
// identity in the platform: there is no separate store entity, so a
// depot-scoped principal is scoped by depot code even when the depot trades
// as a store.
type ScopeLevel string

const (
	// ScopeLevelGlobal sees every country, market, and depot.
	ScopeLevelGlobal ScopeLevel = "global"
	// ScopeLevelCountry sees one country and everything inside it.
	ScopeLevelCountry ScopeLevel = "country"
	// ScopeLevelMarket sees the granted markets only.
	ScopeLevelMarket ScopeLevel = "market"
	// ScopeLevelDepot sees the granted depots only.
	ScopeLevelDepot ScopeLevel = "depot"
)

// IsValid reports whether l is a known ScopeLevel value.
func (l ScopeLevel) IsValid() bool {
	switch l {
	case ScopeLevelGlobal, ScopeLevelCountry, ScopeLevelMarket, ScopeLevelDepot:
		return true
	}
	return false
}

// String returns the wire value for l.
func (l ScopeLevel) String() string { return string(l) }
