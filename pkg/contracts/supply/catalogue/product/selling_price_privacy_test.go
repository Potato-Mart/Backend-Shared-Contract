package product_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v35/pkg/contracts/supply/catalogue/product"
)

func TestSellingProductPriceGraphExcludesPrivateEvidence(t *testing.T) {
	price, ok := reflect.TypeOf(product.SellingProduct{}).FieldByName("Price")
	if !ok {
		t.Fatal("SellingProduct has no canonical price projection")
	}
	forbiddenKeys := map[string]bool{
		"price_book_code": true, "price_book_revision": true, "price_entry_id": true,
		"price_entry_revision": true, "actor": true, "actor_user_id": true,
		"source_approved_price": true, "compared_cost": true, "cost_comparison": true,
		"below_cost_warning": true, "custom_override": true, "manual_lock": true,
		"source_id": true, "source_revision": true, "net_goods_amount": true,
	}
	seen := make(map[reflect.Type]bool)
	var inspect func(reflect.Type)
	inspect = func(kind reflect.Type) {
		if seen[kind] {
			return
		}
		seen[kind] = true
		if strings.Contains(kind.PkgPath(), "/pricing/quote") || strings.Contains(kind.PkgPath(), "/supply/procurement") {
			t.Errorf("customer price graph imports private checkout/procurement type %s", kind)
		}
		switch kind.Kind() {
		case reflect.Pointer, reflect.Slice, reflect.Array:
			inspect(kind.Elem())
		case reflect.Map:
			inspect(kind.Key())
			inspect(kind.Elem())
		case reflect.Struct:
			for i := 0; i < kind.NumField(); i++ {
				field := kind.Field(i)
				if !field.IsExported() {
					continue
				}
				key := strings.Split(field.Tag.Get("json"), ",")[0]
				if forbiddenKeys[key] {
					t.Errorf("customer price graph exposes private %s.%s (%s)", kind, field.Name, key)
				}
				inspect(field.Type)
			}
		}
	}
	inspect(price.Type)
}
