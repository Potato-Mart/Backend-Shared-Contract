# Pricing market-listing adoption

Pin `github.com/Potato-Mart/Backend-Shared-Contract/v39 v39.4.0` in the service module and matching CI input before implementation.

Canonical commercial listing models are `pkg/contracts/pricing/listing.MarketListing` and `SaleRestriction`, with `MarketListingStatus` and `SaleRestrictionKind` in `pricing/listing/listing_enums`. Pricing owns these definitions alongside markets, tax categories and price books. Listings contain availability and commercial configuration references, not authoritative price amounts.

The existing `supply/catalogue/listing` structs and enums remain frozen compatibility representations. Their imports, fields, optionality and enum wire values are unchanged. They are distinct Go types, not aliases; services map them explicitly at their boundaries. Contract tests enforce JSON shape and enum parity. Neither package is removed in this release. Removal requires a future major release after consumer adoption and rollback windows close.

Supply retains `SaleEligibilitySnapshot`, `DamageSaleApproval`, DepotMarket, catalogue and inventory evidence. Moving commercial listing ownership does not transfer physical-stock authority or reinterpret eligibility tokens. Existing `pubsub/supply.CatalogListingChangedEvent` and its wire remain unchanged for staged consumers; producer ownership, runtime event versions and routes are service-owned.

Supply and Pricing agreed this additive surface before their implementation. Upgrade both pins first, implement Pricing's listing authority and explicit compatibility mapping, then migrate Supply and consumer reads while retaining the current behavior until the replacement is ready. Orders and other existing consumers can continue using their current Supply imports. Frontend APIs, migration/query DTOs, validation, authorization, persistence, event publication and deployment belong to the owning service sessions. This model release alone does not establish runtime cutover or consumer readiness.
