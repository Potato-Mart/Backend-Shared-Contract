# V6 Customer Contract Migration

V6 deliberately separates retail/general customer profiles from wholesale
customer profiles. Portal admission continues to use `AccountType` and
`PortalAccess`; customer profile contracts describe business details after the
account/persona has been selected.

## Type Moves

| V5 contract | V6 contract |
| --- | --- |
| `customers.Customer` | `customers.RetailCustomer` |
| `customers.CompanyCustomer` | `wholesale.WholesaleCustomer` |
| `customers.WholesaleTerms` | `wholesale.WholesaleTerms` |
| `customers.CustomerActivity` | `customers.RetailCustomerActivity` |
| `customers.CustomerIdentity` | `customers.RetailCustomerExternalIdentity` |
| `enums.CustomerProfileStatus` | `enums.CustomerStatus` |
| `enums.CustomerSegment` | Removed from customer profile contracts |
| `common.CompanyDetail` | `common.OrganisationDetail` |

## Organisation Detail Shape

V6 uses `common.OrganisationDetail` as the single shared company/organisation
profile. It is embedded by contracts such as `purchase.Supplier` and
`wholesale.WholesaleOrganisation`.

Old company-prefixed fields move to organisation-neutral JSON keys:

| V5 field / JSON key | V6 field / JSON key |
| --- | --- |
| `CompanyDescription` / `company_description` | `Description` / `description` |
| `CompanyABN` / `company_abn` | `ABN` / `abn` |
| `CompanyACN` / `company_acn` | `ACN` / `acn` |
| `CompanyWebsite` / `company_website` | `Website` / `website` |
| `CompanyContactPerson` / `company_contact_person` | `ContactPerson` / `contact_person` |
| `CompanyAddress` / `company_address` | `RegisteredAddress` / `registered_address` |
| `CompanyLogoURL` / `logo_url` | `LogoURL` / `logo_url` |

## Retail Customer Shape

`RetailCustomer` represents a `generalCustomer` account/persona profile. It is
grouped by responsibility:

- `identity`: canonical user, account, and auth identity references.
- `basic_info`: customer number, name, contact channels, date of birth, and acquisition source.
- `lifecycle`: retail customer profile status and lifecycle timestamps.
- `management`: staff-managed CRM notes, tags, and sales owner.
- `loyalty`: loyalty points, tier, and tier evaluation counters.
- `marketing`: per-channel marketing consent.
- `commerce`: aggregated order statistics.
- `analytics`: optional RFM/churn analytics.
- `referral`: optional referral-programme state.

## Wholesale Customer Shape

`WholesaleCustomer` represents a `wholesaleCustomer` account/persona profile
tied to a wholesale organisation and membership. It is grouped by
responsibility:

- `identity`: canonical user, account, and auth identity references.
- `organisation_id`: owning wholesale organisation.
- `membership_id`: wholesale membership used for organisation-scoped access.
- `basic_info`: contact person details and acquisition source.
- `commercial`: sales owner, CRM tier, payment terms, tax ID, notes, and tags.
- `account_profile`: customer profile status, role key, and membership-facing flags.
- `terms`: pricing tier, rebate, shipping fee, and freight-rule references.

Organisation business details are inherited from `common.OrganisationDetail`.
Wholesale approval remains on `WholesaleOrganisation`. Membership status
remains on `WholesaleMembership`.

## Common Primitives

V6 adds reusable primitives in `pkg/common`:

- `OrganisationDetail`
- `IdentityLink`
- `PersonName`
- `ContactChannels`

Customer-specific business groups such as loyalty, marketing, CRM, RFM, and
commerce stats remain in their domain packages.
