package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	operations "github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations/user"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/viewmodel"
)

// CreateUser :
func CreateUser(params operations.CreateUserParams) middleware.Responder {
	// todo: validation
	return operations.NewCreateUserCreated().WithPayload(&viewmodel.User{
		ID:        "1", // xxx
		Name:      params.Body.Name,
		CreatedAt: viewmodel.Datetime(time.Now().Format(time.RFC3339)), // xxx
	})
}
