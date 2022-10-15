-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    "birthDate" date NOT NULL,
    password varchar(255) NOT NULL,
    "createdAt" timestamp NOT NULL,
    "updatedAt" timestamp NOT NULL,
    "deletedAt" timestamp
);

-- +goose Down
DROP TABLE users;
