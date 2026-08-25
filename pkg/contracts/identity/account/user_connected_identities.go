package account

// UserConnectedIdentities is an account-connection projection: every login
// provider linked to one canonical user. The backend assembles it from the
// user's AuthIdentity rows; the frontend renders it on the "connected accounts"
// screen (link Google / Apple / Line / Microsoft / Discord, etc.).
type UserConnectedIdentities struct {
	UserID     string                `json:"user_id"`
	Identities []AuthIdentitySummary `json:"identities"`
}
