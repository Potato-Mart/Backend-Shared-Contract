# Package Boundaries

This repository is a model-only Go library. It has no runtime server, routes, HTTP DTOs, validation policy, database or provider adapters.

- `pkg/contracts/common` owns shared value objects and primitives.
- Service-owned packages under `pkg/contracts` own their domain records and snapshots.
- `pkg/contracts/pubsub` owns typed event envelopes and payload models; services own publication, runtime event versions and delivery.
- Finite enums live in leaf `_enums` packages.
- `pkg/test` owns model shape, inventory, enum, ownership and model-only boundary checks.

Services retain their own business workflows, transport DTOs, validation, authorisation and persistence. Shared models must not create cross-service database or runtime coupling.
