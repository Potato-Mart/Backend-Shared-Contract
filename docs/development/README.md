# Development Notes

Read [Git Workflow](git-workflow.md) for branches, PRs and checks, and [Package Boundaries](package-boundaries.md) for code ownership and dependency direction.

## Verification and versioning

Run `./scripts/powershell/Test-ReleaseAlignment.ps1` and `./scripts/powershell/Test-Contract.ps1` from the repository root. Tests run independently with `GOWORK=off`. Release metadata is declared in `go.mod`; module majors follow semantic versioning and published tags remain immutable. Consumers pin released versions in their own repositories.

## Event payload reference

This reference covers the catalogue, inventory, order and analytics payload set checked by the repository. Other payloads are defined alongside these under `pkg/contracts/pubsub`.

| # | Payload | Event type | Topic |
| --- | --- | --- | --- |
| 1 | `StockLocationAvailabilityChangedEvent` | `stock.location_availability_changed` | `stock-events` |
| 2 | `InventoryLotReceivedEvent` | `inventory.lot_received` | `stock-events` |
| 3 | `InventoryStockBucketChangedEvent` | `inventory.stock_bucket_changed` | `stock-events` |
| 4 | `InventoryPackageConvertedEvent` | `inventory.package_converted` | `stock-events` |
| 5 | `InventoryQualityAssessedEvent` | `inventory.quality_assessed` | `stock-events` |
| 6 | `InventoryReservationChangedEvent` | `inventory.reservation_changed` | `stock-events` |
| 7 | `StockStagingChangedEvent` | `inventory.staged` | `stock-events` |
| 8 | `InventorySaleCommittedEvent` | `inventory.sold` | `stock-events` |
| 9 | `InventoryDateMarkThresholdEvent` | `inventory.date_mark_threshold_reached` | `stock-events` |
| 10 | `OrderPaidEvent` | `order.paid` | `order-events` |
| 11 | `RefundCompletedEvent` | `refund.completed` | `refund-events` |
| 12 | `OrderFact` | `analytics.order_fact` | `order-events` |
| 13 | `RefundFact` | `analytics.refund_fact` | `refund-events` |
| 14 | `ProductSalesRollup` | `product.sales_performance_updated` | `product-stats` |
| 15 | `CatalogBaseCostChangedEvent` | `catalog.base_cost_changed` | `catalog-events` |
| 16 | `CatalogListingChangedEvent` | `catalog.listing_changed` | `catalog-events` |
| 17 | `OrderPackingProjection` | `fulfilment.packing_updated` | `fulfilment-events` |

## Release index

| Version | Release date | Type |
| --- | --- | --- |
| `v42.0.1` | 2026-10-04 | Patch |
| `v42.0.0` | 2026-10-02 | Major |

## v42.0.1 (2026-10-04)

Consolidate repository documentation into OpenAPI status and three development files. Update documentation-dependent tests and release tooling to the surviving paths. Shared production model files and their JSON shapes are unchanged from v42.0.0. The patch metadata preserves immutable release payload checks without rewriting the existing tag.

Historical source and complete release records remain available through published Git tags and GitHub releases.
