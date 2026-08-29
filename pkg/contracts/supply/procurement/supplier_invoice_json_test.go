package procurement_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/packaging/packaging_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/common/temporal"
	purchase "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/procurement"
	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/supply/procurement/purchase_enums"
)

func supplierComposition(baseUnits int64) packaging.PackageCompositionSnapshot {
	return packaging.PackageCompositionSnapshot{
		TotalBaseUnits: baseUnits,
		Components: []packaging.PackageComponentSnapshot{{
			PackageOptionCode: "pkg_case_12", HandlingUnit: packaging_enums.PackageHandlingUnitCase,
			PackageCount: baseUnits / 12, UnitsPerPackage: 12, BaseUnits: baseUnits,
		}},
	}
}

func TestSupplierTaxIdentityRecordsRegistrationAsAnEntityFact(t *testing.T) {
	registeredFrom := temporal.Date("2020-07-20")
	payload, err := json.Marshal(purchase.SupplierTaxIdentity{
		SupplierCode: "SUP-1", SupplierName: "Example Supplier",
		BusinessNumberScheme: "abn", BusinessNumber: "00000000000",
		TaxRegistered: true, TaxRegisteredFrom: &registeredFrom,
	})
	if err != nil {
		t.Fatalf("marshal supplier tax identity: %v", err)
	}
	for _, want := range []string{
		`"supplier_code":"SUP-1"`, `"business_number_scheme":"abn"`,
		`"business_number":"00000000000"`, `"tax_registered":true`,
		`"tax_registered_from":"2020-07-20"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("SupplierTaxIdentity JSON = %s, want %s", payload, want)
		}
	}

	// An unregistered supplier omits the optional registration date rather
	// than emitting a zero value that could read as a real date.
	unregistered, err := json.Marshal(purchase.SupplierTaxIdentity{
		SupplierCode: "SUP-2", SupplierName: "Unregistered Supplier", TaxRegistered: false,
	})
	if err != nil {
		t.Fatalf("marshal unregistered supplier: %v", err)
	}
	for _, omitted := range []string{`"tax_registered_from"`, `"business_number"`, `"business_number_scheme"`} {
		if strings.Contains(string(unregistered), omitted) {
			t.Fatalf("unregistered supplier must omit %s: %s", omitted, unregistered)
		}
	}
	if !strings.Contains(string(unregistered), `"tax_registered":false`) {
		t.Fatalf("registration must be explicit even when false: %s", unregistered)
	}
}

func TestSupplierInvoiceDocumentIsImmutableAndHashed(t *testing.T) {
	receivedAt := time.Date(2026, 8, 12, 2, 3, 4, 0, time.UTC)
	value := purchase.SupplierInvoiceDocument{
		Reference:     "INV-2026-0001.pdf",
		ContentSHA256: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		MediaCode:     "media_1",
		ReceivedAt:    receivedAt,
	}
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal invoice document: %v", err)
	}
	var decoded purchase.SupplierInvoiceDocument
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal invoice document: %v", err)
	}
	if decoded != value {
		t.Fatalf("invoice document did not round-trip: %+v", decoded)
	}
	if !strings.Contains(string(payload), `"content_sha256":"e3b0c442`) {
		t.Fatalf("document hash must be carried: %s", payload)
	}
}

