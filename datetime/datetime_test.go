package datetime

import (
	"testing"
	"time"
)

var fixed = time.Date(2026, time.September, 10, 12, 30, 45, 0, time.UTC)

// go test -v ./datetime -run '^TestFormatDate$'
func TestFormatDate(t *testing.T) {
	if got := FormatDate(fixed); got != "2026-09-10" {
		t.Errorf("FormatDate() = %q, want %q", got, "2026-09-10")
	}
}

// go test -v ./datetime -run '^TestFormatDateTime$'
func TestFormatDateTime(t *testing.T) {
	if got := FormatDateTime(fixed); got != "2026-09-10 12:30:45" {
		t.Errorf("FormatDateTime() = %q, want %q", got, "2026-09-10 12:30:45")
	}
}

// go test -v ./datetime -run '^TestFormatCustomLayout$'
func TestFormatCustomLayout(t *testing.T) {
	if got := Format(fixed, "2006/01/02"); got != "2026/09/10" {
		t.Errorf("Format() = %q, want %q", got, "2026/09/10")
	}
}

// go test -v ./datetime -run '^TestUnixRoundTrip$'
func TestUnixRoundTrip(t *testing.T) {
	got := UnixToTime(TimeToUnix(fixed))
	if !got.Equal(fixed) {
		t.Errorf("Unix round trip = %v, want %v", got, fixed)
	}
}

// go test -v ./datetime -run '^TestCurrentTimestamp$'
func TestCurrentTimestamp(t *testing.T) {
	before := time.Now().Unix()
	got := CurrentTimestamp()
	after := time.Now().Unix()
	if got < before || got > after {
		t.Errorf("CurrentTimestamp() = %d, want between %d and %d", got, before, after)
	}
}

// go test -v ./datetime -run '^TestParseDate$'
func TestParseDate(t *testing.T) {
	got, err := ParseDate("2026-09-10")
	if err != nil {
		t.Fatalf("ParseDate() error = %v", err)
	}
	if got.Format(DateLayout) != "2026-09-10" {
		t.Errorf("ParseDate() = %v, want date 2026-09-10", got)
	}
}

// go test -v ./datetime -run '^TestParseDateTime$'
func TestParseDateTime(t *testing.T) {
	got, err := ParseDateTime("2026-09-10 12:30:45")
	if err != nil {
		t.Fatalf("ParseDateTime() error = %v", err)
	}
	if got.Format(DateTimeLayout) != "2026-09-10 12:30:45" {
		t.Errorf("ParseDateTime() = %v, want 2026-09-10 12:30:45", got)
	}
}

// go test -v ./datetime -run '^TestDateTimeToTimestamp$'
func TestDateTimeToTimestamp(t *testing.T) {
	got, err := DateTimeToTimestamp("2026-09-10 12:30:45")
	if err != nil {
		t.Fatalf("DateTimeToTimestamp() error = %v", err)
	}
	if want := TimeToUnix(fixed); got != want {
		t.Errorf("DateTimeToTimestamp() = %d, want %d", got, want)
	}
}

// go test -v ./datetime -run '^TestStartAndEndOfDay$'
func TestStartAndEndOfDay(t *testing.T) {
	start := StartOfDay(fixed)
	if want := time.Date(2026, time.September, 10, 0, 0, 0, 0, time.UTC); !start.Equal(want) {
		t.Errorf("StartOfDay() = %v, want %v", start, want)
	}
	end := EndOfDay(fixed)
	if want := time.Date(2026, time.September, 10, 23, 59, 59, 999999999, time.UTC); !end.Equal(want) {
		t.Errorf("EndOfDay() = %v, want %v", end, want)
	}
}

// go test -v ./datetime -run '^TestStartAndEndOfMonth$'
func TestStartAndEndOfMonth(t *testing.T) {
	start := StartOfMonth(fixed)
	if want := time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC); !start.Equal(want) {
		t.Errorf("StartOfMonth() = %v, want %v", start, want)
	}
	end := EndOfMonth(fixed)
	if want := time.Date(2026, time.September, 30, 23, 59, 59, 999999999, time.UTC); !end.Equal(want) {
		t.Errorf("EndOfMonth() = %v, want %v", end, want)
	}
}
