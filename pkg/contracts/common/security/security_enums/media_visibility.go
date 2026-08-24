package security_enums

// MediaVisibility controls whether an object has a stable public URL or must
// be accessed through an authenticated signed URL.
type MediaVisibility string

const (
	MediaVisibilityPublic  MediaVisibility = "public"
	MediaVisibilityPrivate MediaVisibility = "private"
)

func (v MediaVisibility) IsValid() bool {
	switch v {
	case MediaVisibilityPublic, MediaVisibilityPrivate:
		return true
	default:
		return false
	}
}

func (v MediaVisibility) String() string { return string(v) }
