# Permission Catalogue Handoff

## Purpose

Contract v32 removes every hard-coded permission key from the shared contract.
The contract now publishes the permission *types*; the concrete key catalogues,
their validation, and their retired-key deny lists move to the backend service
that owns each catalogue and are seeded there.

This document is the implementation brief for that move. No service repository
is modified by the contract change, so the seeding work described below must
land in the owning services **before** they adopt `/v32`.

Two catalogues are in scope, with two different owners:

| Catalogue | Contract type after v32 | Owning service |
| --- | --- | --- |
| Workforce (staff/admin) permissions, 102 keys | `role.PermissionKey` | Backend-Identity |
| Buyer-portal (wholesale) permissions, 18 keys | `wholesale.WholesalePermission` | Backend-Customers |

## Why this changed

`docs/identity-access-model.md` already stated that the contract does not
define the permission strings and that Backend-Identity owns that matrix. A
closed 102-value enum in the contract contradicted that, and it contradicted
the repository's model-only governance: `pkg/test/model_only_boundary_test.go`
rejects package-level variables, non-enum constants, and free functions, so a
catalogue can only ever have entered the contract through the enum loophole.
Both catalogues are now open typed strings, matching `money.CurrencyCode` and
`geography.CountryCode`, whose vocabularies are likewise service-owned.

A repository gate, `TestPermissionKeysRemainOpenCodes` in
`pkg/test/domain_boundary_test.go`, fails the build if either catalogue
reappears in the contract as constants or as fixed dotted wire literals.

## Contract surface after v32

```go
// pkg/contracts/identity/role/permission_key.go
type PermissionKey string        // open code, no constants, no methods

// pkg/contracts/customers/wholesale/wholesale_permission.go
type WholesalePermission string  // open code, no constants, no methods
```

Still shipped by the contract, unchanged:

- `role.Role` — `Permissions []role.PermissionKey`, `Key role_enums.UserRole`.
- `role.PermissionDefinition` — the catalogue-metadata record described below.
- `role_enums.PermissionClassification` — closed enum: `ui`, `field-level`,
  `service-only`, `intentionally-reserved`.
- `role_enums.UserRole` — closed enum of the six built-in workforce roles.
- `security_enums.SecurityRiskLevel` — closed enum: `low`, `medium`, `high`,
  `critical`.
- `access.LoginSession.Permissions` remains `[]string`, deliberately: one
  session can carry both workforce and buyer-portal permissions, so it is not
  typed to either catalogue.

`PermissionKey` and `WholesalePermission` no longer expose `IsValid()` or
`String()`. Call sites that validated a key against the contract must now
validate against the seeded catalogue in the owning service; call sites that
formatted a key use a `string(...)` conversion.

---

## Backend-Identity: the workforce catalogue

### Seed record shape

Seed one `role.PermissionDefinition`-shaped document per key:

```go
type PermissionDefinition struct {
	Key            PermissionKey                       `json:"key"`
	Label          string                              `json:"label"`
	Description    string                              `json:"description,omitempty"`
	Module         string                              `json:"module"`
	RiskLevel      security_enums.SecurityRiskLevel    `json:"risk_level"`
	RequiresMFA    bool                                `json:"requires_mfa"`
	Classification role_enums.PermissionClassification `json:"classification"`
}
```

The contract never shipped any populated instance of this struct, so
`Label`, `Description`, `Module`, `RiskLevel`, `RequiresMFA`, and
`Classification` are **new metadata Identity must author**. Only `Key` is
carried over from the deleted enum. Treat authoring that metadata as part of
this task, not as a mechanical copy.

### Where the seed belongs

Backend-Identity has no `migrations/` directory; bootstrap is the established
seeding surface. Recommended shape:

1. Keep the compile-time catalogue in `internal/security/rbac/permissions.go`,
   typed on the contract's `role.PermissionKey`, as the in-process source of
   truth for validation and claim minting.
2. Have `internal/security/bootstrap/bootstrap.go` upsert one definition
   document per key into a seed collection (for example `mgmt_permissions`)
   through `internal/platform/mongox`. The upsert must be idempotent so repeated
   startups converge rather than duplicate.
3. Let `internal/modules/roles` read the catalogue from that collection when
   serving role administration.

### Validation ownership

`PermissionKey.IsValid()` no longer exists. Identity owns membership validation
against the seeded catalogue for: role create/update writes, catalogue seeding
itself, and claim minting. Every other service treats a permission key as an
opaque string and must not attempt local validation.

### Retired-key deny list

These keys were previously rejected by the contract's `IsValid()` and must
continue to be rejected. They must never be seeded, never validate, and should
be stripped when found in a stored role document or a live token:

```
customer.read
customer.write
wholesale_customer.read
wholesale_customer.write
regular_customer.read
regular_customer.write
pos.shift.manage
```

### The 102 workforce keys

