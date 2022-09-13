-- +goose Up
-- +goose StatementBegin
SELECT 'CREATE TABLE users (
    id int NOT NULL PRIMARY KEY IDENTITY(1,1),
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL
)';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'DROP TABLE users';
-- +goose StatementEnd
