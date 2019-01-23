package configuration

import (
	"sync"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/components"
)

var (
	current *Configurator
	mu      sync.Mutex
)

// GetConfigurator :
func GetConfigurator() *Configurator {
	return current
}

// Configurator : components configurator
type Configurator struct {
	Configuration func(api *operations.UsegoswaggerAPI)
	Registry      *components.Registry
}

// Configure :
func (c *Configurator) Configure(api *operations.UsegoswaggerAPI) {
	ConfigureComponents(c, api)
	ConfigureHandlers(c, api)
}

// WithRegistry for testing
func WithRegistry(callback func(*components.Registry)) {
	mu.Lock()
	defer mu.Unlock()
	original := current
	defer func() {
		current = original // rollback
	}()

	// copy
	c := *current
	r := *current.Registry
	c.Registry = &r

	current = &c
	callback(&r)
}

func init() {
	defer mu.Unlock()
	mu.Lock()
	if current == nil {
		current = &Configurator{Registry: &components.Registry{}}
	}
}
