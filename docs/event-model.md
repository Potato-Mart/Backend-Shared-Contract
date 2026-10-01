# Event models

`pkg/contracts/pubsub/envelope.EventEnvelope` contains event identity, type, version, occurrence time, aggregate identity and typed JSON payload evidence. Event payloads live under their owning domain; commercial listing facts are Pricing-owned, while physical inventory facts are Supply-owned.

Runtime event versions, publication, authorization, idempotency and delivery are defined by the owning services. A module version or payload table does not prescribe an envelope version.

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
