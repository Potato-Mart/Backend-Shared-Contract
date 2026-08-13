package account_enums

// NotificationTopicGroup identifies one fixed customer preference-centre
// topic group. The five groups form the rows of the topic-group x channel
// preference matrix.
type NotificationTopicGroup string

const (
	NotificationTopicGroupPayment        NotificationTopicGroup = "payment"
	NotificationTopicGroupOrder          NotificationTopicGroup = "order"
	NotificationTopicGroupReceipt        NotificationTopicGroup = "receipt"
	NotificationTopicGroupProductUpdates NotificationTopicGroup = "product_updates"
	NotificationTopicGroupPromotions     NotificationTopicGroup = "promotions"
)

// IsValid reports whether g is a known NotificationTopicGroup value.
func (g NotificationTopicGroup) IsValid() bool {
	switch g {
	case NotificationTopicGroupPayment, NotificationTopicGroupOrder,
		NotificationTopicGroupReceipt, NotificationTopicGroupProductUpdates,
		NotificationTopicGroupPromotions:
		return true
	}
	return false
}

// String returns the wire value for g.
func (g NotificationTopicGroup) String() string { return string(g) }
