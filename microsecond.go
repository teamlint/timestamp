package timestamp

import (
	"strconv"
	"time"
)

// Microsecond is the Unix timestamp, the number of microseconds
// elapsed since January 1, 1970 UTC.
type Microsecond int64

// FromTimeToMicrosecond returns t as a Microsecond.
func FromTimeToMicrosecond(t time.Time) Microsecond {
	return Microsecond(t.UnixNano() / 1e3)
}

// ToTime returns m as a time.Time.
func (m Microsecond) ToTime() time.Time {
	return time.Unix(int64(m)/1e6, (int64(m)%1e6)*1e3)
}

// Add returns m+d.
func (m Microsecond) Add(d time.Duration) Microsecond {
	return m + Microsecond(d/time.Microsecond)
}

// Sub returns the duration m-t.
// To compute m-d for duration, use m.Add(-d).
func (m Microsecond) Sub(t Microsecond) time.Duration {
	return time.Duration(m-t) * time.Microsecond
}

// Int64 int64 microsecond
func (m Microsecond) Int64() int64 {
	return int64(m)
}

// String string microsecond
func (m Microsecond) String() string {
	return strconv.FormatInt(int64(m), 10)
}

// Format format microsecond
func (m Microsecond) Format(layout ...string) string {
	l := RFC3339MicroZ
	if len(l) > 0 {
		l = layout[0]
	}
	return m.ToTime().Format(l)
}
