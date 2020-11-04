package core

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	if err := Validate.RegisterValidation("not_blank", validators.NotBlank); err != nil {
		log.Fatalln(err)
	}
}
