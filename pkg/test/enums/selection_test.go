package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v52/pkg/contracts/supply/catalogue/selection/selection_enums"
)

func TestSelectionEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "selectionenum.ProductSelectionKind", valid: []stringEnum{selection_enums.ProductSelectionKindBadge, selection_enums.ProductSelectionKindSection}, invalid: selection_enums.ProductSelectionKind("__invalid__")},
		{name: "selectionenum.ProductSelectionStatus", valid: []stringEnum{selection_enums.ProductSelectionStatusDraft, selection_enums.ProductSelectionStatusPublished, selection_enums.ProductSelectionStatusArchived}, invalid: selection_enums.ProductSelectionStatus("__invalid__")},
	})
}
