package identityenum

type UserPreferredLanguage string

const (
	PreferredLanguageEnglish            UserPreferredLanguage = "english"
	PreferredLanguageTraditionalChinese UserPreferredLanguage = "traditional_chinese"
	PreferredLanguageSimplifiedChinese  UserPreferredLanguage = "simplified_chinese"
)

// IsValid reports whether s is a known UserPreferredLanguage.
func (l UserPreferredLanguage) IsValid() bool {
	switch l {
	case PreferredLanguageEnglish, PreferredLanguageTraditionalChinese, PreferredLanguageSimplifiedChinese:
		return true
	}
	return false
}

func (l UserPreferredLanguage) String() string { return string(l) }
