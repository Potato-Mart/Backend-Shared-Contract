package common

// Date is a calendar date in YYYY-MM-DD format.
//
// Use Date for date-only contract fields where a timestamp, timezone, or
// database-specific date type would change the meaning at service boundaries.
type Date string

// TimeOfDay is a wall-clock time in 24-hour "HH:MM" format.
//
// Use TimeOfDay for time-only contract fields (cut-off windows, quiet
// hours) where a full timestamp or timezone-qualified type would change
// the meaning at service boundaries.
type TimeOfDay string

