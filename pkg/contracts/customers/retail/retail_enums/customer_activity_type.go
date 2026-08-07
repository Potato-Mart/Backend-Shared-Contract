package retail_enums

// CustomerIdentityKind is the channel or identifier type stored in
// customer_identities for cross-channel customer matching.
type CustomerIdentityKind string

const (
	CustomerIdentityKindPhone      CustomerIdentityKind = "PHONE"
	CustomerIdentityKindEmail      CustomerIdentityKind = "EMAIL"
	CustomerIdentityKindLine       CustomerIdentityKind = "LINE"
	CustomerIdentityKindMemberCard CustomerIdentityKind = "MEMBER_CARD"
	CustomerIdentityKindPOSID      CustomerIdentityKind = "POS_ID"
	CustomerIdentityKindExternal   CustomerIdentityKind = "EXTERNAL"
)

func (c CustomerIdentityKind) IsValid() bool {
	switch c {
	case CustomerIdentityKindPhone, CustomerIdentityKindEmail, CustomerIdentityKindLine,
		CustomerIdentityKindMemberCard, CustomerIdentityKindPOSID, CustomerIdentityKindExternal:
		return true
	}
	return false
}

func (c CustomerIdentityKind) String() string { return string(c) }
