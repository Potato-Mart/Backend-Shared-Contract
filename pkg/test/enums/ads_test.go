package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/insights/ads/ads_enums"
)

func TestPaidAdsEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "adsenum.AdPlatform", valid: []stringEnum{ads_enums.AdPlatformMeta, ads_enums.AdPlatformGoogle}, invalid: ads_enums.AdPlatform("__invalid__")},
		{name: "adsenum.AdLevel", valid: []stringEnum{ads_enums.AdLevelCampaign, ads_enums.AdLevelAdSet, ads_enums.AdLevelAd}, invalid: ads_enums.AdLevel("__invalid__")},
		{name: "adsenum.AdEntityStatus", valid: []stringEnum{ads_enums.AdEntityStatusActive, ads_enums.AdEntityStatusPaused, ads_enums.AdEntityStatusArchived, ads_enums.AdEntityStatusDeleted}, invalid: ads_enums.AdEntityStatus("__invalid__")},
		{name: "adsenum.AdReviewStatus", valid: []stringEnum{ads_enums.AdReviewStatusPending, ads_enums.AdReviewStatusApproved, ads_enums.AdReviewStatusRejected}, invalid: ads_enums.AdReviewStatus("__invalid__")},
		{name: "adsenum.AdBidStrategy", valid: []stringEnum{ads_enums.AdBidStrategyHighestVolume, ads_enums.AdBidStrategyCostPerActionTarget, ads_enums.AdBidStrategyBidCap, ads_enums.AdBidStrategyMinimumROAS, ads_enums.AdBidStrategyMaximiseConversions, ads_enums.AdBidStrategyMaximiseConversionValue}, invalid: ads_enums.AdBidStrategy("__invalid__")},
		{name: "adsenum.AdChangeStatus", valid: []stringEnum{ads_enums.AdChangeStatusPending, ads_enums.AdChangeStatusApplied, ads_enums.AdChangeStatusFailed, ads_enums.AdChangeStatusConflict}, invalid: ads_enums.AdChangeStatus("__invalid__")},
		{name: "adsenum.AdChangeSource", valid: []stringEnum{ads_enums.AdChangeSourceAdmin, ads_enums.AdChangeSourcePlatform}, invalid: ads_enums.AdChangeSource("__invalid__")},
		{name: "adsenum.AdDraftStatus", valid: []stringEnum{ads_enums.AdDraftStatusInReview, ads_enums.AdDraftStatusApproved, ads_enums.AdDraftStatusRejected, ads_enums.AdDraftStatusActive}, invalid: ads_enums.AdDraftStatus("__invalid__")},
		{name: "adsenum.AdHealthSeverity", valid: []stringEnum{ads_enums.AdHealthSeverityOK, ads_enums.AdHealthSeverityWarning, ads_enums.AdHealthSeverityError}, invalid: ads_enums.AdHealthSeverity("__invalid__")},
		{name: "adsenum.AdHealthCheckCode", valid: []stringEnum{ads_enums.AdHealthCheckCodeAdMissingTracking, ads_enums.AdHealthCheckCodeConversionGap, ads_enums.AdHealthCheckCodeClickIDMissing, ads_enums.AdHealthCheckCodeCaptureRate, ads_enums.AdHealthCheckCodeMetaConversionsAPI, ads_enums.AdHealthCheckCodeGoogleConversionUpload, ads_enums.AdHealthCheckCodeToken, ads_enums.AdHealthCheckCodeTimeZoneMismatch, ads_enums.AdHealthCheckCodeCurrencyMismatch, ads_enums.AdHealthCheckCodeSyncStale}, invalid: ads_enums.AdHealthCheckCode("__invalid__")},
		{name: "adsenum.AdPostIneligibilityReason", valid: []stringEnum{ads_enums.AdPostIneligibilityReasonLicensedMusic, ads_enums.AdPostIneligibilityReasonSharedContent, ads_enums.AdPostIneligibilityReasonUnsupportedType}, invalid: ads_enums.AdPostIneligibilityReason("__invalid__")},
		{name: "adsenum.AdPostType", valid: []stringEnum{ads_enums.AdPostTypeImage, ads_enums.AdPostTypeCarousel, ads_enums.AdPostTypeVideo, ads_enums.AdPostTypeReel}, invalid: ads_enums.AdPostType("__invalid__")},
		{name: "adsenum.AdCallToAction", valid: []stringEnum{ads_enums.AdCallToActionShopNow, ads_enums.AdCallToActionLearnMore, ads_enums.AdCallToActionOrderNow, ads_enums.AdCallToActionGetOffer, ads_enums.AdCallToActionSignUp, ads_enums.AdCallToActionSeeMore}, invalid: ads_enums.AdCallToAction("__invalid__")},
		{name: "adsenum.AdLandingKind", valid: []stringEnum{ads_enums.AdLandingKindCampaign, ads_enums.AdLandingKindProduct, ads_enums.AdLandingKindSelection, ads_enums.AdLandingKindURL}, invalid: ads_enums.AdLandingKind("__invalid__")},
		{name: "adsenum.AdAttributionModel", valid: []stringEnum{ads_enums.AdAttributionModelLastClick, ads_enums.AdAttributionModelFirstClick}, invalid: ads_enums.AdAttributionModel("__invalid__")},
		{name: "adsenum.AdConsentState", valid: []stringEnum{ads_enums.AdConsentStateGranted, ads_enums.AdConsentStateDenied}, invalid: ads_enums.AdConsentState("__invalid__")},
		{name: "adsenum.AdSourcePlatform", valid: []stringEnum{ads_enums.AdSourcePlatformWeb, ads_enums.AdSourcePlatformIOS, ads_enums.AdSourcePlatformAndroid}, invalid: ads_enums.AdSourcePlatform("__invalid__")},
		{name: "adsenum.AdAppTrackingStatus", valid: []stringEnum{ads_enums.AdAppTrackingStatusAuthorized, ads_enums.AdAppTrackingStatusDenied, ads_enums.AdAppTrackingStatusRestricted, ads_enums.AdAppTrackingStatusNotDetermined}, invalid: ads_enums.AdAppTrackingStatus("__invalid__")},
		{name: "adsenum.AdCreativeSource", valid: []stringEnum{ads_enums.AdCreativeSourceMedia, ads_enums.AdCreativeSourcePagePost}, invalid: ads_enums.AdCreativeSource("__invalid__")},
		{name: "adsenum.AdPlacementMode", valid: []stringEnum{ads_enums.AdPlacementModeNew, ads_enums.AdPlacementModeExisting}, invalid: ads_enums.AdPlacementMode("__invalid__")},
		{name: "adsenum.AdConnectionStatus", valid: []stringEnum{ads_enums.AdConnectionStatusConnected, ads_enums.AdConnectionStatusNotConfigured, ads_enums.AdConnectionStatusNeedsAttention}, invalid: ads_enums.AdConnectionStatus("__invalid__")},
		{name: "adsenum.AdCredentialStatus", valid: []stringEnum{ads_enums.AdCredentialStatusOK, ads_enums.AdCredentialStatusExpiring, ads_enums.AdCredentialStatusInvalid}, invalid: ads_enums.AdCredentialStatus("__invalid__")},
	})
}

