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

`role_enums.UserRole` holds the six built-in workforce roles. The wire values
are camelCase and stable; `role.Role.rank` carries the hierarchy position, 1
(widest) through 6 (narrowest), and is omitted on custom roles.

| Rank | Role | Scope level | Responsibility |
| --- | --- | --- | --- |
| 1 | `superAdmin` | `global` | every country, market, and depot |
| 2 | `countryAdmin` | `country` | one country and everything inside it |
| 3 | `depotManager` | `depot` | the depots it is granted |
| 4 | `marketing` | `market` | the markets it is granted |
| 5 | `warehouseManager` | `depot` | depot warehouse operations |
| 6 | `warehouseOperator` | `depot` | depot warehouse tasks |

`role.Role` records the role key, its label, its permission list, its rank, and
whether it is one of the six system roles. System roles cannot be deleted,
but a `superAdmin` may adjust their permissions. The contract does not define
the permission strings themselves — Backend-Identity owns that matrix, and each
backend enforces it independently. `role.PermissionKey` is the open typed code
those strings travel as; it carries no constants, so the contract names the
shape of a permission key without naming a single one. `role.PermissionDefinition`
is the catalogue-metadata record Backend-Identity seeds — key, label,
description, module, risk level, MFA requirement, and classification — for one
permission it owns. A consumer therefore treats a permission key as an opaque
string and never validates one locally; Identity validates a key against the
seeded catalogue, including its retired-key deny list.

**There is no selling rank.** Point-of-sale sign-in is a permission every rank
may hold, bounded by the geographic scope that rank is granted, so neither a
`cashier` nor a `sales` role exists. Who may work a register is answered by
scope — which depots or markets the principal holds — not by a separate rank,
and a dedicated selling rank therefore carried no authority the remaining six
did not already carry.

The four retired keys `admin`, `warehouse`, `cashier`, and `sales` never
validate. A stored role document or a live token carrying one is not a role.

## Geographic scope

Ranks 2 and below are geographically scoped. `access_enums.ScopeLevel` names
the breadth, widest to narrowest — country, then market, then depot:

| Level | Meaning | Fields that carry the grant | Ranks that hold it |
| --- | --- | --- | --- |
| `global` | everything | none | 1 `superAdmin` |
| `country` | one country | `country_code` | 2 `countryAdmin` |
| `market` | the granted markets | `market_codes` (with their `country_code`) | 4 `marketing` |
| `depot` | the granted depots | `depot_codes` | 3 `depotManager`, 5 `warehouseManager`, 6 `warehouseOperator` |

Rank orders **authority**, not geographic breadth, and the two ladders do not
line up: rank 3 `depotManager` is depot-scoped while rank 4 `marketing` is
market-scoped, so the lower rank holds the narrower geography. A consumer
resolves what a principal may see from `scope_level` and its grant fields, never
by inferring breadth from the rank number.

`access.StaffGeoScope` is the persisted grant,
`{level, country_code, market_codes, depot_codes}`. It hangs off a workforce
profile as `account.UserProfile.geo_scope` and is absent on customer profiles.

**Depots are the only site identity in the platform.** Some depots trade as
stores, but there is no store entity, no store code, and no store scope level.
A site-scoped principal is scoped by depot code.

Staff-visible records across the domains carry denormalized `market_code`,
`country_code`, and — where a site applies — `depot_code`, so a scoped query is
a plain indexed match rather than a join through market and depot tables. The
contract declares those fields; each service owns the indexes and the filter
injection.

## Sessions and access records

`access.LoginSession`, `PortalAccess`, and relevant security/audit records
carry the stable identifiers needed to preserve domain, account, and portal
isolation, including:

- `auth_identity_id`
- `identity_domain`
- `user_id`
- `account_id`
- `account_type`
- `portal`
- `audience`
- scoped role and organisation identifiers when applicable

Token-claim schemas, refresh-token storage, request correlation, and
authorization evaluation are Backend-Identity implementation concerns. The
shared contract deliberately does not define those transport or persistence
shapes. Each backend independently applies `StaffGeoScope` to its own
staff-facing reads and writes.

## Customer and organisation records

Retail customer business details use `retail.RetailCustomer`.

Wholesale organisation business details use
`wholesale.WholesaleOrganisation`, with people linked through
`wholesale.OrganisationAccess`. Organisation approval, portal grants, account
state, and organisation access remain distinct persisted records.

`party.OrganisationDetail` provides the shared organisation value shape used
by suppliers and wholesale organisations. It is a data model, not an onboarding
or authorization workflow.
