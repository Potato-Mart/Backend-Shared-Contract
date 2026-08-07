package enums_test

import (
	"testing"

	favouriteenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/favourite"
)

func TestFavouriteEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{
			name: "favouriteenum.FavouriteListOwnerType",
			valid: []stringEnum{
				favouriteenum.FavouriteListOwnerTypeRetailUser,
				favouriteenum.FavouriteListOwnerTypeWholesaleOrganisation,
			},
			invalid: favouriteenum.FavouriteListOwnerType("__invalid__"),
		},
		{
			name: "favouriteenum.FavouriteListErrorCode",
			valid: []stringEnum{
				favouriteenum.FavouriteListErrorCodeLimitReached,
				favouriteenum.FavouriteListErrorCodeNameConflict,
				favouriteenum.FavouriteListErrorCodeProductLimitReached,
			},
			invalid: favouriteenum.FavouriteListErrorCode("__invalid__"),
		},
	})
}
