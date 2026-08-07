package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v24/pkg/contracts/supply/favourite/favourite_enums"
)

func TestFavouriteEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "favouriteenum.FavouriteListOwnerType",
			valid: []stringEnum{
				favourite_enums.FavouriteListOwnerTypeRetailUser,
				favourite_enums.FavouriteListOwnerTypeWholesaleOrganisation,
			},
			invalid: favourite_enums.FavouriteListOwnerType("__invalid__"),
		},
		{
			name: "favouriteenum.FavouriteListErrorCode",
			valid: []stringEnum{
				favourite_enums.FavouriteListErrorCodeLimitReached,
				favourite_enums.FavouriteListErrorCodeNameConflict,
				favourite_enums.FavouriteListErrorCodeProductLimitReached,
			},
			invalid: favourite_enums.FavouriteListErrorCode("__invalid__"),
		},
	})
}
