package configuration

import (
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations/user"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/handlers"
)

// ConfigureHandlers :
func ConfigureHandlers(c *Configurator, api *operations.UsegoswaggerAPI) {
	registry := c.Registry
	{
		h := &handlers.UserHandler{Registry: registry}
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(h.CreateUser)
	}
}