// TestPaidAdsEnumWireValues locks the wire strings the storefront, the admin
// and both advertising platforms already agreed on, so a later rename cannot
// silently change what is sent or stored.
func TestPaidAdsEnumWireValues(t *testing.T) {
	for _, tc := range []struct {
		got  string
		want string
	}{
		{string(ads_enums.AdPlatformMeta), "meta"},
		{string(ads_enums.AdPlatformGoogle), "google"},
		{string(ads_enums.AdLevelAdSet), "ad_set"},
		{string(ads_enums.AdChangeStatusConflict), "conflict"},
		{string(ads_enums.AdChangeSourcePlatform), "platform"},
		{string(ads_enums.AdDraftStatusInReview), "in_review"},
		{string(ads_enums.AdHealthSeverityOK), "ok"},
		{string(ads_enums.AdHealthCheckCodeAdMissingTracking), "ad_missing_tracking"},
		{string(ads_enums.AdPostIneligibilityReasonLicensedMusic), "licensed_music"},
		{string(ads_enums.AdPostTypeReel), "reel"},
		{string(ads_enums.AdCallToActionShopNow), "shop_now"},
		{string(ads_enums.AdLandingKindURL), "url"},
		{string(ads_enums.AdAttributionModelLastClick), "last_click"},
		{string(ads_enums.AdConsentStateGranted), "granted"},
		{string(ads_enums.AdSourcePlatformIOS), "ios"},
		{string(ads_enums.AdAppTrackingStatusNotDetermined), "not_determined"},
		{string(ads_enums.AdCreativeSourcePagePost), "page_post"},
		{string(ads_enums.AdPlacementModeExisting), "existing"},
		{string(ads_enums.AdConnectionStatusNotConfigured), "not_configured"},
		{string(ads_enums.AdCredentialStatusExpiring), "expiring"},
	} {
		if tc.got != tc.want {
			t.Errorf("paid-ads enum wire value = %q, want %q", tc.got, tc.want)
		}
	}

	// Consent is denied by default, so the empty string is never a granted
	// state and never validates.
	if ads_enums.AdConsentState("").IsValid() {
		t.Error("an absent advertising consent answer must not validate as consent")
	}
	// Only Meta and Google are advertising platforms in this contract.
	for _, candidate := range []ads_enums.AdPlatform{"other", "tiktok", "Meta", "GOOGLE"} {
		if candidate.IsValid() {
			t.Errorf("AdPlatform(%q) must not validate", candidate)
		}
	}
}
