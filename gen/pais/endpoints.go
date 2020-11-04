// Code generated by goa v3.2.5, DO NOT EDIT.
//
// pais endpoints
//
// Command:
// $ goa gen github.com/selmison/dojo-1/design

package pais

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "pais" service endpoints.
type Endpoints struct {
	CreatePais endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "pais" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreatePais: NewCreatePaisEndpoint(s),
	}
}

// Use applies the given middleware to all the "pais" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.CreatePais = m(e.CreatePais)
}

// NewCreatePaisEndpoint returns an endpoint function that calls the method
// "create_pais" of service "pais".
func NewCreatePaisEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatePaisDTO)
		return s.CreatePais(ctx, p)
	}
}
