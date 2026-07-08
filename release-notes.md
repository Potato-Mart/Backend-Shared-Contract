# v14.1.0

Changes since `v14.0.0`.

## Added

- Added optional product collection/category slug fields:
  - `product.CollectionRef.slug`
  - `product.Collection.slug`
  - `product.CategoryTag.slug`
- Added optional account avatar projection fields:
  - `identity.UserProfile.avatar_media_id`
  - `identity.UserProfile.avatar_url`
- Added JSON-shape tests for slug and avatar field inclusion/omission.

## Compatibility Notes

- Minor additive release; no `/v14` module path migration.
- All new fields are tagged `omitempty`.
- Consumers should upgrade deliberately and run contract serialization/deserialization tests.

## Contract Files Changed

- `pkg/contracts/identity/user.go`
- `pkg/contracts/identity/user_json_test.go`
- `pkg/contracts/product/category_tag.go`
- `pkg/contracts/product/collection.go`
- `pkg/contracts/product/json_shape_test.go`
- `pkg/versioning/version.go`
