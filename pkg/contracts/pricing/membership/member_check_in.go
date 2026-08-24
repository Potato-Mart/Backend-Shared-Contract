package membership

import "time"

// MemberCheckIn records a daily programme check-in for point awards.
type MemberCheckIn struct {
	ID             string    `json:"id"`
	CustomerNumber string    `json:"customer_number"`
	CheckInDate    time.Time `json:"check_in_date"`
	StreakCount    int       `json:"streak_count"`
	PointsAwarded  int       `json:"points_awarded"`
	CreatedAt      time.Time `json:"created_at"`
}
