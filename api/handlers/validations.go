package handlers

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Constraint  string
}

func ValidateStruct(object interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(object)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ErrorResponse{
				FailedField: err.StructNamespace(),
				Constraint:  err.Tag(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
