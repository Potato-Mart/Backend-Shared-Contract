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
