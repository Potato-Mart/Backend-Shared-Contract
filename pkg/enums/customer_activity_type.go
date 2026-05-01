package enums

// CustomerActivityType classifies an entry in the customer activity timeline.
type CustomerActivityType string

const (
	CustomerActivityTypeNote         CustomerActivityType = "NOTE"
	CustomerActivityTypeCall         CustomerActivityType = "CALL"
	CustomerActivityTypeEmail        CustomerActivityType = "EMAIL"
	CustomerActivityTypeSMS          CustomerActivityType = "SMS"
	CustomerActivityTypeLine         CustomerActivityType = "LINE"
	CustomerActivityTypeOrder        CustomerActivityType = "ORDER"
	CustomerActivityTypeComplaint    CustomerActivityType = "COMPLAINT"
	CustomerActivityTypeReturn       CustomerActivityType = "RETURN"
	CustomerActivityTypeRefund       CustomerActivityType = "REFUND"
	CustomerActivityTypePointsAdjust CustomerActivityType = "POINTS_ADJUST"
	CustomerActivityTypeTierChange   CustomerActivityType = "TIER_CHANGE"
	CustomerActivityTypeStatusChange CustomerActivityType = "STATUS_CHANGE"
	CustomerActivityTypeReferral     CustomerActivityType = "REFERRAL"
	CustomerActivityTypeCampaign     CustomerActivityType = "CAMPAIGN"
)

func (c CustomerActivityType) IsValid() bool {
	switch c {
	case CustomerActivityTypeNote, CustomerActivityTypeCall, CustomerActivityTypeEmail,
		CustomerActivityTypeSMS, CustomerActivityTypeLine, CustomerActivityTypeOrder,
		CustomerActivityTypeComplaint, CustomerActivityTypeReturn, CustomerActivityTypeRefund,
		CustomerActivityTypePointsAdjust, CustomerActivityTypeTierChange,
		CustomerActivityTypeStatusChange, CustomerActivityTypeReferral, CustomerActivityTypeCampaign:
		return true
	}
	return false
}

func (c CustomerActivityType) String() string { return string(c) }

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
