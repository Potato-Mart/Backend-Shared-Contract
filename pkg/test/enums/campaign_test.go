package enums_test

import (
	ae "github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/marketing/audience/audience_enums"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v46/pkg/contracts/marketing/campaign/campaign_enums"
)

func TestCampaignEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "audience.CustomerType", valid: []stringEnum{ae.CustomerTypeGuest, ae.CustomerTypeRetail, ae.CustomerTypeWholesale}, invalid: ae.CustomerType("__invalid__")},
		{name: "audience.Platform", valid: []stringEnum{ae.PlatformWeb, ae.PlatformMobile}, invalid: ae.Platform("__invalid__")},
		{name: "campaignenum.CampaignCustomerType", valid: []stringEnum{campaign_enums.CampaignCustomerTypeRetail, campaign_enums.CampaignCustomerTypeWholesale}, invalid: campaign_enums.CampaignCustomerType("__invalid__")},
		{name: "campaignenum.CampaignPlacement", valid: []stringEnum{campaign_enums.CampaignPlacementAnnouncementBar, campaign_enums.CampaignPlacementHomeBanner, campaign_enums.CampaignPlacementAccountBanner, campaign_enums.CampaignPlacementHomeModal}, invalid: campaign_enums.CampaignPlacement("__invalid__")},
		{name: "campaignenum.CampaignPlatform", valid: []stringEnum{campaign_enums.CampaignPlatformWeb, campaign_enums.CampaignPlatformMobile}, invalid: campaign_enums.CampaignPlatform("__invalid__")},
	})
}
