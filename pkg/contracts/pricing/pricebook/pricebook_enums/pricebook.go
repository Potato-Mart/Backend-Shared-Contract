package pricebook_enums

// PriceBookStatus is the admin-controlled lifecycle state of a price book.
type PriceBookStatus string

const (
	// PriceBookStatusDraft is a book being prepared; its entries may be
	// stored but never resolve into a quote.
	PriceBookStatusDraft PriceBookStatus = "draft"
	// PriceBookStatusActive is a book that resolves for its market,
	// channel, and audience inside its validity window.
	PriceBookStatusActive PriceBookStatus = "active"
	// PriceBookStatusInactive is a book withheld from resolution while its
	// entries are retained.
	PriceBookStatusInactive PriceBookStatus = "inactive"
	// PriceBookStatusArchived is a book permanently withdrawn and kept for
	// history only.
	PriceBookStatusArchived PriceBookStatus = "archived"
)

// IsValid reports whether s is a known PriceBookStatus.
func (s PriceBookStatus) IsValid() bool {
	switch s {
	case PriceBookStatusDraft, PriceBookStatusActive, PriceBookStatusInactive, PriceBookStatusArchived:
		return true
	}
	return false
}

func (s PriceBookStatus) String() string { return string(s) }

// PriceEntryStatus is the lifecycle state of one price entry. Only an approved
// entry inside its validity window is resolvable; there is no silent fallback
// to a draft, to acquisition cost, to a converted currency, or to another
// market when no approved entry exists.
type PriceEntryStatus string

const (
	// PriceEntryStatusDraft is a stored suggestion or work in progress.
	PriceEntryStatusDraft PriceEntryStatus = "draft"
	// PriceEntryStatusPendingApproval is a draft submitted for admin
	// approval. It is still not resolvable.
	PriceEntryStatusPendingApproval PriceEntryStatus = "pending_approval"
	// PriceEntryStatusApproved is an admin-approved entry. Inside its
	// validity window it is the live sale price.
	PriceEntryStatusApproved PriceEntryStatus = "approved"
	// PriceEntryStatusRejected is a draft an admin declined.
	PriceEntryStatusRejected PriceEntryStatus = "rejected"
	// PriceEntryStatusSuperseded is a previously approved entry replaced by
	// a newer approved entry.
	PriceEntryStatusSuperseded PriceEntryStatus = "superseded"
	// PriceEntryStatusWithdrawn is an entry deliberately taken out of sale.
	PriceEntryStatusWithdrawn PriceEntryStatus = "withdrawn"
	// PriceEntryStatusExpired is an entry whose validity window has closed.
	PriceEntryStatusExpired PriceEntryStatus = "expired"
)

// IsValid reports whether s is a known PriceEntryStatus.
func (s PriceEntryStatus) IsValid() bool {
	switch s {
	case PriceEntryStatusDraft, PriceEntryStatusPendingApproval, PriceEntryStatusApproved,
		PriceEntryStatusRejected, PriceEntryStatusSuperseded, PriceEntryStatusWithdrawn,
		PriceEntryStatusExpired:
		return true
	}
	return false
}

func (s PriceEntryStatus) String() string { return string(s) }

// PriceTaxInclusion records whether the amounts in a price book already
// include the market's consumption tax. All Australian customer sale books are
// tax inclusive; supplier and cost values are not.
type PriceTaxInclusion string

const (
	PriceTaxInclusionInclusive PriceTaxInclusion = "tax_inclusive"
	PriceTaxInclusionExclusive PriceTaxInclusion = "tax_exclusive"
)

// IsValid reports whether i is a known PriceTaxInclusion.
func (i PriceTaxInclusion) IsValid() bool {
	switch i {
	case PriceTaxInclusionInclusive, PriceTaxInclusionExclusive:
		return true
	}
	return false
}

func (i PriceTaxInclusion) String() string { return string(i) }

// PriceEndingPolicy names the rounding presentation a book applies when a
// draft price is generated. It applies only to draft generation, never to
// group, expiry, damage, tax, or cash-settlement arithmetic.
type PriceEndingPolicy string

const (
	// PriceEndingPolicyNone keeps ordinary currency-precision rounding.
	PriceEndingPolicyNone PriceEndingPolicy = "none"
	// PriceEndingPolicyCharmNine keeps an exact ten-cent value as it is and
	// otherwise rounds upward to one cent below the next ten-cent boundary,
	// producing an x.x9 ending.
	PriceEndingPolicyCharmNine PriceEndingPolicy = "charm_nine"
)

// IsValid reports whether p is a known PriceEndingPolicy.
func (p PriceEndingPolicy) IsValid() bool {
	switch p {
	case PriceEndingPolicyNone, PriceEndingPolicyCharmNine:
		return true
	}
	return false
}

func (p PriceEndingPolicy) String() string { return string(p) }

// PriceBookAssignmentKind names how a book is bound to buyers in a market.
// Resolution runs exact organisation override, then organisation category,
// then channel default; a missing or inactive level falls through.
type PriceBookAssignmentKind string

const (
	// PriceBookAssignmentKindChannelDefault binds a book to every buyer on
	// one channel of one market.
	PriceBookAssignmentKindChannelDefault PriceBookAssignmentKind = "channel_default"
	// PriceBookAssignmentKindOrganisationCategory binds a book to one
	// wholesale organisation category.
	PriceBookAssignmentKindOrganisationCategory PriceBookAssignmentKind = "organisation_category"
	// PriceBookAssignmentKindOrganisationOverride binds a book to one exact
	// organisation.
	PriceBookAssignmentKindOrganisationOverride PriceBookAssignmentKind = "organisation_override"
)

// IsValid reports whether k is a known PriceBookAssignmentKind.
func (k PriceBookAssignmentKind) IsValid() bool {
	switch k {
	case PriceBookAssignmentKindChannelDefault, PriceBookAssignmentKindOrganisationCategory,
		PriceBookAssignmentKindOrganisationOverride:
		return true
	}
	return false
}

func (k PriceBookAssignmentKind) String() string { return string(k) }

// PriceDerivation records where a price entry's amount came from, so an
// approved price is always distinguishable from a generated suggestion.
type PriceDerivation string

const (
	// PriceDerivationManual is an amount an administrator entered.
	PriceDerivationManual PriceDerivation = "manual"
	// PriceDerivationSuggestedFromBaseCost is an amount a suggestion policy
	// generated from the SKU's base acquisition cost. The book's channel
	// selects which markup policy produced it.
	PriceDerivationSuggestedFromBaseCost PriceDerivation = "suggested_from_base_cost"
	// PriceDerivationImported is an amount loaded from an external price
	// list.
	PriceDerivationImported PriceDerivation = "imported"
)

// IsValid reports whether d is a known PriceDerivation.
func (d PriceDerivation) IsValid() bool {
	switch d {
	case PriceDerivationManual, PriceDerivationSuggestedFromBaseCost, PriceDerivationImported:
		return true
	}
	return false
}

func (d PriceDerivation) String() string { return string(d) }
