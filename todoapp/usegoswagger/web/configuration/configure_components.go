package configuration

import (
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/mylib/mytime"
)

// ConfigureComponents :
func ConfigureComponents(c *Configurator, api *operations.UsegoswaggerAPI) {
	defer c.Registry.Unlock()
	c.Registry.Lock()
	if c.Registry.Now == nil {
		c.Registry.Now = mytime.NewNowProvider().Now
	}
}
