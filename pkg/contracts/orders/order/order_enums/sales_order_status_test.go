package order_enums

import (
	"encoding/json"
	"testing"
)

func TestExpiredSalesOrderStatus(t *testing.T) {
	status := SalesOrderStatusExpired
	if !status.IsValid() || status.String() != "expired" {
		t.Fatal(status)
	}
	payload, err := json.Marshal(status)
	if err != nil || string(payload) != `"expired"` {
		t.Fatalf("marshal = %s, %v", payload, err)
	}
	var decoded SalesOrderStatus
	if err := json.Unmarshal(payload, &decoded); err != nil || decoded != status {
		t.Fatalf("decode = %s, %v", decoded, err)
	}
	if SalesOrderStatus("EXPIRED").IsValid() {
		t.Fatal("noncanonical status accepted")
	}
}
