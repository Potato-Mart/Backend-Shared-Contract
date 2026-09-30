package courier

// DeliveryCapabilities contains admin-authored capability flags. They describe
// supported operations, not live availability, verified readiness or coverage.
// False values never imply support. ProviderSlots means the configured API
// supplies transient authoritative selectable windows; no slot catalogue is stored.
type DeliveryCapabilities struct {
	Booking          bool `json:"booking"`
	Tracking         bool `json:"tracking"`
	ProofOfDelivery  bool `json:"proof_of_delivery"`
	Refrigerated     bool `json:"refrigerated"`
	ProviderCoverage bool `json:"provider_coverage"`
	ProviderSlots    bool `json:"provider_slots"`
}
