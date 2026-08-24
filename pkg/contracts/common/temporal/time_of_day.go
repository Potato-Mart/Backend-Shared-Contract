package temporal

// TimeOfDay is a wall-clock time in 24-hour "HH:MM" format.
//
// Use TimeOfDay for time-only contract fields (cut-off windows and recurring
// schedules) where a full timestamp or timezone-qualified type would change
// the meaning at service boundaries.
type TimeOfDay string
