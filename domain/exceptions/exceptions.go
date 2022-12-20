package exceptions

type GofinError struct {
	Code         int
	InternalCode string
	Message      string
	Description  string
}

var (
	DEFAULT = newError(500, "0001", "Erro interno.", "Um erro interno do servidor aconteceu. Muito triste =(")

	// 1xxx - Generic validations
	UNAUTHORIZED = newError(401, "1001", "Credenciais Inválidas!", "As credenciais que você digitou são inválidas. Veja se digitou tudo corretamente e tente novamente.")

	// 2xxx - Email
	EMAIL_ALREADY_EXIST = newError(400, "2001", "Email já está em uso.", "Caso já possua uma conta, tente recuperar a senha.")

	// 3xxx - Vault
	VAULT_NOT_FOUND = newError(404, "3001", "Não encontrado.", "O cofre procurado não foi encontrado.")

	// 4xxx - Transaction
	TRANSACTION_NOT_FOUND            = newError(404, "4001", "Não encontrado.", "A transação procurada não foi encontrada.")
	TRANSACTION_WRONG_TYPE           = newError(400, "4002", "Transação não permitida.", "Verifique se o tipo de cofre permite essa transação.")
	TRANSACTION_NOT_BELONGS_TO_VAULT = newError(400, "4003", "Transação não pertence a esse cofre.", "Verifique se essa transação realmente pertence a esse cofre.")

	// 5xxx - User
	USER_NOT_FOUND = newError(404, "5001", "Não encontrado.", "O usuário procurado não foi encontrado.")
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
