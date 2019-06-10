package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/useoapi-codegen/cmd/useoapi-server/api"
)

// API :
type API struct {
}

// CreateUser :
func (api *API) CreateUser(ctx echo.Context) error {
	return nil
}

func main() {
	e := echo.New()
	var myapi API
	api.RegisterHandlers(e, &myapi)
}
