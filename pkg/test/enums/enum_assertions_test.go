package enums_test

import (
	"reflect"
	"testing"
)

type stringEnum interface {
	IsValid() bool
	String() string
}

type enumCase struct {
	name    string
	valid   []stringEnum
	invalid stringEnum
}

func assertStringEnums(t *testing.T, cases []enumCase) {
	t.Helper()

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			for _, value := range tt.valid {
				if !value.IsValid() {
					t.Fatalf("%T(%q) should be valid", value, reflect.ValueOf(value).String())
				}
				if got, want := value.String(), reflect.ValueOf(value).String(); got != want {
					t.Fatalf("%T.String() = %q, want %q", value, got, want)
				}
			}

			if tt.invalid.IsValid() {
				t.Fatalf("%T(%q) should be invalid", tt.invalid, reflect.ValueOf(tt.invalid).String())
			}
			if got, want := tt.invalid.String(), reflect.ValueOf(tt.invalid).String(); got != want {
				t.Fatalf("%T.String() = %q for invalid value, want %q", tt.invalid, got, want)
			}
		})
	}
}
