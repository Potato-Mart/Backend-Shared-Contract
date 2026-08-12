package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/classification/classification_enums"
)

func TestFavouriteEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "classificationenum.FavouriteListOwnerType",
			valid: []stringEnum{
				classification_enums.FavouriteListOwnerTypeRetailUser,
				classification_enums.FavouriteListOwnerTypeWholesaleOrganisation,
			},
			invalid: classification_enums.FavouriteListOwnerType("__invalid__"),
		},
		{
			name: "classificationenum.FavouriteListErrorCode",
			valid: []stringEnum{
				classification_enums.FavouriteListErrorCodeLimitReached,
				classification_enums.FavouriteListErrorCodeNameConflict,
				classification_enums.FavouriteListErrorCodeProductLimitReached,
			},
			invalid: classification_enums.FavouriteListErrorCode("__invalid__"),
		},
	})
}
