package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/pasannissanka/learning-golang/go-crud-rest-api/types"
)

func ExtractValidationErrors(err error) []*types.IError {
	var errors []*types.IError
	for _, err := range err.(validator.ValidationErrors) {
		var el types.IError
		el.Field = err.Field()
		el.Tag = err.Tag()
		el.Value = err.Param()
		errors = append(errors, &el)
	}
	return errors
}
