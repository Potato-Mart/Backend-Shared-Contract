package pkg_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/pricing/membership"
	"github.com/Potato-Mart/Backend-Shared-Contract/v42/pkg/contracts/pubsub/notification"
)

func TestV38PackageReferencesPreserveSKUAndPhysicalVersion(t *testing.T) {
	refs := []packaging.PackageOptionRef{
		{SKUCode: "A00001", Code: "CASE12", Version: 1},
		{SKUCode: "A00001", Code: "CASE12", Version: 2},
		{SKUCode: "A00002", Code: "CASE12", Version: 1},
	}
	data, err := json.Marshal(refs)
	if err != nil {
		t.Fatal(err)
	}
	var decoded []packaging.PackageOptionRef
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(refs, decoded) {
		t.Fatal("physical package identity was lost")
	}
	model := reflect.TypeOf(packaging.PackageOptionRef{})
	for field, tag := range map[string]string{"SKUCode": "sku_code", "Code": "code", "Version": "version"} {
		value, ok := model.FieldByName(field)
		if !ok || value.Tag.Get("json") != tag {
			t.Fatalf("required package reference field %s changed", field)
		}
	}
}

func TestV38TierRetiredFieldsStayAbsent(t *testing.T) {
	model := reflect.TypeOf(membership.MembershipTier{})
	for _, name := range []string{"IsActive", "DiscountPercent", "FreeShippingThreshold", "BirthdayBonusPoints", "Rank"} {
		if _, found := model.FieldByName(name); found {
			t.Errorf("retired or service-derived field %s survives", name)
		}
	}
}

func TestV38GeographicPathKeepsOfficialStateSeparateFromDepotRegion(t *testing.T) {
	model := reflect.TypeOf(geography.GeographicPath{})
	for field, tag := range map[string]string{"CountryCode": "country_code", "MarketCode": "market_code,omitempty", "StateCode": "state_code,omitempty", "DepotCode": "depot_code,omitempty"} {
		value, ok := model.FieldByName(field)
		if !ok || value.Tag.Get("json") != tag {
			t.Fatalf("geographic path field %s changed", field)
		}
	}
	if model.NumField() != 4 {
		t.Fatal("path must contain only the four official hierarchy levels")
	}
}

func TestV38GiftDeliveryFactContainsOnlySafeCorrelation(t *testing.T) {
	fact := notification.GiftCardDeliveredEvent{IssuanceID: "issuance-1", DeliveredAt: time.Date(2026, 9, 29, 0, 0, 0, 0, time.UTC)}
	data, err := json.Marshal(fact)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"issuance_id":"issuance-1","delivered_at":"2026-09-29T00:00:00Z"}` {
		t.Fatalf("unsafe or changed delivery fact: %s", data)
	}
}
