package quote

import "github.com/Potato-Mart/Backend-Shared-Contract/v53/pkg/contracts/common/money"

// MixedAdminLineWitness binds one mixed-source line's tax-group accounting to
// MixedAdminDiscountWitness. GlobalWitnessHash is lowercase SHA-256 of the
// complete global witness encoded as structural canonical JSON: object keys
// sorted, exact JSON integers preserved, array order preserved, and no quote
// revision embedded in the witness. The transaction service validates the
// global and group proofs against authenticated predecessor and quote evidence.
// This records accounting reductions, not physical unit tax reclassification.
type MixedAdminLineWitness struct {
	Version           int32                     `json:"version"`
	Operation         string                    `json:"operation"`
	LineID            string                    `json:"line_id"`
	GlobalWitnessHash string                    `json:"global_witness_hash"`
	PreAdminGross     money.Money               `json:"pre_admin_gross"`
	PreAdminTax       money.Money               `json:"pre_admin_tax"`
	CarriedGross      money.Money               `json:"carried_gross"`
	CarriedTax        money.Money               `json:"carried_tax"`
	AppliedGross      money.Money               `json:"applied_gross"`
	AppliedTax        money.Money               `json:"applied_tax"`
	Buckets           []MixedAdminBucketWitness `json:"buckets"`
}
