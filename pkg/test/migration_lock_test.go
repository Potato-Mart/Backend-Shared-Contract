package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/cart"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/orders/order"
	pos "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/payments/receipt"
	event_enums "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/pubsub/routing"
	purchase "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/procurement"
)

// TestTransactionEvidenceUsesOneImmutableSKUCode keeps one business key on
// every transaction-evidence model.
func TestTransactionEvidenceUsesOneImmutableSKUCode(t *testing.T) {
	for _, model := range []reflect.Type{
		reflect.TypeOf(cart.CartItem{}),
		reflect.TypeOf(order.OrderItem{}),
		reflect.TypeOf(order.OrderLineSummary{}),
		reflect.TypeOf(pos.ReceiptLine{}),
		reflect.TypeOf(purchase.PurchaseOrderItem{}),
		reflect.TypeOf(purchase.PurchaseReceiptItem{}),
		reflect.TypeOf(purchase.SupplierInvoiceLine{}),
	} {
		field, ok := model.FieldByName("SKUCode")
		if !ok || field.Type.Kind() != reflect.String || field.Tag.Get("json") != "sku_code" {
			t.Errorf("%s must carry exactly one scalar sku_code", model)
		}
	}
}

// TestCatalogEventTopologyIsDeclared locks catalog-events routing
// vocabulary so the contract, the cloud topology, and the local emulator script
// can be compared against one source of truth.
func TestCatalogEventTopologyIsDeclared(t *testing.T) {
	if got := event_enums.EventTopicCatalogEvents; got.String() != "catalog-events" || !got.IsValid() {
		t.Fatalf("catalog topic = %q valid=%v", got, got.IsValid())
	}
	for _, eventType := range []event_enums.EventType{
		event_enums.EventTypeCatalogBaseCostChanged,
		event_enums.EventTypeCatalogListingChanged,
	} {
		if !eventType.IsValid() {
			t.Fatalf("catalog event type %q must validate", eventType)
		}
		if !strings.HasPrefix(eventType.String(), "catalog.") {
			t.Fatalf("catalog event type %q must be namespaced under catalog.", eventType)
		}
	}

	topics := map[string]struct{}{}
	for _, topic := range []event_enums.EventTopic{
		event_enums.EventTopicOrderEvents,
		event_enums.EventTopicPaymentEvents,
		event_enums.EventTopicRefundEvents,
		event_enums.EventTopicStockEvents,
		event_enums.EventTopicFulfilmentEvents,
		event_enums.EventTopicCustomerEvents,
		event_enums.EventTopicProductStats,
		event_enums.EventTopicStorefrontEvents,
		event_enums.EventTopicCatalogEvents,
	} {
		if !topic.IsValid() {
			t.Fatalf("canonical topic %q must validate", topic)
		}
		topics[topic.String()] = struct{}{}
	}
	if len(topics) != 9 {
		t.Fatalf("canonical topic count = %d, want 9", len(topics))
	}
}

// TestMonetaryFieldsUseMinorUnitIntegers keeps every money-shaped field an
// exact integer. A float in a price, cost, tax, or amount field would silently
// drift across JSON and service boundaries.
func TestMonetaryFieldsUseMinorUnitIntegers(t *testing.T) {
	monetaryFragments := []string{"amount", "_minor", "price", "cost", "tax"}
	pkgRoot := sharedContractPkgRoot(t)
	var violations []string

	err := filepath.WalkDir(pkgRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			return nil
		}
		fset := token.NewFileSet()
		file, parseErr := parser.ParseFile(fset, path, nil, 0)
		if parseErr != nil {
			return parseErr
		}
		ast.Inspect(file, func(node ast.Node) bool {
			field, ok := node.(*ast.Field)
			if !ok {
				return true
			}
			jsonKey, present := jsonFieldName(t, path, field)
			if !present {
				return true
			}
			monetary := false
			for _, fragment := range monetaryFragments {
				if strings.Contains(jsonKey, fragment) {
					monetary = true
					break
				}
			}
			if !monetary {
				return true
			}
			ast.Inspect(field.Type, func(inner ast.Node) bool {
				identifier, ok := inner.(*ast.Ident)
				if ok && (identifier.Name == "float32" || identifier.Name == "float64") {
					violations = append(violations, fset.Position(field.Pos()).String()+": monetary field "+jsonKey+" uses "+identifier.Name)
				}
				return true
			})
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("scan monetary fields: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("monetary precision violations:\n%s", strings.Join(violations, "\n"))
	}
}
