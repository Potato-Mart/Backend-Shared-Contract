package order_test

import (
	"encoding/json"

	sales "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/commerce/commerce_enums"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/customers/retail/retail_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/orders/order/order_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/product/product_enums"
)

// TestOrderBuyerAndItemPricingRoundTrip checks the additive buyer/commercial
// context round-trips on a sales order and its line item.
func TestOrderBuyerAndItemPricingRoundTrip(t *testing.T) {
	capturedAt := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	order := sales.Order{
		ID:          "ord_b2b_1",
		OrderNumber: "B2B-1001",
		Channel:     commerce_enums.OrderTypeB2B,
		Buyer: &sales.BuyerContext{
			Type:                      retail_enums.BuyerTypeWholesaleOrganisation,
			WholesaleOrganisationCode: "org_1",
			OrganisationAccessID:      "oacc_1",
		},
		Items: []sales.OrderItem{
			{
				Components:     []sales.PricedPackageComponent{{RequestedPackageCount: 10, RequestedBaseUnits: 10, PackagePrice: money.Money{AmountMinor: 8000, Currency: "AUD"}}},
				TotalBaseUnits: 10,
				Pricing: &sales.PricingContext{
					Audience:   product_enums.PriceAudienceWholesale,
					Visibility: product_enums.PriceVisibilityWholesaleApprovedOnly,
				},
				SubstitutionPolicy: sales.LooseSubstitutionPolicySnapshot{Allowed: true, Source: order_enums.LooseSubstitutionPolicySourceChannelDefault, CapturedAt: capturedAt},
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

	if decoded.Channel != commerce_enums.OrderTypeB2B {
		t.Fatalf("channel = %q, want b2b", decoded.Channel)
	}
	if decoded.Buyer == nil {
		t.Fatalf("buyer did not round-trip: %s", payload)
	}
	if decoded.Buyer.Type != retail_enums.BuyerTypeWholesaleOrganisation {
		t.Fatalf("buyer.type = %q, want wholesale_organisation", decoded.Buyer.Type)
	}
	if decoded.Buyer.WholesaleOrganisationCode != "org_1" || decoded.Buyer.OrganisationAccessID != "oacc_1" {
		t.Fatalf("buyer org references did not round-trip: %+v", decoded.Buyer)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].Pricing == nil || decoded.Items[0].Pricing.Audience != product_enums.PriceAudienceWholesale {
		t.Fatalf("item pricing.audience did not round-trip: %+v", decoded.Items)
	}
}

// TestCartChannelAndBuyerRoundTrip checks a POS walk-in cart round-trips.
func TestCartChannelAndBuyerRoundTrip(t *testing.T) {
	capturedAt := time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC)
	cart := sales.Cart{
		ID:        "cart_1",
		SessionID: "sess_1",
		Channel:   commerce_enums.OrderTypePOS,
		Buyer:     &sales.BuyerContext{Type: retail_enums.BuyerTypeGuestRetail},
		Items: []sales.CartItem{
			{
				Components:         []sales.PricedPackageComponent{{RequestedPackageCount: 1, RequestedBaseUnits: 1, PackagePrice: money.Money{AmountMinor: 500, Currency: "AUD"}}},
				TotalBaseUnits:     1,
				Pricing:            &sales.PricingContext{Audience: product_enums.PriceAudienceRetail},
				SubstitutionPolicy: sales.LooseSubstitutionPolicySnapshot{Allowed: false, Source: order_enums.LooseSubstitutionPolicySourceBuyerSelected, CapturedAt: capturedAt},
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
	if decoded.Channel != commerce_enums.OrderTypePOS {
		t.Fatalf("cart channel = %q, want pos", decoded.Channel)
	}
	if decoded.Buyer == nil || decoded.Buyer.Type != retail_enums.BuyerTypeGuestRetail {
		t.Fatalf("cart buyer.type did not round-trip: %+v", decoded.Buyer)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].Pricing == nil || decoded.Items[0].Pricing.Audience != product_enums.PriceAudienceRetail {
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
			order:       sales.Order{Channel: commerce_enums.OrderTypeB2B, Buyer: &sales.BuyerContext{Type: retail_enums.BuyerTypeWholesaleOrganisation}},
			wantChannel: "b2b",
			wantBuyer:   "wholesale_organisation",
		},
		{
			name: "wholesale customer in physical shop",
			order: sales.Order{
				Channel: commerce_enums.OrderTypePOS,
				Buyer:   &sales.BuyerContext{Type: retail_enums.BuyerTypeWholesaleOrganisation},
				Items:   []sales.OrderItem{{Pricing: &sales.PricingContext{Audience: product_enums.PriceAudienceWholesale}}},
			},
			wantChannel:  "pos",
			wantBuyer:    "wholesale_organisation",
			wantAudience: "wholesale",
		},
		{
			name:        "walk-in normal customer",
			order:       sales.Order{Channel: commerce_enums.OrderTypePOS, Buyer: &sales.BuyerContext{Type: retail_enums.BuyerTypeGuestRetail}},
			wantChannel: "pos",
			wantBuyer:   "guest_retail",
		},
		{
			name:        "retail website member",
			order:       sales.Order{Channel: commerce_enums.OrderTypeOnline, Buyer: &sales.BuyerContext{Type: retail_enums.BuyerTypeRetailCustomer}},
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
