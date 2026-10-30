package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
)

// PromotionExactMinor is an exact rational amount in minor units of the
// enclosing run's currency. Denominator is positive and the fraction is reduced.
// It is not itself a rounded payable amount.
type PromotionExactMinor struct {
	Numerator   int64 `json:"numerator"`
	Denominator int64 `json:"denominator"`
}

// PromotionGroupSpan identifies complete repeated candidate groups. CandidateID
// is stable within the immutable Pricing decision and shared by all qualifier
// and target peers. Group ordinals and FirstGroupUnitOffset are one-based;
// GroupCount is the number of root groups touched and UnitsPerGroup is this
// role's number of units in each group. A fresh allocation contains complete
// groups. A predecessor-bound retained slice may start or end within a group;
// its omitted peers are authenticated from the immutable predecessor, not
// asserted as retained. For zero-based run offset j, the original group is
// FirstGroup + (FirstGroupUnitOffset - 1 + j) / UnitsPerGroup; the within-role
// position is 1 + (FirstGroupUnitOffset - 1 + j) % UnitsPerGroup. GroupCount
// equals the number of groups touched by BaseUnits. For a complete fresh run,
// the offset is 1 and BaseUnits = GroupCount * UnitsPerGroup. Consumers
// authenticate the entire original peer set.
type PromotionGroupSpan struct {
	CandidateID          string `json:"candidate_id"`
	FirstGroup           int64  `json:"first_group"`
	GroupCount           int64  `json:"group_count"`
	UnitsPerGroup        int64  `json:"units_per_group"`
	FirstGroupUnitOffset int64  `json:"first_group_unit_offset"`
}

// PromotionStageRoundingPool binds the common exact total, rounded total,
// rounding mode and exponent of the complete original-line peer population.
// Every member repeats the same pool proof, ID and policy; consumers must
// authenticate all peers and recompute their ranks, awards and stable tie keys.
// Version 1 permits only TieBreakRule "line_unit_ordinal_v1": a pool covers
// one original line and ranks exact fractional remainders descending, then
// eligible original unit ordinals ascending. Ineligible peers still contribute
// their floors to the pool total, but receive no award, rank or tie key.
// The tie key for eligible peers is the root unit ordinal
// rendered as a zero-padded 20-digit decimal string. RootLineID binds the line. A
// different peer order cannot be labelled with this rule and must be rejected
// until another negotiated version defines it. Pools of settled book prices
// and of promotion contributions are distinct and cannot be substituted.
type PromotionStageRoundingPool struct {
	ID           string           `json:"id"`
	Total        RoundingEvidence `json:"total"`
	TieBreakRule string           `json:"tie_break_rule"`
	RootLineID   string           `json:"root_line_id"`
}

// PromotionStageLineage binds a retained run to its original immutable quote
// and ordinal interval. The current run's FirstUnit is rebased; RootFirstUnit
// identifies the same first native unit in the root line. Root quote identity
// and the complete predecessor chain must be independently authenticated.
// Lineage with a nonempty root quote, fingerprint and line is required for
// every clipped or rebased retained run.
type PromotionStageLineage struct {
	RootQuoteID          string `json:"root_quote_id"`
	RootQuoteFingerprint string `json:"root_quote_fingerprint"`
	RootLineID           string `json:"root_line_id"`
	RootFirstUnit        int64  `json:"root_first_unit"`
}

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
// differ from exact AfterUnit. Contribution
// rounding is optional only when no unit contribution is asserted. Consumers
// that need per-promotion unit attribution must require it together with the
// complete peer set, stable rank, tie key and award, and reject unproven cases.
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
	ContributionRoundingPool *PromotionStageRoundingPool       `json:"contribution_rounding_pool,omitempty"`
	UnitContributionRounding *RoundingEvidence                 `json:"unit_contribution_rounding,omitempty"`
	Lineage                  *PromotionStageLineage            `json:"lineage,omitempty"`
}
