package template

import "github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notification/template/template_enums"

// TemplatePlaceholder declares one required scalar substitution without a value.
// Schema 1 tokens are {{key}} with keys matching
// [a-z][a-z0-9_]*(\.[a-z][a-z0-9_]*)*. There is no evaluation or control flow.
// Money resolves to the common minor-unit/currency value; timestamps are UTC.
// URL values belong only in nontranslatable URL slots. Sensitive classifies
// runtime material; it never authorizes exposure to a translator or preview.
// Notification enforces types, contexts and token preservation before rendering.
type TemplatePlaceholder struct {
	Key       string                                 `json:"key"`
	ValueType template_enums.TemplatePlaceholderType `json:"value_type"`
	Sensitive bool                                   `json:"sensitive"`
}
