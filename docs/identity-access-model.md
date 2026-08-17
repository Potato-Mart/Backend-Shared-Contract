# Identity Data Model

This module defines shared identity data only. Authentication flows, password
handling, token issuance, portal admission, permission resolution, HTTP
transport, persistence, and authorization policy belong to Backend-Identity
and to each backend's local enforcement layer.

## Core records

`User` is the canonical principal profile. It contains no credential or global
role policy.

`AuthIdentity` records a provider identity in an explicit identity domain. It
stores provider and verification metadata, but never password hashes, refresh
token hashes, OAuth tokens, passkey private material, or private IdP material.
The same normalized email can therefore be represented independently by
different identity-domain records without implying a link between them.

`UserAccount` is a portal-facing account/persona owned by a user. One user can
have multiple account records, including both customer account types.

`PortalAccess` is the persisted grant or revocation record between an account
and a portal. It is data consumed by Management's portal policy; this module
does not map portals to domains or account types.

`RoleAssignment` is a persisted scoped role grant. Role-to-permission policy
and access evaluation are backend-owned business rules.

## Workforce roles

`role_enums.UserRole` holds the seven built-in workforce roles. The wire values
are camelCase and stable; `role.Role.rank` carries the hierarchy position, 1
(widest) through 7 (narrowest), and is omitted on custom roles.

| Rank | Role | Scope level | Responsibility |
| --- | --- | --- | --- |
| 1 | `superAdmin` | `global` | every country, market, and depot |
| 2 | `countryAdmin` | `country` | one country and everything inside it |
| 3 | `depotManager` | `depot` | the depots it is granted |
| 4 | `sales` | `depot` | selling from the depots it is granted |
| 5 | `marketing` | `market` | the markets it is granted |
| 6 | `warehouseManager` | `depot` | depot warehouse operations |
| 7 | `warehouseOperator` | `depot` | depot warehouse tasks |

`role.Role` records the role key, its label, its permission list, its rank, and
whether it is one of the seven system roles. System roles cannot be deleted,
but a `superAdmin` may adjust their permissions. The contract does not define
the permission strings themselves — Backend-Identity owns that matrix, and each
backend enforces it independently.

Point-of-sale sign-in is a permission held by every staff role rather than a
role of its own; there is no cashier role.

## Geographic scope

Ranks 2 and below are geographically scoped. `access_enums.ScopeLevel` names
the breadth:

| Level | Meaning | Fields that carry the grant |
| --- | --- | --- |
| `global` | everything | none |
| `country` | one country | `country_code` |
| `market` | the granted markets | `market_ids` (with their `country_code`) |
| `depot` | the granted depots | `depot_codes` |

`access.StaffGeoScope` is the persisted grant,
`{level, country_code, market_ids, depot_codes}`. It hangs off a workforce
profile as `account.UserProfile.geo_scope` and is absent on customer profiles.

**Depots are the only site identity in the platform.** Some depots trade as
stores, but there is no store entity, no store code, and no store scope level.
A site-scoped principal is scoped by depot code.

Staff-visible records across the domains carry denormalized `market_id`,
`country_code`, and — where a site applies — `depot_code`, so a scoped query is
a plain indexed match rather than a join through market and depot tables. The
contract declares those fields; each service owns the indexes and the filter
injection.

## Sessions and claims

Identity claims, login sessions, refresh records, portal decisions, and
security/audit events carry the stable identifiers needed to preserve domain,
account, and portal isolation, including:

- `auth_identity_id`
- `identity_domain`
- `user_id`
- `account_id`
- `account_type`
- `portal`
- `audience`
- scoped role and organisation identifiers when applicable

`access.AccessTokenClaims` additionally carries the workforce scope as four
flat, optional claims — `scope_level`, `country_code`, `market_ids`, and
`depot_codes` — mirroring `StaffGeoScope`. They are absent on customer and
service-to-service tokens.

Consumers **fail closed**: a workforce token that carries no valid
`scope_level` is rejected, never treated as global. `superAdmin` and
service-to-service callers resolve to global scope. Events and
service-to-service calls are never geographically truncated; scoping applies to
staff-facing reads and writes.

The shared models record those values. Issuers and consumers independently
enforce their required claim tuples and audience policy.

## Customer and organisation records

Retail customer business details use `customers.RetailCustomer`.

Wholesale organisation business details use
`wholesale.WholesaleOrganisation`, with people linked through
`wholesale.OrganisationAccess`. Organisation approval, portal grants, account
state, and organisation access remain distinct persisted records.

`common.OrganisationDetail` provides the shared organisation value shape used
by suppliers and wholesale organisations. It is a data model, not an onboarding
or authorization workflow.
