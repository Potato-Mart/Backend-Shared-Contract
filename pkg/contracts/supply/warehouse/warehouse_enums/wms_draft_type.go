package warehouse_enums

// WMSDraftType classifies a WMS draft as an inbound or outbound operation.
type WMSDraftType string

const (
	WMSDraftTypeInbound  WMSDraftType = "INBOUND"
	WMSDraftTypeOutbound WMSDraftType = "OUTBOUND"
)

func (w WMSDraftType) IsValid() bool {
	switch w {
	case WMSDraftTypeInbound, WMSDraftTypeOutbound:
		return true
	}
	return false
}

func (w WMSDraftType) String() string { return string(w) }
