package review

// RatingSummary is the public aggregate for a review subject. ScoreCounts is
// ordered by score: index 0 represents one-star reviews through index 4 for
// five-star reviews.
type RatingSummary struct {
	Subject      PublishedReviewSubject `json:"subject"`
	AverageScore float64                `json:"average_score"`
	RatingCount  int64                  `json:"rating_count"`
	ScoreCounts  [5]int64               `json:"score_counts"`
}
