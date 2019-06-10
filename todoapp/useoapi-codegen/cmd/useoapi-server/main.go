package main

import (
	echo "github.com/labstack/echo/v4"
	api "github.com/podhmo-sandbox/apiserver-examples/todoapp/useoapi-codegen/Openapi3"
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
	e.Logger.Fatal(e.Start(":8080"))
}
