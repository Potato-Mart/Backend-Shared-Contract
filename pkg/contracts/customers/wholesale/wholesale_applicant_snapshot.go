package wholesale

// WholesaleApplicantSnapshot freezes the applicant identity reviewed by an
// administrator without making Identity data part of the application key.
type WholesaleApplicantSnapshot struct {
	UserID        string `json:"user_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone,omitempty"`
	Role          string `json:"role,omitempty"`
	EmailVerified bool   `json:"email_verified"`
}
