package courier

// CarrierBookingRequestSummaryEntry is one redacted, non-secret fact about the
// request sent to a carrier, kept for staff diagnosis. Key is an open,
// Supply-owned identifier. Value is display text already redacted by Supply:
// credentials, configuration values marked sensitive and full contact details
// must never be recorded. It is not the raw provider payload.
type CarrierBookingRequestSummaryEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
