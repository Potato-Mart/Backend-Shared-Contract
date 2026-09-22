package courier

// DeliveryCapabilities describes supported delivery operations, not current
// availability or confirmed coverage. False values never imply support.
// ConfiguredSlots means the backend can publish configured service windows;
// ProviderSlots means the provider supplies authoritative selectable windows.
type DeliveryCapabilities struct {
	Booking          bool `json:"booking"`
	Tracking         bool `json:"tracking"`
	ProofOfDelivery  bool `json:"proof_of_delivery"`
	Refrigerated     bool `json:"refrigerated"`
	ProviderCoverage bool `json:"provider_coverage"`
	ProviderSlots    bool `json:"provider_slots"`
	ConfiguredSlots  bool `json:"configured_slots"`
}
