package import_compliance

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v31/pkg/contracts/common/temporal"
)

type DeclarationLine struct {
	ID                        string        `json:"id"`
	SourceLineID              string        `json:"source_line_id,omitempty"`
	SourceLabelID             string        `json:"source_label_id,omitempty"`
	SourceLabelRevisionNumber *int64        `json:"source_label_revision_number,omitempty"`
	SKUCode                   string        `json:"sku_code,omitempty"`
	EnglishName               string        `json:"english_name"`
	ChineseName               string        `json:"chinese_name,omitempty"`
	OrderedQuantity           int64         `json:"ordered_quantity"`
	CartonCount               int64         `json:"carton_count"`
	SingleNetWeightGrams      int64         `json:"single_net_weight_grams"`
	TotalNetWeightGrams       int64         `json:"total_net_weight_grams"`
	TotalGrossWeightGrams     int64         `json:"total_gross_weight_grams"`
	ExpiryDate                temporal.Date `json:"expiry_date,omitempty"`
	Ingredients               string        `json:"ingredients,omitempty"`
	ManufacturingProcess      string        `json:"manufacturing_process,omitempty"`
	Note                      string        `json:"note,omitempty"`
}
