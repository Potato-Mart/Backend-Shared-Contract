# Identity data model

Shared Contract defines identity records only. Backend-Identity owns authentication, credentials, token issuance, permission catalogues and role policy. Each backend owns authorization enforcement, transport and persistence.

## Principal and account records

- `account.UserProfile` is the principal profile, without credentials or global role policy.
- `AuthIdentity` records a provider identity in an explicit identity domain, without password hashes, OAuth tokens or private authentication material. Matching email addresses in different domains do not imply linked identities.
- `UserAccount` is a user-owned portal persona; a user can hold multiple accounts.
- `PortalAccess` records account-to-portal grants or revocations. Portal policy is service-owned.
- `RoleAssignment` records a scoped role grant; evaluation is service-owned.

## Roles and permissions

`identity/authorisation.Role` uses an open `RoleCode`. Built-in keys use the closed `role_enums.UserRole` vocabulary; custom roles can use other nonempty codes. Built-in ranks are omitted for custom roles.

| Rank | Built-in key | Geographic authority |
| --- | --- | --- |
| 1 | `superAdmin` | Global |
| 2 | `countryAdmin` | Country |
| 3 | `depotManager` | Granted depots |
| 4 | `marketing` | Granted markets |
| 5 | `warehouseManager` | Granted depots |
| 6 | `warehouseOperator` | Granted depots |

Rank orders authority, not geographic breadth. System roles cannot be deleted; a super administrator can adjust their permissions. Register access is a scoped permission, without a separate selling role.

`authorisation.PermissionKey` is an open opaque string. Identity owns its catalogue and validation. `PermissionDefinition` carries display, module, risk, MFA and classification metadata. Buyer-portal permissions use the separate Customers-owned `wholesale.WholesalePermission` catalogue. See [permission catalogue](permission-catalogue.md).

## Geographic selections

`access.StaffGeoScope` contains `mode` (`GLOBAL` or `TARGETED`) and `targets`, an array of `geography.GeographicPath`. Each path is a contiguous prefix of Country → Market → State → Depot, using `country_code`, `market_code`, `state_code` and `depot_code`. State is an official subdivision. Independent paths are separate selections, not Cartesian combinations.

Services validate ancestry and intersect selections with role-derived grants. Selecting `GLOBAL` cannot elevate a role. `UserProfile.geo_scope` is optional and absent for customer workforce grants. Depots are the platform site identity, including depots trading as stores.

## Session and business records

Login sessions and access records retain domain, principal, account and portal identifiers needed for isolation. Token claims, refresh storage, correlation and access evaluation remain service-owned.

Retail business details use `retail.RetailCustomer`. Wholesale details use `wholesale.WholesaleOrganisation`, linked to people through `OrganisationAccess`. Organisation approval, account state and portal access are distinct records. `party.OrganisationDetail` is the shared organisation value shape, without onboarding or authorization workflow.
