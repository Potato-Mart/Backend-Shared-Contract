package temporal

// Date is a calendar date in YYYY-MM-DD format.
//
// Use Date for date-only JSON fields where a timestamp or timezone would
// change the meaning at service boundaries.
type Date string
