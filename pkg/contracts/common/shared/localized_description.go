package shared

// LocalizedDescription records a description together with the language it is written in.
// Language should use a BCP 47 language tag such as "en", "zh-TW", or "ja".
type LocalizedDescription struct {
	Language    string `json:"language"`
	Description string `json:"description"`
}