func TestSupplierInvoiceLineRecordsExplicitHeaderCalculatedAndUnknownTax(t *testing.T) {
	declared := money.Money{AmountMinor: 600, Currency: "AUD"}
	calculated := money.Money{AmountMinor: 600, Currency: "AUD"}

	explicit := purchase.SupplierInvoiceLine{
		ID: "line_1", SKUCode: "A00001",
		Composition: supplierComposition(12), BaseUnits: 12,
		UnitPrice:      money.Money{AmountMinor: 550, Currency: "AUD"},
		LineAmount:     money.Money{AmountMinor: 6600, Currency: "AUD"},
		TaxTreatment:   purchase_enums.LineTaxTreatmentTaxable,
		PriceBasis:     purchase_enums.TaxPriceBasisExclusive,
		DeclaredTax:    &declared,
		TaxSource:      purchase_enums.SupplierTaxSourceExplicitLine,
		RecoverableTax: &declared,
		InputTaxClaim:  purchase_enums.InputTaxClaimStatusClaimable,
	}
	payload, err := json.Marshal(explicit)
	if err != nil {
		t.Fatalf("marshal explicit line: %v", err)
	}
	var decodedExplicit purchase.SupplierInvoiceLine
	if err := json.Unmarshal(payload, &decodedExplicit); err != nil {
		t.Fatalf("unmarshal explicit line: %v", err)
	}
	if decodedExplicit.TaxSource != purchase_enums.SupplierTaxSourceExplicitLine ||
		decodedExplicit.DeclaredTax == nil || decodedExplicit.DeclaredTax.AmountMinor != 600 {
		t.Fatalf("explicit supplier tax did not round-trip: %+v", decodedExplicit)
	}
	if decodedExplicit.SKUCode != "A00001" {
		t.Fatalf("invoice line must carry the immutable SKU code: %+v", decodedExplicit)
	}
	if strings.Contains(string(payload), `"calculated_tax"`) {
		t.Fatalf("an explicitly declared line must not also carry a calculated tax: %s", payload)
	}

	// Missing tax may be calculated only for an explicitly taxable line with
	// a known basis.
	calculatedLine, err := json.Marshal(purchase.SupplierInvoiceLine{
		ID: "line_2", SKUCode: "A00002",
		TaxTreatment:  purchase_enums.LineTaxTreatmentTaxable,
		PriceBasis:    purchase_enums.TaxPriceBasisInclusive,
		CalculatedTax: &calculated,
		TaxSource:     purchase_enums.SupplierTaxSourceCalculated,
		InputTaxClaim: purchase_enums.InputTaxClaimStatusPendingReview,
	})
	if err != nil {
		t.Fatalf("marshal calculated line: %v", err)
	}
	if !strings.Contains(string(calculatedLine), `"tax_source":"calculated"`) ||
		strings.Contains(string(calculatedLine), `"declared_tax"`) {
		t.Fatalf("calculated line JSON = %s", calculatedLine)
	}

	// Unknown taxability or basis carries no tax figure at all and cannot be
	// claimed; the invoice status is what blocks confirmation.
	unknown, err := json.Marshal(purchase.SupplierInvoiceLine{
		ID: "line_3", SKUCode: "A00003",
		TaxTreatment:  purchase_enums.LineTaxTreatmentUnknown,
		PriceBasis:    purchase_enums.TaxPriceBasisUnknown,
		TaxSource:     purchase_enums.SupplierTaxSourceAbsent,
		InputTaxClaim: purchase_enums.InputTaxClaimStatusInsufficientEvidence,
	})
	if err != nil {
		t.Fatalf("marshal unknown line: %v", err)
	}
	for _, omitted := range []string{`"declared_tax"`, `"calculated_tax"`, `"recoverable_tax"`} {
		if strings.Contains(string(unknown), omitted) {
			t.Fatalf("an unknown-taxability line must omit %s: %s", omitted, unknown)
		}
	}
	if !strings.Contains(string(unknown), `"input_tax_claim":"insufficient_evidence"`) {
		t.Fatalf("unknown line JSON = %s", unknown)
	}
}

