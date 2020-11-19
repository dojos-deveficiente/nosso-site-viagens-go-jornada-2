// Code generated by goa v3.2.5, DO NOT EDIT.
//
// decolar HTTP server types
//
// Command:
// $ goa gen github.com/selmison/dojo-1/design

package server

import (
	"unicode/utf8"

	decolar "github.com/selmison/dojo-1/gen/decolar"
	goa "goa.design/goa/v3/pkg"
)

// CreatePaisRequestBody is the type of the "decolar" service "create_pais"
// endpoint HTTP request body.
type CreatePaisRequestBody struct {
	Nome *string `form:"nome,omitempty" json:"nome,omitempty" xml:"nome,omitempty"`
}

// CreateCompaniaRequestBody is the type of the "decolar" service
// "create_compania" endpoint HTTP request body.
type CreateCompaniaRequestBody struct {
	Nome   *string `form:"nome,omitempty" json:"nome,omitempty" xml:"nome,omitempty"`
	PaisID *string `form:"pais_id,omitempty" json:"pais_id,omitempty" xml:"pais_id,omitempty"`
}

// CreateAeroportoRequestBody is the type of the "decolar" service
// "create_aeroporto" endpoint HTTP request body.
type CreateAeroportoRequestBody struct {
	Nome   *string `form:"nome,omitempty" json:"nome,omitempty" xml:"nome,omitempty"`
	PaisID *string `form:"pais_id,omitempty" json:"pais_id,omitempty" xml:"pais_id,omitempty"`
}

// CreatePaisResponseBody is the type of the "decolar" service "create_pais"
// endpoint HTTP response body.
type CreatePaisResponseBody struct {
	ID   string `form:"id" json:"id" xml:"id"`
	Nome string `form:"nome" json:"nome" xml:"nome"`
}