Grouped as the deleted contract file grouped them. Wire values are exact and
must not be re-spelled.

**Identity and directory (9)**
`user.read`, `user.write`, `user.delete`,
`retail_customer.read`, `retail_customer.write`,
`supplier.read`, `supplier.write`,
`role.read`, `role.write`

**Membership and wallet (6)**
`membership.read`, `membership.write`,
`wallet.read`, `wallet.write`,
`gift_card_policy.read`, `gift_card_policy.write`

**Promotion and coupon (5)**
`promotion.read`, `promotion.write`, `promotion.publish`,
`coupon.read`, `coupon.write`

**Pricing (2)**
`price.read`, `price.approve`

**Marketing and analytics (5)**
`marketing.read`, `marketing.send`,
`analytics.sales.read`, `analytics.product.read`, `analytics.export`

**Media (3)**
`media.read`, `media.upload`, `media.delete`

**Settings (2)**
`settings.read`, `settings.write`

**Audit (1)**
`audit.read`

**Security (2)**
`security.read`, `security.write`

**Access log (1)**
`access_log.read`

**Wholesale administration (2)**
`wholesale.read`, `wholesale.write`

**Catalog (10)**
`product.read`, `product.write`, `product.delete`,
`sku.read`, `sku.write`, `sku.products.read`,
`category.read`, `category.write`,
`collection.read`, `collection.write`

**Depot and location (4)**
`depot.read`, `depot.write`, `location.read`, `location.write`

**Stock (4)**
`stock.read`, `stock.adjust`, `stock.reserve`, `stock.transfer`

**Fulfilment (6)**
`picking.read`, `picking.write`,
`packing.read`, `packing.write`,
`shipment.read`, `shipment.write`

**Purchase and receipt (5)**
`purchase.read`, `purchase.write`, `purchase.publish`,
`receipt.read`, `receipt.write`

**Warehouse operations (12)**
`expiry.read`, `expiry.run`,
`forecast.read`, `forecast.write`,
`inbound.read`, `inbound.write`,
`damage.read`, `damage.write`,
`wmsdraft.read`, `wmsdraft.write`,
`layout.read`, `layout.write`

**Review and wish (4)**
`review.read`, `review.moderate`, `wish.read`, `wish.manage`

**Commerce (17)**
`order.read`, `order.write`, `order.cancel`,
`payment.read`, `payment.capture`,
`refund.read`, `refund.write`, `refund.request`, `refund.approve`,
`invoice.read`, `invoice.write`,
`cart.manage`,
`terminal.manage`, `terminal.transact`,
`commerce.config`,
`preorder.read`, `preorder.write`

**Point of sale (2)**
`pos.access`, `pos.session.manage`

### Tests to port

The deleted `permission_key_test.go` locked the catalogue. Reproduce its
guarantees in `internal/security/rbac/permissions_test.go`:

- every key's exact wire literal;
- the catalogue size is 102;
- no duplicate wire value;
- every deny-list key is rejected;
- an unknown key is rejected.

---

## Backend-Customers: the buyer-portal catalogue

`WholesalePermission` follows the same treatment. Backend-Customers owns the
catalogue, the buyer-role permission matrix (keyed by
`wholesale_enums.WholesaleBuyerRole`, which remains a closed contract enum),
and forbidden-permission policy.

The 18 keys, wire values exact:

```
products.view
cart.write
checkout.submit
orders.view_own
orders.view_org
orders.reorder
invoices.view_own
invoices.view_org
invoices.pay
account.view
team.view
favourite_lists.view_org
favourite_lists.write_org
group_orders.view_org
group_orders.manage_org
group_orders.invite
group_orders.submit
group_order_discount.apply
```

Seed and validate these the same way: an idempotent catalogue seed, a
service-owned membership check, and a test locking the wire values, the count
of 18, and uniqueness.

Identity must **not** validate buyer-portal strings against the workforce
catalogue. Identity only transports whatever Customers resolves, which is
exactly why `access.LoginSession.Permissions` stays `[]string`.

---

## Rollout

1. Land the Identity and Customers catalogue seeds and their tests, still
   compiled against `/v31`.
2. Verify the seeded catalogues in each service's datastore.
3. Confirm claim minting reads from the seeded catalogue rather than a contract
   constant.
4. Adopt the `/v32` module in each service, in the order given by the v32
   Consumer Migration Matrix in `docs/release-notes.md` (Identity first).
5. Run each service's suite.

## Impact on the other services

The remaining services compile against the open types and treat keys as opaque
strings. Because v32 was never tagged, no released consumer ever depended on
the deleted constants, so there is no migration shim to write.

One consequence worth noting: the v32 follow-ups call for removing the unused
`analytics.export` permission. With the catalogue seeded rather than compiled
into the contract, that removal is now a seed-data decision in Backend-Identity
and needs no contract release at all.
