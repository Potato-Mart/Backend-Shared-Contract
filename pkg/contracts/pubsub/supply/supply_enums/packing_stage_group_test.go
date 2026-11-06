package supply_enums

import "testing"

func TestPackingStageGroup(t *testing.T) {
	for _, v := range []PackingStageGroup{PackingStageGroupAmbientChilled, PackingStageGroupFrozen} {
		if !v.IsValid() || v.String() != string(v) {
			t.Fatalf("invalid stage group %q", v)
		}
	}
	for _, v := range []PackingStageGroup{"", "AMBIENT", "CHILLED", "frozen"} {
		if v.IsValid() {
			t.Fatalf("unexpected valid stage group %q", v)
		}
	}
}
