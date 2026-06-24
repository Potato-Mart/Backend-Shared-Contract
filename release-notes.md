# v8.1.0

Changes since `v8.0.0`.

## Added
- Added `taxed` to `product.Product` and `product.Snapshot` for GST/FRE invoice rendering.
- Added product JSON shape coverage for the `taxed` field.

## Contract Files Changed
- `pkg/contracts/product/product.go`
- `pkg/contracts/product/snapshot.go`
- `pkg/contracts/product/json_shape_test.go`
- `pkg/versioning/version.go`

## Compatibility Notes
- Consumers should upgrade deliberately and run contract serialization/deserialization tests.
- Major versions may include removed fields, renamed JSON keys, enum value changes, or changed primitive shapes.
- Minor versions may add fields, contracts, enum values, or helper types in a backward-compatible way.
- Patch versions should contain documentation, comments, helper fixes, or non-breaking corrections.
