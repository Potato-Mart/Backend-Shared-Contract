package compliance

// TariffLineSnapshot preserves every purchase-order/product field displayed by
// the tariff worksheet before classification begins.
type TariffLineSnapshot struct {
	ID                  string `json:"id"`
	PurchaseOrderLineID string `json:"purchase_order_line_id,omitempty"`
	SKUCode             string `json:"sku_code,omitempty"`
	Barcode             string `json:"barcode,omitempty"`
	ProductName         string `json:"product_name,omitempty"`
	AlternateNames      string `json:"alternate_names,omitempty"`
	Brand               string `json:"brand,omitempty"`
	OrderedQuantity     *int64 `json:"ordered_quantity,omitempty"`
	ProductTaxed        *bool  `json:"product_taxed,omitempty"`
}
