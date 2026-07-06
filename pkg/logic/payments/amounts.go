// Package paymentslogic holds derived helpers for payments contract structs.
// It lives outside pkg/contracts so the contract files stay pure data shapes.
package paymentslogic

import "github.com/Potato-Mart/Backend-Shared-Contract/v12/pkg/contracts/payments"

// TotalMinor returns the sum of all positive components (purchase + tip +
// surcharge + cashout + moto) minus refund. It is a convenience helper; read
// individual components and AuthorizedMinor for receipts and reconciliation.
func TotalMinor(a payments.Amounts) int64 {
	return a.PurchaseMinor + a.TipMinor + a.SurchargeMinor + a.CashoutMinor + a.MOTOMinor - a.RefundMinor
}
