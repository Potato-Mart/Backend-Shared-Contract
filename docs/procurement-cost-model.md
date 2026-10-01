# Procurement cost models

Supply owns purchase-order records under `supply/procurement` and validates supplier cost basis, currency, quantities and totals.

`PurchaseOrderItem.package_option` identifies the frozen SKU/package code/version. Product package facts and ordered package composition preserve the selected units-per-package, package count and total base units.

`unit_cost` is optional Money for one base unit. `selected_package_cost` is optional Money for ONE package of the frozen selected package option. These are mutually exclusive supplier-price bases. An absent field is missing evidence; an explicit zero amount with currency is zero-cost evidence. A package price must not be divided and rounded into a fabricated per-base-unit supplier price.

Two CASE12 packages priced at AUD 1,000 minor units each carry `selected_package_cost = {amount_minor:1000,currency:"AUD"}`, package count 2, total base units 24, and `line_total = {amount_minor:2000,currency:"AUD"}`. `unit_cost` is absent in that example. Supply verifies all counts, package references, arithmetic and currency; no calculations execute in this module.

`line_total` retains its line-total meaning. `net_goods_amount` remains the frozen goods total after goods discounts and excluding tax, freight and duty. Receipt and invoice valuation evidence is unchanged. Base-unit valuations do not redefine the selected-package supplier price.
