package timestamp

import (
	"strconv"
	"time"
)

// Nanosecond is the Unix timestamp, the number of nanoseconds
// elapsed since January 1, 1970 UTC.
type Nanosecond int64

// FromTimeToNanosecond returns t as a Nanosecond.
func FromTimeToNanosecond(t time.Time) Nanosecond {
	return Nanosecond(t.UnixNano())
}

// ToTime returns n as a time.Time.
func (n Nanosecond) ToTime() time.Time {
	return time.Unix(int64(n)/1e9, int64(n)%1e9)
}

// Add returns n+d.
func (n Nanosecond) Add(d time.Duration) Nanosecond {
	return n + Nanosecond(d/time.Nanosecond)
}

// Sub returns the duration n-t.
// To compute n-d for duration, use n.Add(-d).
func (n Nanosecond) Sub(t Nanosecond) time.Duration {
	return time.Duration(n-t) * time.Nanosecond
}

// Int64 int64 Nanosecond
func (m Nanosecond) Int64() int64 {
	return int64(m)
}

// String string nanosecond
func (m Nanosecond) String() string {
	return strconv.FormatInt(int64(m), 10)
}

// Format format nanosecond
func (m Nanosecond) Format(layout ...string) string {
	l := RFC3339NanoZ
	if len(l) > 0 {
		l = layout[0]
	}
	return m.ToTime().Format(l)
}
