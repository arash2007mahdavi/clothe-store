package validation

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IdValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(`^[a-zA-Z0-9_]{3,20}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}