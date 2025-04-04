package helpers

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var alphaNumSpaceRegex = regexp.MustCompile(`^[\pL\pN\s]+$`)

func alphaNumSpace(fl validator.FieldLevel) bool {
    return alphaNumSpaceRegex.MatchString(fl.Field().String())
}

func RegisterCustomValidations(v *validator.Validate) {
    v.RegisterValidation("alphaNumSpace", alphaNumSpace)
}
