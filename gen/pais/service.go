// Code generated by goa v3.2.5, DO NOT EDIT.
//
// pais service
//
// Command:
// $ goa gen github.com/selmison/dojo-1/design

package pais

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// O serviço pais executa operações em pais
type Service interface {
	// CreatePais implements create_pais.
	CreatePais(context.Context, *CreatePaisDTO) (res *PaisDTO, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "pais"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"create_pais"}

// CreatePaisDTO is the payload type of the pais service create_pais method.
type CreatePaisDTO struct {
	Nome string
}

// PaisDTO is the result type of the pais service create_pais method.
type PaisDTO struct {
	ID   string
	Nome string
}

// MakeInvalidFields builds a goa.ServiceError from an error.
func MakeInvalidFields(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid_fields",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
