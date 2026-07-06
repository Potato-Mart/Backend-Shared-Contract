# v13.0.0

Changes since `v12.0.0`.

## Breaking Contract Changes

- Go module path changes to `github.com/Potato-Mart/Backend-Shared-Contract/v13`.
- Consumers must update imports and dependency pins from `/v12` to `/v13`.
- No enum JSON wire values change in this release.

## Added

- `membership.Reward` adds `trigger_tier_key` for rewards tied to a specific membership tier achievement.
- `membership.Reward` adds `issue_on_tier_achievement` so services can mark rewards for automatic issue when the tier is achieved.

## Release Engineering

- Split the previous large `pkg/enums/enums_test.go` into small domain-focused enum test files under `pkg/enums`.
- Added shared enum test assertions in `pkg/enums/enum_assertions_test.go`.
- Added standalone test scripts:
  - `scripts/Test-Contract.ps1`
  - `scripts/test-contract.sh`
- Added PR/push contract test workflow and configured release tests with `GOWORK=off`.

## Compatibility Notes

- Existing reward JSON remains readable because the new reward fields are `omitempty`.
- Run `GOWORK=off go test ./...` or one of the test scripts after upgrading.
