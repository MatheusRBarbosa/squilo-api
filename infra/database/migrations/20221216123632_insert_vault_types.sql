-- +goose Up
-- +goose StatementBegin
INSERT INTO vault_types(name, description, "createdAt", "updatedAt", "allowPositiveTransactions", "allowNegativeTransactions") values ('Geral', 'Este é o tipo de cofre padrão, onde possui o resultado das suas transações.', now(), now(), true, true);
INSERT INTO vault_types(name, description, "createdAt", "updatedAt", "allowPositiveTransactions", "allowNegativeTransactions") values ('Cota', 'Este tipo de cofre, todo mês, reinicia para um valor pré-definido em uma data a sua escolha. Esse tipo de cofre aceita apenas transações negativas.', now(), now(), false, true);
INSERT INTO vault_types(name, description, "createdAt", "updatedAt", "allowPositiveTransactions", "allowNegativeTransactions") values ('Bolsa', 'Este tipo de cofre, todo mês, adiciona um valor pré-definio em uma data a sua escolha, ele também vai acumular o resultado das suas transações.', now(), now(), true, true);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'nothing to down';
-- +goose StatementEnd
