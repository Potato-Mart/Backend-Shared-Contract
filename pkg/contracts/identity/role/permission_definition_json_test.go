package role

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/common/security/security_enums"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/identity/access"
	"github.com/Potato-Mart/Backend-Shared-Contract/v32/pkg/contracts/identity/role/role_enums"
)

func TestPermissionDefinitionJSONUsesTypedWireValues(t *testing.T) {
	definition := PermissionDefinition{
		Key:            PermissionKey("role.write"),
		Label:          "Manage roles",
		Module:         "roles",
		RiskLevel:      security_enums.SecurityRiskLevelCritical,
		RequiresMFA:    true,
		Classification: role_enums.PermissionClassificationUI,
	}

	payload, err := json.Marshal(definition)
	if err != nil {
		t.Fatalf("marshal PermissionDefinition: %v", err)
	}
	const want = `{"key":"role.write","label":"Manage roles","module":"roles","risk_level":"critical","requires_mfa":true,"classification":"ui"}`
	if got := string(payload); got != want {
		t.Fatalf("PermissionDefinition JSON = %s, want %s", got, want)
	}

	var decoded PermissionDefinition
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal PermissionDefinition: %v", err)
	}
	if decoded.Key != PermissionKey("role.write") || decoded.RiskLevel != security_enums.SecurityRiskLevelCritical || decoded.Classification != role_enums.PermissionClassificationUI {
		t.Fatalf("typed PermissionDefinition values = %#v", decoded)
	}
}

func TestRolePermissionsMarshalAsTypedStringArray(t *testing.T) {
	value := Role{
		Key:         role_enums.UserRoleSuperAdmin,
		Label:       "Platform administrators",
		Permissions: []PermissionKey{PermissionKey("user.read"), PermissionKey("role.write")},
		IsSystem:    true,
	}

	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal Role: %v", err)
	}
	const want = `{"key":"superAdmin","label":"Platform administrators","permissions":["user.read","role.write"],"is_system":true,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`
	if got := string(payload); got != want {
		t.Fatalf("Role JSON = %s, want %s", got, want)
	}

	var decoded Role
	if err := json.Unmarshal([]byte(`{"permissions":["user.read","role.write"]}`), &decoded); err != nil {
		t.Fatalf("unmarshal Role permissions: %v", err)
	}
	if len(decoded.Permissions) != 2 || decoded.Permissions[0] != PermissionKey("user.read") || decoded.Permissions[1] != PermissionKey("role.write") {
		t.Fatalf("decoded typed permissions = %#v", decoded.Permissions)
	}
}

func TestLoginSessionPermissionsRemainMixedStringSlice(t *testing.T) {
	permissionsField, ok := reflect.TypeFor[access.LoginSession]().FieldByName("Permissions")
	if !ok {
		t.Fatal("LoginSession.Permissions is missing")
	}
	if got, want := permissionsField.Type.String(), "[]string"; got != want {
		t.Fatalf("LoginSession.Permissions type = %s, want %s", got, want)
	}
}
