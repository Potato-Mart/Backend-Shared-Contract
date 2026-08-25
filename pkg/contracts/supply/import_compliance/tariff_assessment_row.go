package import_compliance

type TariffAssessmentRow struct {
	ID                       string               `json:"id"`
	Product                  TariffLineSnapshot   `json:"product"`
	Taiwan                   TariffClassification `json:"taiwan"`
	Australia                TariffClassification `json:"australia"`
	Trademark                string               `json:"trademark,omitempty"`
	TrademarkEvidenceIDs     []string             `json:"trademark_evidence_ids,omitempty"`
	Source                   string               `json:"source"`
	Notes                    string               `json:"notes,omitempty"`
	PublishedProfileRevision *int64               `json:"published_profile_revision,omitempty"`
}
