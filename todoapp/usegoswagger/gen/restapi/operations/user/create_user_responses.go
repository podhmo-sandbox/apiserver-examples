// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	viewmodel "github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/gen/viewmodel"
)

// CreateUserCreatedCode is the HTTP code returned for type CreateUserCreated
const CreateUserCreatedCode int = 201

/*CreateUserCreated created

swagger:response createUserCreated
*/
type CreateUserCreated struct {

	/*
	  In: Body
	*/
	Payload *viewmodel.User `json:"body,omitempty"`
}

// NewCreateUserCreated creates CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {

	return &CreateUserCreated{}
}

// WithPayload adds the payload to the create user created response
func (o *CreateUserCreated) WithPayload(payload *viewmodel.User) *CreateUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user created response
func (o *CreateUserCreated) SetPayload(payload *viewmodel.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
