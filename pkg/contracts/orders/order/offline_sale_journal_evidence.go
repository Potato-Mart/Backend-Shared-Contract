package order

import (
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/common/temporal"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/payments/merchant"
	"github.com/Potato-Mart/Backend-Shared-Contract/v54/pkg/contracts/pricing/pricebook"
)

// OfflineSaleJournalEvidence captures the device's sale and cash provenance
// for protected exchange between Orders and Payments. It is not an HTTP DTO
// or proof that a journal, cashier, price book, or issuer has been verified.
// CachedIssuer is the legal snapshot actually used by the till; services can
// compare it to authoritative issuer evidence after synchronization without
// changing the original receipt or asserting a review outcome here.
// CanonicalEntryBase64 preserves the exact hashed journal body, including its
// recorded sale and receipt snapshot, so owners can verify the chain without
// reconstructing those bytes from a different representation.
type OfflineSaleJournalEvidence struct {
	ClientSaleID  string `json:"client_sale_id"`
	DeviceKey     string `json:"device_key"`
	ReceiptNumber string `json:"receipt_number"`
	ReceiptPrefix string `json:"receipt_prefix"`
	// ReceiptCounter is the device's persistent counter as a canonical
	// unsigned decimal string, preserving its full uint64 range in JSON.
	ReceiptCounter       string                             `json:"receipt_counter"`
	MarketCode           string                             `json:"market_code"`
	DepotCode            string                             `json:"depot_code"`
	RegisterID           string                             `json:"register_id"`
	SessionID            string                             `json:"session_id"`
	BusinessDate         temporal.Date                      `json:"business_date"`
	CashierUserID        string                             `json:"cashier_user_id"`
	CashierAuth          OfflineCashierAuth                 `json:"cashier_auth"`
	SoldAt               time.Time                          `json:"sold_at"`
	OfflineSince         time.Time                          `json:"offline_since"`
	Sequence             int64                              `json:"sequence"`
	JournalSchema        int64                              `json:"journal_schema"`
	RecordedAt           time.Time                          `json:"recorded_at"`
	PreviousHash         string                             `json:"previous_hash"`
	EntryHash            string                             `json:"entry_hash"`
	CanonicalEntryBase64 string                             `json:"canonical_entry_base64"`
	PriceBook            pricebook.OfflinePriceBookEvidence `json:"price_book"`
	CachedIssuer         merchant.MerchantLegalSnapshot     `json:"cached_issuer"`
	CashTender           OfflineCashTenderEvidence          `json:"cash_tender"`
}
