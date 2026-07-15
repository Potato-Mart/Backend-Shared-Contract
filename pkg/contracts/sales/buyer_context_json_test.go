package sales_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/contracts/sales"
	customerenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/customer"
	productenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/product"
	salesenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/sales"
	shippingenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/shipping"
)

// TestOrderBuyerAndItemPricingRoundTrip checks the additive buyer/commercial
// context round-trips on a sales order and its line item.
func TestOrderBuyerAndItemPricingRoundTrip(t *testing.T) {
	order := sales.Order{
		ID:          "ord_b2b_1",
		OrderNumber: "B2B-1001",
		Channel:     salesenum.OrderTypeB2B,
		Buyer: &sales.BuyerContext{
			Type:                      customerenum.BuyerTypeWholesaleOrganisation,
			WholesaleOrganisationCode: "org_1",
			OrganisationAccessID:      "oacc_1",
			FulfilmentIntent:          shippingenum.FulfilmentIntentDelivery,
		},
		Items: []sales.OrderItem{
			{
				UnitPrice: common.Money{AmountMinor: 8000, Currency: "AUD"},
				Pricing: &sales.PricingContext{
					Audience:   productenum.PriceAudienceWholesale,
					Visibility: productenum.PriceVisibilityWholesaleApprovedOnly,
				},
				Quantity: 10,
			},
		},
	}

	payload, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}

	var decoded sales.Order
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal order: %v", err)
	}

	if decoded.Channel != salesenum.OrderTypeB2B {
		t.Fatalf("channel = %q, want b2b", decoded.Channel)
	}
	if decoded.Buyer == nil {
		t.Fatalf("buyer did not round-trip: %s", payload)
	}
	if decoded.Buyer.Type != customerenum.BuyerTypeWholesaleOrganisation {
		t.Fatalf("buyer.type = %q, want wholesale_organisation", decoded.Buyer.Type)
	}
	if decoded.Buyer.WholesaleOrganisationCode != "org_1" || decoded.Buyer.OrganisationAccessID != "oacc_1" {
		t.Fatalf("buyer org references did not round-trip: %+v", decoded.Buyer)
	}
	if decoded.Buyer.FulfilmentIntent != shippingenum.FulfilmentIntentDelivery {
		t.Fatalf("buyer.fulfilment_intent = %q, want delivery", decoded.Buyer.FulfilmentIntent)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].Pricing == nil || decoded.Items[0].Pricing.Audience != productenum.PriceAudienceWholesale {
		t.Fatalf("item pricing.audience did not round-trip: %+v", decoded.Items)
	}
}

// TestCartChannelAndBuyerRoundTrip checks a POS walk-in cart round-trips.
func TestCartChannelAndBuyerRoundTrip(t *testing.T) {
	cart := sales.Cart{
		ID:        "cart_1",
		SessionID: "sess_1",
		Channel:   salesenum.OrderTypePOS,
		Buyer:     &sales.BuyerContext{Type: customerenum.BuyerTypeGuestRetail},
		Items: []sales.CartItem{
			{
				Price:   common.Money{AmountMinor: 500, Currency: "AUD"},
				Pricing: &sales.PricingContext{Audience: productenum.PriceAudienceRetail},
			},
		},
	}

	payload, err := json.Marshal(cart)
	if err != nil {
		t.Fatalf("marshal cart: %v", err)
	}

	var decoded sales.Cart
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal cart: %v", err)
	}
	if decoded.Channel != salesenum.OrderTypePOS {
		t.Fatalf("cart channel = %q, want pos", decoded.Channel)
	}
	if decoded.Buyer == nil || decoded.Buyer.Type != customerenum.BuyerTypeGuestRetail {
		t.Fatalf("cart buyer.type did not round-trip: %+v", decoded.Buyer)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].Pricing == nil || decoded.Items[0].Pricing.Audience != productenum.PriceAudienceRetail {
		t.Fatalf("cart item pricing.audience did not round-trip: %+v", decoded.Items)
	}
	if strings.Contains(string(payload), `"properties"`) || strings.Contains(string(payload), `"fulfilment_mode"`) {
		t.Fatalf("removed cart item properties leaked into JSON: %s", payload)
	}
}

