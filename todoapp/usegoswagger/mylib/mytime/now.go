package mytime

import "time"

// NowProvider :
type NowProvider interface {
	Now() Time
}

type nowProvider struct {
	nowFunc func() time.Time
}

// NewNowProvider :
func NewNowProvider(options ...func(NowProvider)) NowProvider {
	p := &nowProvider{
		nowFunc: time.Now,
	}
	for _, op := range options {
		op(p)
	}
	return p
}

// WithRawTime :
func WithRawTime(now time.Time) func(p NowProvider) {
	return func(p NowProvider) {
		p.(*nowProvider).nowFunc = func() time.Time {
			return now
		}
	}
}

// Now :
func (p *nowProvider) Now() Time {
	return Time{Time: p.nowFunc()}
}
