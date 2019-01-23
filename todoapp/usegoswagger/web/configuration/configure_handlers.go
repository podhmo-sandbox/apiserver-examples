package configuration

import (
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations/user"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/handlers"
)

// ConfigureHandlers :
func ConfigureHandlers(c *Configurator, api *operations.UsegoswaggerAPI) {
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(handlers.CreateUser)
}