func TestSupplierInvoiceKeepsFreightAndDutySeparateAndProtectsAgainstDuplicates(t *testing.T) {
	issueDate := temporal.Date("2026-08-10")
	receivedAt := time.Date(2026, 8, 12, 2, 3, 4, 0, time.UTC)
	declaredTotal := money.Money{AmountMinor: 600, Currency: "AUD"}
	freight := money.Money{AmountMinor: 1500, Currency: "AUD"}
	duty := money.Money{AmountMinor: 300, Currency: "AUD"}
	lineTax := money.Money{AmountMinor: 600, Currency: "AUD"}

	value := purchase.SupplierInvoice{
		ID: "supplier_invoice_1", InvoiceNumber: "INV-2026-0001",
		Supplier: purchase.SupplierTaxIdentity{
			SupplierCode: "SUP-1", SupplierName: "Example Supplier",
			BusinessNumberScheme: "abn", BusinessNumber: "00000000000", TaxRegistered: true,
		},
		IssueDate:            issueDate,
		Currency:             "AUD",
		CurrencyExponent:     money.CurrencyExponent{Currency: "AUD", Exponent: 2},
		PurchaseOrderNumbers: []string{"PO-1"},
		ReceiptIDs:           []string{"receipt_1"},
		Lines: []purchase.SupplierInvoiceLine{{
			ID: "line_1", SKUCode: "A00001",
			Composition: supplierComposition(12), BaseUnits: 12,
			UnitPrice:           money.Money{AmountMinor: 550, Currency: "AUD"},
			LineAmount:          money.Money{AmountMinor: 6600, Currency: "AUD"},
			TaxTreatment:        purchase_enums.LineTaxTreatmentTaxable,
			PriceBasis:          purchase_enums.TaxPriceBasisExclusive,
			DeclaredTax:         &lineTax,
			TaxSource:           purchase_enums.SupplierTaxSourceExplicitLine,
			RecoverableTax:      &lineTax,
			InputTaxClaim:       purchase_enums.InputTaxClaimStatusClaimable,
			FreightAmount:       &freight,
			DutyAmount:          &duty,
			PurchaseOrderNumber: "PO-1", ReceiptID: "receipt_1",
		}},
		Subtotal:             money.Money{AmountMinor: 6600, Currency: "AUD"},
		DeclaredTaxTotal:     &declaredTotal,
		CalculatedTaxTotal:   money.Money{AmountMinor: 600, Currency: "AUD"},
		FreightTotal:         &freight,
		DutyTotal:            &duty,
		Total:                money.Money{AmountMinor: 9000, Currency: "AUD"},
		TaxSource:            purchase_enums.SupplierTaxSourceExplicitLine,
		InputTaxClaim:        purchase_enums.InputTaxClaimStatusClaimable,
		QualifyingTaxInvoice: true,
		Document: purchase.SupplierInvoiceDocument{
			Reference: "INV-2026-0001.pdf", ContentSHA256: "e3b0c442", ReceivedAt: receivedAt,
		},
		DuplicateKey: "SUP-1|INV-2026-0001",
		Status:       purchase_enums.SupplierInvoiceStatusConfirmed,
		Revision:     1,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal supplier invoice: %v", err)
	}
	for _, want := range []string{
		`"invoice_number":"INV-2026-0001"`, `"duplicate_key":"SUP-1|INV-2026-0001"`,
		`"issue_date":"2026-08-10"`, `"qualifying_tax_invoice":true`,
		`"status":"confirmed"`, `"freight_total"`, `"duty_total"`,
		`"purchase_order_numbers":["PO-1"]`, `"receipt_ids":["receipt_1"]`,
		`"content_sha256":"e3b0c442"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("SupplierInvoice JSON = %s, want %s", payload, want)
		}
	}

	var decoded purchase.SupplierInvoice
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal supplier invoice: %v", err)
	}
	if len(decoded.Lines) != 1 || decoded.Lines[0].SKUCode != "A00001" {
		t.Fatalf("invoice lines did not round-trip: %+v", decoded.Lines)
	}
	if decoded.FreightTotal == nil || decoded.DutyTotal == nil {
		t.Fatal("freight and duty must stay separately recorded, never blended into cost")
	}
	if decoded.DeclaredTaxTotal == nil || decoded.DeclaredTaxTotal.AmountMinor != decoded.CalculatedTaxTotal.AmountMinor {
		t.Fatalf("declared and calculated tax must both survive for reconciliation: %+v", decoded)
	}
	if decoded.Document.ContentSHA256 != "e3b0c442" || decoded.DuplicateKey == "" {
		t.Fatal("document hash and duplicate key must survive the round trip")
	}
}

func TestSupplierInvoiceBlockedWithoutQualifyingEvidence(t *testing.T) {
	payload, err := json.Marshal(purchase.SupplierInvoice{
		ID: "supplier_invoice_2", InvoiceNumber: "INV-2026-0002",
		Supplier:  purchase.SupplierTaxIdentity{SupplierCode: "SUP-2", SupplierName: "Unknown Supplier"},
		IssueDate: temporal.Date("2026-08-11"),
		Currency:  "AUD",
		Lines: []purchase.SupplierInvoiceLine{{
			ID: "line_1", SKUCode: "A00004",
			TaxTreatment:  purchase_enums.LineTaxTreatmentUnknown,
			PriceBasis:    purchase_enums.TaxPriceBasisUnknown,
			TaxSource:     purchase_enums.SupplierTaxSourceAbsent,
			InputTaxClaim: purchase_enums.InputTaxClaimStatusInsufficientEvidence,
		}},
		TaxSource:            purchase_enums.SupplierTaxSourceAbsent,
		InputTaxClaim:        purchase_enums.InputTaxClaimStatusInsufficientEvidence,
		QualifyingTaxInvoice: false,
		DuplicateKey:         "SUP-2|INV-2026-0002",
		Status:               purchase_enums.SupplierInvoiceStatusBlocked,
	})
	if err != nil {
		t.Fatalf("marshal blocked invoice: %v", err)
	}
	for _, want := range []string{
		`"status":"blocked"`, `"qualifying_tax_invoice":false`,
		`"input_tax_claim":"insufficient_evidence"`, `"tax_source":"absent"`,
	} {
		if !strings.Contains(string(payload), want) {
			t.Fatalf("blocked invoice JSON = %s, want %s", payload, want)
		}
	}
	for _, omitted := range []string{`"declared_tax_total"`, `"freight_total"`, `"duty_total"`} {
		if strings.Contains(string(payload), omitted) {
			t.Fatalf("blocked invoice must omit %s: %s", omitted, payload)
		}
	}
}
