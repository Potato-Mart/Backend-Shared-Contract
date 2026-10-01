# Pricing listing models

Pricing owns `pricing/listing.MarketListing` and `SaleRestriction`, their `listing_enums`, and `pubsub/pricing.CatalogListingChangedEvent`. Each has one canonical definition.

A listing links `market_code` and `sku_code` with lifecycle status, tax-category reference, localized display names, sale restrictions, availability window, expiry lead-time override, unit-pricing requirement, revision and audit evidence. It contains no authoritative price amount. Prices are held in Pricing price books.

Listing statuses are `draft`, `coming_soon`, `active`, `suspended`, `unavailable` and `delisted`. Sale-restriction kinds are `age_verification`, `quantity_limit`, `channel_excluded`, `delivery_excluded` and `prescription`. Validation and enforcement belong to services.

Supply owns `supply/catalogue/listing.SaleEligibilitySnapshot` and `DamageSaleApproval`, DepotMarket, catalogue and physical inventory evidence. Listing revision and tax-category fields in stock evidence reference Pricing facts; inventory revisions and validity tokens reference Supply facts.