// TestOrderOmitsEmptyBuyer verifies the additive fields stay out of the
// JSON when unset.
func TestOrderOmitsEmptyBuyer(t *testing.T) {
	payload, err := json.Marshal(sales.Order{ID: "ord_1"})
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}
	if strings.Contains(string(payload), `"buyer"`) {
		t.Fatalf("empty buyer should be omitted, got %s", payload)
	}
}

// TestBuyerCommercialTargetMappings asserts the four documented sales
// scenarios serialise to the expected channel / buyer.type / pricing.audience
// combinations. POS is a channel, never a buyer type.
func TestBuyerCommercialTargetMappings(t *testing.T) {
	tests := []struct {
		name         string
		order        sales.Order
		wantChannel  string
		wantBuyer    string
		wantAudience string // checked on items[0].pricing.audience when non-empty
	}{
		{
			name:        "wholesale portal order",
			order:       sales.Order{Channel: salesenum.OrderTypeB2B, Buyer: &sales.BuyerContext{Type: customerenum.BuyerTypeWholesaleOrganisation}},
			wantChannel: "b2b",
			wantBuyer:   "wholesale_organisation",
		},
		{
			name: "wholesale customer in physical shop",
			order: sales.Order{
				Channel: salesenum.OrderTypePOS,
				Buyer:   &sales.BuyerContext{Type: customerenum.BuyerTypeWholesaleOrganisation},
				Items:   []sales.OrderItem{{Pricing: &sales.PricingContext{Audience: productenum.PriceAudienceWholesale}}},
			},
			wantChannel:  "pos",
			wantBuyer:    "wholesale_organisation",
			wantAudience: "wholesale",
		},
		{
			name:        "walk-in normal customer",
			order:       sales.Order{Channel: salesenum.OrderTypePOS, Buyer: &sales.BuyerContext{Type: customerenum.BuyerTypeGuestRetail}},
			wantChannel: "pos",
			wantBuyer:   "guest_retail",
		},
		{
			name:        "retail website member",
			order:       sales.Order{Channel: salesenum.OrderTypeOnline, Buyer: &sales.BuyerContext{Type: customerenum.BuyerTypeRetailCustomer}},
			wantChannel: "online",
			wantBuyer:   "retail_customer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := json.Marshal(tt.order)
			if err != nil {
				t.Fatalf("marshal order: %v", err)
			}
			var got map[string]any
			if err := json.Unmarshal(payload, &got); err != nil {
				t.Fatalf("unmarshal order JSON: %v", err)
			}
			if got["channel"] != tt.wantChannel {
				t.Fatalf("channel = %v, want %q (%s)", got["channel"], tt.wantChannel, payload)
			}
			buyer, ok := got["buyer"].(map[string]any)
			if !ok {
				t.Fatalf("buyer missing or not an object: %s", payload)
			}
			if buyer["type"] != tt.wantBuyer {
				t.Fatalf("buyer.type = %v, want %q (%s)", buyer["type"], tt.wantBuyer, payload)
			}
			if tt.wantAudience != "" {
				items, ok := got["items"].([]any)
				if !ok || len(items) == 0 {
					t.Fatalf("items missing: %s", payload)
				}
				item, _ := items[0].(map[string]any)
				pricing, ok := item["pricing"].(map[string]any)
				if !ok {
					t.Fatalf("item pricing missing: %s", payload)
				}
				if pricing["audience"] != tt.wantAudience {
					t.Fatalf("item pricing.audience = %v, want %q (%s)", pricing["audience"], tt.wantAudience, payload)
				}
			}
		})
	}
}
