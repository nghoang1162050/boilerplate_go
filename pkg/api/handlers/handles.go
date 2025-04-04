package handlers

import (
	constants "boilerplate_go/pkg/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Success(payload interface{}) *ApiResponse {
	return BuildResponse("200", "Success", []string{}, payload)
}

func Accepted() *ApiResponse {
	return BuildResponse(
		"200",
		"Success",
		[]string{},
		nil)
}

func Error(status_code string, err error) *ApiResponse {
	return BuildResponse(
		status_code,
		"Error",
		[]string{err.Error()},
		nil)
}

func ValidationErrors(errs error) *ApiResponse {
	payload := []FieldValidationError{}
	for _, err := range errs.(validator.ValidationErrors) {
		errObj := &FieldValidationError{}
		errObj.Field = err.Field()
		errObj.Namespace = err.Namespace()
		errObj.Kind = err.Kind().String()
		errObj.Value = err.Value()
		errObj.Error = fmt.Sprintf("%s %s", err.Tag(), err.Param())
		payload = append(payload, *errObj)
	}
	
	return BuildResponse(
		constants.STATUS_CODE_VALIDATION_ERROR,
		constants.MSG_VALIDATION_ERROR,
		[]string{constants.MSG_VALIDATION_ERROR},
		payload)
}