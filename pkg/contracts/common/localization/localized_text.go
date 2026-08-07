package localization

// LocalizedText records generic text together with the language it is written
// in. Language should use a BCP 47 language tag such as "en", "zh-TW", or
// "ja".
type LocalizedText struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}
