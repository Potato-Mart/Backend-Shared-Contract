package security_enums

// DataClassification labels information according to its confidentiality
// and handling requirements.
type DataClassification string

const (
	DataClassificationPublic       DataClassification = "public"
	DataClassificationInternal     DataClassification = "internal"
	DataClassificationConfidential DataClassification = "confidential"
	DataClassificationRestricted   DataClassification = "restricted"
)

func (c DataClassification) IsValid() bool {
	switch c {
	case DataClassificationPublic, DataClassificationInternal,
		DataClassificationConfidential, DataClassificationRestricted:
		return true
	}
	return false
}

func (c DataClassification) String() string { return string(c) }
