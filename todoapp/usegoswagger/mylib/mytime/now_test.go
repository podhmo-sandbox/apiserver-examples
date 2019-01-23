package mytime_test

import (
	"testing"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/mylib/mytime"
)

func useNow(p mytime.NowProvider) mytime.Time {
	return p.Now()
}

func TestFixedNow(t *testing.T) {
	now := mytime.MustParseRawTime("2000-01-01T00:00:00Z")
	p := mytime.NewNowProvider(mytime.WithRawTime(now))

	got := useNow(p).String()
	if expected := "2000-01-01T00:00:00Z"; got != expected {
		t.Errorf("string()	expected:%s	but:%s", expected, got)
	}
}
