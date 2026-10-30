package quote

import (
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/packaging"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/pricebook/pricebook_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/pricing/quote/quote_enums"
)

// PromotionManifestEvidence is quote-level, not a per-line allocation. An
// adopting service must include this complete object in its authoritative
// immutable decision and expose it only after independent capability
// negotiation. Omitted evidence retains historical wire and hash semantics.
type PromotionManifestEvidence struct {
	Version                   int32                               `json:"version"`
	Manifests                 []PromotionCandidateManifest        `json:"manifests"`
	SettlementPools           []PromotionSettlementPool           `json:"settlement_pools"`
	ApplicationRoundingProofs []PromotionApplicationRoundingProof `json:"application_rounding_proofs"`
	GroupedNetEvents          []PromotionGroupedNetEvent          `json:"grouped_net_events,omitempty"`
	StageAudits               []PromotionStageAudit               `json:"stage_audits"`
}

// PromotionCandidateManifest freezes the complete original candidate role
// vector for a single actual stage. It is carried in separately negotiated
// quote-level evidence and bound to the immutable Pricing decision; it is not
// inferred from the runs a consumer happened to receive. Version 1 requires
// contiguous one-based requirement indices, complete group coverage for each
// declared peer and bounded counts. Each stage run names ManifestID and its
// requirement index. A candidate with different peers per repetition must be
// represented by distinct manifests, never silently regrouped.
// MinimumRepetitions zero means no extra minimum beyond one selected group;
// MaximumRepetitions is a strict positive upper bound. Requirement UnitsPerGroup
// is the sum of that requirement's peer multiplicities, not a per-line guess.
// Version 1 permits competition_class "product_competing" for grouped BOGO
// and bundle candidates only. Other mechanics or unknown classes require a
// separately specified proof and fail closed.
//
// ManifestID is an opaque Pricing-owned identifier, authenticated against the
// complete immutable root manifest through the owner. Consumers must not
// recompute it from this wire struct or assume it is a cross-service canonical
// digest. Its owner preimage excludes the finalized quote revision/fingerprint:
// the quote decision binds the manifest, so including either creates a cycle.
// EligibilityEvidenceID identifies a separately frozen Pricing-owned decision
// bound to the root quote, candidate and original normalized request, including
// buyer, channel, platform, geography, line/package quantities, POS selection,
// evaluation time and the immutable definition revisions. The owner must
// authenticate the approved original source, tax and exact unit net for every
// original cart component used by the decision, including nonparticipating
// lines used by a cart-wide threshold. An excluded exclusive-path line must
// instead have its exclusion independently bound to the original request,
// actual pre-promotion exclusive rule and immutable root snapshot; it cannot
// participate in the candidate. Missing input or exclusion proof fails closed.
// The owner must authenticate this record and the decision; an unrelated quote digest or
// request fingerprint alone is not that proof. UsageEvidenceID identifies
// separately frozen authoritative usage evidence. Both IDs are computed before
// and independently of the containing quote revision/fingerprint; the final
// quote binds them. Atomic reservation remains a
// separate final race guard. A manifest alone does not prove audience,
// channel, geography, time, promotion availability or a usage reservation.
// Missing or unverifiable facts fail closed. This evidence does not alter the
// existing line-level proportional refund policy.
type PromotionCandidateManifest struct {
	Version               int32                           `json:"version"`
	ManifestID            string                          `json:"manifest_id"`
	PromotionID           string                          `json:"promotion_id"`
	PromotionKind         string                          `json:"promotion_kind"`
	CompetitionClass      string                          `json:"competition_class"`
	PromotionRevision     int64                           `json:"promotion_revision"`
	ScopeRevision         int64                           `json:"scope_revision"`
	RuleRevision          int64                           `json:"rule_revision"`
	PublicationRevision   int64                           `json:"publication_revision"`
	CandidateID           string                          `json:"candidate_id"`
	StageOrdinal          int32                           `json:"stage_ordinal"`
	FirstGroup            int64                           `json:"first_group"`
	GroupCount            int64                           `json:"group_count"`
	MinimumRepetitions    int64                           `json:"minimum_repetitions"`
	MaximumRepetitions    int64                           `json:"maximum_repetitions"`
	RequirementCount      int32                           `json:"requirement_count"`
	Requirements          []PromotionCandidateRequirement `json:"requirements"`
	EligibilityEvidenceID string                          `json:"eligibility_evidence_id"`
	UsageEvidenceID       string                          `json:"usage_evidence_id"`
}

// PromotionCandidateRequirement preserves an authored role-vector index even
// when two requirements use the same line and role. WholePackage states the
// actual requirement, not a guess from the chosen unit count. Every peer is
// necessary for every indicated original group; zero-reduction qualifiers are
// explicit and physically disjoint from other competing participants.
type PromotionCandidateRequirement struct {
	Index         int32                    `json:"index"`
	Role          string                   `json:"role"`
	WholePackage  bool                     `json:"whole_package"`
	UnitsPerGroup int64                    `json:"units_per_group"`
	Peers         []PromotionCandidatePeer `json:"peers"`
}

