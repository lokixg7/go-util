package datetime

import "time"

// Common date and date-time layouts used by this package.
const (
	// DateLayout is the layout of a date like 2026-09-10.
	DateLayout = "2006-01-02"
	// DateTimeLayout is the layout of a date-time like 2026-09-10 12:30:45.
	DateTimeLayout = "2006-01-02 15:04:05"
)

// CurrentTimestamp returns the current Unix time in seconds.
func CurrentTimestamp() int64 {
	return time.Now().Unix()
}

// UnixToTime converts a Unix timestamp in seconds into a time.Time.
func UnixToTime(seconds int64) time.Time {
	return time.Unix(seconds, 0)
}

// TimeToUnix converts a time.Time into a Unix timestamp in seconds.
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

// Format formats t using layout. Layout uses the Go reference time format,
// e.g. time.RFC3339 or DateTimeLayout.
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatDate formats t as a date, e.g. 2026-09-10.
func FormatDate(t time.Time) string {
	return t.Format(DateLayout)
}

// FormatDateTime formats t as a date-time, e.g. 2026-09-10 12:30:45.
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeLayout)
}

// TimestampToDateTime formats a Unix timestamp in seconds as a date-time
// string, e.g. 2026-09-10 12:30:45.
func TimestampToDateTime(seconds int64) string {
	return FormatDateTime(UnixToTime(seconds))
}

// Parse parses value using layout. Layout uses the Go reference time format.
// As with time.Parse, a layout without a time zone indicator parses the value
// as UTC.
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// ParseDate parses a date string like 2026-09-10. See Parse for time zone
// behavior.
func ParseDate(value string) (time.Time, error) {
	return time.Parse(DateLayout, value)
}

// ParseDateTime parses a date-time string like 2026-09-10 12:30:45. See Parse
// for time zone behavior.
func ParseDateTime(value string) (time.Time, error) {
	return time.Parse(DateTimeLayout, value)
}

// DateTimeToTimestamp parses a date-time string like 2026-09-10 12:30:45 as
// UTC (see Parse) and returns its Unix timestamp in seconds.
func DateTimeToTimestamp(value string) (int64, error) {
	t, err := ParseDateTime(value)
	if err != nil {
		return 0, err
	}
	return TimeToUnix(t), nil
}

// StartOfDay returns the first instant of t's local day (00:00:00).
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the last instant of t's local day (23:59:59.999999999).
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// StartOfMonth returns the first instant of t's local month.
func StartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the last instant of t's local month.
func EndOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
}
