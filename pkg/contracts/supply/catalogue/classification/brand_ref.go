package classification

// BrandRef is the stable relationship embedded in product records and
// snapshots. Display data is resolved from the brand master by its immutable
// code; mutable names, logos, slugs, and database identifiers are excluded.
type BrandRef struct {
	Code string `json:"code"`
}
