-- +goose Up
-- +goose StatementBegin
CREATE TABLE vaults(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text,
    configs json NOT NULL,
    "userId" int NOT NULL,
    "typeId" int NOT NULL,
    "createdAt" timestamp NOT NULL,
    "updatedAt" timestamp NOT NULL,
    "deleted_at" timestamp,
    FOREIGN KEY ("userId") REFERENCES users(id),
    FOREIGN KEY ("typeId") REFERENCES vault_types(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vaults;
-- +goose StatementEnd
