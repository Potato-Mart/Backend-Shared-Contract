package import_compliance_enums

type ArtifactKind string

const (
	ArtifactKindPDF  ArtifactKind = "pdf"
	ArtifactKindDOCX ArtifactKind = "docx"
	ArtifactKindXLSX ArtifactKind = "xlsx"
	ArtifactKindZIP  ArtifactKind = "zip"
	ArtifactKindEML  ArtifactKind = "eml"
)

func (k ArtifactKind) IsValid() bool {
	switch k {
	case ArtifactKindPDF, ArtifactKindDOCX, ArtifactKindXLSX, ArtifactKindZIP, ArtifactKindEML:
		return true
	}
	return false
}

func (k ArtifactKind) String() string { return string(k) }
