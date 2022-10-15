package validators

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type apiError struct {
	Param   string
	Message string
}

func ParseError(err error) []apiError {
	var validationErros validator.ValidationErrors
	if errors.As(err, &validationErros) {
		out := make([]apiError, len(validationErros))
		for i, fe := range validationErros {
			out[i] = apiError{fe.Field(), parseErrorMessage(fe)}
		}
		return out
	}

	return nil
}

func parseErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "gt":
		return fmt.Sprintf("This field must be greater than %s", fieldError.Param())
	case "gte":
		return fmt.Sprintf("This field must be greater or equal than %s", fieldError.Param())
	case "lt":
		return fmt.Sprintf("This field must be less than %s", fieldError.Param())
	case "lte":
		return fmt.Sprintf("This field must be less or equal than %s", fieldError.Param())
	}

	return fieldError.Error() // default error
}
