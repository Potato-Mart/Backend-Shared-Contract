package enums_test

import (
	"testing"

	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v20/pkg/enums/campaign"
)

func TestCampaignEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "campaignenum.CampaignCustomerType", valid: []stringEnum{campaignenum.CampaignCustomerTypeGuest, campaignenum.CampaignCustomerTypeRetail, campaignenum.CampaignCustomerTypeWholesale}, invalid: campaignenum.CampaignCustomerType("__invalid__")},
		{name: "campaignenum.CampaignPlacement", valid: []stringEnum{campaignenum.CampaignPlacementTopBanner, campaignenum.CampaignPlacementHomeHero, campaignenum.CampaignPlacementModal, campaignenum.CampaignPlacementCheckoutNotice, campaignenum.CampaignPlacementProductNotice}, invalid: campaignenum.CampaignPlacement("__invalid__")},
		{name: "campaignenum.CampaignPlatform", valid: []stringEnum{campaignenum.CampaignPlatformWeb, campaignenum.CampaignPlatformMobile}, invalid: campaignenum.CampaignPlatform("__invalid__")},
		{name: "campaignenum.CampaignSeverity", valid: []stringEnum{campaignenum.CampaignSeverityInfo, campaignenum.CampaignSeveritySuccess, campaignenum.CampaignSeverityWarning, campaignenum.CampaignSeverityCritical}, invalid: campaignenum.CampaignSeverity("__invalid__")},
	})
}
