package handlers

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	FailedField string
	Constraint  string
}

func ValidateStruct(object interface{}) []*ValidationError {
	var errors []*ValidationError
	validate := validator.New()
	err := validate.Struct(object)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ValidationError{
				FailedField: err.StructField(),
				Constraint:  err.Tag(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
