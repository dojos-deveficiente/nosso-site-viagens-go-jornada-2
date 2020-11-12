package decolar

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	decolarGen "github.com/selmison/dojo-1/gen/decolar"

	"github.com/selmison/dojo-1/pkg/core"
)

// Compania represents a single companhia.
type Compania struct {
	gorm.Model
	ID     string `gorm:"primarykey"`
	Nome   string `validate:"required,not_blank"`
	PaisID string `validate:"required,not_blank"`
}

func (a *Compania) Validate() error {
	err := core.Validate.Struct(a)
	if err != nil {
		vErrs := err.(validator.ValidationErrors)
		err := decolarGen.MakeInvalidFields(
			fmt.Errorf("the '%s' field %w", vErrs[0].StructField(), core.ErrIsNotValid),
		)
		return err
	}
	return nil
}
