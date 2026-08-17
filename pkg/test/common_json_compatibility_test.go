package pkg_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/audit"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/device"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geography/geography_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/geometry"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/identity"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/localization"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/metadata"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/party"
	"github.com/Potato-Mart/Backend-Shared-Contract/v28/pkg/contracts/common/temporal"
)

func TestV27MovedCommonModelsMatchV23GoldenJSON(t *testing.T) {
	// These byte-for-byte fixtures are the v23.0.0 wire baseline for every
	// common model moved by the v27 package reorganization.
	createdAt := time.Date(2026, time.August, 7, 3, 4, 5, 0, time.UTC)
	address := geography.Address{
		Line1:      "1 Market Lane",
		Locality:   "Dandenong South",
		PostalCode: "3175",
		AdministrativeArea: &geography.AdministrativeAreaRef{
			Code: "AU-VIC",
			Name: "Victoria",
			Type: geography_enums.AdministrativeAreaState,
		},
		Country: geography.CountryRef{Code: "AU", Name: "Australia"},
	}

	fixtures := []struct {
		name  string
		value any
		want  string
	}{
		{"Address", address, `{"line1":"1 Market Lane","locality":"Dandenong South","administrative_area":{"code":"AU-VIC","name":"Victoria","type":"STATE"},"postal_code":"3175","country":{"code":"AU","name":"Australia"}}`},
		{"AuditFields", audit.AuditFields{CreatedAt: createdAt, UpdatedAt: createdAt}, `{"created_at":"2026-08-07T03:04:05Z","updated_at":"2026-08-07T03:04:05Z"}`},
		{"LifecycleAction", audit.LifecycleAction{By: "operator-1"}, `{"by":"operator-1"}`},
		{"ContactAddress", party.ContactAddress{ID: "address-1", Contact: &party.Recipient{Name: "Ada", Email: "ada@example.test"}, Address: &address}, `{"id":"address-1","contact":{"name":"Ada","email":"ada@example.test"},"address":{"line1":"1 Market Lane","locality":"Dandenong South","administrative_area":{"code":"AU-VIC","name":"Victoria","type":"STATE"},"postal_code":"3175","country":{"code":"AU","name":"Australia"}}}`},
		{"OrganisationDetail", party.OrganisationDetail{PartyRef: party.PartyRef{ID: "org-1", Code: "POTATO", Name: "Potato Mart"}, TradingName: "Potato Mart", Metadata: metadata.Metadata{"source": "v23.0.0"}}, `{"id":"org-1","code":"POTATO","name":"Potato Mart","trading_name":"Potato Mart","metadata":{"source":"v23.0.0"}}`},
		{"PartyRef", party.PartyRef{ID: "party-1", Code: "SUP-1", Name: "Supplier", Phone: "+61", Email: "supplier@example.test"}, `{"id":"party-1","code":"SUP-1","name":"Supplier","phone":"+61","email":"supplier@example.test"}`},
		{"Recipient", party.Recipient{Name: "Ada", Phone: "+61", Email: "ada@example.test"}, `{"name":"Ada","phone":"+61","email":"ada@example.test"}`},
		{"PersonName", party.PersonName{FirstName: "Ada", PreferredName: "A"}, `{"first_name":"Ada","preferred_name":"A"}`},
		{"ContactChannels", party.ContactChannels{Email: "ada@example.test", ExternalHandles: map[string]string{"wechat": "ada"}}, `{"email":"ada@example.test","external_handles":{"wechat":"ada"}}`},
		{"IdentityLink", identity.IdentityLink{UserID: "user-1", AuthIdentityIDs: []string{"auth-1", "auth-2"}}, `{"user_id":"user-1","auth_identity_ids":["auth-1","auth-2"]}`},
		{"Date", temporal.Date("2026-08-07"), `"2026-08-07"`},
		{"TimeOfDay", temporal.TimeOfDay("09:30"), `"09:30"`},
		{"DeviceRecord", device.DeviceRecord{DeviceKey: "device-1", IPAddress: "127.0.0.1"}, `{"device_key":"device-1","ip_address":"127.0.0.1"}`},
		{"Vector3", geometry.Vector3{X: 1, Y: 2, Z: 3}, `{"x":1,"y":2,"z":3}`},
		{"Rotation3", geometry.Rotation3{X: 1, Y: 2, Z: 3}, `{"x":1,"y":2,"z":3}`},
		{"Size3D", geometry.Size3D{WidthMM: 10, HeightMM: 20, DepthMM: 30}, `{"width_mm":10,"height_mm":20,"depth_mm":30}`},
		{"BoundingBox", geometry.BoundingBox{Min: geometry.Vector3{X: 1}, Max: geometry.Vector3{X: 2}}, `{"min":{"x":1,"y":0,"z":0},"max":{"x":2,"y":0,"z":0}}`},
		{"Transform", geometry.Transform{Position: geometry.Vector3{X: 1}, Scale: &geometry.Vector3{X: 2}}, `{"position":{"x":1,"y":0,"z":0},"scale":{"x":2,"y":0,"z":0}}`},
		{"LocalizedDescription", localization.LocalizedDescription{Language: "en", Description: "Potatoes"}, `{"language":"en","description":"Potatoes"}`},
		{"LocalizedName", localization.LocalizedName{Language: "en", Name: "Potato"}, `{"language":"en","name":"Potato"}`},
		{"LocalizedText", localization.LocalizedText{Language: "en", Text: "Hello"}, `{"language":"en","text":"Hello"}`},
		{"Dimensions", measurement.Dimensions{WidthMM: 1, LengthMM: 2, HeightMM: 3}, `{"width_mm":1,"length_mm":2,"height_mm":3}`},
		{"Weight", measurement.Weight{Grams: 500}, `{"grams":500}`},
		{"PhysicalPackage", packaging.PhysicalPackage{Dimensions: &measurement.Dimensions{WidthMM: 1, LengthMM: 2, HeightMM: 3}, Weight: &measurement.Weight{Grams: 500}}, `{"dimensions":{"width_mm":1,"length_mm":2,"height_mm":3},"weight":{"grams":500}}`},
		{"PackageComponentSnapshot", packaging.PackageComponentSnapshot{PackageOptionID: "each", HandlingUnit: packaging_enums.PackageHandlingUnitEach, PackageCount: 1, UnitsPerPackage: 2, BaseUnits: 2}, `{"package_option_id":"each","handling_unit":"EACH","package_count":1,"units_per_package":2,"base_units":2}`},
		{"PackageCompositionSnapshot", packaging.PackageCompositionSnapshot{TotalBaseUnits: 2, Components: []packaging.PackageComponentSnapshot{{PackageOptionID: "each", HandlingUnit: packaging_enums.PackageHandlingUnitEach, PackageCount: 1, UnitsPerPackage: 2, BaseUnits: 2}}}, `{"total_base_units":2,"components":[{"package_option_id":"each","handling_unit":"EACH","package_count":1,"units_per_package":2,"base_units":2}]}`},
		{"Metadata", metadata.Metadata{"source": "v23.0.0", "unchanged": true}, `{"source":"v23.0.0","unchanged":true}`},
		{"Money", money.Money{AmountMinor: 1234, Currency: "AUD"}, `{"amount_minor":1234,"currency":"AUD"}`},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			got, err := json.Marshal(fixture.value)
			if err != nil {
				t.Fatalf("marshal v23 golden fixture: %v", err)
			}
			if string(got) != fixture.want {
				t.Fatalf("JSON changed:\n got: %s\nwant: %s", got, fixture.want)
			}
		})
	}
}
