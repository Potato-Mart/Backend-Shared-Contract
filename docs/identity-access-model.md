# Identity Access Model

This module is contract-layer only. It defines shared enums, DTOs, error
codes, and event payloads. It does not implement authentication, token signing,
password hashing, HTTP middleware, database access, or portal admission logic.

## Core Model

`User` is the canonical human or principal. A user can have multiple login
identities and multiple account/persona records.

`AuthIdentity` is a non-secret login identity attached to a user. It records the
provider, domain, provider subject, email projection, verification state, and
status. It never carries password hashes, refresh token hashes, OAuth tokens,
passkey private material, or private IdP material.

`UserAccount` is the account/persona record attached to a user. It carries an
`AccountType`, lifecycle status, public display fields, and account-level audit
metadata. A user can have more than one account persona.

`PortalAccess` controls front-door platform admission. It answers whether a
specific account/persona may enter a specific portal.

`RoleAssignment` controls capabilities after portal admission. It assigns a
role key inside an account, portal, and optional business scope such as a
wholesale organisation.

## Account Type Is Not User Role

`AccountType` answers: is this account/persona allowed into this portal?

`UserRole`, `Role`, `Permission`, and `RoleAssignment` answer: what can this
account do after it is inside the portal?

`UserRole` remains exported for backward compatibility, but new services must
not use it as the sole platform gate. Use `AccountType` and `PortalAccess` for
portal admission.

## Portal Mapping

Each current portal accepts exactly one account type:

| Portal | Wire value | Required account type |
| --- | --- | --- |
| `PortalControl` | `control` | `adminUser` |
| `PortalStore` | `store` | `generalCustomer` |
| `PortalPartner` | `partner` | `wholesaleCustomer` |

`PortalPartner` is the wholesale/partner portal for the current v6 contract.

## Sessions And Claims

`LoginSession` is scoped to one portal and, for user sessions, one account
persona. Session and token-claim contracts can carry:

- `account_id`
- `account_type`
- `portal`
- `audience`
- `roles`
- `permissions`
- `wholesale_organisation_id`
- `organisation_access_id`
- `role_key`

Services that issue or validate sessions must check portal, audience, and
account type consistently. This repository only defines the shared shape.

## Customer Profile Contracts

Retail/general customer business details live in `customers.RetailCustomer` and
belong to the `generalCustomer` account/persona model.

Wholesale customer business details live in `wholesale.WholesaleCustomer` and
belong to the `wholesaleCustomer` account/persona model. Organisation approval
and organisation access lifecycle remain separate wholesale contracts.

Wholesale organisation business details use `common.OrganisationDetail`
embedded by `wholesale.WholesaleOrganisation`, so company/organisation identity,
registration, contact, address, branding, and metadata fields have one shared
shape across suppliers and wholesale organisations.

## Wholesale Access

Wholesale portal access is organisation and membership aware:

- `WholesaleOrganisation` records approval status and B2B organisation
  references.
- `OrganisationAccess` links a user account to a wholesale organisation and
  carries the organisation-scoped role key.
- A wholesale portal session can identify the account, organisation,
  organisation access grant, and role key used for that session.

Wholesale organisation status and organisation access status are separate from
general account lifecycle status and portal access status.