// PromotionCandidatePeer identifies original eligible facts and every selected
// ordinal interval for this (requirement, root line, role). Spans partition the
// full manifest group range without gaps or overlapping units. The exact unit
// basis is the original approved book basis, before any promotion. Source and
// tax identities are authenticated against frozen root evidence, not inferred
// from an adjusted scalar line. The stage-zero proof additionally requires an
// independently authenticated initial source and absence of an earlier
// monetary stage on these units. SourceEvidenceID is a Pricing-owned opaque
// identity for the complete original source, authenticated by the owner against
// the immutable root decision. It excludes the finalized quote revision and
// must match the corresponding source-before evidence when that is present.
// TaxRate describes the recorded original tax calculation; the separate
// ExclusiveRate operands are the actual frozen rule values used by Pricing's
// inclusive from-net conversion. Consumers must not infer one rate from the
// other. Denominators and arithmetic are verified against the owner source.
type PromotionCandidatePeer struct {
	RootLineID                string                            `json:"root_line_id"`
	SourceEvidenceID          string                            `json:"source_evidence_id"`
	SKUCode                   string                            `json:"sku_code"`
	PackageOption             packaging.PackageOptionRef        `json:"package_option"`
	UnitsPerPackage           int64                             `json:"units_per_package"`
	UnitsPerGroup             int64                             `json:"units_per_group"`
	OriginalLineBaseUnits     int64                             `json:"original_line_base_units"`
	OriginalLinePackageCount  int64                             `json:"original_line_package_count"`
	ExactOriginalLineBasis    PromotionExactMinor               `json:"exact_original_line_basis"`
	RoundedOriginalLineBasis  money.Money                       `json:"rounded_original_line_basis"`
	SourceBefore              *PromotionSourceBeforeEvidence    `json:"source_before,omitempty"`
	QualifierTransitions      []PromotionQualifierIdentityProof `json:"qualifier_transitions,omitempty"`
	Currency                  money.CurrencyCode                `json:"currency"`
	CurrencyExponent          money.CurrencyExponent            `json:"currency_exponent"`
	BookTaxInclusion          pricebook_enums.PriceTaxInclusion `json:"book_tax_inclusion"`
	TaxCategoryCode           string                            `json:"tax_category_code"`
	TaxRuleID                 string                            `json:"tax_rule_id"`
	TaxRuleRevision           int64                             `json:"tax_rule_revision"`
	TaxCalculationSource      quote_enums.TaxCalculationSource  `json:"tax_calculation_source"`
	TaxRateNumerator          int64                             `json:"tax_rate_numerator"`
	TaxRateDenominator        int64                             `json:"tax_rate_denominator"`
	ExclusiveRateNumerator    int64                             `json:"exclusive_rate_numerator"`
	ExclusiveRateDenominator  int64                             `json:"exclusive_rate_denominator"`
	EligibleUnitBasis         PromotionExactMinor               `json:"eligible_unit_basis"`
	PriceBookCode             string                            `json:"price_book_code"`
	PriceBookRevision         int64                             `json:"price_book_revision"`
	PriceEntryID              string                            `json:"price_entry_id"`
	PriceEntryRevision        int64                             `json:"price_entry_revision"`
	PackagePriceEntryID       string                            `json:"package_price_entry_id,omitempty"`
	PackagePriceEntryRevision int64                             `json:"package_price_entry_revision,omitempty"`
	Spans                     []PromotionCandidatePeerSpan      `json:"spans"`
}

// PromotionQualifierIdentityProof records an actual no-change qualifier
// transition where Pricing performed no settlement on the qualifier line.
// It must not be used as a fabricated rounding pool. The owner authenticates
// the full original source, complete earlier-stage observation, actual qualifier
// event, candidate requirement and root ordinal/group spans against the root
// quote. The consumer verifies whole-package boundaries, full peer coverage,
// disjoint competing units, equal original/before/after exact unit values and
// absence of prior monetary stages or locks. Retained slices additionally
// require authenticated predecessor lineage. Usage and eligibility evidence
// remain independent. In the supported `no_earlier_events_v1` mode the
// Pricing owner authenticates its captured observation that no other selected
// candidate or prior/same-stage mechanic event used the line, and the root
// snapshot has no applied rules. This is not a separate audit record. Unknown
// modes fail closed. Identity is opaque and computed before quote revision.
type PromotionQualifierIdentityProof struct {
	ID                string              `json:"id"`
	SourceEvidenceID  string              `json:"source_evidence_id"`
	PriorStageMode    string              `json:"prior_stage_mode"`
	ManifestID        string              `json:"manifest_id"`
	PromotionID       string              `json:"promotion_id"`
	PromotionRevision int64               `json:"promotion_revision"`
	StageOrdinal      int32               `json:"stage_ordinal"`
	SubstageOrdinal   int32               `json:"substage_ordinal"`
	RequirementIndex  int32               `json:"requirement_index"`
	FirstGroup        int64               `json:"first_group"`
	GroupCount        int64               `json:"group_count"`
	RootLineID        string              `json:"root_line_id"`
	FirstUnit         int64               `json:"first_unit"`
	BaseUnits         int64               `json:"base_units"`
	BeforeUnit        PromotionExactMinor `json:"before_unit"`
	AfterUnit         PromotionExactMinor `json:"after_unit"`
}

// PromotionCandidatePeerSpan binds a complete root group interval to its
// original contiguous one-based line ordinals. A span represents exactly
// GroupCount*UnitsPerGroup units, ordered by group then within-role ordinal.
// Retained clipped slices reference this root span and the authenticated
// predecessor chain; they do not claim omitted peers are still present.
type PromotionCandidatePeerSpan struct {
	FirstGroup int64 `json:"first_group"`
	GroupCount int64 `json:"group_count"`
	FirstUnit  int64 `json:"first_unit"`
}
