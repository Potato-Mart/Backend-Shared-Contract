package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// PricePromotionStageRun is an ordered, compressed homogeneous run of units
// participating in one promotion stage. Version 1 is used only after explicit
// consumer negotiation. FirstUnit is one-based on the current requested line;
// BaseUnits may span only units sharing every stated fact, including actual
// awards and exact per-unit values. A compressed rounding proof's rank and
// tie key are for FirstUnit. Later units use consecutive ranks and the pool's
// supported original-ordinal tie rule for both settled and contribution
// rounding; otherwise each unit must be split into
// a separate run. Role is target, qualifier or both; zero-reduction qualifier
// participation remains explicit. PromotionID/revision and StageOrdinal bind
// the immutable promotion and actual application order. A group span requires
// all groups' peers, including peers on other lines, and must respect the
// stated package boundaries and repetition limits. Retained slices bind their
// clipped group fragments and original peers through authenticated lineage.
//
// EligibleUnitBasis is the original eligible book basis, not a previous
// promotion's reduced price. BeforeUnit, ReductionUnit and AfterUnit are exact
// book-basis rationals, with AfterUnit = BeforeUnit - ReductionUnit. Optional
// settled amounts record only actual prices assigned by the original scorer;
// SettledUnitRounding proves SettledAfterUnit. A stated SettledBeforeUnit must
// equal the authenticated previous stage's settled after value or an
// independently proved original source price. Rounding can make settled after
// differ from exact AfterUnit. For settled_difference_v1, SettledUnitContribution is the independently
// authenticated same-unit settled-before less settled-after, nonnegative and
// in the same currency and exponent. Complete peers, group and root ordinals
// must reconcile to the authoritative application and quote discount without
// double counting; qualifier-only roles contribute zero. Retained amendments
// authenticate predecessor proofs without rerounding. Unknown or missing modes
// fail closed when attribution is required. Legacy contribution rounding
// evidence remains independent and requires complete peers and ranks.
// SettledRoundingPool and ContributionRoundingPool bind distinct original
// populations and totals; neither substitutes for the other or for complete
// peer evidence. SettledAwardEligible explicitly states whether this run can
// receive a settled-price remainder award. Zero-valued and previously locked
// units are ineligible; SettledLockStage identifies the authenticated prior
// stage that locked such units: its positive ordinal must be less than this
// stage and refer to the authenticated same original source units. A false
// value requires proof of a zero exact value or that earlier lock; true requires
// a nonzero exact value and no effective earlier lock. Consumers must prove
// eligibility and all peer floors before ranking the eligible peers. A nil
// eligibility value cannot support a complete settlement proof. Pools
// never combine incompatible book or tax-inclusion bases.
// Unit rank and fixed-width root ordinal tie identity in an inherited
// amendment remain those of the authenticated original, not freshly
// computed ranks; clipping advances them by the retained native offset. Consumers must bind
// these runs, source intersections, original facts, peers, and lineage to the
// authoritative quote fingerprint, enforce bounded counts, and fail closed
// if any proof is unavailable. This is not a refund allocation policy.
type PricePromotionStageRun struct {
	PromotionID              string                            `json:"promotion_id"`
	PromotionRevision        int64                             `json:"promotion_revision"`
	StageOrdinal             int32                             `json:"stage_ordinal"`
	Role                     string                            `json:"role"`
	FirstUnit                int64                             `json:"first_unit"`
	BaseUnits                int64                             `json:"base_units"`
	PackageOption            packaging.PackageOptionRef        `json:"package_option"`
	UnitsPerPackage          int64                             `json:"units_per_package"`
	FirstPackageOrdinal      int64                             `json:"first_package_ordinal"`
	FirstWithinPackageUnit   int64                             `json:"first_within_package_unit"`
	Currency                 money.CurrencyCode                `json:"currency"`
	CurrencyExponent         money.CurrencyExponent            `json:"currency_exponent"`
	BookTaxInclusion         pricebook_enums.PriceTaxInclusion `json:"book_tax_inclusion"`
	EligibleUnitBasis        PromotionExactMinor               `json:"eligible_unit_basis"`
	BeforeUnit               PromotionExactMinor               `json:"before_unit"`
	ReductionUnit            PromotionExactMinor               `json:"reduction_unit"`
	AfterUnit                PromotionExactMinor               `json:"after_unit"`
	SettledBeforeUnit        *money.Money                      `json:"settled_before_unit,omitempty"`
	SettledAfterUnit         *money.Money                      `json:"settled_after_unit,omitempty"`
	GroupSpan                *PromotionGroupSpan               `json:"group_span,omitempty"`
	SettledRoundingPool      *PromotionStageRoundingPool       `json:"settled_rounding_pool,omitempty"`
	SettledUnitRounding      *RoundingEvidence                 `json:"settled_unit_rounding,omitempty"`
	SettledAwardEligible     *bool                             `json:"settled_award_eligible,omitempty"`
	SettledLockStage         *int32                            `json:"settled_lock_stage,omitempty"`
	ContributionProofMode    string                            `json:"contribution_proof_mode,omitempty"`
	SettledUnitContribution  *money.Money                      `json:"settled_unit_contribution,omitempty"`
	ContributionRoundingPool *PromotionStageRoundingPool       `json:"contribution_rounding_pool,omitempty"`
	UnitContributionRounding *RoundingEvidence                 `json:"unit_contribution_rounding,omitempty"`
	Lineage                  *PromotionStageLineage            `json:"lineage,omitempty"`
}
