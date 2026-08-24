package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v30/pkg/contracts/pricing/pricebook/pricebook_enums"
)

func TestPriceBookEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "pricebookenum.PriceBookStatus", valid: []stringEnum{pricebook_enums.PriceBookStatusDraft, pricebook_enums.PriceBookStatusActive, pricebook_enums.PriceBookStatusInactive, pricebook_enums.PriceBookStatusArchived}, invalid: pricebook_enums.PriceBookStatus("__invalid__")},
		{name: "pricebookenum.PriceEntryStatus", valid: []stringEnum{pricebook_enums.PriceEntryStatusDraft, pricebook_enums.PriceEntryStatusPendingApproval, pricebook_enums.PriceEntryStatusApproved, pricebook_enums.PriceEntryStatusRejected, pricebook_enums.PriceEntryStatusSuperseded, pricebook_enums.PriceEntryStatusWithdrawn, pricebook_enums.PriceEntryStatusExpired}, invalid: pricebook_enums.PriceEntryStatus("__invalid__")},
		{name: "pricebookenum.PriceTaxInclusion", valid: []stringEnum{pricebook_enums.PriceTaxInclusionInclusive, pricebook_enums.PriceTaxInclusionExclusive}, invalid: pricebook_enums.PriceTaxInclusion("__invalid__")},
		{name: "pricebookenum.PriceEndingPolicy", valid: []stringEnum{pricebook_enums.PriceEndingPolicyNone, pricebook_enums.PriceEndingPolicyCharmNine}, invalid: pricebook_enums.PriceEndingPolicy("__invalid__")},
		{name: "pricebookenum.PriceBookAssignmentKind", valid: []stringEnum{pricebook_enums.PriceBookAssignmentKindChannelDefault, pricebook_enums.PriceBookAssignmentKindOrganisationCategory, pricebook_enums.PriceBookAssignmentKindOrganisationOverride}, invalid: pricebook_enums.PriceBookAssignmentKind("__invalid__")},
		{name: "pricebookenum.PriceDerivation", valid: []stringEnum{pricebook_enums.PriceDerivationManual, pricebook_enums.PriceDerivationSuggestedFromBaseCost, pricebook_enums.PriceDerivationImported}, invalid: pricebook_enums.PriceDerivation("__invalid__")},
	})
}
