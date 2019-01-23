package testutil

import (
	"sync"

	"github.com/go-openapi/loads"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations"
)

// this package is tentative

var (
	api *operations.UsegoswaggerAPI
	mu  sync.Mutex
)

// CreateAPI :
func CreateAPI() *operations.UsegoswaggerAPI {
	defer mu.Unlock()
	mu.Lock()
	if api != nil {
		return api
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err) // xxx
	}
	api = operations.NewUsegoswaggerAPI(swaggerSpec)
	return api
}

// CreateServer :
func CreateServer() (*restapi.Server, func()) {
	api := CreateAPI()
	server := restapi.NewServer(api)
	server.ConfigureAPI()
	return server, func() {
		// xxx:
		if err := server.Shutdown(); err != nil {
			panic(err)
		}
	}
}
