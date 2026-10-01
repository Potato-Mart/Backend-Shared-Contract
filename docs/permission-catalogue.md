# Permission catalogue

Shared Contract publishes permission shapes; owning services publish catalogue contents and validate keys.

| Contract type | Catalogue owner |
| --- | --- |
| `identity/authorisation.PermissionKey` | Backend-Identity workforce permissions |
| `customers/wholesale.WholesalePermission` | Backend-Customers buyer-portal permissions |

Both types are open strings without fixed permission constants or membership-validation methods. Consumers treat keys as opaque strings. Identity validates workforce keys against its catalogue; Customers owns buyer-role policy and buyer-portal validation. Identity must not validate buyer permissions against the workforce catalogue.

`authorisation.Role.permissions` uses `[]PermissionKey`. `access.LoginSession.permissions` uses `[]string` because a session can carry permissions from both catalogues.

## Workforce metadata

`authorisation.PermissionDefinition` contains:

| Field | Meaning |
| --- | --- |
| `key` | Opaque permission identity |
| `label` | Display label |
| `description` | Optional description |
| `module` | Owning functional module |
| `risk_level` | `low`, `medium`, `high` or `critical` |
| `requires_mfa` | MFA requirement metadata |
| `classification` | `ui`, `field-level`, `service-only` or `intentionally-reserved` |

Identity owns these records and role-to-permission policy. Shared models supply no seeded keys, grants, datastore schema or authorization implementation. Catalogue contents belong in service documentation, where their owner can keep them authoritative.
