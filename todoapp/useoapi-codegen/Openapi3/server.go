// Package Openapi3 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package Openapi3

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"strings"
)

// Datetime defines component schema for datetime.
type Datetime string

// User defines component schema for user.
type User struct {
	CreatedAt *Datetime `json:"createdAt,omitempty"`
	Id        *string   `json:"id,omitempty"`
	Name      string    `json:"name"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	//  (POST /users)
	CreateUser(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST("/users", wrapper.CreateUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RSwW7UMBD9ldXA0Y2dcln5BNy40wvVHgZ7uusqsY09W1FF/nfkSRMoIHHiNsk8v/f8",
	"nhdwac4pUuQKdoHqLjSjjB6ZOMzUZ/qOc54ILNya8XhjxhszfjbGGmO/gAJ+zn1XuYR4hqbgWqn0c7mk",
	"TIUDCaMrhEz+A/ePt4UewMIb/dOAflHXu3RTEHxH/6EQcXX226IpKPTtGgp5sPcr6rQbTF8fyTG0Dgvx",
	"IYmpFBmdWKIZwwQWZnr/cuEhEsMmBjOBgmvpkAtztlr/CmsKOLCEdFfpkDCHg0uezhRBwROVGlIEC2YY",
	"B9PRKVPEHMDCu8EMBhRk5IskpXuAMuVUxZun6krIvHKsSR466uCREYStYF9/8jvgrtewJkKVPyb/vN2Y",
	"orBizlNwckw/1k69PYF/VSQNt9d5c7mS/Kg5xbp2fmvG/6D5tzS8tM94rr15wZ4EXKk8SZr3y+v2puRw",
	"uqTK9miORvcu2mmnWLbaV9lT+xEAAP//QTqY1SwDAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

