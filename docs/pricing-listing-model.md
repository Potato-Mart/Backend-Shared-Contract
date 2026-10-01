# Pricing listing models

Pricing owns `pricing/listing.MarketListing` and `SaleRestriction`, their `listing_enums`, and `pubsub/pricing.CatalogListingChangedEvent`. Each has one canonical definition.

A listing links `market_code` and `sku_code` with lifecycle status, tax-category reference, sale restrictions, availability window, expiry lead-time override, revision and audit evidence. Product display names are catalogue-owned. Listings contain no display-name override, unit-pricing flag or authoritative price amount; prices are held in Pricing price books.

Listing statuses are `draft`, `coming_soon`, `active`, `suspended`, `unavailable` and `delisted`. Restriction kinds are `age`, `quantity_limit`, `channel_excluded`, `delivery_excluded` and `prescription`.

For `age`, `age_years` is an integer threshold in completed calendar years and `age_comparison` is `below` or `above`. Below blocks an age strictly less than the threshold; above blocks strictly greater. Equality is allowed. The pointer preserves explicit zero versus an absent threshold; services validate the configured value.

Age rules intrinsically apply only to retail buyers, independent of order channel. Services use the authenticated customer's saved Customers DOB and the applicable market calendar date. Guests and retail buyers with missing or unusable DOB may view products but cannot purchase age-restricted items. Wholesale buyers bypass the age rule. No proof verification or buyer DOB/exact-age field is included in listing policy.

For `delivery_excluded`, `excluded_delivery_methods` enumerates the existing `delivery`, `pickup` and `outsourced` methods. Empty or missing collections do not stand for an implicit method. Quantity limits retain `value`, channel exclusions retain `channels`, and prescription restrictions retain their meaning. `note` is descriptive text, not an encoded policy payload. Services own kind-specific payload validation and enforcement.

The listing event carries identity, status, tax reference, expiry override, availability and revisions. It carries no display override or unit-pricing flag. Services retrieve authoritative current restrictions through their owner interfaces; the event is not a purchase decision.

Supply owns `supply/catalogue/listing.SaleEligibilitySnapshot` and `DamageSaleApproval`, DepotMarket, catalogue and physical inventory evidence. Listing revision and tax-category fields in stock evidence reference Pricing facts; inventory revisions and validity tokens reference Supply facts.
