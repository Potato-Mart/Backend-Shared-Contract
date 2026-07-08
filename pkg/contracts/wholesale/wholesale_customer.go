package wholesale

// WholesaleCustomer is a compatibility name for the canonical wholesale
// organisation/business account. People inside the organisation are managed via
// OrganisationAccess records rather than separate wholesale customer groups.
type WholesaleCustomer = WholesaleOrganisation

// WholesaleCustomerSummary is a compatibility name for the compact wholesale
// organisation/business-account projection.
type WholesaleCustomerSummary = WholesaleOrganisationSummary
