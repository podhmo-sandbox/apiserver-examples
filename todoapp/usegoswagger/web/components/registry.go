package components

import (
	"sync"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/mylib/mytime"
)

// Registry : components registry
type Registry struct {
	Now mytime.NowProvider
	sync.Mutex
}

// WithNow :
func WithNow(now mytime.Time) func(*Registry) {
	return func(r *Registry) {
		defer r.Unlock()
		r.Lock()
		r.Now = mytime.NewNowProvider(mytime.WithRawTime(now.Time))
	}
}
