package packaging

// PackageOptionRef identifies one immutable physical package version. Code is
// EACH for one base unit or CASE followed by UnitsPerPackage (e.g. CASE6).
// Version is a positive immutable ordinal within SKUCode+Code. Changed physical
// specifications create a new version; financial and stock evidence retains
// the original tuple. Supply owns code derivation and version allocation.
type PackageOptionRef struct {
	SKUCode string `json:"sku_code"`
	Code    string `json:"code"`
	Version int64  `json:"version"`
}
