package timestamp

import (
	"strconv"
	"time"
)

// Second is the Unix timestamp, the number of seconds
// elapsed since January 1, 1970 UTC.
type Second int64

// FromTimeToSecond returns t as a Second.
func FromTimeToSecond(t time.Time) Second {
	return Second(t.Unix())
}

// ToTime returns s as a time.Time.
func (s Second) ToTime() time.Time {
	return time.Unix(int64(s), 0)
}

// Add returns s+d.
func (s Second) Add(d time.Duration) Second {
	return s + Second(d/time.Second)
}

// Sub returns the duration s-t.
// To compute s-d for duration, use s.Add(-d).
func (s Second) Sub(t Second) time.Duration {
	return time.Duration(s-t) * time.Second
}

// Int64 int64 second
func (m Second) Int64() int64 {
	return int64(m)
}

// String string second
func (m Second) String() string {
	return strconv.FormatInt(int64(m), 10)
}

// Format format second
func (m Second) Format(layout ...string) string {
	l := RFC3339Z
	if len(l) > 0 {
		l = layout[0]
	}
	return m.ToTime().Format(l)
}
