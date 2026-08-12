package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v27/pkg/contracts/supply/purchase/purchase_enums"
)

func TestPurchaseEnumsValidateKnownValues(t *testing.T) {
	assertStringEnums(t, []enumCase{
		{name: "purchaseenum.LineTaxTreatment", valid: []stringEnum{purchase_enums.LineTaxTreatmentTaxable, purchase_enums.LineTaxTreatmentGSTFree, purchase_enums.LineTaxTreatmentInputTaxed, purchase_enums.LineTaxTreatmentOutOfScope, purchase_enums.LineTaxTreatmentUnknown}, invalid: purchase_enums.LineTaxTreatment("__invalid__")},
		{name: "purchaseenum.TaxPriceBasis", valid: []stringEnum{purchase_enums.TaxPriceBasisInclusive, purchase_enums.TaxPriceBasisExclusive, purchase_enums.TaxPriceBasisUnknown}, invalid: purchase_enums.TaxPriceBasis("__invalid__")},
		{name: "purchaseenum.SupplierTaxSource", valid: []stringEnum{purchase_enums.SupplierTaxSourceExplicitLine, purchase_enums.SupplierTaxSourceInvoiceHeader, purchase_enums.SupplierTaxSourceCalculated, purchase_enums.SupplierTaxSourceAbsent}, invalid: purchase_enums.SupplierTaxSource("__invalid__")},
		{name: "purchaseenum.InputTaxClaimStatus", valid: []stringEnum{purchase_enums.InputTaxClaimStatusClaimable, purchase_enums.InputTaxClaimStatusNotClaimable, purchase_enums.InputTaxClaimStatusInsufficientEvidence, purchase_enums.InputTaxClaimStatusPendingReview}, invalid: purchase_enums.InputTaxClaimStatus("__invalid__")},
		{name: "purchaseenum.SupplierInvoiceStatus", valid: []stringEnum{purchase_enums.SupplierInvoiceStatusDraft, purchase_enums.SupplierInvoiceStatusBlocked, purchase_enums.SupplierInvoiceStatusConfirmed, purchase_enums.SupplierInvoiceStatusDisputed, purchase_enums.SupplierInvoiceStatusCancelled}, invalid: purchase_enums.SupplierInvoiceStatus("__invalid__")},
		{name: "purchaseenum.PurchaseOrderStatus", valid: []stringEnum{purchase_enums.PurchaseOrderStatusDraft, purchase_enums.PurchaseOrderStatusSubmitted, purchase_enums.PurchaseOrderStatusConfirmed, purchase_enums.PurchaseOrderStatusPartiallyReceived, purchase_enums.PurchaseOrderStatusReceived, purchase_enums.PurchaseOrderStatusCancelled, purchase_enums.PurchaseOrderStatusRefunded}, invalid: purchase_enums.PurchaseOrderStatus("__invalid__")},
	})
}
