package membership_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/membership/membership_enums"
)

// V53.54 adds the optional tier chip colour token to the tier progress
// summary. Progress written before the release reads back unchanged.
func TestV5354TierProgressTierDisplayColour(t *testing.T) {
	field, ok := reflect.TypeOf(membership.TierProgressTier{}).FieldByName("DisplayColour")
	if !ok || field.Type != reflect.TypeOf(membership_enums.TierDisplayColour("")) || field.Tag.Get("json") != "display_colour,omitempty" {
		t.Fatalf("DisplayColour = %+v, want membership_enums.TierDisplayColour `json:\"display_colour,omitempty\"`", field)
	}

	tier := membership.TierProgressTier{TierKey: "gold", DisplayColour: membership_enums.TierDisplayColourGold}
	encoded, err := json.Marshal(tier)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `{"tier_key":"gold","display_colour":"gold"}` {
		t.Fatalf("unexpected wire shape: %s", encoded)
	}
	var decoded membership.TierProgressTier
	if err := json.Unmarshal(encoded, &decoded); err != nil || !reflect.DeepEqual(decoded, tier) {
		t.Fatalf("round trip changed tier: %v %#v", err, decoded)
	}

	var legacy membership.TierProgressTier
	if err := json.Unmarshal([]byte(`{"tier_key":"silver"}`), &legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.DisplayColour != "" {
		t.Fatalf("legacy tier gained a colour: %q", legacy.DisplayColour)
	}
	empty, err := json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	if string(empty) != `{"tier_key":"silver"}` {
		t.Fatalf("absent colour must be omitted: %s", empty)
	}
}
