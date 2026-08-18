package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v29/pkg/contracts/pricing/promotion/promotion_enums"
)

func TestPromotionEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "promotion_enums.PromotionStatus", valid: []stringEnum{promotion_enums.PromotionStatusDraft, promotion_enums.PromotionStatusActive, promotion_enums.PromotionStatusInactive, promotion_enums.PromotionStatusArchived}, invalid: promotion_enums.PromotionStatus("__invalid__")},
		{name: "promotion_enums.PromotionMatchMode", valid: []stringEnum{promotion_enums.PromotionMatchModeAll, promotion_enums.PromotionMatchModeAny}, invalid: promotion_enums.PromotionMatchMode("__invalid__")},
	})
}
