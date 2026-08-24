package warehouse_enums

// LayoutNodeType is the level of a node in the warehouse 3D hierarchy.
// The tree is conceptually Zone > Aisle > Rack > Shelf > Bin, but a
// concrete depot may skip levels (e.g., go straight from Aisle to Bin).
type LayoutNodeType string

const (
	LayoutNodeZone  LayoutNodeType = "ZONE"
	LayoutNodeAisle LayoutNodeType = "AISLE"
	LayoutNodeRack  LayoutNodeType = "RACK"
	LayoutNodeShelf LayoutNodeType = "SHELF"
	LayoutNodeBin   LayoutNodeType = "BIN"
)

// IsValid reports whether t is a known LayoutNodeType.
func (t LayoutNodeType) IsValid() bool {
	switch t {
	case LayoutNodeZone, LayoutNodeAisle, LayoutNodeRack,
		LayoutNodeShelf, LayoutNodeBin:
		return true
	}
	return false
}

func (t LayoutNodeType) String() string { return string(t) }
