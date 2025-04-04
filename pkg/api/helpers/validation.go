package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validation *validator.Validate

func init() {
	validation = validator.New()
}

func Validate(s interface{}) error {
	return validation.Struct(s)
}

func Error(c echo.Context, err error, originalError error) error {
	c.Logger().Error(err, originalError)
	return err
}