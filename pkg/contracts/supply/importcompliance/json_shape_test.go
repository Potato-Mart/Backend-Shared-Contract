package importcompliance_test

import (
	"encoding/json"
	geography "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/geography"
	common "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/common/shared"
	"github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/importcompliance"
	importcomplianceenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/importcompliance"
	purchaseenum "github.com/Potato-Mart/Backend-Shared-Contract/v23/pkg/contracts/supply/purchase"
	"strings"
	"testing"
	"time"
)

func TestImportSettingsJSONUsesFixedPointFields(t *testing.T) {
	settings := importcompliance.ImportSettings{
		ID:                      "import_settings",
		Revision:                importcompliance.RevisionMetadata{Number: 1, State: importcomplianceenum.ReviewStateDraft},
		TWDPerAUDMicros:         21_250_000,
		NextInvoiceNumber:       1001,
		ExportMarginBasisPoints: 1500,
		DefaultSignatoryName:    "A. Signatory",
		TaiwanTaxBasisPoints:    500,
		AirCargo: importcompliance.AirCargoSettings{
			ReferenceWeightGrams:                         100_000,
			TaiwanInspectionCost:                         common.Money{AmountMinor: 1500, Currency: "TWD"},
			StorageThresholdGrams:                        50_000,
			StorageUnderThresholdCost:                    common.Money{AmountMinor: 200, Currency: "TWD"},
			StorageAtOrOverThresholdCost:                 common.Money{AmountMinor: 300, Currency: "TWD"},
			VolumetricDivisorCubicCentimetresPerKilogram: 6000,
		},
		AmbientSea: importcompliance.AmbientSeaSettings{ReferenceVolumeCubicCentimetres: 1_000_000},
		FrozenSea:  importcompliance.FrozenSeaSettings{ReferenceVolumeCubicCentimetres: 2_000_000},
		IngredientRules: []importcompliance.IngredientDeclarationRule{
			{ID: "rule_1", Ingredients: "milk", Statement: "Contains milk"},
		},
	}

	payload, err := json.Marshal(settings)
	if err != nil {
		t.Fatalf("marshal import settings: %v", err)
	}
	text := string(payload)
	for _, key := range []string{
		`"twd_per_aud_micros":21250000`,
		`"export_margin_basis_points":1500`,
		`"taiwan_tax_basis_points":500`,
		`"reference_weight_grams":100000`,
		`"volumetric_divisor_cubic_centimetres_per_kilogram":6000`,
		`"reference_volume_cubic_centimetres":1000000`,
	} {
		if !strings.Contains(text, key) {
			t.Fatalf("settings JSON missing fixed-point field %s: %s", key, payload)
		}
	}
	for _, forbidden := range []string{`"exchange_rate"`, `"export_margin"`, `"air_ref_weight"`, `"taiwan_base_storage_cost"`} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("settings JSON exposed ambiguous floating-point field %s: %s", forbidden, payload)
		}
	}
}

