-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    "vaultId" int NOT NULL,
    value float NOT NULL,
    "date" date NOT NULL,
    observation VARCHAR(255), 
    "createdAt" timestamp NOT NULL,
    "updatedAt" timestamp NOT NULL,
    FOREIGN KEY ("vaultId") REFERENCES vaults(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
