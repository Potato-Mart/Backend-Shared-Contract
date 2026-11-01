package courier_enums

// CarrierBookingErrorCode is the provider-neutral vocabulary Supply maps a
// carrier's booking error onto. Raw carrier codes and messages are not wire
// values. Timeout accompanies the unreachable outcome: the carrier did not
// answer within Supply's deadline, so it may or may not have received the
// request, and the sweeper re-checks the attempt. UnsupportedArea means the
// carrier reported that it does not serve the delivery address. CarrierError
// is the fallback for a carrier failure that maps to no more specific code.
type CarrierBookingErrorCode string

const (
	CarrierBookingErrorCodeInvalidAddress  CarrierBookingErrorCode = "invalid_address"
	CarrierBookingErrorCodeInvalidContact  CarrierBookingErrorCode = "invalid_contact"
	CarrierBookingErrorCodeDuplicateJob    CarrierBookingErrorCode = "duplicate_job"
	CarrierBookingErrorCodeInvalidDate     CarrierBookingErrorCode = "invalid_date"
	CarrierBookingErrorCodeInvalidParcel   CarrierBookingErrorCode = "invalid_parcel"
	CarrierBookingErrorCodeAuthFailed      CarrierBookingErrorCode = "auth_failed"
	CarrierBookingErrorCodeRateLimited     CarrierBookingErrorCode = "rate_limited"
	CarrierBookingErrorCodeTimeout         CarrierBookingErrorCode = "timeout"
	CarrierBookingErrorCodeUnsupportedArea CarrierBookingErrorCode = "unsupported_area"
	CarrierBookingErrorCodeCarrierError    CarrierBookingErrorCode = "carrier_error"
)

func (c CarrierBookingErrorCode) IsValid() bool {
	switch c {
	case CarrierBookingErrorCodeInvalidAddress, CarrierBookingErrorCodeInvalidContact,
		CarrierBookingErrorCodeDuplicateJob, CarrierBookingErrorCodeInvalidDate,
		CarrierBookingErrorCodeInvalidParcel, CarrierBookingErrorCodeAuthFailed,
		CarrierBookingErrorCodeRateLimited, CarrierBookingErrorCodeTimeout,
		CarrierBookingErrorCodeUnsupportedArea, CarrierBookingErrorCodeCarrierError:
		return true
	default:
		return false
	}
}

func (c CarrierBookingErrorCode) String() string { return string(c) }