func TestDeclarationAndLabelRoundTripManagedMediaAndMeasurements(t *testing.T) {
	now := time.Date(2026, 7, 16, 3, 4, 5, 0, time.UTC)
	baseRevision := int64(1)
	declaration := importcompliance.ManufacturerDeclaration{
		ID:       "decl_1",
		Revision: importcompliance.RevisionMetadata{Number: 2, BaseRevisionNumber: &baseRevision, State: importcomplianceenum.ReviewStateDraft},
		PurchaseOrder: importcompliance.PurchaseOrderSnapshot{
			ID: "po_1", OrderNumber: "PO-1001", Status: purchaseenum.PurchaseOrderStatusConfirmed,
			SupplierCode: "SUP-1", SupplierName: "Supplier", ExpectedArrival: common.Date("2026-08-01"), CapturedAt: now,
		},
		DeclarationReference: "DEC-1001",
		Shipment: importcompliance.DeclarationShipment{
			DeclarationDate: common.Date("2026-07-16"), ImportMode: importcomplianceenum.ImportModeAirCargo,
			ShippingMethod: "Air", TransportReference: "BR-315", ConsignmentIdentifier: "AWB-1",
			PortOfLoading: "TPE", DestinationPort: "SYD",
		},
		Manufacturer: importcompliance.ManufacturerDetails{Name: "Maker", Address: &common.Address{Label: "Manufacturer", Line1: "1 Zhongxiao Road", Locality: "Taipei", PostalCode: "100", Country: geography.CountryRef{Code: "TW", Name: "Taiwan"}, FormattedAddress: "1 Zhongxiao Road, Taipei 100, Taiwan", PlaceID: "place-2"}, Phone: "+886"},
		Signatory:    importcompliance.DeclarationSignatory{Name: "Signer", Title: "Director", SignatureMediaID: "media_signature"},
		Lines: []importcompliance.DeclarationLine{{
			ID: "line_1", SourceLineID: "po_line_1", ProductReference: "SKU-1", EnglishName: "Product", ChineseName: "產品",
			OrderedQuantity: 10, CartonCount: 2, SingleNetWeightGrams: 500, TotalNetWeightGrams: 5000,
			TotalGrossWeightGrams: 5500, ExpiryDate: common.Date("2027-01-01"), Ingredients: "Milk",
			ManufacturingProcess: "Cooked", Note: "Keep dry",
		}},
	}
	declarationPayload, err := json.Marshal(declaration)
	if err != nil {
		t.Fatalf("marshal declaration: %v", err)
	}
	for _, key := range []string{`"signature_media_id":"media_signature"`, `"single_net_weight_grams":500`, `"consignment_identifier":"AWB-1"`} {
		if !strings.Contains(string(declarationPayload), key) {
			t.Fatalf("declaration JSON missing %s: %s", key, declarationPayload)
		}
	}
	if strings.Contains(string(declarationPayload), "signature_image") || strings.Contains(string(declarationPayload), "data:image") {
		t.Fatalf("declaration embedded a signature payload: %s", declarationPayload)
	}

	label := importcompliance.LabelMaster{
		ID: "label_1", Revision: importcompliance.RevisionMetadata{Number: 1, State: importcomplianceenum.ReviewStateDraft},
		SourceProduct: importcompliance.ProductSnapshot{SKUCode: "SKU-1", CapturedAt: now},
		SKUCode:       "SKU-1", SKU: "retail-sku", VariantCode: "au-65x45", Brand: "Brand", EnglishName: "Product", ChineseName: "產品",
		Barcode: "930000000001", NetWeightGrams: 500, PackageDimensions: common.Dimensions{WidthMM: 65, LengthMM: 45, HeightMM: 10},
		Ingredients: "Milk", Allergens: "Milk", ManufacturingProcess: "Cooked", BestBefore: "Shown on package (YYYY/MM/DD)", ShelfLife: "12 months",
		Importer: importcompliance.LabelImporter{Name: "Potato Mart", Address: &common.Address{Label: "Importer", Line1: "1 Market Street", Locality: "Sydney", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-NSW"}, PostalCode: "2000", Country: geography.CountryRef{Code: "AU", Name: "Australia"}, FormattedAddress: "1 Market Street, Sydney NSW 2000, Australia", PlaceID: "place-1"}, Phone: "+61"}, CountryOfOrigin: "Taiwan",
		SecondNutritionEnabled: true,
		NutritionPanels: []importcompliance.NutritionPanel{
			{Title: "Prepared", ServingsPerPack: "2", ServingSize: "250 g", EnergyPerServe: "500", SodiumPer100Grams: "30"},
			{Title: "As sold"},
		},
		PackagePhotoMediaID: "media_package", PackagePhotoName: "package.jpg",
		Layout: importcompliance.LabelLayout{Size: importcomplianceenum.LabelSize65x45, Orientation: importcomplianceenum.LabelOrientationPortrait, FontScaleBasisPoints: 10_000, IncludeBarcode: true},
	}
	labelPayload, err := json.Marshal(label)
	if err != nil {
		t.Fatalf("marshal label: %v", err)
	}
	var decoded importcompliance.LabelMaster
	if err := json.Unmarshal(labelPayload, &decoded); err != nil {
		t.Fatalf("unmarshal label: %v", err)
	}
	if decoded.VariantCode != "au-65x45" || decoded.NetWeightGrams != 500 || decoded.PackageDimensions.WidthMM != 65 ||
		decoded.PackagePhotoMediaID != "media_package" || len(decoded.NutritionPanels) != 2 || !decoded.Layout.IncludeBarcode {
		t.Fatalf("label fields did not round-trip: %+v", decoded)
	}
}

func TestTariffRatesPreserveOfficialTextAndParsedBasisPoints(t *testing.T) {
	percentage := int64(500)
	classification := importcompliance.TariffClassification{
		Jurisdiction: importcomplianceenum.JurisdictionAustralia,
		Code:         "1905.90",
		DutyRate:     importcompliance.RateValue{Raw: "5%", BasisPoints: &percentage},
		GSTRate:      &importcompliance.RateValue{Raw: "10%"},
		Catalogue: importcompliance.CatalogueReference{
			ID: "abf-2026", Jurisdiction: importcomplianceenum.JurisdictionAustralia, Version: "2026-07-16",
			EntryID: "1905.90", SourceURL: "https://www.abf.gov.au/", SourceChecksumSHA256: "abc",
			RetrievedAt: time.Date(2026, 7, 16, 0, 0, 0, 0, time.UTC),
		},
	}
	payload, err := json.Marshal(classification)
	if err != nil {
		t.Fatalf("marshal tariff classification: %v", err)
	}
	if !strings.Contains(string(payload), `"raw":"5%","basis_points":500`) ||
		!strings.Contains(string(payload), `"gst_rate":{"raw":"10%"}`) {
		t.Fatalf("tariff rate shape lost raw or optional parsed value: %s", payload)
	}
}

func TestRFIRecordRoundTripKeepsExternalEventsExplicit(t *testing.T) {
	now := time.Date(2026, 7, 16, 6, 0, 0, 0, time.UTC)
	record := importcompliance.RFIRecord{
		ID: "rfi_1", Revision: importcompliance.RevisionMetadata{Number: 1, State: importcomplianceenum.ReviewStateApproved},
		Channel: importcomplianceenum.RFIChannelEmailException, CurrentSubmissionState: importcomplianceenum.RFISubmissionStateSubmitted,
		QuarantineNumber: "Q-1", AirwayBill: "AWB-1", AvailableFrom: common.Date("2026-07-20"), RequestedDate: common.Date("2026-07-21"),
		Comments: "Handle frozen", BookingAgent: importcompliance.RFIBookingAgent{Name: "Agent", Phone: "+61", Email: "agent@example.com"},
		InspectionLocation: importcompliance.RFIInspectionLocation{
			BusinessNameAndAANumber: "Potato Mart / AA-1",
			PremiseAddress:          &common.Address{Label: "Inspection premise", Line1: "1 Market Street", Locality: "Sydney", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-NSW"}, PostalCode: "2000", Country: geography.CountryRef{Code: "AU", Name: "Australia"}, FormattedAddress: "1 Market Street, Sydney NSW 2000, Australia", PlaceID: "place-rfi-1"},
			OpeningHours:            "09:00-17:00", ContactName: "Receiver", ContactPhone: "+61", PrivateResidence: false,
		},
		InspectionDirection: "Dock 2", RequestedTime: importcomplianceenum.RFIRequestedTimeAM, Overtime: true,
		EmailSubjectPrefix: "RFI", EmailBody: "Please inspect", AttachmentMediaIDs: []string{"media_1"},
		SubmissionEvents: []importcompliance.RFIExternalEvent{{ID: "event_1", State: importcomplianceenum.RFISubmissionStateSubmitted, ExternalReference: "DAFF-1", OccurredAt: now, RecordedBy: "admin_1"}},
	}
	payload, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("marshal RFI record: %v", err)
	}
	var decoded importcompliance.RFIRecord
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal RFI record: %v", err)
	}
	if decoded.CurrentSubmissionState != importcomplianceenum.RFISubmissionStateSubmitted || len(decoded.SubmissionEvents) != 1 ||
		decoded.SubmissionEvents[0].ExternalReference != "DAFF-1" || decoded.EmailSubjectPrefix != "RFI" ||
		decoded.InspectionLocation.PremiseAddress == nil || decoded.InspectionLocation.PremiseAddress.PlaceID != "place-rfi-1" {
		t.Fatalf("RFI fields did not round-trip: %+v", decoded)
	}
}
