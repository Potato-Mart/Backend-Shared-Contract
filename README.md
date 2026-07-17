# Backend-Shared-Contract v17

Potato Mart's shared Go module for cross-service data models.

Version 17 is deliberately model-only. It contains domain entities, records,
snapshots, events, value objects, typed enums, field names, JSON/BSON
serialization, shared error-code enums, and module-version metadata.

It does not contain endpoint request or response DTOs, paths, route inventories,
service-scope catalogues, HTTP envelopes, validation policy, authorization
policy, state transitions, calculations, normalizers, or business workflows.
Each backend owns those concerns and publishes its transport contract in its own
OpenAPI document.

## Version and module path

```text
v17.4.0
github.com/Potato-Mart/Backend-Shared-Contract/v17
```

Consumers pin the released module directly:

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v17 v17.4.0
```

For example:

```go
import (
    walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v17/pkg/enums/wallet"
)
```

Version 17 is a breaking wallet/checkout cutover. Wallet-export contracts are
gone; coupon ownership supports retail and wholesale owners; gift cards expose
committed, reserved, and available balances; and checkout benefit reservations,
redemption snapshots, and gift-card payment references are first-class models.

Version 17.2 adds model-only import-compliance records for revisioned settings,
manufacturer declarations, product labels, tariff assessments/profiles,
trademark evidence, Requests for Inspection, source evidence, catalogue pins,
and generated-artifact references. Monetary and calculated values use fixed
precision (`common.Money`, basis points, micros, grams, and cubic centimetres),
while HTTP workflows and regulatory decisions remain backend-owned.

Version 17.2.1 republishes the same model surface after the repository release
alignment checks were strengthened. Consumers do not need model or wire-shape
changes when upgrading from 17.2.0.

Version 17.3 adds a stable optional identifier for persisted contact-address
book entries and a customer-safe storefront promotion projection. Backend-owned
HTTP routes and request DTOs remain outside this module.

Version 17.4 adds session-bound access-token and exact last-login-IP fields,
provider-neutral membership-pass content, and wholesale fixed/on-request price
mode enums. Provider routes, signing payloads, service scopes, and media types
remain backend-owned.

## Boundary governance

The package manifest and AST boundary tests reject unclassified exported
models, transport tags, endpoint DTO naming, paths, scopes, free business
functions, and non-intrinsic receiver methods. Approved receiver behavior is
limited to serialization plus single-value `String` and `IsValid` methods.

## Verification

To avoid a parent workspace changing dependency resolution, run:

```powershell
./scripts/Test-Contract.ps1
```

or:

```bash
bash scripts/test-contract.sh
```

The equivalent command is:

```bash
GOWORK=off go test ./...
```

Published model changes follow semantic versioning. Breaking exported shape or
wire-value changes require a new major module path.
