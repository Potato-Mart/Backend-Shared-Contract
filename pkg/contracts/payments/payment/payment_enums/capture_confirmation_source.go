package payment_enums

// CaptureConfirmationSource identifies authenticated evidence that capture
// had already succeeded by an upper bound; it does not identify exact capture time.
type CaptureConfirmationSource string

const (
	CaptureConfirmationSourceProviderEvent             CaptureConfirmationSource = "provider_event"
	CaptureConfirmationSourceProviderStatusObservation CaptureConfirmationSource = "provider_status_observation"
)

func (s CaptureConfirmationSource) IsValid() bool {
	switch s {
	case CaptureConfirmationSourceProviderEvent, CaptureConfirmationSourceProviderStatusObservation:
		return true
	default:
		return false
	}
}
func (s CaptureConfirmationSource) String() string { return string(s) }
