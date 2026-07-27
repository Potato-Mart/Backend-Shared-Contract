package productenum

// MediaStatus controls visibility and lifecycle of a media row.
//
//	pending  – SAS URL issued but the upload hasn't been finalized yet.
//	active   – Upload finalized; safe to reference.
//	deleted  – Soft-deleted; the orphan reaper will purge the blob
//	           once no documents still reference the row.
type MediaStatus string

const (
	MediaStatusPending MediaStatus = "pending"
	MediaStatusActive  MediaStatus = "active"
	MediaStatusDeleted MediaStatus = "deleted"
)

// IsValid reports whether s is a known MediaStatus.
func (s MediaStatus) IsValid() bool {
	switch s {
	case MediaStatusPending, MediaStatusActive, MediaStatusDeleted:
		return true
	}
	return false
}

func (s MediaStatus) String() string { return string(s) }

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
