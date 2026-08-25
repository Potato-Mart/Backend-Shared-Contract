package import_compliance_test

import (
	"encoding/json"

	geography "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/geography"

	import_compliance "github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/import_compliance"

	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/measurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/import_compliance/import_compliance_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/supply/purchase/purchase_enums"
)

func TestDeclarationAndLabelRoundTripManagedMediaAndMeasurements(t *testing.T) {
	now := time.Date(2026, 7, 16, 3, 4, 5, 0, time.UTC)
	baseRevision := int64(1)
	declaration := import_compliance.ManufacturerDeclaration{
		ID:       "decl_1",
		Revision: import_compliance.RevisionMetadata{Number: 2, BaseRevisionNumber: &baseRevision, State: import_compliance_enums.ReviewStateDraft},
		PurchaseOrder: import_compliance.PurchaseOrderSnapshot{
			ID: "po_1", OrderNumber: "PO-1001", Status: purchase_enums.PurchaseOrderStatusConfirmed,
			SupplierCode: "SUP-1", SupplierName: "Supplier", ExpectedArrival: temporal.Date("2026-08-01"), CapturedAt: now,
		},
		DeclarationReference: "DEC-1001",
		Shipment: import_compliance.DeclarationShipment{
			DeclarationDate: temporal.Date("2026-07-16"), ImportMode: import_compliance_enums.ImportModeAirCargo,
			ShippingMethod: "Air", TransportReference: "BR-315", ConsignmentIdentifier: "AWB-1",
			PortOfLoading: "TPE", DestinationPort: "SYD",
		},
		Manufacturer: import_compliance.ManufacturerDetails{Name: "Maker", Address: &geography.Address{Label: "Manufacturer", Line1: "1 Zhongxiao Road", Locality: "Taipei", PostalCode: "100", Country: geography.CountryRef{Code: "TW", Name: "Taiwan"}, FormattedAddress: "1 Zhongxiao Road, Taipei 100, Taiwan", PlaceID: "place-2"}, Phone: "+886"},
		Signatory:    import_compliance.DeclarationSignatory{Name: "Signer", Title: "Director", SignatureMediaCode: "media_signature"},
		Lines: []import_compliance.DeclarationLine{{
			ID: "line_1", SourceLineID: "po_line_1", SKUCode: "SKU-1", EnglishName: "Product", ChineseName: "產品",
			OrderedQuantity: 10, CartonCount: 2, SingleNetWeightGrams: 500, TotalNetWeightGrams: 5000,
			TotalGrossWeightGrams: 5500, ExpiryDate: temporal.Date("2027-01-01"), Ingredients: "Milk",
			ManufacturingProcess: "Cooked", Note: "Keep dry",
		}},
	}
	declarationPayload, err := json.Marshal(declaration)
	if err != nil {
		t.Fatalf("marshal declaration: %v", err)
	}
	for _, key := range []string{`"signature_media_code":"media_signature"`, `"single_net_weight_grams":500`, `"consignment_identifier":"AWB-1"`} {
		if !strings.Contains(string(declarationPayload), key) {
			t.Fatalf("declaration JSON missing %s: %s", key, declarationPayload)
		}
	}
	if strings.Contains(string(declarationPayload), "signature_image") || strings.Contains(string(declarationPayload), "data:image") {
		t.Fatalf("declaration embedded a signature payload: %s", declarationPayload)
	}

	label := import_compliance.LabelMaster{
		ID: "label_1", Revision: import_compliance.RevisionMetadata{Number: 1, State: import_compliance_enums.ReviewStateDraft},
		SourceProductEvidence: import_compliance.LabelProductEvidence{SKUCode: "SKU-1", CapturedAt: now},
		SKUCode:               "SKU-1", VariantCode: "au-65x45", Brand: "Brand", EnglishName: "Product", ChineseName: "產品",
		Barcode: "930000000001", NetWeightGrams: 500, PackageDimensions: measurement.Dimensions{WidthMM: 65, LengthMM: 45, HeightMM: 10},
		Ingredients: "Milk", Allergens: "Milk", ManufacturingProcess: "Cooked", BestBefore: "Shown on package (YYYY/MM/DD)", ShelfLife: "12 months",
		Importer: import_compliance.LabelImporter{Name: "Potato Mart", Address: &geography.Address{Label: "Importer", Line1: "1 Market Street", Locality: "Sydney", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-NSW"}, PostalCode: "2000", Country: geography.CountryRef{Code: "AU", Name: "Australia"}, FormattedAddress: "1 Market Street, Sydney NSW 2000, Australia", PlaceID: "place-1"}, Phone: "+61"}, CountryOfOrigin: "Taiwan",
		SecondNutritionEnabled: true,
		NutritionPanels: []import_compliance.NutritionPanel{
			{Title: "Prepared", ServingsPerPack: "2", ServingSize: "250 g", EnergyPerServe: "500", SodiumPer100Grams: "30"},
			{Title: "As sold"},
		},
		PackagePhotoMediaCode: "media_package", PackagePhotoName: "package.jpg",
		Layout: import_compliance.LabelLayout{Size: import_compliance_enums.LabelSize65x45, Orientation: import_compliance_enums.LabelOrientationPortrait, FontScaleBasisPoints: 10_000, IncludeBarcode: true},
	}
	labelPayload, err := json.Marshal(label)
	if err != nil {
		t.Fatalf("marshal label: %v", err)
	}
	var decoded import_compliance.LabelMaster
	if err := json.Unmarshal(labelPayload, &decoded); err != nil {
		t.Fatalf("unmarshal label: %v", err)
	}
	if decoded.VariantCode != "au-65x45" || decoded.NetWeightGrams != 500 || decoded.PackageDimensions.WidthMM != 65 ||
		decoded.SKUCode != "SKU-1" || decoded.SourceProductEvidence.SKUCode != "SKU-1" ||
		decoded.PackagePhotoMediaCode != "media_package" || len(decoded.NutritionPanels) != 2 || !decoded.Layout.IncludeBarcode {
		t.Fatalf("label fields did not round-trip: %+v", decoded)
	}
	for _, removed := range []string{`"source_product":`, `"sku":`} {
		if strings.Contains(string(labelPayload), removed) {
			t.Fatalf("label retained legacy product link %s: %s", removed, labelPayload)
		}
	}
}

func TestTariffRatesPreserveOfficialTextAndParsedBasisPoints(t *testing.T) {
	percentage := int64(500)
	classification := import_compliance.TariffClassification{
		Jurisdiction: import_compliance_enums.JurisdictionAustralia,
		Code:         "1905.90",
		DutyRate:     import_compliance.RateValue{Raw: "5%", BasisPoints: &percentage},
		GSTRate:      &import_compliance.RateValue{Raw: "10%"},
		Catalogue: import_compliance.CatalogueReference{
			ID: "abf-2026", Jurisdiction: import_compliance_enums.JurisdictionAustralia, Version: "2026-07-16",
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
	record := import_compliance.RFIRecord{
		ID: "rfi_1", Revision: import_compliance.RevisionMetadata{Number: 1, State: import_compliance_enums.ReviewStateApproved},
		Channel: import_compliance_enums.RFIChannelEmailException, CurrentSubmissionState: import_compliance_enums.RFISubmissionStateSubmitted,
		QuarantineNumber: "Q-1", AirwayBill: "AWB-1", AvailableFrom: temporal.Date("2026-07-20"), RequestedDate: temporal.Date("2026-07-21"),
		Comments: "Handle frozen", BookingAgent: import_compliance.RFIBookingAgent{Name: "Agent", Phone: "+61", Email: "agent@example.com"},
		InspectionLocation: import_compliance.RFIInspectionLocation{
			BusinessNameAndAANumber: "Potato Mart / AA-1",
			PremiseAddress:          &geography.Address{Label: "Inspection premise", Line1: "1 Market Street", Locality: "Sydney", AdministrativeArea: &geography.AdministrativeAreaRef{Code: "AU-NSW"}, PostalCode: "2000", Country: geography.CountryRef{Code: "AU", Name: "Australia"}, FormattedAddress: "1 Market Street, Sydney NSW 2000, Australia", PlaceID: "place-rfi-1"},
			OpeningHours:            "09:00-17:00", ContactName: "Receiver", ContactPhone: "+61", PrivateResidence: false,
		},
		InspectionDirection: "Dock 2", RequestedTime: import_compliance_enums.RFIRequestedTimeAM, Overtime: true,
		EmailSubjectPrefix: "RFI", EmailBody: "Please inspect", AttachmentMediaCodes: []string{"media_1"},
		SubmissionEvents: []import_compliance.RFIExternalEvent{{ID: "event_1", State: import_compliance_enums.RFISubmissionStateSubmitted, ExternalReference: "DAFF-1", OccurredAt: now, RecordedBy: "admin_1"}},
	}
	payload, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("marshal RFI record: %v", err)
	}
	var decoded import_compliance.RFIRecord
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal RFI record: %v", err)
	}
	if decoded.CurrentSubmissionState != import_compliance_enums.RFISubmissionStateSubmitted || len(decoded.SubmissionEvents) != 1 ||
		decoded.SubmissionEvents[0].ExternalReference != "DAFF-1" || decoded.EmailSubjectPrefix != "RFI" ||
		decoded.InspectionLocation.PremiseAddress == nil || decoded.InspectionLocation.PremiseAddress.PlaceID != "place-rfi-1" {
		t.Fatalf("RFI fields did not round-trip: %+v", decoded)
	}
}
