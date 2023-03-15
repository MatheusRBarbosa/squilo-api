package validators

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type apiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func ParseError(err error) []apiError {
	var validationErros validator.ValidationErrors
	var timeError *time.ParseError

	if errors.As(err, &validationErros) {
		out := make([]apiError, len(validationErros))
		for i, fe := range validationErros {
			out[i] = apiError{fe.Field(), parseErrorMessage(fe)}
		}
		return out
	} else if errors.As(err, &timeError) {
		out := make([]apiError, 1)
		out[0] = apiError{timeError.Value, fmt.Sprintf("Should have %s format", timeError.Layout)}

		return out
	}

	return nil
}

func parseErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "Este campo é obrigatório"
	case "email":
		return "E-mail inválido"
	case "gt":
		return fmt.Sprintf("Este campo deve ser maior que %s", fieldError.Param())
	case "gte":
		return fmt.Sprintf("Este campo deve ser maior ou igual a %s", fieldError.Param())
	case "lt":
		return fmt.Sprintf("Este campo deve ser menor que %s", fieldError.Param())
	case "lte":
		return fmt.Sprintf("Este campo deve ser menor ou igual a %s", fieldError.Param())
	}

	return fieldError.Error() // default error
}
