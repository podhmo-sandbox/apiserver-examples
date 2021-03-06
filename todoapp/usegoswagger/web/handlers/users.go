package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	operations "github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/restapi/operations/user"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/viewmodel"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/components"
)

// UserHandler :
type UserHandler struct {
	*components.Registry
}

// CreateUser :
func (h *UserHandler) CreateUser(params operations.CreateUserParams) middleware.Responder {
	// todo: validation
	return operations.NewCreateUserCreated().WithPayload(&viewmodel.User{
		ID:        "1", // xxx
		Name:      params.Body.Name,
		CreatedAt: viewmodel.Datetime(h.Now().String()),
	})
}
