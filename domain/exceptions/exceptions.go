package exceptions

type GofinError struct {
	Code         int
	InternalCode string
	Message      string
	Description  string
}

var (
	UNAUTHORIZED        = newError(401, "1000", "Credenciais Inválidas!", "As credenciais que você digitou são inválidas. Veja se digitou tudo corretamente e tente novamente.")
	DEFAULT             = newError(500, "0001", "Erro interno", "Um erro interno do servidor aconteceu. Muito triste =(")
	EMAIL_ALREADY_EXIST = newError(400, "2000", "Email já está em uso.", "Caso já possua uma conta, tente recuperar a senha.")
)

func (e *GofinError) Error() string {
	return e.Message
}

func SlimError(code int, message string) *GofinError {
	return &GofinError{
		Message: message,
		Code:    code,
	}
}

func newError(code int, internalCode string, msg string, description string) *GofinError {
	return &GofinError{
		Message:      msg,
		Code:         code,
		InternalCode: internalCode,
		Description:  description,
	}
}

func GetError(e *GofinError) *GofinError {
	return &GofinError{
		Message:     e.Message,
		Code:        e.Code,
		Description: e.Description,
	}
}
