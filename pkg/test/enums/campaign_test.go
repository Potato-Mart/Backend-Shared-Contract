package enums_test

import (
	"testing"

	campaignenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/customers/campaign"
)

func TestCampaignEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "campaignenum.CampaignCTADestinationType", valid: []stringEnum{campaignenum.CampaignCTADestinationProduct, campaignenum.CampaignCTADestinationCollection, campaignenum.CampaignCTADestinationCategory, campaignenum.CampaignCTADestinationCart, campaignenum.CampaignCTADestinationPromotions}, invalid: campaignenum.CampaignCTADestinationType("__invalid__")},
		{name: "campaignenum.CampaignCustomerType", valid: []stringEnum{campaignenum.CampaignCustomerTypeGuest, campaignenum.CampaignCustomerTypeRetail, campaignenum.CampaignCustomerTypeWholesale}, invalid: campaignenum.CampaignCustomerType("__invalid__")},
		{name: "campaignenum.CampaignPlacement", valid: []stringEnum{campaignenum.CampaignPlacementTopBanner, campaignenum.CampaignPlacementHomeHero, campaignenum.CampaignPlacementModal, campaignenum.CampaignPlacementCheckoutNotice, campaignenum.CampaignPlacementProductNotice}, invalid: campaignenum.CampaignPlacement("__invalid__")},
		{name: "campaignenum.CampaignPlatform", valid: []stringEnum{campaignenum.CampaignPlatformWeb, campaignenum.CampaignPlatformMobile}, invalid: campaignenum.CampaignPlatform("__invalid__")},
		{name: "campaignenum.CampaignSeverity", valid: []stringEnum{campaignenum.CampaignSeverityInfo, campaignenum.CampaignSeveritySuccess, campaignenum.CampaignSeverityWarning, campaignenum.CampaignSeverityCritical}, invalid: campaignenum.CampaignSeverity("__invalid__")},
	})
}
