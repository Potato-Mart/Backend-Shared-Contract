package enums_test

import (
	"testing"

	classification_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/catalogue/favourite/favourite_enums"
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
