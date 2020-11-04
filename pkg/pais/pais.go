package pais

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/selmison/dojo-1/gen/pais"
	"github.com/selmison/dojo-1/pkg/core"
)

// Pais represents a single pais.
type Pais struct {
	gorm.Model
	ID   string `gorm:"primarykey"`
	Nome string `validate:"required,not_blank"`
}

func (a *Pais) Validate() error {
	err := core.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		err := pais.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), core.ErrIsNotValid),
		)
		return err
	}
	return nil
}
