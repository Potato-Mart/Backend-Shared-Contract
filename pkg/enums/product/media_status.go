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
