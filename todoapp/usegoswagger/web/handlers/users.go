package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/k0kubun/pp"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations/user"
)

// CreateUser :
func CreateUser(params user.CreateUserParams) middleware.Responder {
	pp.Println(params)
	return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
}
