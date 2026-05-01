package common

// Date is a calendar date in YYYY-MM-DD format.
//
// Use Date for date-only contract fields where a timestamp, timezone, or
// database-specific date type would change the meaning at service boundaries.
type Date string

