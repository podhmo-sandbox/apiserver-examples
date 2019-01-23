package mytime

import "time"

// Time :
type Time struct {
	time.Time
}

// String :
func (t Time) String() string {
	return t.Time.Format(time.RFC3339)
}
