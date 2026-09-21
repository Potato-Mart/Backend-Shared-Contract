package enums_test

import (
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v33/pkg/contracts/notification/template/template_enums"
)

func TestNotificationTemplateEnumWireValues(t *testing.T) {
	cases := []struct {
		name  string
		value stringEnum
		want  string
	}{
		{"TemplateContentOriginAuthored", template_enums.TemplateContentOriginAuthored, "authored"},
		{"TemplateContentOriginTranslated", template_enums.TemplateContentOriginTranslated, "translated"},
		{"TemplatePlaceholderTypeText", template_enums.TemplatePlaceholderTypeText, "text"},
		{"TemplatePlaceholderTypeInteger", template_enums.TemplatePlaceholderTypeInteger, "integer"},
		{"TemplatePlaceholderTypeMoney", template_enums.TemplatePlaceholderTypeMoney, "money"},
		{"TemplatePlaceholderTypeTimestamp", template_enums.TemplatePlaceholderTypeTimestamp, "timestamp"},
		{"TemplatePlaceholderTypeURL", template_enums.TemplatePlaceholderTypeURL, "url"},
		{"EmailTemplateBlockKindHeading", template_enums.EmailTemplateBlockKindHeading, "heading"},
		{"EmailTemplateBlockKindParagraph", template_enums.EmailTemplateBlockKindParagraph, "paragraph"},
		{"EmailTemplateBlockKindButton", template_enums.EmailTemplateBlockKindButton, "button"},
		{"EmailTemplateBlockKindImage", template_enums.EmailTemplateBlockKindImage, "image"},
		{"EmailTemplateBlockKindDivider", template_enums.EmailTemplateBlockKindDivider, "divider"},
		{"EmailTemplateBlockKindSystemSection", template_enums.EmailTemplateBlockKindSystemSection, "system_section"},
		{"EmailFontFamilySans", template_enums.EmailFontFamilySans, "sans"},
		{"EmailFontFamilySerif", template_enums.EmailFontFamilySerif, "serif"},
		{"EmailTextSizeSmall", template_enums.EmailTextSizeSmall, "small"},
		{"EmailTextSizeNormal", template_enums.EmailTextSizeNormal, "normal"},
		{"EmailTextSizeLarge", template_enums.EmailTextSizeLarge, "large"},
		{"EmailSpacingCompact", template_enums.EmailSpacingCompact, "compact"},
		{"EmailSpacingNormal", template_enums.EmailSpacingNormal, "normal"},
		{"EmailSpacingRelaxed", template_enums.EmailSpacingRelaxed, "relaxed"},
		{"EmailButtonShapeSquare", template_enums.EmailButtonShapeSquare, "square"},
		{"EmailButtonShapeRounded", template_enums.EmailButtonShapeRounded, "rounded"},
		{"EmailAlignmentStart", template_enums.EmailAlignmentStart, "start"},
		{"EmailAlignmentCenter", template_enums.EmailAlignmentCenter, "center"},
		{"EmailAlignmentEnd", template_enums.EmailAlignmentEnd, "end"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.value.IsValid() || tc.value.String() != tc.want {
				t.Fatalf("enum valid=%v wire=%q, want true and %q", tc.value.IsValid(), tc.value.String(), tc.want)
			}
		})
	}
	for _, raw := range []string{"", "__invalid__", "EN", " html ", "script"} {
		invalid := []stringEnum{
			template_enums.TemplateContentOrigin(raw),
			template_enums.TemplatePlaceholderType(raw),
			template_enums.EmailTemplateBlockKind(raw),
			template_enums.EmailFontFamily(raw),
			template_enums.EmailTextSize(raw),
			template_enums.EmailSpacing(raw),
			template_enums.EmailButtonShape(raw),
			template_enums.EmailAlignment(raw),
		}
		for _, value := range invalid {
			if value.IsValid() || value.String() != raw {
				t.Errorf("%T(%q) must preserve but reject an unsupported value", value, raw)
			}
		}
	}
}
