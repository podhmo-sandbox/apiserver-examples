package testutil

import (
	"github.com/go-openapi/loads"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/components"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/configuration"
)

// this package is tentative

// CreateAPI :
func CreateAPI() *operations.UsegoswaggerAPI {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err) // xxx
	}
	return operations.NewUsegoswaggerAPI(swaggerSpec)
}

// CreateServer :
func CreateServer(api *operations.UsegoswaggerAPI, options ...func(*components.Registry)) (*restapi.Server, func()) {
	server := restapi.NewServer(api)
	configuration.WithRegistry(func(r *components.Registry) {
		for _, op := range options {
			op(r)
		}
		server.ConfigureAPI()
	})
	return server, func() {
		// xxx:
		if err := server.Shutdown(); err != nil {
			panic(err)
		}
	}
}
