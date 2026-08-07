package localization

// LocalizedName records a name together with the language it is written in.
// Language should use a BCP 47 language tag such as "en", "zh-TW", or "ja".
type LocalizedName struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}
