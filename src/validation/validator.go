package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Validationerror struct {
	Property string
	Tag      string
	Value    string
}

func MakeValidationError(err error) *[]Validationerror {
	var validationerrors []Validationerror
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			element := Validationerror{}
			element.Property = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			validationerrors = append(validationerrors, element)
		}
		return &validationerrors
	}
	return nil
}