package role_enums

// PermissionClassification identifies where a workforce permission is
// enforced. It is catalogue metadata, not a separate authorization policy.
type PermissionClassification string

const (
	PermissionClassificationUI       PermissionClassification = "ui"
	PermissionClassificationField    PermissionClassification = "field-level"
	PermissionClassificationService  PermissionClassification = "service-only"
	PermissionClassificationReserved PermissionClassification = "intentionally-reserved"
)

func (c PermissionClassification) IsValid() bool {
	switch c {
	case PermissionClassificationUI,
		PermissionClassificationField,
		PermissionClassificationService,
		PermissionClassificationReserved:
		return true
	default:
		return false
	}
}

func (c PermissionClassification) String() string { return string(c) }
