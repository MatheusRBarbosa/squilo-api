-- +goose Up
-- +goose StatementBegin
CREATE TABLE vault_types(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text NOT NULL,
    "allowPositiveTransactions" boolean NOT NULL,
    "allowNegativeTransactions" boolean NOT NULL,
    "createdAt" timestamp NOT NULL,
    "updatedAt" timestamp NOT NULL,
    "deletedAt" timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vault_types
-- +goose StatementEnd
