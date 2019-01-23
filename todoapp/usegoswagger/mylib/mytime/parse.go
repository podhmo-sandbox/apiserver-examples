package mytime

import "time"

// MustParse :
func MustParse(s string) Time {
	return Time{Time: MustParseRawTime(s)}
}

// MustParseRawTime :
func MustParseRawTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}
