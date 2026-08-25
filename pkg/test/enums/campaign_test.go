package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/marketing/campaign/campaign_enums"
)

func TestCampaignEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "campaignenum.CampaignCTADestinationType", valid: []stringEnum{campaign_enums.CampaignCTADestinationProduct, campaign_enums.CampaignCTADestinationCollection, campaign_enums.CampaignCTADestinationCategory, campaign_enums.CampaignCTADestinationCart, campaign_enums.CampaignCTADestinationPromotions}, invalid: campaign_enums.CampaignCTADestinationType("__invalid__")},
		{name: "campaignenum.CampaignCustomerType", valid: []stringEnum{campaign_enums.CampaignCustomerTypeGuest, campaign_enums.CampaignCustomerTypeRetail, campaign_enums.CampaignCustomerTypeWholesale}, invalid: campaign_enums.CampaignCustomerType("__invalid__")},
		{name: "campaignenum.CampaignPlacement", valid: []stringEnum{campaign_enums.CampaignPlacementTopBanner, campaign_enums.CampaignPlacementHomeHero, campaign_enums.CampaignPlacementModal, campaign_enums.CampaignPlacementCheckoutNotice, campaign_enums.CampaignPlacementProductNotice}, invalid: campaign_enums.CampaignPlacement("__invalid__")},
		{name: "campaignenum.CampaignPlatform", valid: []stringEnum{campaign_enums.CampaignPlatformWeb, campaign_enums.CampaignPlatformMobile}, invalid: campaign_enums.CampaignPlatform("__invalid__")},
		{name: "campaignenum.CampaignSeverity", valid: []stringEnum{campaign_enums.CampaignSeverityInfo, campaign_enums.CampaignSeveritySuccess, campaign_enums.CampaignSeverityWarning, campaign_enums.CampaignSeverityCritical}, invalid: campaign_enums.CampaignSeverity("__invalid__")},
	})
}
