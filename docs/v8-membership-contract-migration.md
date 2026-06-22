# V8 Membership Contract Migration

V8 separates two concepts that were previously easy to confuse:

- **Global membership**: points, tiers, wallet balances, rewards, and recurring member subscriptions.
- **Wholesale organisation access**: a user/account grant that allows a person to act inside a wholesale organisation.

## Package Moves

| V7 contract | V8 contract |
| --- | --- |
| `loyalty.LoyaltyLedgerEntry` | `membership.PointLedgerEntry` |
| `loyalty.LoyaltyPointAllocation` | `membership.PointAllocation` |
| `loyalty.LoyaltyPointBucket` | `membership.PointBucket` |
| `loyalty.LoyaltyBalanceBreakdown` | `membership.PointBalanceBreakdown` |
| `loyalty.LoyaltyTier` | `membership.MembershipTier` |
| `subscription.SubscriptionPlan` | `membership.SubscriptionPlan` |
| `subscription.CustomerSubscription` | `membership.MemberSubscription` |
| `wholesale.WholesaleMembership` | `wholesale.OrganisationAccess` |
| `wholesale.WholesaleMembershipSummary` | `wholesale.OrganisationAccessSummary` |

## JSON Key Changes

Wholesale organisation access identifiers now use `organisation_access_id`.

| V7 JSON key | V8 JSON key |
| --- | --- |
| `membership_id` | `organisation_access_id` |
| `default_membership_id` | `default_organisation_access_id` |

Membership programme identifiers use `membership_account_id`.

## Points Ownership

Retail points are owned by the retail customer's `MembershipAccount`.

Wholesale and bulk-order points are owned by the wholesale organisation's
`MembershipAccount`. A contact spending organisation-owned points must have a
valid `OrganisationAccess` and a permission such as `membership.points.spend`.

## Redemption Flow

The membership ledger is the source of truth. Wallet balance fields are
projections for display and quote responses.

1. Quote the requested points spend.
2. Reserve points using FIFO by earliest expiry, then oldest earn row.
3. Commit the reservation when the order or reward redemption is final.
4. Cancel or expire the reservation if checkout does not complete.

Committed spends create a negative `PointLedgerEntry` with `allocations` that
show which earn rows were consumed.
