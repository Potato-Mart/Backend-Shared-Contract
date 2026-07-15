package warehouse_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/warehouse"
)

func TestLayoutNodeUsesBusinessKeys(t *testing.T) {
	body, err := json.Marshal(warehouse.LayoutNode{
		ID: "node-storage-id", DepotCode: "MEL", LayoutVersion: 3,
		Code: "A-1", ParentCode: "ZONE-A", PathCodes: []string{"ZONE-A"}, LocationCode: "BIN-A-1",
	})
	if err != nil {
		t.Fatalf("marshal layout node: %v", err)
	}
	for _, want := range []string{`"depot_code":"MEL"`, `"layout_version":3`, `"parent_code":"ZONE-A"`, `"path_codes":["ZONE-A"]`, `"location_code":"BIN-A-1"`} {
		if !strings.Contains(string(body), want) {
			t.Fatalf("layout node JSON missing %s: %s", want, body)
		}
	}
	for _, removed := range []string{`"layout_id"`, `"parent_id"`, `"path"`, `"location_id"`} {
		if strings.Contains(string(body), removed) {
			t.Fatalf("layout node JSON retained %s: %s", removed, body)
		}
	}
}
