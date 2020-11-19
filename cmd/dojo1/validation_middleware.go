package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	goa "goa.design/goa/v3/pkg"
	"gorm.io/gorm"

	decolarGen "github.com/selmison/dojo-1/gen/decolar"
	"github.com/selmison/dojo-1/pkg/core"
	"github.com/selmison/dojo-1/pkg/decolar"
)

func ValidationMiddleware(repo *gorm.DB) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			switch dto := request.(type) {
			case *decolarGen.CreatePaisDTO:
				if err := fieldShouldBeUnique(repo, "nome", dto.Nome, &decolar.Pais{}); err != nil {
					return nil, err
				}
			case *decolarGen.CreateCompaniaDTO:
				if err := fieldShouldBeUnique(repo, "nome", dto.Nome, &decolar.Compania{}); err != nil {
					return nil, err
				}
			case *decolarGen.CreateAeroportoDTO:
				if err := fieldShouldBeUnique(repo, "nome", dto.Nome, &decolar.Aeroporto{}); err != nil {
					return nil, err
				}
			}
			return next(ctx, request)
		}
	}
}

func fieldShouldBeUnique(repo *gorm.DB, fieldName, fieldValue string, iface interface{}) error {
	query := fmt.Sprintf("%s = ?", fieldName)
	v := reflect.ValueOf(iface).Interface()
	rows := repo.First(v, query, fieldValue).RowsAffected
	if rows > 0 {
		return &goa.ServiceError{
			Name:    "invalid_fields",
			ID:      goa.NewErrorID(),
			Message: fmt.Sprintf("the 'body.%s' %v", fieldName, core.ErrShouldBeUnique),
		}
	}
	return nil
}
