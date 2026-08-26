package wholesale

// WholesalePermission is a persisted or signed buyer-portal permission value
// such as "products.view" or "checkout.submit". It is an open typed string:
// the concrete key catalogue, buyer-role resolution, and forbidden-permission
// policy are configuration owned and seeded by Backend-Customers, not a closed
// contract enum.
type WholesalePermission string
