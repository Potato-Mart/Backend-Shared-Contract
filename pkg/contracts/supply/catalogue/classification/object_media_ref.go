package classification

// ObjectMediaRef is the media relationship persisted by catalog records: the
// managed media code plus optional copy for this placement. Render URLs are
// resolved from the managed media asset.
//
// AltText and Caption are staff-authored copy for this use of the image, set
// for product cover, gallery and detail images. AltText describes the image
// for screen readers and when it fails to load; Caption is visible text shown
// with the image. Both are optional and omitted when empty; records whose
// owners do not author image copy leave them empty.
type ObjectMediaRef struct {
	Code    string `json:"code"`
	AltText string `json:"alt_text,omitempty"`
	Caption string `json:"caption,omitempty"`
}
