package product

// SellingProductClassification contains resolved customer-safe catalogue
// references. Its entries intentionally omit catalogue IDs, audit data, and
// administrative relationships.
type SellingProductClassification struct {
	Brands       []SellingProductClassificationRef `json:"brands,omitempty"`
	Collection   *SellingProductClassificationRef  `json:"collection,omitempty"`
	CategoryTags []SellingProductClassificationRef `json:"category_tags,omitempty"`
}
