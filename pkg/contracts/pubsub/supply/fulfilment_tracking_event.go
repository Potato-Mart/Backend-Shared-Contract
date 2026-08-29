package supply

// FulfilmentTrackingEvent is emitted on the fulfilment-events topic when
// carrier tracking details are assigned or corrected.
type FulfilmentTrackingEvent struct {
	OrderNumber    string `json:"order_number"`
	ShipmentID     string `json:"shipment_id,omitempty"`
	Carrier        string `json:"carrier,omitempty"`
	TrackingNumber string `json:"tracking_number"`
	TrackingURL    string `json:"tracking_url,omitempty"`
	RequestID      string `json:"request_id,omitempty"`
}
