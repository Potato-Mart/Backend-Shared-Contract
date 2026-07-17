# Backend-Shared-Contract v18

Potato Mart's shared Go module for cross-service data models.

Version 18 is deliberately model-only. It contains domain entities, records,
snapshots, events, value objects, typed enums, field names, JSON/BSON
serialization, shared error-code enums, and module-version metadata.

It does not contain endpoint request or response DTOs, paths, route inventories,
service-scope catalogues, HTTP envelopes, validation policy, authorization
policy, state transitions, calculations, normalizers, or business workflows.
Each backend owns those concerns and publishes its transport contract in its own
OpenAPI document.

## Version and module path

```text
v18.0.0
github.com/Potato-Mart/Backend-Shared-Contract/v18
```

Consumers pin the released module directly:

```go
require github.com/Potato-Mart/Backend-Shared-Contract/v18 v18.0.0
```

For example:

```go
import (
    walletenum "github.com/Potato-Mart/Backend-Shared-Contract/v18/pkg/enums/wallet"
)
```

Version 18 is the customer-commerce contract cutover. It replaces manually
authored product velocity fields with computed historical sales statistics and
category ranks, adds Operations-owned retail and organisation favourite lists,
adds terminal notification dismissal/expiry state, and introduces one
customer-safe storefront product projection with audience-filtered pricing and
promotion badges. Legacy Commerce wholesale list permissions are removed.

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
