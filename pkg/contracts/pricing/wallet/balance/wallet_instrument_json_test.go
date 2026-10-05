package balance_test

import (
	"encoding/json"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/balance"
	"github.com/Potato-Mart/Backend-Shared-Contract/v43/pkg/contracts/pricing/wallet/wallet_enums"
	"testing"
)

func TestWalletCustomerStatusJSON(t *testing.T) {
	for _, status := range []wallet_enums.WalletInstrumentCustomerStatus{"available", "used", "expired"} {
		t.Run(string(status), func(t *testing.T) {
			wire := `{"type":"reward","code":"reward-1","customer_status":"` + string(status) + `","status":"REDEEMED"}`
			var model balance.WalletInstrument
			if err := json.Unmarshal([]byte(wire), &model); err != nil {
				t.Fatal(err)
			}
			if model.CustomerStatus != status || model.Status != "REDEEMED" {
				t.Fatal("customer and lifecycle evidence changed")
			}
			encoded, err := json.Marshal(model)
			if err != nil {
				t.Fatal(err)
			}
			if string(encoded) != wire {
				t.Fatalf("got %s want %s", encoded, wire)
			}
		})
	}
	legacy := balance.WalletInstrument{Type: wallet_enums.WalletInstrumentTypeReward, Code: "reward-1", Status: "REDEEMED"}
	wire, err := json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	if string(wire) != `{"type":"reward","code":"reward-1","status":"REDEEMED"}` {
		t.Fatalf("legacy shape changed: %s", wire)
	}
}
