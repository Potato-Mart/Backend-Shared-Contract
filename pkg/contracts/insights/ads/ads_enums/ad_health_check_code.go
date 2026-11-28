package ads_enums

// AdHealthCheckCode names one tracking health check. Each code fixes which
// numbers the check reports and which objects it names; the thresholds and the
// staff-facing wording stay service- and admin-owned.
type AdHealthCheckCode string

const (
	// AdHealthCheckCodeAdMissingTracking reports delivering ads whose
	// landing URLs carry none of the agreed tracking parameters.
	AdHealthCheckCodeAdMissingTracking AdHealthCheckCode = "ad_missing_tracking"
	// AdHealthCheckCodeConversionGap reports campaigns where the platform
	// claims many purchases but we matched far fewer paid orders.
	AdHealthCheckCodeConversionGap AdHealthCheckCode = "conversion_gap"
	// AdHealthCheckCodeClickIDMissing reports the share of ad sessions that
	// arrived without a platform click id.
	AdHealthCheckCodeClickIDMissing AdHealthCheckCode = "click_id_missing"
	// AdHealthCheckCodeCaptureRate reports the share of ad sessions for
	// which the storefront stored a click.
	AdHealthCheckCodeCaptureRate AdHealthCheckCode = "capture_rate"
	// AdHealthCheckCodeMetaConversionsAPI reports recent server-side Meta
	// purchase sends, failures and pixel deduplications.
	AdHealthCheckCodeMetaConversionsAPI AdHealthCheckCode = "meta_conversions_api"
	// AdHealthCheckCodeGoogleConversionUpload reports recent Google
	// conversion uploads, failures and withdrawals.
	AdHealthCheckCodeGoogleConversionUpload AdHealthCheckCode = "google_conversion_upload"
	// AdHealthCheckCodeToken reports a platform credential that is expiring
	// or no longer valid. It never carries the credential itself.
	AdHealthCheckCodeToken AdHealthCheckCode = "token"
	// AdHealthCheckCodeTimeZoneMismatch reports an ad account whose time
	// zone differs from the market it serves.
	AdHealthCheckCodeTimeZoneMismatch AdHealthCheckCode = "time_zone_mismatch"
	// AdHealthCheckCodeCurrencyMismatch reports an ad account whose currency
	// differs from the market it serves.
	AdHealthCheckCodeCurrencyMismatch AdHealthCheckCode = "currency_mismatch"
	// AdHealthCheckCodeSyncStale reports that no sync has succeeded recently
	// enough for the numbers to be trusted.
	AdHealthCheckCodeSyncStale AdHealthCheckCode = "sync_stale"
)

// IsValid reports whether c is a known AdHealthCheckCode.
func (c AdHealthCheckCode) IsValid() bool {
	switch c {
	case AdHealthCheckCodeAdMissingTracking, AdHealthCheckCodeConversionGap,
		AdHealthCheckCodeClickIDMissing, AdHealthCheckCodeCaptureRate,
		AdHealthCheckCodeMetaConversionsAPI, AdHealthCheckCodeGoogleConversionUpload,
		AdHealthCheckCodeToken, AdHealthCheckCodeTimeZoneMismatch,
		AdHealthCheckCodeCurrencyMismatch, AdHealthCheckCodeSyncStale:
		return true
	}
	return false
}

func (c AdHealthCheckCode) String() string { return string(c) }
