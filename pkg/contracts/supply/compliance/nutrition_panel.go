package compliance

// NutritionPanel intentionally preserves author-entered display text and
// units. It is regulatory label copy, not a value used in shared arithmetic.
type NutritionPanel struct {
	Title                   string `json:"title,omitempty"`
	ServingsPerPack         string `json:"servings_per_pack,omitempty"`
	ServingSize             string `json:"serving_size,omitempty"`
	EnergyPerServe          string `json:"energy_per_serve,omitempty"`
	EnergyPer100Grams       string `json:"energy_per_100_grams,omitempty"`
	ProteinPerServe         string `json:"protein_per_serve,omitempty"`
	ProteinPer100Grams      string `json:"protein_per_100_grams,omitempty"`
	FatPerServe             string `json:"fat_per_serve,omitempty"`
	FatPer100Grams          string `json:"fat_per_100_grams,omitempty"`
	SaturatedFatPerServe    string `json:"saturated_fat_per_serve,omitempty"`
	SaturatedFatPer100Grams string `json:"saturated_fat_per_100_grams,omitempty"`
	CarbohydratePerServe    string `json:"carbohydrate_per_serve,omitempty"`
	CarbohydratePer100Grams string `json:"carbohydrate_per_100_grams,omitempty"`
	SugarsPerServe          string `json:"sugars_per_serve,omitempty"`
	SugarsPer100Grams       string `json:"sugars_per_100_grams,omitempty"`
	SodiumPerServe          string `json:"sodium_per_serve,omitempty"`
	SodiumPer100Grams       string `json:"sodium_per_100_grams,omitempty"`
}
