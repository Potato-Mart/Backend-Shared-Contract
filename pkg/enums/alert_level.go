package enums

// AlertLevel classifies the urgency of a stock-out forecast.
type AlertLevel string

const (
	AlertLevelOK       AlertLevel = "OK"
	AlertLevelWarning  AlertLevel = "WARNING"
	AlertLevelCritical AlertLevel = "CRITICAL"
	AlertLevelExpired  AlertLevel = "EXPIRED"
)

// IsValid reports whether a is a known AlertLevel.
func (a AlertLevel) IsValid() bool {
	switch a {
	case AlertLevelOK, AlertLevelWarning, AlertLevelCritical, AlertLevelExpired:
		return true
	}
	return false
}

func (a AlertLevel) String() string { return string(a) }
