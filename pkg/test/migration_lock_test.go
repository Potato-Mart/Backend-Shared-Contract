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

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/order"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/orders/pos"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/pubsub/event/event_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/supply/purchase"
)

// TestV27FrozenSKUCodeAllowlistIsExactlyTheTransactionEvidenceTypes freezes the
// allowlist itself. Widening it is a contract decision, not an implementation
// detail: every other model links to a SKU by sku_id only.
func TestV27FrozenSKUCodeAllowlistIsExactlyTheTransactionEvidenceTypes(t *testing.T) {
	want := []string{
		"orders/order.CartItem",
		"orders/order.OrderItem",
		"orders/order.OrderLineSummary",
		"orders/pos.ReceiptLine",
		"supply/purchase.OrderItem",
		"supply/purchase.ReceiptItem",
		"supply/purchase.SupplierInvoiceLine",
	}
	got := make([]string, 0, len(v27FrozenSKUCodeTypes))
	for typeKey := range v27FrozenSKUCodeTypes {
		got = append(got, typeKey)
	}
	sort.Strings(got)
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("frozen sku_code allowlist = %v, want %v", got, want)
	}

	for _, model := range []reflect.Type{
		reflect.TypeOf(order.CartItem{}),
		reflect.TypeOf(order.OrderItem{}),
		reflect.TypeOf(order.OrderLineSummary{}),
		reflect.TypeOf(pos.ReceiptLine{}),
		reflect.TypeOf(purchase.OrderItem{}),
		reflect.TypeOf(purchase.ReceiptItem{}),
		reflect.TypeOf(purchase.SupplierInvoiceLine{}),
	} {
		for name, wantTag := range map[string]string{"SKUID": "sku_id", "SKUCode": "sku_code"} {
			field, ok := model.FieldByName(name)
			if !ok {
				t.Errorf("%s must carry %s", model, name)
				continue
			}
			if field.Type.Kind() != reflect.String {
				t.Errorf("%s.%s must be a string scalar, got %s", model, name, field.Type)
			}
			if got := field.Tag.Get("json"); got != wantTag {
				t.Errorf("%s.%s JSON tag = %q, want %q", model, name, got, wantTag)
			}
		}
	}
}

// TestV27CatalogEventTopologyIsDeclared locks the new catalog-events routing
// vocabulary so the contract, the cloud topology, and the local emulator script
// can be compared against one source of truth.
func TestV27CatalogEventTopologyIsDeclared(t *testing.T) {
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

// TestV27MonetaryFieldsUseMinorUnitIntegers keeps every money-shaped field an
// exact integer. A float in a price, cost, tax, or amount field would silently
// drift across JSON and service boundaries.
func TestV27MonetaryFieldsUseMinorUnitIntegers(t *testing.T) {
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
			jsonKey, present := v27JSONFieldName(t, path, field)
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
		t.Fatalf("scan v27 monetary fields: %v", err)
	}
	if len(violations) > 0 {
		sort.Strings(violations)
		t.Fatalf("monetary precision violations:\n%s", strings.Join(violations, "\n"))
	}
}