// CreateCompaniaResponseBody is the type of the "decolar" service
// "create_compania" endpoint HTTP response body.
type CreateCompaniaResponseBody struct {
	ID        string `form:"id" json:"id" xml:"id"`
	Nome      string `form:"nome" json:"nome" xml:"nome"`
	PaisID    string `form:"pais_id" json:"pais_id" xml:"pais_id"`
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// CreateAeroportoResponseBody is the type of the "decolar" service
// "create_aeroporto" endpoint HTTP response body.
type CreateAeroportoResponseBody struct {
	ID        string `form:"id" json:"id" xml:"id"`
	Nome      string `form:"nome" json:"nome" xml:"nome"`
	PaisID    string `form:"pais_id" json:"pais_id" xml:"pais_id"`
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// CreatePaisInvalidFieldsResponseBody is the type of the "decolar" service
// "create_pais" endpoint HTTP response body for the "invalid_fields" error.
type CreatePaisInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateCompaniaInvalidFieldsResponseBody is the type of the "decolar" service
// "create_compania" endpoint HTTP response body for the "invalid_fields" error.
type CreateCompaniaInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateAeroportoInvalidFieldsResponseBody is the type of the "decolar"
// service "create_aeroporto" endpoint HTTP response body for the
// "invalid_fields" error.
type CreateAeroportoInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewCreatePaisResponseBody builds the HTTP response body from the result of
// the "create_pais" endpoint of the "decolar" service.
func NewCreatePaisResponseBody(res *decolar.PaisDTO) *CreatePaisResponseBody {
	body := &CreatePaisResponseBody{
		ID:   res.ID,
		Nome: res.Nome,
	}
	return body
}

// NewCreateCompaniaResponseBody builds the HTTP response body from the result
// of the "create_compania" endpoint of the "decolar" service.
func NewCreateCompaniaResponseBody(res *decolar.CompaniaDTO) *CreateCompaniaResponseBody {
	body := &CreateCompaniaResponseBody{
		ID:        res.ID,
		Nome:      res.Nome,
		PaisID:    res.PaisID,
		CreatedAt: res.CreatedAt,
	}
	return body
}

// NewCreateAeroportoResponseBody builds the HTTP response body from the result
// of the "create_aeroporto" endpoint of the "decolar" service.
func NewCreateAeroportoResponseBody(res *decolar.AeroportoDTO) *CreateAeroportoResponseBody {
	body := &CreateAeroportoResponseBody{
		ID:        res.ID,
		Nome:      res.Nome,
		PaisID:    res.PaisID,
		CreatedAt: res.CreatedAt,
	}
	return body
}

// NewCreatePaisInvalidFieldsResponseBody builds the HTTP response body from
// the result of the "create_pais" endpoint of the "decolar" service.
func NewCreatePaisInvalidFieldsResponseBody(res *goa.ServiceError) *CreatePaisInvalidFieldsResponseBody {
	body := &CreatePaisInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateCompaniaInvalidFieldsResponseBody builds the HTTP response body
// from the result of the "create_compania" endpoint of the "decolar" service.
func NewCreateCompaniaInvalidFieldsResponseBody(res *goa.ServiceError) *CreateCompaniaInvalidFieldsResponseBody {
	body := &CreateCompaniaInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateAeroportoInvalidFieldsResponseBody builds the HTTP response body
// from the result of the "create_aeroporto" endpoint of the "decolar" service.
func NewCreateAeroportoInvalidFieldsResponseBody(res *goa.ServiceError) *CreateAeroportoInvalidFieldsResponseBody {
	body := &CreateAeroportoInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreatePaisDTO builds a decolar service create_pais endpoint payload.
func NewCreatePaisDTO(body *CreatePaisRequestBody) *decolar.CreatePaisDTO {
	v := &decolar.CreatePaisDTO{
		Nome: *body.Nome,
	}

	return v
}

// NewCreateCompaniaDTO builds a decolar service create_compania endpoint
// payload.
func NewCreateCompaniaDTO(body *CreateCompaniaRequestBody) *decolar.CreateCompaniaDTO {
	v := &decolar.CreateCompaniaDTO{
		Nome:   *body.Nome,
		PaisID: *body.PaisID,
	}

	return v
}

// NewCreateAeroportoDTO builds a decolar service create_aeroporto endpoint
// payload.
func NewCreateAeroportoDTO(body *CreateAeroportoRequestBody) *decolar.CreateAeroportoDTO {
	v := &decolar.CreateAeroportoDTO{
		Nome:   *body.Nome,
		PaisID: *body.PaisID,
	}

	return v
}

// ValidateCreatePaisRequestBody runs the validations defined on
// create_pais_request_body
func ValidateCreatePaisRequestBody(body *CreatePaisRequestBody) (err error) {
	if body.Nome == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("nome", "body"))
	}
	return
}

// ValidateCreateCompaniaRequestBody runs the validations defined on
// create_compania_request_body
func ValidateCreateCompaniaRequestBody(body *CreateCompaniaRequestBody) (err error) {
	if body.Nome == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("nome", "body"))
	}
	if body.PaisID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pais_id", "body"))
	}
	if body.Nome != nil {
		if utf8.RuneCountInString(*body.Nome) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.nome", *body.Nome, utf8.RuneCountInString(*body.Nome), 1, true))
		}
	}
	if body.PaisID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.pais_id", *body.PaisID, goa.FormatUUID))
	}
	return
}

// ValidateCreateAeroportoRequestBody runs the validations defined on
// create_aeroporto_request_body
func ValidateCreateAeroportoRequestBody(body *CreateAeroportoRequestBody) (err error) {
	if body.Nome == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("nome", "body"))
	}
	if body.PaisID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pais_id", "body"))
	}
	if body.Nome != nil {
		if utf8.RuneCountInString(*body.Nome) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.nome", *body.Nome, utf8.RuneCountInString(*body.Nome), 1, true))
		}
	}
	if body.PaisID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.pais_id", *body.PaisID, goa.FormatUUID))
	}
	return
}
