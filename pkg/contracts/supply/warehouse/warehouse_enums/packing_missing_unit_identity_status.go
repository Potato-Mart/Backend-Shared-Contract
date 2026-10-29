package warehouse_enums

// PackingMissingUnitIdentityStatus describes identification evidence for
// tracked units missing from an order during packing.
type PackingMissingUnitIdentityStatus string

const PackingMissingUnitIdentityUnidentified PackingMissingUnitIdentityStatus = "unidentified"

func (s PackingMissingUnitIdentityStatus) IsValid() bool {
	return s == PackingMissingUnitIdentityUnidentified
}

func (s PackingMissingUnitIdentityStatus) String() string { return string(s) }
